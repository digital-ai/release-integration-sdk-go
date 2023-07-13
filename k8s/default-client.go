package k8s

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
	"sync"
)

var lock = &sync.Mutex{}
var configLock = &sync.Mutex{}
var _clientset *kubernetes.Clientset
var _config *rest.Config

// GetConfig returns the Kubernetes REST configuration.
// If the configuration has not been initialized, it retrieves the in-cluster config.
// It returns the config or an error if there is any issue.
func GetConfig() (*rest.Config, error) {
	if _config == nil {
		configLock.Lock()
		defer configLock.Unlock()
		var err error

		_config, err = rest.InClusterConfig()
		if err != nil {
			return nil, err
		}
	}
	return _config, nil
}

// GetClientset returns the Kubernetes clientset.
// If the clientset has not been initialized, it creates a new one based on the config obtained from GetConfig.
// It returns the clientset or an error if there is any issue.
func GetClientset() (*kubernetes.Clientset, error) {
	if _clientset == nil {
		lock.Lock()
		defer lock.Unlock()
		if _clientset == nil {
			klog.Info("Creating Clientset singleton.")
			config, err := GetConfig()
			if err != nil {
				return nil, err
			}
			_clientset, err = kubernetes.NewForConfig(config)
			if err != nil {
				return nil, err
			}
		}
	}
	return _clientset, nil
}
