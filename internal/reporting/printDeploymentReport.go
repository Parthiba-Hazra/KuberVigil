package reporting

import (
	"fmt"

	"github.com/Parthiba-Hazra/kubervigil/internal/shared"
)

// PrintDeploymentHealthReports prints the Deployment health reports in a readable format.
func PrintDeploymentHealthReports(healthReports []shared.DeploymentHealthReport) {
	fmt.Println("Deployment Health Reports:")
	fmt.Println("==========================")

	for _, report := range healthReports {
		fmt.Printf("Deployment Name: %s\n", report.DeploymentName)
		fmt.Printf("Namespace: %s\n", report.Namespace)
		fmt.Printf("Replicas: %d\n", report.Replicas)
		fmt.Printf("Ready Replicas: %d\n", report.ReadyReplicas)
		fmt.Printf("Updated Replicas: %d\n", report.UpdatedReplicas)
		fmt.Printf("Available Replicas: %d\n", report.AvailableReplicas)
		fmt.Printf("Overall Status: %s\n", report.OverallStatus)
		fmt.Printf("Overall Message: %s\n", report.OverallMessage)

		fmt.Println("Conditions:")
		for _, condition := range report.Conditions {
			fmt.Printf("  Type: %s\n", condition.Type)
			fmt.Printf("  Status: %s\n", condition.Status)
			fmt.Printf("  Message: %s\n", condition.Message)
			fmt.Printf("  Last Update Time: %s\n", condition.LastUpdateTime)
			fmt.Println("--------------")
		}

		fmt.Println("==========================")
	}
}
