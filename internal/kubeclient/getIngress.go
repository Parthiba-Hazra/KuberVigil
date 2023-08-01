package kubeclient

import (
	"context"

	v1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *KubeClient) GetIngress1(namespace string) (*v1.IngressList, error) {
	if namespace == "" {
		namespace = metav1.NamespaceDefault
	}

	ingressClient := c.clientset.NetworkingV1().Ingresses(namespace)
	ingresses, err := ingressClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return ingresses, nil
}
