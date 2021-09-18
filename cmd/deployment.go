package cmd

import (
	"github.com/anilkusc/kinject/kapi"
	"github.com/anilkusc/kinject/logic"
	"github.com/spf13/cobra"
)

func initDeploymentCommand() {
	deploymentCmd.Flags().StringP("environment", "e", "", "Environment Key Value(Key:Value)")
	deploymentCmd.Flags().StringP("mode", "m", "add", "Mode of tool(add,delete).Add will be edit if exist.")
	rootCmd.AddCommand(deploymentCmd)
	var deployCmd = deploymentCmd
	deployCmd.Use = "deploy"
	rootCmd.AddCommand(deployCmd)
}

var deploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "inject smt to deployment",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		client := kapi.CreateClient(cmd.Flag("kubeconfig").Value.String())
		switch cmd.Flag("mode").Value.String() {
		case "add":
			logic.DeploymentEnvironmentSetter(client, cmd.Flag("namespace").Value.String(), cmd.Flag("environment").Value.String())
		case "delete":
			logic.DeploymentEnvironmentRemover(client, cmd.Flag("namespace").Value.String(), cmd.Flag("environment").Value.String())
		}

	},
}
