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

var inputContextUpdateChan = make(chan bool, 1)

type inputContextUpdatedFunc func()

func StartInputContextWatcher(onInputContextUpdateFunc inputContextUpdatedFunc) {
	stop := make(chan struct{})
	defer close(stop)

	go func() {
		for {
			select {
			case <-inputContextUpdateChan:
				onInputContextUpdateFunc()
			}
		}
	}()

	err := startInputSecretWatcher(stop)
	if err != nil {
		klog.Info("Failed to start secret watcher: ", err)
		return
	}

	<-time.After(1 * time.Minute) //TODO read from configuration
}

func startInputSecretWatcher(stop chan struct{}) error {
	clientset, err := k8s.GetClientset()
	if err != nil {
		klog.Warningf("Cannot get clientset for fetching Secret: %s", err)
		return err
	}

	watchlist := cache.NewListWatchFromClient(
		clientset.CoreV1().RESTClient(),
		"secrets",
		os.Getenv(task.RunnerNamespace),
		fields.OneTermEqualSelector("metadata.name", os.Getenv(task.InputContextSecretName)),
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
					klog.Infof("Input values have been updated, starting new execution")
					inputContextUpdateChan <- true
				}
			},
		},
	)

	go controller.Run(stop)

	return nil
}
