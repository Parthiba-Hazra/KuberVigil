package analyzer

import (
	"github.com/Parthiba-Hazra/kubervigil/internal/kubeclient"
	"github.com/Parthiba-Hazra/kubervigil/internal/reporting"
	"github.com/Parthiba-Hazra/kubervigil/internal/shared"
)

func CompareAPIVersions(preferredAPIVersions map[string]string, resourceInfo []*shared.Package) []*shared.CompareVersion {
	var resources []*shared.CompareVersion

	for _, info := range resourceInfo {
		resource := &shared.CompareVersion{
			ResourceName: info.Metadata.Name,
			ResourceKind: info.Kind,
			Namespace:    info.Metadata.Namespace,
			CurrentAPI:   info.APIVersion,
		}

		if preferredAPI, ok := preferredAPIVersions[info.Kind]; ok {
			resource.PreferredAPI = preferredAPI
			resource.PreferredEquals = resource.CurrentAPI == preferredAPI
		}

		resources = append(resources, resource)
	}

	return resources
}

func GetComparisonAPIversion(configPath string, namespace string) ([]*shared.CompareVersion, error) {

	discoveryClient, err := kubeclient.NewDiscoveryClient(namespace, configPath)
	if err != nil {
		return nil, err
	}

	preferredAPIVersions, err := discoveryClient.PreferredAPIversion()
	if err != nil {
		return nil, err
	}

	resourceInfo, err := discoveryClient.GetResourceInfo()
	if err != nil {
		return nil, err
	}

	results := CompareAPIVersions(preferredAPIVersions, resourceInfo)

	reporting.PrintComparisonResults(results)

	return results, nil
}
