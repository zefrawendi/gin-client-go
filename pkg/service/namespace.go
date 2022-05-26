package service

import (
	"context"

	"github.com/zefrawendi/gin-client-go/pkg/client"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
)

func GetNamespaces() ([]v1.Namespace, error) {
	clientSet, err := client.GetK8sClientSet()
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	ctx := context.Background()
	namespaceList, err := clientSet.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		klog.Error()
		return nil, err
	}
	return namespaceList.Items, nil
}
