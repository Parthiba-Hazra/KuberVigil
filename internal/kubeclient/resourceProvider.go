// analyzer/resource_provider.go

package kubeclient

import (
	"context"

	"github.com/Parthiba-Hazra/kubervigil/internal/shared"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetResources implements Client.
func (c *kubeClient) GetResources(namespace string) ([]shared.Resource, error) {
	if namespace == "" {
		namespace = metav1.NamespaceDefault
	}

	var resources []shared.Resource

	// Fetch Deployments
	deployments, err := c.clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, deployment := range deployments.Items {
		resources = append(resources, shared.Resource{
			Kind:       "Deployment",
			Name:       deployment.Name,
			Namespace:  deployment.Namespace,
			APIVersion: deployment.APIVersion,
			CreatedAt:  deployment.CreationTimestamp.Time,
		})
	}

	// Fetch Pods
	pods, err := c.clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, pod := range pods.Items {
		resources = append(resources, shared.Resource{
			Kind:       "Pod",
			Name:       pod.Name,
			Namespace:  pod.Namespace,
			APIVersion: pod.APIVersion,
			CreatedAt:  pod.CreationTimestamp.Time,
		})
	}

	// Fetch Services
	services, err := c.clientset.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, service := range services.Items {
		resources = append(resources, shared.Resource{
			Kind:       "Service",
			Name:       service.Name,
			Namespace:  service.Namespace,
			APIVersion: service.APIVersion,
			CreatedAt:  service.CreationTimestamp.Time,
		})
	}

	// Fetch StatefulSets
	statefulSets, err := c.clientset.AppsV1().StatefulSets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, statefulSet := range statefulSets.Items {
		resources = append(resources, shared.Resource{
			Kind:       "StatefulSet",
			Name:       statefulSet.Name,
			Namespace:  statefulSet.Namespace,
			APIVersion: statefulSet.APIVersion,
			CreatedAt:  statefulSet.CreationTimestamp.Time,
		})
	}

	// Fetch DaemonSets
	daemonSets, err := c.clientset.AppsV1().DaemonSets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, daemonSet := range daemonSets.Items {
		resources = append(resources, shared.Resource{
			Kind:       "DaemonSet",
			Name:       daemonSet.Name,
			Namespace:  daemonSet.Namespace,
			APIVersion: daemonSet.APIVersion,
			CreatedAt:  daemonSet.CreationTimestamp.Time,
		})
	}

	// Fetch Jobs
	jobs, err := c.clientset.BatchV1().Jobs(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, job := range jobs.Items {
		resources = append(resources, shared.Resource{
			Kind:       "Job",
			Name:       job.Name,
			Namespace:  job.Namespace,
			APIVersion: job.APIVersion,
			CreatedAt:  job.CreationTimestamp.Time,
		})
	}

	// Fetch CronJobs
	cronJobs, err := c.clientset.BatchV1().CronJobs(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, cronJob := range cronJobs.Items {
		resources = append(resources, shared.Resource{
			Kind:       "CronJob",
			Name:       cronJob.Name,
			Namespace:  cronJob.Namespace,
			APIVersion: cronJob.APIVersion,
			CreatedAt:  cronJob.CreationTimestamp.Time,
		})
	}

	// Add other resource types similarly.

	return resources, nil
}
