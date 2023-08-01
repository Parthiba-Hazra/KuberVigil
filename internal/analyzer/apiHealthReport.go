package analyzer

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Parthiba-Hazra/kubervigil/internal/shared"
)

// CreateAPIServerHealthReport creates a detailed health report for the Kubernetes API Server.
func CreateAPIServerHealthReport(apiServerEndpoints []string, clusterConditions []shared.ClusterConditionStatus) (*shared.APIServerHealthReport, error) {
	apiServerHealth := &shared.APIServerHealthReport{
		APIEndpoints:      make([]shared.EndpointHealthStatus, 0),
		ClusterConditions: make([]shared.ClusterConditionStatus, 0),
	}

	for _, endpoint := range apiServerEndpoints {
		status, message := checkEndpointHealth(endpoint)
		apiServerHealth.APIEndpoints = append(apiServerHealth.APIEndpoints, shared.EndpointHealthStatus{
			Endpoint: endpoint,
			Status:   status,
			Message:  message,
		})
	}

	for _, condition := range clusterConditions {
		apiServerHealth.ClusterConditions = append(apiServerHealth.ClusterConditions, shared.ClusterConditionStatus{
			Type:    condition.Type,
			Status:  condition.Status,
			Message: condition.Message,
		})
	}

	// Determine overall API Server health
	apiServerHealth.OverallStatus, apiServerHealth.OverallMessage = getOverallAPIServerHealth(apiServerHealth.APIEndpoints, apiServerHealth.ClusterConditions)

	return apiServerHealth, nil
}

func checkEndpointHealth(endpoint string) (string, string) {
	// Create a HTTP client with a timeout
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	// Perform an HTTP GET request to the API server endpoint
	resp, err := client.Get(fmt.Sprintf("http://%s/healthz", endpoint))
	if err != nil {
		return "Error", fmt.Sprintf("Failed to reach API Server endpoint: %s", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode == http.StatusOK {
		return "Healthy", "API Server endpoint is reachable"
	}

	return "Error", fmt.Sprintf("API Server endpoint returned status code: %d", resp.StatusCode)
}

func getOverallAPIServerHealth(endpoints []shared.EndpointHealthStatus, conditions []shared.ClusterConditionStatus) (string, string) {
	// Check the health status of all endpoints
	for _, endpoint := range endpoints {
		if endpoint.Status != "Healthy" {
			return "Error", "API Server endpoint is not healthy"
		}
	}

	// Check the health status of all cluster conditions
	for _, condition := range conditions {
		if condition.Status != "Healthy" {
			return "Error", "Cluster condition not met"
		}
	}

	// If all endpoints and cluster conditions are healthy, return "Healthy"
	return "Healthy", "API Server is healthy"
}
