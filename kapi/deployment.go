package kapi

import (
	"context"
	"log"
	"strings"

	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
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
func ListDeployment(client *kubernetes.Clientset, namespace string) *v1.DeploymentList {
	deployments, err := client.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Println(err)
	}
	return deployments
}
func PatchDeploymentEnv(client *kubernetes.Clientset, namespace string, name string, env string) error {
	envs := strings.Split(env, ":")
	key := envs[0]
	value := envs[1]
	deployment := GetDeployment(client, namespace, name)
	isKeyExist := false
	envNumber := -1
	for i, environment := range deployment.Spec.Template.Spec.Containers[0].Env {
		if environment.Name == key {
			isKeyExist = true
			envNumber = i
			break
		}
	}
	if isKeyExist {
		deployment.Spec.Template.Spec.Containers[0].Env[envNumber].Value = value
	} else {
		deployment.Spec.Template.Spec.Containers[0].Env = append(deployment.Spec.Template.Spec.Containers[0].Env, corev1.EnvVar{Name: key, Value: value})
	}
	_, err := client.AppsV1().Deployments(namespace).Update(context.TODO(), deployment, metav1.UpdateOptions{})

	return err
}
