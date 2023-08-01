package reporting

import (
	"fmt"

	"github.com/Parthiba-Hazra/kubervigil/internal/shared"
)

// PrintServiceHealthReports prints the Service health reports in a readable format.
func PrintServiceHealthReports(healthReports []shared.ServiceHealthReport) {
	fmt.Println("Service Health Reports:")
	fmt.Println("=======================")

	for _, report := range healthReports {
		fmt.Printf("ServiceName: %s\n", report.ServiceName)
		fmt.Printf("Namespace: %s\n", report.Namespace)
		fmt.Printf("IP: %s\n", report.IP)
		fmt.Printf("Port: %d\n", report.Port)
		fmt.Printf("TargetPort: %d\n", report.TargetPort)
		fmt.Printf("Selector: %v\n", report.Selector)
		fmt.Printf("EndpointCount: %d\n", report.EndpointCount)

		fmt.Println("Conditions:")
		for _, condition := range report.Conditions {
			fmt.Printf("\tType: %s\n", condition.Type)
			fmt.Printf("\tStatus: %s\n", condition.Status)
			fmt.Printf("\tMessage: %s\n", condition.Message)
			fmt.Printf("\tLastUpdateTime: %s\n", condition.LastUpdateTime)
			fmt.Println("--------------")
		}

		fmt.Printf("Overall Status: %s\n", report.OverallStatus)
		fmt.Printf("Overall Message: %s\n", report.OverallMessage)
		fmt.Println("=======================")
	}
}
