// shared/shared.go

package shared

import "time"

// Resource represents a Kubernetes resource.
type Resource struct {
	Kind       string
	Name       string
	Namespace  string
	APIVersion string
	CreatedAt  time.Time
}

// ResourceProvider provides the necessary resources to the analyzers.
type ResourceProvider interface {
	GetResources(namespace string) ([]Resource, error)
}
