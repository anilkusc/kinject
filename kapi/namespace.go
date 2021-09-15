package kapi

import (
	"context"
	"log"

	"github.com/anilkusc/kinject/models"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListNamespaces(client *kubernetes.Clientset) []models.MyNamespace {
	var namespaces []models.MyNamespace
	listedNamespaces, err := client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Println(err)
	}
	for _, listedNamespace := range listedNamespaces.Items {
		if listedNamespace.Name == "kube-node-lease" || listedNamespace.Name == "kube-public" || listedNamespace.Name == "kube-system" {
			continue
		} else {
			namespace := models.MyNamespace{Name: listedNamespace.Name}
			namespaces = append(namespaces, namespace)
		}
	}
	return namespaces
}
