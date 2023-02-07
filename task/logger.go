package task

import "k8s.io/klog"

func Comment(comment string) {
	klog.Infof("##[start: comment]%s##[end: comment]", comment)
}

func Status(status string) {
	klog.Infof("##[start: status]%s##[end: status]", status)
}
