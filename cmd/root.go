/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/anilkusc/kinject/kapi"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kinject",
	Short: "A brief description of your application",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		switch cmd.Flag("kind").Value.String() {
		case "deployment", "deploy", "sts", "statefulset", "pod":
			return nil
		default:
			return errors.New("requires a valid kind parameter")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		client := kapi.CreateClient(cmd.Flag("kubeconfig").Value.String())
		deployment := kapi.GetDeployment(client, cmd.Flag("namespace").Value.String(), cmd.Flag("name").Value.String())
		fmt.Printf("%v", deployment)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().StringP("kind", "k", "", "Kubernetes Kind Object like Deployment,Pod etc.")
	rootCmd.MarkFlagRequired("kind")
	rootCmd.PersistentFlags().StringP("kubeconfig", "c", "~/.kube/config", "Kubeconfig path")
	rootCmd.Flags().StringP("namespace", "n", "all", "Kubernetes Namespace That Will Be Affect")
	rootCmd.Flags().StringP("name", "N", "", "Kubernetes Workload Name")

}
