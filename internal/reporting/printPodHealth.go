package reporting

import (
	"fmt"

	"github.com/Parthiba-Hazra/kubervigil/internal/shared"
)

// PrintHealthReports prints the health reports in a readable format.
func PrintHealthReports(healthReports []shared.HealthReport) {
	fmt.Println("Health Reports:")
	fmt.Println("===================")

	for _, report := range healthReports {
		fmt.Printf("Pod Name: %s\n", report.PodName)
		fmt.Printf("Namespace: %s\n", report.Namespace)
		fmt.Println("Containers:")

		for _, containerHealth := range report.Containers {
			fmt.Printf("\tContainer Name: %s\n", containerHealth.ContainerName)
			fmt.Printf("\tContainer State: %s\n", containerHealth.ContainerState)
			fmt.Printf("\tMessage: %s\n", containerHealth.Message)
			fmt.Println("--------------")
		}

		fmt.Printf("Overall Status: %s\n", report.OverallStatus)
		fmt.Printf("Overall Message: %s\n", report.OverallMessage)
		fmt.Println("===================")
	}
}
