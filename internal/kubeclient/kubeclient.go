package kubeclient

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// kubeClient is a concrete implementation of the Client interface.
type kubeClient struct {
	clientset *kubernetes.Clientset
}

// NewClient creates a new Kubernetes API client and returns it as the Client interface.
func NewClient(kubeconfigPath string) (*kubeClient, error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		return nil, err
	}

	// Create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &kubeClient{clientset: clientset}, nil
}
