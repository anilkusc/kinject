package kapi

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func CreateClient(configPath string) *kubernetes.Clientset {
	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	return clientset
}
