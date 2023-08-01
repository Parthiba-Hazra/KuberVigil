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

// statefulsetsCmd represents the statefulsets command
var statefulsetsCmd = &cobra.Command{
	Use:   "statefulsets",
	Short: "Analyze the health of Kubernetes StatefulSets in a given namespace",
	Long: `The 'statefulsets' command analyzes the health of Kubernetes StatefulSets in a specified namespace.
It compares the preferred API version with the current API version of StatefulSets and prints the analysis result.
The preferred API version for StatefulSets is fetched from the Kubernetes cluster and then compared with the
current API version of each StatefulSet in the specified namespace. The result of the analysis is printed to the console.`,
	Run: func(cmd *cobra.Command, args []string) {
		configPath, _ := cmd.Flags().GetString("config")
		namespacce, _ := cmd.Flags().GetString("ns")

		err := analyzer.AnalyzeResourceHealth(configPath, namespacce, "STATEFULSETS")
		if err != nil {
			log.Print(err)
		}
	},
}

func init() {
	kubercmd.CheckCmd.AddCommand(statefulsetsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statefulsetsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statefulsetsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
