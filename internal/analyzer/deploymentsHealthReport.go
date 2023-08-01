package analyzer

import (
	"github.com/Parthiba-Hazra/kubervigil/internal/shared"
	appsv1 "k8s.io/api/apps/v1"
)

// CreateDeploymentHealthReport creates a detailed health report for the given list of deployments.
func CreateDeploymentHealthReport(deployments *appsv1.DeploymentList) ([]shared.DeploymentHealthReport, error) {

	var healthReports []shared.DeploymentHealthReport

	for _, deployment := range deployments.Items {
		healthReport := shared.DeploymentHealthReport{
			DeploymentName:    deployment.Name,
			Namespace:         deployment.Namespace,
			Replicas:          *deployment.Spec.Replicas,
			ReadyReplicas:     deployment.Status.ReadyReplicas,
			UpdatedReplicas:   deployment.Status.UpdatedReplicas,
			AvailableReplicas: deployment.Status.AvailableReplicas,
		}

		healthReport.Conditions = getDeploymentConditions(deployment.Status.Conditions)

		healthReport.OverallStatus, healthReport.OverallMessage = getOverallDeploymentHealthStatus(healthReport.Conditions, deployment.Status.ReadyReplicas, *deployment.Spec.Replicas)

		healthReports = append(healthReports, healthReport)
	}

	return healthReports, nil
}

func getDeploymentConditions(conditions []appsv1.DeploymentCondition) []shared.DeploymentCondition {
	var deploymentConditions []shared.DeploymentCondition

	for _, condition := range conditions {
		deploymentConditions = append(deploymentConditions, shared.DeploymentCondition{
			Type:           string(condition.Type),
			Status:         string(condition.Status),
			Message:        condition.Message,
			LastUpdateTime: condition.LastUpdateTime,
		})
	}

	return deploymentConditions
}

func getOverallDeploymentHealthStatus(conditions []shared.DeploymentCondition, readyReplicas int32, replicas int32) (string, string) {
	for _, condition := range conditions {
		if condition.Type == "Progressing" && condition.Status != "True" {
			return "Degraded", "Deployment is not progressing as expected"
		} else if condition.Type == "Available" && condition.Status != "True" {
			return "Degraded", "Deployment is not available"
		}
	}

	if readyReplicas == 0 {
		return "Error", "No replicas are ready"
	}

	if readyReplicas < replicas {
		return "Degraded", "Not all replicas are ready"
	}

	return "Healthy", "All checks passed"
}
