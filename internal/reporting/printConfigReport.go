package reporting

import (
	"fmt"

	"github.com/Parthiba-Hazra/kubervigil/internal/shared"
)

// PrintConfigMapHealthReports prints the ConfigMap health reports in a readable format.
func PrintConfigMapHealthReports(healthReports []shared.ConfigMapHealthReport) {
	fmt.Println("ConfigMap Health Reports:")
	fmt.Println("==========================")

	for _, report := range healthReports {
		fmt.Printf("ConfigMap Name: %s\n", report.ConfigMapName)
		fmt.Printf("Namespace: %s\n", report.Namespace)
		fmt.Println("Data Keys:")
		for _, key := range report.DataKeys {
			fmt.Printf("  %s\n", key)
		}
		fmt.Printf("Overall Status: %s\n", report.OverallStatus)
		fmt.Printf("Overall Message: %s\n", report.OverallMessage)
		fmt.Println("==========================")
	}
}
