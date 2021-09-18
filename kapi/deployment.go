package kapi

import (
	"context"
	"log"
	"strings"

	"github.com/anilkusc/kinject/models"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetDeployment(client *kubernetes.Clientset, namespace string, name string) models.MyDeployment {
	gotDeployment, err := client.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		log.Println(err)
	}

	deployment := models.MyDeployment{
		Name:      gotDeployment.Name,
		Namespace: gotDeployment.Namespace,
		Replicas:  int(*gotDeployment.Spec.Replicas),
	}
	var environment models.MyEnvironment
	var environments []models.MyEnvironment
	for _, env := range gotDeployment.Spec.Template.Spec.Containers[0].Env {
		environment.Key = env.Name
		environment.Value = env.Value
		environments = append(environments, environment)
	}
	deployment.Env = environments
	return deployment
}

func ListDeployments(client *kubernetes.Clientset, namespace string) []models.MyDeployment {
	var deployments []models.MyDeployment
	var deployment models.MyDeployment
	var environment models.MyEnvironment
	var environments []models.MyEnvironment
	listedDeployments, err := client.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Println(err)
	}
	for _, listedDeployment := range listedDeployments.Items {
		deployment.Name = listedDeployment.Name
		deployment.Namespace = listedDeployment.Namespace
		deployment.Replicas = int(*listedDeployment.Spec.Replicas)
		for _, listedEnv := range listedDeployment.Spec.Template.Spec.Containers[0].Env {
			environment.Key = listedEnv.Name
			environment.Value = listedEnv.Value
			environments = append(environments, environment)
		}
		deployment.Env = environments
		deployments = append(deployments, deployment)
	}
	return deployments
}

func PatchDeploymentEnv(client *kubernetes.Clientset, namespace string, deploymentName string, env string) error {
	envs := strings.Split(env, ":")
	key := envs[0]
	value := envs[1]
	gotDeployment, err := client.AppsV1().Deployments(namespace).Get(context.TODO(), deploymentName, metav1.GetOptions{})
	if err != nil {
		log.Println(err)
	}
	isKeyExist := false
	envNumber := -1
	for i, environment := range gotDeployment.Spec.Template.Spec.Containers[0].Env {
		if environment.Name == key {
			isKeyExist = true
			envNumber = i
			break
		}
	}
	if isKeyExist {
		gotDeployment.Spec.Template.Spec.Containers[0].Env[envNumber].Value = value
	} else {
		gotDeployment.Spec.Template.Spec.Containers[0].Env = append(gotDeployment.Spec.Template.Spec.Containers[0].Env, corev1.EnvVar{Name: key, Value: value})
	}
	client.AppsV1().Deployments(namespace).Update(context.TODO(), gotDeployment, metav1.UpdateOptions{})

	return err
}
func DeleteDeploymentEnv(client *kubernetes.Clientset, namespace string, deploymentName string, env string) error {
	envs := strings.Split(env, ":")
	key := envs[0]
	//value := envs[1]
	gotDeployment, err := client.AppsV1().Deployments(namespace).Get(context.TODO(), deploymentName, metav1.GetOptions{})
	if err != nil {
		log.Println(err)
	}
	for i, environment := range gotDeployment.Spec.Template.Spec.Containers[0].Env {
		if environment.Name == key {
			//append(gotDeployment.Spec.Template.Spec.Containers[0].Env[:i], gotDeployment.Spec.Template.Spec.Containers[0].Env[i+1:]...)
			gotDeployment.Spec.Template.Spec.Containers[0].Env[i] = gotDeployment.Spec.Template.Spec.Containers[0].Env[len(gotDeployment.Spec.Template.Spec.Containers[0].Env)-1]
			gotDeployment.Spec.Template.Spec.Containers[0].Env = gotDeployment.Spec.Template.Spec.Containers[0].Env[:len(gotDeployment.Spec.Template.Spec.Containers[0].Env)-1]
			break
		}
	}
	client.AppsV1().Deployments(namespace).Update(context.TODO(), gotDeployment, metav1.UpdateOptions{})

	return err
}
