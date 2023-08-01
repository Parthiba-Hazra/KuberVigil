/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package apiversion

import (
	"github.com/Parthiba-Hazra/kubervigil/cmd/kuberVigil/kubercmd"
	"github.com/Parthiba-Hazra/kubervigil/internal/analyzer"
	"github.com/spf13/cobra"
)

// apiversionCmd represents the apiversion command
var apiversionCmd = &cobra.Command{
	Use:   "apiversion",
	Short: "Check prefered Kubernetes API versions",
	Long: `This command analyzes the health of Kubernetes API versions for resources in a given namespace.
It checks the preferred API version of resources and prints the result.`,
	Run: func(cmd *cobra.Command, args []string) {
		configPath, _ := cmd.Flags().GetString("config")

		analyzer.GetPreferedAPIversion(configPath)
	},
}

func init() {
	kubercmd.CheckCmd.AddCommand(apiversionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apiversionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apiversionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
