package analyzer

import (
	"fmt"
	"strings"

	"github.com/Parthiba-Hazra/kubervigil/internal/shared"
	corev1 "k8s.io/api/core/v1"
)

// CreatePodHealthReport creates a detailed health report for the given list of pods.
func CreatePodHealthReport(pods *corev1.PodList) ([]shared.HealthReport, error) {

	var healthReports []shared.HealthReport

	for _, pod := range pods.Items {
		healthReport := shared.HealthReport{
			PodName:   pod.Name,
			Namespace: pod.Namespace,
		}

		var containerHealths []shared.ContainerHealth

		for _, containerStatus := range pod.Status.ContainerStatuses {
			containerHealth := shared.ContainerHealth{
				ContainerName:  containerStatus.Name,
				ContainerState: getContainerState(containerStatus.State),
			}

			if containerStatus.State.Waiting != nil {
				containerHealth.Message = containerStatus.State.Waiting.Message
			} else if containerStatus.State.Terminated != nil {
				containerHealth.Message = containerStatus.State.Terminated.Message
			}

			containerHealths = append(containerHealths, containerHealth)
		}

		healthReport.Containers = containerHealths

		// Additional health checks
		healthReport.CheckDNSResolution = checkDNSResolution(pod.Spec)
		healthReport.CheckResourceLimits = checkResourceLimits(pod.Spec)

		healthReport.OverallStatus, healthReport.OverallMessage = getOverallHealthStatus(containerHealths, healthReport.CheckDNSResolution, healthReport.CheckResourceLimits)

		healthReports = append(healthReports, healthReport)
	}

	return healthReports, nil
}

func getContainerState(state corev1.ContainerState) string {
	if state.Waiting != nil {
		return fmt.Sprintf("Waiting (%s)", state.Waiting.Reason)
	} else if state.Running != nil {
		return "Running"
	} else if state.Terminated != nil {
		return fmt.Sprintf("Terminated (%s)", state.Terminated.Reason)
	}

	return "Unknown"
}

func getOverallHealthStatus(containerHealths []shared.ContainerHealth, dnsResolutionOK bool, resourceLimitsOK bool) (string, string) {
	for _, containerHealth := range containerHealths {
		if strings.Contains(containerHealth.ContainerState, "Waiting") {
			return "Degraded", "Some containers are in a Waiting state"
		} else if strings.Contains(containerHealth.ContainerState, "Terminated") {
			return "Error", "Some containers have terminated unexpectedly"
		}
	}

	if !dnsResolutionOK {
		return "Degraded", "DNS resolution issue detected"
	}

	if !resourceLimitsOK {
		return "Degraded", "Resource limits exceeded for some containers"
	}

	return "Healthy", "All checks passed"
}

func checkDNSResolution(podSpec corev1.PodSpec) bool {
	for _, container := range podSpec.Containers {
		for _, envVar := range container.Env {
			if envVar.Name == "HTTP_PROXY" || envVar.Name == "HTTPS_PROXY" {
				// Check if DNS resolution is blocked by a proxy
				return false
			}
		}
	}
	return true
}

func checkResourceLimits(podSpec corev1.PodSpec) bool {
	for _, container := range podSpec.Containers {
		if container.Resources.Limits != nil {
			if container.Resources.Limits.Cpu().IsZero() || container.Resources.Limits.Memory().IsZero() {
				// Check if resource limits are not set
				return false
			}
		}
	}
	return true
}
