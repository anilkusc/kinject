package cmd

import (
	"github.com/anilkusc/kinject/kapi"
	"github.com/anilkusc/kinject/logic"
	"github.com/spf13/cobra"
)

func initDeploymentCommand() {
	rootCmd.AddCommand(deploymentCmd)
}

var deploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "Print the version number of Hugo",
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
