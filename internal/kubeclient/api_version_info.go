package kubeclient

import (
	"github.com/Parthiba-Hazra/kubervigil/internal/shared"
)

// Client is the interface that represents the Kubernetes API client.
type Client interface {
	// GetResources fetches a list of Kubernetes resources from the cluster.
	GetResources() ([]shared.Resource, error)

	GetAPIVersionInfo() (shared.Resource, error)
}

func (c *KubeClient) GetAPIVersionInfo() (map[string]string, error) {
	// Fetch the preferred API resources from the Kubernetes cluster.
	groupVersions, err := c.Clientset.Discovery().ServerPreferredResources()
	if err != nil {
		return nil, err
	}

	// Create a map to store resource kinds and their preferred API versions.
	versionInfo := make(map[string]string)

	// Iterate through the API group versions and extract the preferred API versions for each resource kind.
	for _, groupVersion := range groupVersions {
		for _, resourceList := range groupVersion.APIResources {
			apiVersion := groupVersion.GroupVersion
			versionInfo[resourceList.Kind] = apiVersion
		}
	}

	// Return the extracted version info.
	return versionInfo, nil
}
