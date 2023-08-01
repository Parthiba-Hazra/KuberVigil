/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "kubervigil",
	Short: "A Kubernetes Resource Health Checker",
	Long: `KuberVigil is a powerful Kubernetes resource health checker that helps developers and operators maintain the health and stability of their Kubernetes environments. It analyzes Kubernetes resources, including Deployments, StatefulSets, ConfigMaps, Pods, and Services, to identify deprecated or removed APIs, potential issues with Pods, and misconfigured Services.

	With KuberVigil, you can proactively ensure your Kubernetes resources are up-to-date with the latest APIs and best practices, reducing the risk of unexpected failures and enhancing overall cluster reliability. The tool offers user-friendly reports with actionable insights, making it easy to identify and address potential problems.

	Example:
	$ kuberVigil check apis
	$ kuberVigil check pods
	$ kuberVigil check services

	Start using KuberVigil today to keep your Kubernetes resources healthy and optimize your cluster's performance.
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().String("ns", "", "You can provide kubernets namspace (eg: --ns=default)")
	RootCmd.PersistentFlags().String("config", "", "You can provide kubernets namspace (eg: --config=./.kube/config)")
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kubervigil.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
