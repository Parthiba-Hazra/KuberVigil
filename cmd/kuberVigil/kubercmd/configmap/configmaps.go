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

// configmapsCmd represents the configmaps command
var configmapsCmd = &cobra.Command{
	Use:   "configmaps",
	Short: "Analyze the health of Kubernetes ConfigMaps in a given namespace",
	Long: `The 'configmaps' command analyzes the health of Kubernetes ConfigMaps in a specified namespace.
It checks if there is any data defined in the ConfigMaps and prints the analysis result.
The result of the analysis is printed to the console, indicating the health status of each ConfigMap.`,
	Run: func(cmd *cobra.Command, args []string) {
		configPath, _ := cmd.Flags().GetString("config")
		namespacce, _ := cmd.Flags().GetString("ns")

		err := analyzer.AnalyzeResourceHealth(configPath, namespacce, "CONFIGMAPS")
		if err != nil {
			log.Print(err)
		}
	},
}

func init() {
	kubercmd.CheckCmd.AddCommand(configmapsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configmapsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configmapsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
