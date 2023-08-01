package analyzer

import (
	"context"

	"github.com/Parthiba-Hazra/kubervigil/internal/shared"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// CreateConfigMapHealthReport creates a detailed health report for the ConfigMaps in the given namespace.
func CreateConfigMapHealthReport(configMaps *corev1.ConfigMapList) ([]shared.ConfigMapHealthReport, error) {

	var healthReports []shared.ConfigMapHealthReport

	for _, configMap := range configMaps.Items {
		healthReport := shared.ConfigMapHealthReport{
			ConfigMapName: configMap.Name,
			Namespace:     configMap.Namespace,
			DataKeys:      getConfigMapDataKeys(configMap),
		}

		healthReport.OverallStatus, healthReport.OverallMessage = getOverallConfigMapHealthStatus(healthReport.DataKeys)

		healthReports = append(healthReports, healthReport)
	}

	return healthReports, nil
}

// CreateSecretHealthReport creates a detailed health report for the Secrets in the given namespace.
func CreateSecretHealthReport(clientset kubernetes.Interface, namespace string) ([]shared.SecretHealthReport, error) {
	secrets, err := clientset.CoreV1().Secrets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var healthReports []shared.SecretHealthReport

	for _, secret := range secrets.Items {
		healthReport := shared.SecretHealthReport{
			SecretName: secret.Name,
			Namespace:  secret.Namespace,
			DataKeys:   getSecretDataKeys(secret),
		}

		healthReport.OverallStatus, healthReport.OverallMessage = getOverallSecretHealthStatus(healthReport.DataKeys)

		healthReports = append(healthReports, healthReport)
	}

	return healthReports, nil
}

func getConfigMapDataKeys(configMap corev1.ConfigMap) []string {
	var dataKeys []string
	for key := range configMap.Data {
		dataKeys = append(dataKeys, key)
	}
	return dataKeys
}

func getSecretDataKeys(secret corev1.Secret) []string {
	var dataKeys []string
	for key := range secret.Data {
		dataKeys = append(dataKeys, key)
	}
	return dataKeys
}

func getOverallConfigMapHealthStatus(dataKeys []string) (string, string) {
	if len(dataKeys) == 0 {
		return "Error", "ConfigMap has no data"
	}
	return "Healthy", "All checks passed"
}

func getOverallSecretHealthStatus(dataKeys []string) (string, string) {
	if len(dataKeys) == 0 {
		return "Error", "Secret has no data"
	}
	return "Healthy", "All checks passed"
}
