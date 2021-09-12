package kapi

import (
	"context"
	"log"

	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetDeployment(client *kubernetes.Clientset, namespace string, name string) *v1.Deployment {
	deployment, err := client.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		log.Println(err)
	}
	return deployment
}
