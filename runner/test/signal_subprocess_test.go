package test

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/digital-ai/release-integration-sdk-go/runner"
	"github.com/digital-ai/release-integration-sdk-go/task"
	"github.com/digital-ai/release-integration-sdk-go/task/command"
)

// This file provides an isolated-subprocess regression test verifying that
// SIGABRT handling is restored across sequential daemon-mode Run invocations,
// guarding against reintroduction of the signal-handler leak.

// helperEnv gates TestHelperDaemon so it only runs when invoked as a subprocess.
const helperEnv = "RUN_SIGNAL_HELPER"

// TestSignalHandledAcrossDaemonInvocations runs the daemon scenario in a
// subprocess and asserts SIGABRT is handled rather than defaulted.
func TestSignalHandledAcrossDaemonInvocations(t *testing.T) {
	cmd := exec.Command(os.Args[0], "-test.run=TestHelperDaemon", "-test.v")
	cmd.Env = append(os.Environ(), helperEnv+"=1")

	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("helper subprocess failed (SIGABRT handling not restored / abort not delivered): %v\n--- child output ---\n%s", err, out)
	}
	if !strings.Contains(string(out), "HELPER_OK") {
		t.Fatalf("helper subprocess did not report success\n--- child output ---\n%s", out)
	}
}

// TestHelperDaemon is the subprocess body; it runs only when RUN_SIGNAL_HELPER=1.
func TestHelperDaemon(t *testing.T) {
	if os.Getenv(helperEnv) != "1" {
		t.Skip("helper process; only runs when " + helperEnv + "=1")
	}

	// First invocation completes normally: registers then releases its SIGABRT
	// handler via the deferred signal.Stop in Run.
	immediateFactory := abortTestFactory{immediate: true, withAbortCommand: true}
	firstRunner := runner.NewCommandRunner(func(_ task.InputContext) (command.CommandFactory, error) {
		return immediateFactory, nil
	})
	if res := runWithTimeout(t, firstRunner, mainInputContext()); res == nil {
		t.Fatal("first invocation returned nil")
	}

	// Second invocation is long-running and must catch the SIGABRT delivered
	// after Run has entered its select.
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
		t.Fatal("expected active invocation to be aborted; SIGABRT handling not restored")
	} else if _, ok := err.(*task.AbortError); !ok {
		t.Fatalf("expected *task.AbortError, got %T: %v", err, err)
	}

	// Sentinel string the parent scans for to confirm clean success.
	t.Log("HELPER_OK")
}
