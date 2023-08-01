/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Parthiba-Hazra/kubervigil/cmd/kuberVigil/kubercmd"
	"github.com/Parthiba-Hazra/kubervigil/internal/analyzer"
	"github.com/spf13/cobra"
)

// podsCmd represents the pods command
var podsCmd = &cobra.Command{
	Use:   "pods",
	Short: "Analyze the health of Kubernetes Pods in a given namespace",
	Long: `The 'pods' command analyzes the health of Kubernetes Pods in a specified namespace.
It compares the preferred API version with the current API version of Pods and prints the analysis result.
The preferred API version for Pods is fetched from the Kubernetes cluster and then compared with the
current API version of each Pod in the specified namespace. The result of the analysis is printed to the console.`,
	Run: func(cmd *cobra.Command, args []string) {
		configPath, _ := cmd.Flags().GetString("config")
		namespacce, _ := cmd.Flags().GetString("ns")

		analyzer.AnalyzeResourceHealth(configPath, namespacce, "PODS")
	},
}

func init() {
	kubercmd.CheckCmd.AddCommand(podsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// podsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// podsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
