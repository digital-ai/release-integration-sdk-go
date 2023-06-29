package logger

import "k8s.io/klog/v2"

// init sets up the log filter.
func init() {
	klog.InitFlags(nil)
	klog.SetLogFilter(&secretLogFilter{})
}
