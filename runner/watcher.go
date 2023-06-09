package runner

import (
	"bytes"
	"github.com/digital-ai/release-integration-sdk-go/k8s"
	"github.com/digital-ai/release-integration-sdk-go/task"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
	"os"
	"time"
)

const (
	ExecutionMode       = "EXECUTOR_EXECUTION_MODE"
	ExecutionModeDaemon = "daemon"
)

func StartInputContextWatcher(pluginVersion string, buildDate string, commandFactory FactoryBuilder, onInputContextUpdateFunc ExecuteFunction) {
	stop := make(chan struct{})

	err := startInputSecretWatcher(pluginVersion, buildDate, commandFactory, onInputContextUpdateFunc, stop)
	if err != nil {
		klog.Info("Failed to start secret watcher: ", err)
		return
	}

	<-stop
}

func startInputSecretWatcher(pluginVersion string, buildDate string, commandFactory FactoryBuilder, onInputContextUpdateFunc ExecuteFunction, stop chan struct{}) error {
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

				oldInput := oldSecret.Data[task.InputCategory]
				newInput := newSecret.Data[task.InputCategory]

				// Checking if 'input' field has changed
				if !bytes.Equal(oldInput, newInput) {
					klog.Infof("Detected input context value change")
					onInputContextUpdateFunc(pluginVersion, buildDate, commandFactory)
				}
			},
		},
	)

	go controller.Run(stop)

	return nil
}
