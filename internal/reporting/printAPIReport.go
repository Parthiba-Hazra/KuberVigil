package reporting

import (
	"fmt"

	"github.com/Parthiba-Hazra/kubervigil/internal/shared"
)

// PrintAPIServerHealthReport prints the API Server health report in a readable format.
func PrintAPIServerHealthReport(apiServerHealth *shared.APIServerHealthReport) {
	fmt.Println("API Server Health Report:")
	fmt.Println("==========================")

	fmt.Println("API Endpoints:")
	for _, endpoint := range apiServerHealth.APIEndpoints {
		fmt.Printf("Endpoint: %s\n", endpoint.Endpoint)
		fmt.Printf("Status: %s\n", endpoint.Status)
		fmt.Printf("Message: %s\n", endpoint.Message)
		fmt.Println("==========================")
	}

	fmt.Println("Cluster Conditions:")
	for _, condition := range apiServerHealth.ClusterConditions {
		fmt.Printf("Condition Type: %s\n", condition.Type)
		fmt.Printf("Status: %s\n", condition.Status)
		fmt.Printf("Message: %s\n", condition.Message)
		fmt.Println("==========================")
	}

	fmt.Printf("Overall Status: %s\n", apiServerHealth.OverallStatus)
	fmt.Printf("Overall Message: %s\n", apiServerHealth.OverallMessage)
}
