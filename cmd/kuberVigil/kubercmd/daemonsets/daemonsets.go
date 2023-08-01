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

// daemonsetsCmd represents the daemonsets command
var daemonsetsCmd = &cobra.Command{
	Use:   "daemonsets",
	Short: "Analyze the health of Kubernetes DaemonSets in a given namespace",
	Long: `The 'daemonsets' command analyzes the health of Kubernetes DaemonSets in a specified namespace.
It compares the desired number of pods with the current number, ready number, and updated number of pods for each DaemonSet and prints the analysis result.
The result of the analysis is printed to the console, indicating the health status of each DaemonSet.`,
	Run: func(cmd *cobra.Command, args []string) {
		configPath, _ := cmd.Flags().GetString("config")
		namespacce, _ := cmd.Flags().GetString("ns")

		err := analyzer.AnalyzeResourceHealth(configPath, namespacce, "DAEMONSETS")
		if err != nil {
			log.Print(err)
		}
	},
}

func init() {
	kubercmd.CheckCmd.AddCommand(daemonsetsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// daemonsetsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// daemonsetsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
