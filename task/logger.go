package task

import "k8s.io/klog/v2"

func Comment(comment string) {
	klog.Infof("##[start: comment]%s##[end: comment]", comment)
}

func Status(status string) {
	klog.Infof("##[start: status]%s##[end: status]", status)
}
