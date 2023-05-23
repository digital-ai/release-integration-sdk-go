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

var ExecutionChan = make(chan bool, 1)

func StartExecutionRetriggerWatcher() {
	stop := make(chan struct{})
	defer close(stop)
	err := startSecretWatcher(stop)
	if err != nil {
		klog.Info("Failed to start secret watcher: ", err)
		return
	}

	<-time.After(1 * time.Minute)
}

func startSecretWatcher(stop chan struct{}) error {
	secretName := os.Getenv(task.InputContextSecretName)
	runnerNamespace := os.Getenv(task.RunnerNamespace)

	clientset, err := k8s.GetClientset()
	if err != nil {
		klog.Warningf("Cannot get clientset for fetching Secret: %s", err)
		return err
	}

	watchlist := cache.NewListWatchFromClient(
		clientset.CoreV1().RESTClient(),
		"secrets",
		runnerNamespace,
		fields.OneTermEqualSelector("metadata.name", secretName),
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

				// Check if 'input' field has changed
				if !bytes.Equal(oldInput, newInput) {
					klog.Infof("Secret values have been updated, starting new execution")
					ExecutionChan <- true
				}
			},
		},
	)

	go controller.Run(stop)

	return nil
}
