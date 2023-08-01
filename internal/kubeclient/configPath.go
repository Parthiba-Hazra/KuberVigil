package kubeclient

import (
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/client-go/util/homedir"
)

// GetKubeConfigPath returns the path of the kubeconfig file.
// If kubeconfigPath is provided, it will use that path.
// If kubeconfigPath is empty, it will automatically detect the kubeconfig file.
func GetKubeConfigPath(kubeconfigPath string) (string, error) {
	if kubeconfigPath != "" {
		// If kubeconfigPath is provided, use it as is.
		return kubeconfigPath, nil
	}

	// If kubeconfigPath is empty, try to detect the kubeconfig file.
	homeDir := homedir.HomeDir()
	if homeDir != "" {
		kubeconfigPath = filepath.Join(homeDir, ".kube", "config")
		if _, err := os.Stat(kubeconfigPath); err == nil {
			// kubeconfig file exists in the home directory.
			return kubeconfigPath, nil
		}
	}

	// If kubeconfig file was not found in the home directory, check the KUBECONFIG environment variable.
	kubeconfigPath = os.Getenv("KUBECONFIG")
	if kubeconfigPath != "" {
		return kubeconfigPath, nil
	}

	// If kubeconfig file was not found in the home directory or in the KUBECONFIG environment variable,
	// return an error indicating that the kubeconfig file was not found.
	return "", fmt.Errorf("kubeconfig file not found")
}
