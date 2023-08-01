/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Parthiba-Hazra/kubervigil/cmd/kuberVigil/kubercmd"
	"github.com/Parthiba-Hazra/kubervigil/internal/analyzer"
	"github.com/spf13/cobra"
)

// deploymentsCmd represents the deployments command
var deploymentsCmd = &cobra.Command{
	Use:   "deployments",
	Short: "Analyze the health of Kubernetes Deployments in a given namespace",
	Long: `The 'deployments' command analyzes the health of Kubernetes Deployments in a specified namespace.
It compares the desired replicas with the available, updated, and ready replicas for each Deployment and prints the analysis result.
The result of the analysis is printed to the console, indicating the health status of each Deployment.`,
	Run: func(cmd *cobra.Command, args []string) {
		configPath, _ := cmd.Flags().GetString("config")
		namespacce, _ := cmd.Flags().GetString("ns")

		analyzer.AnalyzeResourceHealth(configPath, namespacce, "DEPLOYMENTS")
	},
}

func init() {
	kubercmd.CheckCmd.AddCommand(deploymentsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deploymentsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deploymentsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
