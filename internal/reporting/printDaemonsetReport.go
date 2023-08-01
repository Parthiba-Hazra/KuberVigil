package reporting

import (
	"fmt"

	"github.com/Parthiba-Hazra/kubervigil/internal/shared"
)

// PrintDaemonSetHealthReports prints the DaemonSet health reports in a readable format.
func PrintDaemonSetHealthReports(healthReports []shared.DaemonSetHealthReport) {
	fmt.Println("DaemonSet Health Reports:")
	fmt.Println("==========================")

	for _, report := range healthReports {
		fmt.Printf("DaemonSet Name: %s\n", report.DaemonSetName)
		fmt.Printf("Namespace: %s\n", report.Namespace)
		fmt.Printf("Desired Number: %d\n", report.DesiredNumber)
		fmt.Printf("Current Number: %d\n", report.CurrentNumber)
		fmt.Printf("Ready Number: %d\n", report.ReadyNumber)
		fmt.Printf("Available Number: %d\n", report.AvailableNumber)
		fmt.Printf("Updated Number: %d\n", report.UpdatedNumber)
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
