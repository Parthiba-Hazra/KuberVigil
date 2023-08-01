package analyzer

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Parthiba-Hazra/kubervigil/internal/shared"
	v1 "k8s.io/api/networking/v1"
)

// IngressHealthReport generates a detailed health report for Ingress resources in Kubernetes.
func CreateIngressHealthReport(ingresses *v1.IngressList, namespace string) ([]shared.HealthCheckResult, error) {
	var healthReports []shared.HealthCheckResult

	// Loop through each Ingress and perform health checks
	for _, ingress := range ingresses.Items {
		healthReport := shared.HealthCheckResult{
			Name:        ingress.Name,
			Namespace:   ingress.Namespace,
			Host:        ingress.Spec.Rules[0].Host,
			Path:        "",
			Status:      "Healthy",
			Message:     "",
			HTTPCode:    0,
			LastChecked: time.Now(),
		}

		// Check if the Ingress has any rules defined
		if len(ingress.Spec.Rules) == 0 {
			healthReport.Status = "Error"
			healthReport.Message = "No rules defined in the Ingress"
			healthReports = append(healthReports, healthReport)
			continue
		}

		// Check each rule defined in the Ingress
		for _, rule := range ingress.Spec.Rules {
			// Check if the rule has any HTTP paths defined
			if len(rule.HTTP.Paths) == 0 {
				healthReport.Status = "Error"
				healthReport.Message = "No HTTP paths defined in the rule"
				healthReports = append(healthReports, healthReport)
				continue
			}

			// Loop through each HTTP path and perform an HTTP health check
			for _, path := range rule.HTTP.Paths {
				healthReport.Path = path.Path
				url := fmt.Sprintf("http://%s%s", rule.Host, path.Path)

				// Perform an HTTP GET request to the Ingress endpoint
				resp, err := http.Get(url)
				if err != nil {
					healthReport.Status = "Error"
					healthReport.Message = fmt.Sprintf("Failed to reach the Ingress endpoint: %v", err)
					healthReports = append(healthReports, healthReport)
					continue
				}
				defer resp.Body.Close()

				// Check the HTTP status code
				healthReport.HTTPCode = resp.StatusCode
				if resp.StatusCode < 200 || resp.StatusCode >= 300 {
					healthReport.Status = "Error"
					healthReport.Message = fmt.Sprintf("Received non-successful HTTP status code: %d", resp.StatusCode)
					healthReports = append(healthReports, healthReport)
					continue
				}
			}
		}

		// If no errors were encountered, add the Ingress to the health reports with a "Healthy" status
		if healthReport.Status == "Healthy" {
			healthReport.Message = "Ingress is healthy"
			healthReports = append(healthReports, healthReport)
		}
	}

	return healthReports, nil
}
