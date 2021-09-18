package logic

import (
	"github.com/anilkusc/kinject/kapi"
	"k8s.io/client-go/kubernetes"
)

func DeploymentEnvironmentSetter(client *kubernetes.Clientset, namespace string, env string) {
	if namespace == "all" {
		namespaces := kapi.ListNamespaces(client)
		for _, namespace := range namespaces {
			deployments := kapi.ListDeployments(client, namespace.Name)
			for _, deployment := range deployments {
				kapi.PatchDeploymentEnv(client, namespace.Name, deployment.Name, env)
			}
		}
	} else {
		deployments := kapi.ListDeployments(client, namespace)
		for _, deployment := range deployments {
			kapi.PatchDeploymentEnv(client, deployment.Namespace, deployment.Name, env)
		}
	}
}
func DeploymentEnvironmentRemover(client *kubernetes.Clientset, namespace string, env string) {
	if namespace == "all" {
		namespaces := kapi.ListNamespaces(client)
		for _, namespace := range namespaces {
			deployments := kapi.ListDeployments(client, namespace.Name)
			for _, deployment := range deployments {
				kapi.DeleteDeploymentEnv(client, namespace.Name, deployment.Name, env)
			}
		}
	} else {
		deployments := kapi.ListDeployments(client, namespace)
		for _, deployment := range deployments {
			kapi.DeleteDeploymentEnv(client, deployment.Namespace, deployment.Name, env)
		}
	}
}
