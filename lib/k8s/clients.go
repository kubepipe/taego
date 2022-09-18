package mk8s

import (
	"taego/lib/mlog"

	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var clientset *kubernetes.Clientset

func init() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		mlog.Error("load kubeconfig failed", zap.Error(err))
		return
	}
	// creates the clientset
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		mlog.Error("create k8s clientset failed", zap.Error(err))
		return
	}
}

func GetInClusterClientset() *kubernetes.Clientset {
	if clientset == nil {
		return &kubernetes.Clientset{}
	}
	return clientset
}
