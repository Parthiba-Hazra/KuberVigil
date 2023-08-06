// shared/shared.go

package shared

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Resource represents a Kubernetes resource.
type Resource struct {
	Kind       string
	Name       string
	Namespace  string
	APIVersion string
	CreatedAt  time.Time
}

type VersionInfo struct {
	ResourceName        string
	ResourceNamespace   string
	ResourceKind        string
	CurrentAPIVersion   string
	PrefferedAPIVersion string
	APIhealth           bool
	Message             string
}

// ClusterConditionStatus represents the health status of a cluster condition.
type ClusterConditionStatus struct {
	Type    string
	Status  string
	Message string
}

// ResourceProvider provides the necessary resources to the analyzers.
type ResourceProvider interface {
	GetResources(namespace string) ([]Resource, error)
}

// StatefulSetHealthReport represents the detailed health report for a StatefulSet.
type StatefulSetHealthReport struct {
	StatefulSetName string
	Namespace       string
	Replicas        int32
	ReadyReplicas   int32
	CurrentReplicas int32
	UpdatedReplicas int32
	Conditions      []StatefulSetCondition
	OverallStatus   string
	OverallMessage  string
}

// StatefulSetCondition represents a StatefulSet condition.
type StatefulSetCondition struct {
	Type           string
	Status         string
	Message        string
	LastUpdateTime metav1.Time
}

// ServiceHealthReport represents the detailed health report for a service.
type ServiceHealthReport struct {
	ServiceName    string
	Namespace      string
	IP             string
	Port           int32
	TargetPort     int32
	Selector       map[string]string
	EndpointCount  int
	Conditions     []ServiceCondition
	OverallStatus  string
	OverallMessage string
}

// ServiceCondition represents a service condition.
type ServiceCondition struct {
	Type           string
	Status         string
	Message        string
	LastUpdateTime metav1.Time
}

// HealthReport represents the detailed health report for a pod.
type HealthReport struct {
	PodName             string
	Namespace           string
	Containers          []ContainerHealth
	OverallStatus       string
	OverallMessage      string
	CheckDNSResolution  bool
	CheckResourceLimits bool
}

// ContainerHealth represents the health status of a container within a pod.
type ContainerHealth struct {
	ContainerName  string
	ContainerState string
	Message        string
}

// HealthCheckResult represents the health check result of an Ingress resource.
type HealthCheckResult struct {
	Name        string
	Namespace   string
	Host        string
	Path        string
	Status      string
	Message     string
	HTTPCode    int
	LastChecked time.Time
}

// DeploymentHealthReport represents the detailed health report for a deployment.
type DeploymentHealthReport struct {
	DeploymentName    string
	Namespace         string
	Replicas          int32
	ReadyReplicas     int32
	UpdatedReplicas   int32
	AvailableReplicas int32
	Conditions        []DeploymentCondition
	OverallStatus     string
	OverallMessage    string
}

// DeploymentCondition represents a deployment condition.
type DeploymentCondition struct {
	Type           string
	Status         string
	Message        string
	LastUpdateTime metav1.Time
}

// DaemonSetHealthReport represents the detailed health report for a DaemonSet.
type DaemonSetHealthReport struct {
	DaemonSetName   string
	Namespace       string
	DesiredNumber   int32
	CurrentNumber   int32
	ReadyNumber     int32
	AvailableNumber int32
	UpdatedNumber   int32
	Conditions      []DaemonSetCondition
	OverallStatus   string
	OverallMessage  string
}

// DaemonSetCondition represents a DaemonSet condition.
type DaemonSetCondition struct {
	Type           string
	Status         string
	Message        string
	LastUpdateTime metav1.Time
}

// ConfigMapHealthReport represents the detailed health report for a ConfigMap.
type ConfigMapHealthReport struct {
	ConfigMapName  string
	Namespace      string
	DataKeys       []string
	OverallStatus  string
	OverallMessage string
}

// SecretHealthReport represents the detailed health report for a Secret.
type SecretHealthReport struct {
	SecretName     string
	Namespace      string
	DataKeys       []string
	OverallStatus  string
	OverallMessage string
}

// APIServerHealthReport represents the detailed health report for the Kubernetes API Server.
type APIServerHealthReport struct {
	OverallStatus     string
	OverallMessage    string
	APIEndpoints      []EndpointHealthStatus
	ClusterConditions []ClusterConditionStatus
}

// EndpointHealthStatus represents the health status of an API server endpoint.
type EndpointHealthStatus struct {
	Endpoint string
	Status   string
	Message  string
}

type ClusterCondition struct {
	Type    string
	Status  string
	Message string
}

// ResourceInfo represents the information of a Kubernetes API resource.
type ResourceInfo struct {
	Name       string
	Group      string
	Version    string
	APIPath    string
	Kind       string
	Namespaced bool
}

type Package struct {
	Kind       string    `json:"kind" yaml:"kind"`
	APIVersion string    `json:"apiVersion" yaml:"apiVersion"`
	Metadata   Meta      `json:"metadata" yaml:"metadata"`
	Items      []Package `json:"items" yaml:"items"`
}

type Meta struct {
	Name      string `json:"name" yaml:"name"`
	Namespace string `json:"namespace" yaml:"namespace"`
}
