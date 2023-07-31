/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package kubercmd

import (
	"fmt"

	"github.com/Parthiba-Hazra/kubervigil/cmd"
	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Perform Kubernetes Resource Health Checks",
	Long: `The "check" command allows you to perform comprehensive health checks on your Kubernetes resources. It analyzes Kubernetes Deployments, StatefulSets, ConfigMaps, Pods, and Services to ensure they comply with the latest API versions, assess the health status of Pods, and verify the connectivity and configurations of Services.

	Running the "check" command helps you maintain the overall health and stability of your Kubernetes environment. It provides user-friendly reports with detailed insights into deprecated API usage, potential Pod issues, and misconfigured Services, enabling you to proactively address any concerns.

	Usage Examples:
	$ kuberVigil check apis
	$ kuberVigil check pods
	$ kuberVigil check services

	Start using the "check" command today to keep your Kubernetes resources in optimal condition and promote a reliable cluster performance.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("check called")
	},
}

func init() {
	cmd.RootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
