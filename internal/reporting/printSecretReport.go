package reporting

import (
	"fmt"

	"github.com/Parthiba-Hazra/kubervigil/internal/shared"
)

// PrintSecretHealthReports prints the Secret health reports in a readable format.
func PrintSecretHealthReports(healthReports []shared.SecretHealthReport) {
	fmt.Println("Secret Health Reports:")
	fmt.Println("==========================")

	for _, report := range healthReports {
		fmt.Printf("Secret Name: %s\n", report.SecretName)
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
