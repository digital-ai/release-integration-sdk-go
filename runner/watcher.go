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
	"strconv"
	"time"
)

const (
	EvictionTime = "EXECUTOR_EVICTION_TIME"
)

func StartInputContextWatcher(onInputContextUpdateFunc func()) {
	stop := make(chan struct{})
	defer close(stop)

	evictionTime, err := strconv.Atoi(os.Getenv(EvictionTime))
	if err != nil {
		klog.Errorf("Failed to read executor eviction time, the executor will be marked for shut down: ", err)
		return
	}
	t := time.NewTimer(time.Duration(evictionTime) * time.Second)
	defer t.Stop()

	// Resetting timer when new execution occurs
	wrappedOnInputContextUpdateFunc := func() {
		onInputContextUpdateFunc()
		if !t.Stop() {
			<-t.C
		}
		t.Reset(time.Duration(evictionTime) * time.Second)
	}

	err = startInputSecretWatcher(stop, wrappedOnInputContextUpdateFunc)
	if err != nil {
		klog.Info("Failed to start secret watcher: ", err)
		return
	}

	for {
		select {
		case <-t.C:
			klog.Info("Input Context Watcher reached eviction time: ", err)
			return
		}
	}
}

func startInputSecretWatcher(stop chan struct{}, onInputContextUpdateFunc func()) error {
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
					onInputContextUpdateFunc()
				}
			},
		},
	)

	go controller.Run(stop)

	return nil
}
