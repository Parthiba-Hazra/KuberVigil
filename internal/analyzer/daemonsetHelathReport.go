package analyzer

import (
	"github.com/Parthiba-Hazra/kubervigil/internal/shared"
	appsv1 "k8s.io/api/apps/v1"
)

// CreateDaemonSetHealthReport creates a detailed health report for the given list of DaemonSets.
func CreateDaemonSetHealthReport(daemonSets *appsv1.DaemonSetList) ([]shared.DaemonSetHealthReport, error) {

	var healthReports []shared.DaemonSetHealthReport

	for _, daemonSet := range daemonSets.Items {
		healthReport := shared.DaemonSetHealthReport{
			DaemonSetName:   daemonSet.Name,
			Namespace:       daemonSet.Namespace,
			DesiredNumber:   daemonSet.Status.DesiredNumberScheduled,
			CurrentNumber:   daemonSet.Status.CurrentNumberScheduled,
			ReadyNumber:     daemonSet.Status.NumberReady,
			AvailableNumber: daemonSet.Status.NumberAvailable,
			UpdatedNumber:   daemonSet.Status.UpdatedNumberScheduled,
		}

		healthReport.Conditions = getDaemonSetConditions(daemonSet.Status.Conditions)

		healthReport.OverallStatus, healthReport.OverallMessage = getOverallDaemonSetHealthStatus(healthReport.Conditions)

		healthReports = append(healthReports, healthReport)
	}

	return healthReports, nil
}

func getDaemonSetConditions(conditions []appsv1.DaemonSetCondition) []shared.DaemonSetCondition {
	var daemonSetConditions []shared.DaemonSetCondition

	for _, condition := range conditions {
		daemonSetConditions = append(daemonSetConditions, shared.DaemonSetCondition{
			Type:           string(condition.Type),
			Status:         string(condition.Status),
			Message:        condition.Message,
			LastUpdateTime: condition.LastTransitionTime,
		})
	}

	return daemonSetConditions
}

func getOverallDaemonSetHealthStatus(conditions []shared.DaemonSetCondition) (string, string) {
	for _, condition := range conditions {
		if condition.Type == "Progressing" && condition.Status != "True" {
			return "Error", "DaemonSet is not progressing"
		} else if condition.Type == "Available" && condition.Status != "True" {
			return "Error", "DaemonSet is not available"
		}
	}

	return "Healthy", "All checks passed"
}
