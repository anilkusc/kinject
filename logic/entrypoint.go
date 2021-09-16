package logic

import (
	"github.com/anilkusc/kinject/kapi"
	"github.com/spf13/cobra"
)

func Entrypoint(cmd *cobra.Command) {

	switch cmd.Flag("kind").Value.String() {
	case "deployment", "deploy", "Deploy", "Deployment":
		DeploymentEntrypoint(cmd)
		//	case "Pod", "pod":
		//
		//	default:

	}

}

func DeploymentEntrypoint(cmd *cobra.Command) {
	switch cmd.Flag("type").Value.String() {
	case "env", "environment", "Env", "Environment":
		DeploymentEnvironmentSetter(cmd)

	}
}

func DeploymentEnvironmentSetter(cmd *cobra.Command) {
	client := kapi.CreateClient(cmd.Flag("kubeconfig").Value.String())
	if cmd.Flag("namespace").Value.String() == "all" {
		namespaces := kapi.ListNamespaces(client)
		for _, namespace := range namespaces {
			deployments := kapi.ListDeployments(client, namespace.Name)
			for _, deployment := range deployments {
				kapi.PatchDeploymentEnv(client, namespace.Name, deployment.Name, cmd.Flag("environment").Value.String())
			}
		}
	} else {
		deployments := kapi.ListDeployments(client, cmd.Flag("namespace").Value.String())
		for _, deployment := range deployments {
			kapi.PatchDeploymentEnv(client, deployment.Namespace, deployment.Name, cmd.Flag("environment").Value.String())
		}
	}

}
