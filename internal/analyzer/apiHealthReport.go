package analyzer

import (
	"context"
	"fmt"
	"time"

	"github.com/Parthiba-Hazra/kubervigil/internal/kubeclient"
	"github.com/Parthiba-Hazra/kubervigil/internal/shared"
)

// CreateAPIServerHealthReport creates a detailed health report for the Kubernetes API Server.
func CreateAPIServerHealthReport(apiServerEndpoints []string, clusterConditions []shared.ClusterConditionStatus, client *kubeclient.KubeClient) (*shared.APIServerHealthReport, error) {
	apiServerHealth := &shared.APIServerHealthReport{
		APIEndpoints:      make([]shared.EndpointHealthStatus, 0),
		ClusterConditions: make([]shared.ClusterConditionStatus, 0),
	}

	for _, endpoint := range apiServerEndpoints {
		status, message := checkEndpointHealth(endpoint, client)
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

func checkEndpointHealth(endpoint string, client *kubeclient.KubeClient) (string, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	logMessage := fmt.Sprintf("Checking API Server endpoint: %s...\n", endpoint)
	healthzEndpoint := fmt.Sprintf("https://%s/healthz", endpoint)

	_, err := client.Clientset.CoreV1().RESTClient().Get().AbsPath(healthzEndpoint).DoRaw(ctx)
	if err != nil {
		return fmt.Sprintf("Error: Failed to reach API Server endpoint: %v", err), logMessage
	}

	return "Healthy", logMessage
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
