package kubeclient

import (
	"github.com/Parthiba-Hazra/kubervigil/internal/shared"
)

// Client is the interface that represents the Kubernetes API client.
type Client interface {
	// GetResources fetches a list of Kubernetes resources from the cluster.
	GetResources() ([]shared.Resource, error)

	// GetAPIVersionInfo retrieves information about the deprecation and removal status
	// of Kubernetes APIs for different versions.
	GetAPIVersionInfo() (APIVersionInfo, error)
}

// APIVersionInfo provides information about the deprecation and removal status of Kubernetes APIs.
type APIVersionInfo struct {
	// Map from API version to deprecation status (e.g., "v1": "deprecated", "apps/v1": "removed").
	DeprecationStatus map[string]string

	// Map from API version to removal status (e.g., "v1": "removed", "apps/v1": "not-removed").
	RemovalStatus map[string]string
}

// GetAPIVersionInfo implements Client.
func (c *kubeClient) GetAPIVersionInfo() ([]shared.Resource, error) {
	// Fetch the preferred API resources from the Kubernetes cluster.
	groupVersions, err := c.clientset.Discovery().ServerPreferredResources()
	if err != nil {
		return nil, err
	}

	var resources []shared.Resource

	// Iterate through the API group versions and extract the resources.
	for _, groupVersion := range groupVersions {
		for _, resourceList := range groupVersion.APIResources {
			apiVersion := groupVersion.GroupVersion
			resources = append(resources, shared.Resource{
				Kind:       resourceList.Kind,
				Name:       resourceList.Name,
				Namespace:  "", // No namespace for API group versions
				APIVersion: apiVersion,
			})
		}
	}

	// Return the extracted version info.
	return resources, nil
}
