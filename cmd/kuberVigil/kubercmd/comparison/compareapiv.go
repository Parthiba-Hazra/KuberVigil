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

// compareapivCmd represents the compareapiv command
var compareapivCmd = &cobra.Command{
	Use:   "compareapiv",
	Short: "Compare preferred API versions with current API versions for Kubernetes resources",
	Long: `The 'compare' command compares the preferred API versions obtained from the Kubernetes
	cluster's discovery information with the current API versions of the resources present in the
	specified namespace (or all namespaces if no namespace is specified). It checks if the current
	API versions match the preferred ones or if there are any discrepancies, which might indicate
	that the resources are deprecated or removed.`,
	Run: func(cmd *cobra.Command, args []string) {
		configPath, _ := cmd.Flags().GetString("config")
		namespacce, _ := cmd.Flags().GetString("ns")

		_, err := analyzer.GetComparisonAPIversion(configPath, namespacce)
		if err != nil {
			log.Print(err)
		}
	},
}

func init() {
	kubercmd.CheckCmd.AddCommand(compareapivCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// compareapivCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// compareapivCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
