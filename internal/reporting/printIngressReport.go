package reporting

import (
	"fmt"

	"github.com/Parthiba-Hazra/kubervigil/internal/shared"
)

// PrintIngressHealthReports prints the Ingress health reports in a readable format.
func PrintIngressHealthReports(healthReports []shared.HealthCheckResult) {
	fmt.Println("Ingress Health Reports:")
	fmt.Println("=======================")

	for _, report := range healthReports {
		fmt.Printf("Name: %s\n", report.Name)
		fmt.Printf("Namespace: %s\n", report.Namespace)
		fmt.Printf("Host: %s\n", report.Host)
		fmt.Printf("Path: %s\n", report.Path)
		fmt.Printf("Status: %s\n", report.Status)
		fmt.Printf("Message: %s\n", report.Message)
		fmt.Printf("HTTPCode: %d\n", report.HTTPCode)
		fmt.Printf("LastChecked: %s\n", report.LastChecked)
		fmt.Println("=======================")
	}
}
