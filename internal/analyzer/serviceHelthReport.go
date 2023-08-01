package analyzer

import (
	"github.com/Parthiba-Hazra/kubervigil/internal/shared"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreateServiceHealthReport creates a detailed health report for the given list of services.
func CreateServiceHealthReport(services *v1.ServiceList) ([]shared.ServiceHealthReport, error) {

	var healthReports []shared.ServiceHealthReport

	for _, service := range services.Items {
		healthReport := shared.ServiceHealthReport{
			ServiceName:   service.Name,
			Namespace:     service.Namespace,
			IP:            service.Spec.ClusterIP,
			Port:          service.Spec.Ports[0].Port,
			TargetPort:    service.Spec.Ports[0].TargetPort.IntVal,
			Selector:      service.Spec.Selector,
			EndpointCount: len(service.Status.LoadBalancer.Ingress),
		}

		healthReport.Conditions = getServiceConditions(service.Status.Conditions)

		healthReport.OverallStatus, healthReport.OverallMessage = getOverallServiceHealthStatus(healthReport.Conditions)

		healthReports = append(healthReports, healthReport)
	}

	return healthReports, nil
}

func getServiceConditions(conditions []metav1.Condition) []shared.ServiceCondition {
	var serviceConditions []shared.ServiceCondition

	for _, condition := range conditions {
		serviceConditions = append(serviceConditions, shared.ServiceCondition{
			Type:           string(condition.Type),
			Status:         string(condition.Status),
			Message:        condition.Message,
			LastUpdateTime: condition.LastTransitionTime,
		})
	}

	return serviceConditions
}

func getOverallServiceHealthStatus(conditions []shared.ServiceCondition) (string, string) {
	for _, condition := range conditions {
		if condition.Type == "LoadBalancerReady" && condition.Status != "True" {
			return "Error", "Load balancer is not ready"
		} else if condition.Type == "ServiceReady" && condition.Status != "True" {
			return "Error", "Service is not ready"
		}
	}

	return "Healthy", "All checks passed"
}
