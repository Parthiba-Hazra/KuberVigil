// healthanalyzer.go

// analyzer/healthanalyzer.go

package analyzer

import (
	"fmt"

	"github.com/Parthiba-Hazra/kubervigil/internal/kubeclient"
	"github.com/Parthiba-Hazra/kubervigil/internal/reporting"
)

// HealthAnalyzer defines the interface for all health analysis.
type HealthAnalyzer interface {
	// AnalyzeAPIHealth analyzes the API health for Kubernetes resources and returns the APIVersionInfo.
	AnalyzeAPIHealth(kubeCOnfigPath string, namespace string, resourceKind string) error

	// Add other health analysis methods here...
}

func GetClient(kubeCOnfigPath string) (*kubeclient.KubeClient, error) {

	configPath, err := kubeclient.GetKubeConfigPath(kubeCOnfigPath)
	if err != nil {
		return nil, err
	}
	client, err := kubeclient.NewClient(configPath)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// Function to compare preferred API version with current API version and generate messages/warnings.
func AnalyzeResourceHealth(kubeCOnfigPath string, namespace string, resourceKind string) error {

	client, err := GetClient(kubeCOnfigPath)
	if err != nil {
		return err
	}

	switch resourceKind {
	case "PODS":
		pods, err := client.GetPods(namespace)
		if err != nil {
			return err
		}
		report, err := CreatePodHealthReport(pods)
		if err != nil {
			return err
		}
		reporting.PrintHealthReports(report)

	case "SERVICES":
		services, err := client.GetServices(namespace)
		if err != nil {
			return err
		}
		report, err := CreateServiceHealthReport(services)
		if err != nil {
			return err
		}
		reporting.PrintServiceHealthReports(report)

	case "DEPLOYMENTS":
		deployments, err := client.GetDeployments(namespace)
		if err != nil {
			return err
		}
		report, err := CreateDeploymentHealthReport(deployments)
		if err != nil {
			return err
		}
		reporting.PrintDeploymentHealthReports(report)

	case "STATEFULSETS":
		statefulSets, err := client.GetStatefulsets(namespace)
		if err != nil {
			return err
		}
		report, err := CreateStatefulSetHealthReport(statefulSets)
		if err != nil {
			return err
		}
		reporting.PrintStatefulSetHealthReports(report)

	case "DAEMONSETS":
		daemonsets, err := client.GetDeamonsets(namespace)
		if err != nil {
			return err
		}
		report, err := CreateDaemonSetHealthReport(daemonsets)
		if err != nil {
			return err
		}
		reporting.PrintDaemonSetHealthReports(report)

	case "CONFIGMAPS":
		configMaps, err := client.GetConfigmaps(namespace)
		if err != nil {
			return err
		}
		report, err := CreateConfigMapHealthReport(configMaps)
		if err != nil {
			return err
		}
		reporting.PrintConfigMapHealthReports(report)

	// case "CRDs":
	// 	// Call the function for analyzing Custom Resource Definitions (CRDs) and return error if any.

	case "API":
		apiServerEndpoints, err := client.GetAPIEndpoints(namespace)
		if err != nil {
			return err
		}
		clusterCondition, err := client.GetClusterConditions()
		if err != nil {
			return err
		}
		report, err := CreateAPIServerHealthReport(apiServerEndpoints, clusterCondition)
		if err != nil {
			return err
		}
		reporting.PrintAPIServerHealthReport(report)

	// case "PersistentVolumes":
	// 	// Call the function for analyzing PersistentVolumes and return error if any.

	case "INGRESS":
		ingress, err := client.GetIngress1(namespace)
		if err != nil {
			return err
		}
		report, err := CreateIngressHealthReport(ingress, namespace)
		if err != nil {
			return err
		}
		reporting.PrintIngressHealthReports(report)

	default:
		return fmt.Errorf("unsupported resource type: %v", resourceKind)
	}

	// Use versionInfoList as needed for health analysis or further processing.

	return nil
}

func GetPreferedAPIversion(kubeCOnfigPath string) error {

	client, err := GetClient(kubeCOnfigPath)
	if err != nil {
		return err
	}
	preferredVersions, err := client.GetAPIVersionInfo()
	if err != nil {
		return err
	}
	reporting.PrintAPIVersionInfo(preferredVersions)

	return nil
	// var versionInfoList []VersionInfo
}
