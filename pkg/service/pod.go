package service

import (
	"context"

	"github.com/zefrawendi/gin-client-go/pkg/client"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
)

func GetPods(namespace string) ([]v1.Pod, error) {
	clientSet, err := client.GetK8sClientSet()
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	ctx := context.Background()
	podList, err := clientSet.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		klog.Error()
		return nil, err
	}
	return podList.Items, nil
}
