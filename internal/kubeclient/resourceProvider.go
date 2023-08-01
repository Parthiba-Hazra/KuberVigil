// analyzer/resource_provider.go

package kubeclient

import (
	"context"
	"fmt"

	"github.com/Parthiba-Hazra/kubervigil/internal/shared"
	v1 "k8s.io/api/apps/v1"
	v2 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetResources implements Client.
func (c *KubeClient) GetDeployments(namespace string) (*v1.DeploymentList, error) {
	if namespace == "" {
		namespace = metav1.NamespaceDefault
	}

	// Fetch Deployments
	deployments, err := c.clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return deployments, nil
}

func (c *KubeClient) GetPods(namespace string) (*v2.PodList, error) {
	if namespace == "" {
		namespace = metav1.NamespaceDefault
	}

	pods, err := c.clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return pods, nil
}

func (c *KubeClient) GetServices(namespace string) (*v2.ServiceList, error) {
	if namespace == "" {
		namespace = metav1.NamespaceDefault
	}

	services, err := c.clientset.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return services, nil
}

func (c *KubeClient) GetStatefulsets(namespace string) (*v1.StatefulSetList, error) {
	if namespace == "" {
		namespace = metav1.NamespaceDefault
	}

	statefulSets, err := c.clientset.AppsV1().StatefulSets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return statefulSets, nil
}

func (c *KubeClient) GetDeamonsets(namespace string) (*v1.DaemonSetList, error) {
	if namespace == "" {
		namespace = metav1.NamespaceDefault
	}

	daemonSets, err := c.clientset.AppsV1().DaemonSets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return daemonSets, nil
}

func (c *KubeClient) GetConfigmaps(namespace string) (*v2.ConfigMapList, error) {
	if namespace == "" {
		namespace = metav1.NamespaceDefault
	}

	configMaps, err := c.clientset.CoreV1().ConfigMaps(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return configMaps, nil
}

func (c *KubeClient) GetSecrets(namespace string) (*v2.SecretList, error) {
	if namespace == "" {
		namespace = metav1.NamespaceDefault
	}

	secrets, err := c.clientset.CoreV1().Secrets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return secrets, nil
}

func (c *KubeClient) GetAPIEndpoints(namespace string) ([]string, error) {
	if namespace == "" {
		namespace = metav1.NamespaceDefault
	}

	apiServerEndpoints, err := c.getAPIServerEndpoints()
	if err != nil {
		return nil, err
	}

	return apiServerEndpoints, nil
}

func (c *KubeClient) getAPIServerEndpoints() ([]string, error) {
	endpoints, err := c.clientset.CoreV1().Endpoints("kube-system").Get(context.TODO(), "kubernetes", metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get API server endpoints: %w", err)
	}

	var apiServerEndpoints []string
	for _, subset := range endpoints.Subsets {
		for _, address := range subset.Addresses {
			apiServerEndpoints = append(apiServerEndpoints, fmt.Sprintf("%s:%d", address.IP, subset.Ports[0].Port))
		}
	}

	return apiServerEndpoints, nil
}

func (c *KubeClient) GetClusterConditions() ([]shared.ClusterConditionStatus, error) {
	// Get the list of nodes in the cluster
	nodes, err := c.clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get nodes: %s", err)
	}

	// Check node availability
	allNodesAvailable := true
	for _, node := range nodes.Items {
		for _, condition := range node.Status.Conditions {
			if condition.Type == "Ready" && condition.Status != "True" {
				allNodesAvailable = false
				break
			}
		}
	}

	// Get the list of pods in the cluster
	pods, err := c.clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get pods: %s", err)
	}

	// Check pod status
	allPodsRunning := true
	for _, pod := range pods.Items {
		if pod.Status.Phase != "Running" {
			allPodsRunning = false
			break
		}
	}

	// Create a slice to hold the cluster conditions
	conditions := []shared.ClusterConditionStatus{}

	// Add NodeAvailability condition
	if allNodesAvailable {
		conditions = append(conditions, shared.ClusterConditionStatus{
			Type:    "NodeAvailability",
			Status:  "Healthy",
			Message: "All nodes are available",
		})
	} else {
		conditions = append(conditions, shared.ClusterConditionStatus{
			Type:    "NodeAvailability",
			Status:  "Error",
			Message: "Some nodes are not available",
		})
	}

	// Add PodStatus condition
	if allPodsRunning {
		conditions = append(conditions, shared.ClusterConditionStatus{
			Type:    "PodStatus",
			Status:  "Healthy",
			Message: "All pods are running",
		})
	} else {
		conditions = append(conditions, shared.ClusterConditionStatus{
			Type:    "PodStatus",
			Status:  "Error",
			Message: "Some pods are not running",
		})
	}

	return conditions, nil
}
