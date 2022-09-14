package mk8s

import (
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var clientset *kubernetes.Clientset

func init() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatal(err)
	}
	// creates the clientset
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
}

func GetInClusterClientset() *kubernetes.Clientset {
	if clientset == nil {
		return &kubernetes.Clientset{}
	}
	return clientset
}
