package analyzer

import (
	"github.com/Parthiba-Hazra/kubervigil/internal/shared"
	appsv1 "k8s.io/api/apps/v1"
)

// CreateStatefulSetHealthReport creates a detailed health report for the given list of StatefulSets.
func CreateStatefulSetHealthReport(statefulSets *appsv1.StatefulSetList) ([]shared.StatefulSetHealthReport, error) {

	var healthReports []shared.StatefulSetHealthReport

	for _, statefulSet := range statefulSets.Items {
		healthReport := shared.StatefulSetHealthReport{
			StatefulSetName: statefulSet.Name,
			Namespace:       statefulSet.Namespace,
			Replicas:        *statefulSet.Spec.Replicas,
			ReadyReplicas:   statefulSet.Status.ReadyReplicas,
			CurrentReplicas: statefulSet.Status.CurrentReplicas,
			UpdatedReplicas: statefulSet.Status.UpdatedReplicas,
		}

		healthReport.Conditions = getStatefulSetConditions(statefulSet.Status.Conditions)

		healthReport.OverallStatus, healthReport.OverallMessage = getOverallStatefulSetHealthStatus(healthReport.Conditions)

		healthReports = append(healthReports, healthReport)
	}

	return healthReports, nil
}

func getStatefulSetConditions(conditions []appsv1.StatefulSetCondition) []shared.StatefulSetCondition {
	var statefulSetConditions []shared.StatefulSetCondition

	for _, condition := range conditions {
		statefulSetConditions = append(statefulSetConditions, shared.StatefulSetCondition{
			Type:           string(condition.Type),
			Status:         string(condition.Status),
			Message:        condition.Message,
			LastUpdateTime: condition.LastTransitionTime,
		})
	}

	return statefulSetConditions
}

func getOverallStatefulSetHealthStatus(conditions []shared.StatefulSetCondition) (string, string) {
	for _, condition := range conditions {
		if condition.Type == "Ready" && condition.Status != "True" {
			return "Error", "StatefulSet is not ready"
		}
	}

	return "Healthy", "All checks passed"
}
