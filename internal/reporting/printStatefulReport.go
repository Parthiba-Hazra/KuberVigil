package reporting

import (
	"fmt"

	"github.com/Parthiba-Hazra/kubervigil/internal/shared"
)

// PrintStatefulSetHealthReports prints the StatefulSet health reports in a readable format.
func PrintStatefulSetHealthReports(healthReports []shared.StatefulSetHealthReport) {
	fmt.Println("StatefulSet Health Reports:")
	fmt.Println("==========================")

	for _, report := range healthReports {
		fmt.Printf("StatefulSetName: %s\n", report.StatefulSetName)
		fmt.Printf("Namespace: %s\n", report.Namespace)
		fmt.Printf("Replicas: %d\n", report.Replicas)
		fmt.Printf("ReadyReplicas: %d\n", report.ReadyReplicas)
		fmt.Printf("CurrentReplicas: %d\n", report.CurrentReplicas)
		fmt.Printf("UpdatedReplicas: %d\n", report.UpdatedReplicas)

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
		fmt.Println("==========================")
	}
}
