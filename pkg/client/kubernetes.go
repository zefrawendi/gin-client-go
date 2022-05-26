package client

import (
	"errors"
	"flag"
	"sync"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog"
)

var onceClient sync.Once
var onceConfig sync.Once
var KubeConfig *rest.Config
var KubeClientSet *kubernetes.Clientset

func GetK8sClientSet() (*kubernetes.Clientset, error) {
	onceClient.Do(func() {
		config, err := GetRestConfig()
		if err != nil {
			klog.Fatal(err)
			return
		}
		KubeClientSet, err = kubernetes.NewForConfig(config)
		if err != nil {
			klog.Fatal(err)
			return
		}
	})
	return KubeClientSet, nil
}

func GetRestConfig() (*rest.Config, error) {
	onceConfig.Do(func() {
		var configPath *string
		var err error
		if home := homedir.HomeDir(); home != "" {
			configPath = flag.String("kubeconfig", home+"/.kube/config", "(optional) absolute path to the kubeconfig file")
		} else {
			klog.Fatal("read config error, config file not found in home directory")
			err = errors.New("read config error, config file not found in home directory")
			return
		}
		flag.Parse()
		KubeConfig, err = clientcmd.BuildConfigFromFlags("", *configPath)
		if err != nil {
			klog.Fatal(err)
			return
		}
	})

	return KubeConfig, nil
}
