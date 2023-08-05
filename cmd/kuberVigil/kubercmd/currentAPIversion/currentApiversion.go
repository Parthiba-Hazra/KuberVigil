/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Parthiba-Hazra/kubervigil/cmd/kuberVigil/kubercmd"
	"github.com/Parthiba-Hazra/kubervigil/internal/kubeclient"
	"github.com/spf13/cobra"
)

// currentApiversionCmd represents the currentApiversion command
var currentApiversionCmd = &cobra.Command{
	Use:   "currentApiversion",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		configPath, _ := cmd.Flags().GetString("config")
		namespacce, _ := cmd.Flags().GetString("ns")

		kubeclient.PrintAPIinfo(namespacce, configPath)
	},
}

func init() {
	kubercmd.CheckCmd.AddCommand(currentApiversionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// currentApiversionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// currentApiversionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
