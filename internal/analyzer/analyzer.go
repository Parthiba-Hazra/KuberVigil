// healthanalyzer.go

// analyzer/healthanalyzer.go

package analyzer

import (
	"github.com/Parthiba-Hazra/kubervigil/internal/shared"
)

// HealthAnalyzer defines the interface for all health analysis.
type HealthAnalyzer interface {
	// AnalyzeAPIHealth analyzes the API health for Kubernetes resources and returns the APIVersionInfo.
	AnalyzeAPIHealth(provider shared.Resource, namespace string) (APIVersionInfo, error)

	// Add other health analysis methods here...
}

func AnalyzeAPIHealth(pr)
