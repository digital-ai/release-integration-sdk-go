package test

import (
	"context"
	"errors"
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/digital-ai/release-integration-sdk-go/runner"
	"github.com/digital-ai/release-integration-sdk-go/task"
	"github.com/digital-ai/release-integration-sdk-go/task/command"
)

// blockingCommand blocks until the context is cancelled, then returns.
// It records whether its context was cancelled so tests can assert that the
// main command is stopped on abort.
type blockingCommand struct {
	cancelled chan struct{}
}

func (c *blockingCommand) FetchResult(ctx context.Context) (*task.Result, error) {
	<-ctx.Done()
	if c.cancelled != nil {
		close(c.cancelled)
	}
	return nil, ctx.Err()
}

// abortCommand is the explicit abort handler for the main command.
type abortCommand struct{}

func (abortCommand) FetchResult(_ context.Context) (*task.Result, error) {
	return task.NewResult().String("aborted", "true"), nil
}

// immediateCommand returns immediately, simulating a task that finishes
// before any abort signal is delivered.
type immediateCommand struct{}

func (immediateCommand) FetchResult(_ context.Context) (*task.Result, error) {
	return task.NewResult().String("done", "true"), nil
}

// abortTestFactory wires command types to the fixtures above. The command type
// "main" maps to the blocking command and "main#abort" (see command.AbortCommand)
// maps to the abort handler. Set withAbortCommand=false to exercise the
// "no explicit abort command" path.
type abortTestFactory struct {
	blocking         *blockingCommand
	immediate        bool
	withAbortCommand bool
}

func (f abortTestFactory) InitCommand(commandType command.CommandType) (command.CommandExecutor, error) {
	switch commandType {
	case "main":
		if f.immediate {
			return immediateCommand{}, nil
		}
		return f.blocking, nil
	case command.AbortCommand("main"):
		if f.withAbortCommand {
			return abortCommand{}, nil
		}
		return nil, errors.New("no abort command defined")
	}
	return nil, errors.New("cannot find command type")
}

func mainInputContext() task.InputContext {
	return task.InputContext{
		Task: task.TaskContext{
			Type:       "main",
			Properties: []task.PropertyDefinition{},
		},
	}
}

// sendAbortSoon delivers SIGABRT to the current process after a short delay so
// that Run() has entered its select before the signal arrives.
func sendAbortSoon(t *testing.T) {
	t.Helper()
	go func() {
		time.Sleep(50 * time.Millisecond)
		if err := syscall.Kill(os.Getpid(), syscall.SIGABRT); err != nil {
			t.Errorf("failed to send SIGABRT: %v", err)
		}
	}()
}

func runWithTimeout(t *testing.T, r runner.Runner, ctx task.InputContext) *task.Result {
	t.Helper()
	resultCh := make(chan *task.Result, 1)
	go func() { resultCh <- r.Run(ctx) }()
	select {
	case res := <-resultCh:
		return res
	case <-time.After(5 * time.Second):
		t.Fatal("Run did not return within timeout; possible goroutine/signal leak")
		return nil
	}
}

// TestAbortWithExplicitAbortCommand verifies that a delivered SIGABRT triggers
// the explicit abort command and that the blocking main command's context is
// cancelled (no leaked main goroutine).
func TestAbortWithExplicitAbortCommand(t *testing.T) {
	blocking := &blockingCommand{cancelled: make(chan struct{})}
	factory := abortTestFactory{blocking: blocking, withAbortCommand: true}
	r := runner.NewCommandRunner(func(_ task.InputContext) (command.CommandFactory, error) {
		return factory, nil
	})

	sendAbortSoon(t)
	result := runWithTimeout(t, r, mainInputContext())

	if result == nil {
		t.Fatal("expected an abort result, got nil")
	}
	if _, err := result.Get(); err == nil {
		t.Fatal("expected aborted result to carry an AbortError")
	} else if _, ok := err.(*task.AbortError); !ok {
		t.Fatalf("expected *task.AbortError, got %T: %v", err, err)
	}

	select {
	case <-blocking.cancelled:
		// main command context was cancelled as expected
	case <-time.After(2 * time.Second):
		t.Fatal("main command context was not cancelled after abort (potential context leak)")
	}
}

// TestAbortWithoutExplicitAbortCommand verifies the plain-abort path where no
// explicit abort command is defined: the main command is cancelled and an
// abort result is returned.
func TestAbortWithoutExplicitAbortCommand(t *testing.T) {
	blocking := &blockingCommand{cancelled: make(chan struct{})}
	factory := abortTestFactory{blocking: blocking, withAbortCommand: false}
	r := runner.NewCommandRunner(func(_ task.InputContext) (command.CommandFactory, error) {
		return factory, nil
	})

	sendAbortSoon(t)
	result := runWithTimeout(t, r, mainInputContext())

	if result == nil {
		t.Fatal("expected an abort result, got nil")
	}
	if _, err := result.Get(); err == nil {
		t.Fatal("expected aborted result to carry an AbortError")
	} else if _, ok := err.(*task.AbortError); !ok {
		t.Fatalf("expected *task.AbortError, got %T: %v", err, err)
	}

	select {
	case <-blocking.cancelled:
	case <-time.After(2 * time.Second):
		t.Fatal("main command context was not cancelled after plain abort")
	}
}

// TestDaemonModeSignalNotLeaked simulates daemon-mode behaviour: a prior
// invocation completes normally, then a subsequent long-running invocation is
// aborted. With signal.Stop cleanup, the active invocation must receive the
// abort rather than a stale channel from the completed run swallowing it.
func TestDaemonModeSignalNotLeaked(t *testing.T) {
	// First invocation completes normally (registers then stops its signal channel).
	immediateFactory := abortTestFactory{immediate: true, withAbortCommand: true}
	firstRunner := runner.NewCommandRunner(func(_ task.InputContext) (command.CommandFactory, error) {
		return immediateFactory, nil
	})
	if res := runWithTimeout(t, firstRunner, mainInputContext()); res == nil {
		t.Fatal("first invocation returned nil")
	}

	// Second invocation is long-running and should catch the SIGABRT.
	blocking := &blockingCommand{cancelled: make(chan struct{})}
	secondFactory := abortTestFactory{blocking: blocking, withAbortCommand: true}
	secondRunner := runner.NewCommandRunner(func(_ task.InputContext) (command.CommandFactory, error) {
		return secondFactory, nil
	})

	sendAbortSoon(t)
	result := runWithTimeout(t, secondRunner, mainInputContext())

	if result == nil {
		t.Fatal("expected abort result from active invocation, got nil")
	}
	if _, err := result.Get(); err == nil {
		t.Fatal("expected active invocation to be aborted; signal may have leaked to a stale channel")
	} else if _, ok := err.(*task.AbortError); !ok {
		t.Fatalf("expected *task.AbortError, got %T: %v", err, err)
	}
}
