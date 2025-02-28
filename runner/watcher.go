package runner

import (
	"bytes"
	"context"
	"github.com/digital-ai/release-integration-sdk-go/k8s"
	"github.com/digital-ai/release-integration-sdk-go/task"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
	"os"
	"time"
)

// ExecutionMode is an environment variable that determines the execution mode of the executor.
const ExecutionMode = "EXECUTOR_EXECUTION_MODE"

// ExecutionModeDaemon represents the execution mode for running the executor as a daemon.
const ExecutionModeDaemon = "daemon"

// StartInputContextWatcher starts the input context watcher that monitors changes in the input context secret.
// When a change is detected, the provided onInputContextUpdateFunc is invoked.
func StartInputContextWatcher(onInputContextUpdateFunc ExecuteFunction, pluginVersion string, buildDate string, runner Runner) {
	stop := make(chan struct{})

	err := startInputSecretWatcher(onInputContextUpdateFunc, pluginVersion, buildDate, runner, stop)
	if err != nil {
		klog.Info("Failed to start secret watcher: ", err)
		return
	}

	<-stop
}

// startInputSecretWatcher starts the watcher for changes in the input context secret.
func startInputSecretWatcher(onInputContextUpdateFunc ExecuteFunction, pluginVersion string, buildDate string, runner Runner, stop chan struct{}) error {
	clientset, err := k8s.GetClientset()
	if err != nil {
		klog.Warningf("Cannot get clientset for fetching Secret: %s", err)
		return err
	}

	watchlist := cache.NewListWatchFromClient(
		clientset.CoreV1().RESTClient(),
		"secrets",
		os.Getenv(task.RunnerNamespace),
		fields.OneTermEqualSelector("metadata.name", os.Getenv(task.InputContextSecret)),
	)

	_, controller := cache.NewInformer(
		watchlist,
		&corev1.Secret{},
		time.Second*0,
		cache.ResourceEventHandlerFuncs{
			UpdateFunc: func(oldObj, newObj interface{}) {
				oldSecret := oldObj.(*corev1.Secret)
				newSecret := newObj.(*corev1.Secret)

				oldSessionKey := oldSecret.Data[task.InputContextSecretDataSessionKey]
				newSessionKey := newSecret.Data[task.InputContextSecretDataSessionKey]

				// Checking if 'input' field has changed
				if !bytes.Equal(oldSessionKey, newSessionKey) {
					klog.Infof("Detected input context value change")
					onInputContextUpdateFunc(pluginVersion, buildDate, runner, context.Background())
				}
			},
		},
	)

	go controller.Run(stop)

	return nil
}
