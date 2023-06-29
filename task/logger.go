package task

import "k8s.io/klog/v2"

// Comment logs a comment message.
func Comment(comment string) {
	klog.Infof("##[start: comment]%s##[end: comment]", comment)
}

// Status logs a status message.
func Status(status string) {
	klog.Infof("##[start: status]%s##[end: status]", status)
}
