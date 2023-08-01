/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/Parthiba-Hazra/kubervigil/cmd/kuberVigil/kubercmd"
	"github.com/Parthiba-Hazra/kubervigil/internal/analyzer"
	"github.com/spf13/cobra"
)

// ingressCmd represents the ingress command
var ingressCmd = &cobra.Command{
	Use:   "ingress",
	Short: "Analyze the health of Kubernetes Ingress resources in a given namespace",
	Long: `The 'ingress' command analyzes the health of Kubernetes Ingress resources in a specified namespace.
It performs HTTP health checks for each Ingress rule defined in the namespace and prints the analysis result.
The HTTP health checks are performed by sending GET requests to the endpoints defined in the Ingress rules.
The result of the analysis is printed to the console, indicating the health status of each Ingress rule.`,
	Run: func(cmd *cobra.Command, args []string) {
		configPath, _ := cmd.Flags().GetString("config")
		namespacce, _ := cmd.Flags().GetString("ns")

		err := analyzer.AnalyzeResourceHealth(configPath, namespacce, "INGRESS")
		if err != nil {
			log.Print(err)
		}
	},
}

func init() {
	kubercmd.CheckCmd.AddCommand(ingressCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ingressCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ingressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
