package reporting

import (
	"fmt"

	"github.com/Parthiba-Hazra/kubervigil/internal/shared"
)

func PrintPackages(packages []*shared.Package) {
	fmt.Println("Packages:")
	for _, pkg := range packages {
		fmt.Printf("  Kind: %s\n", pkg.Kind)
		fmt.Printf("  APIVersion: %s\n", pkg.APIVersion)
		fmt.Printf("  Name: %s\n", pkg.Metadata.Name)
		fmt.Printf("  Namespace: %s\n", pkg.Metadata.Namespace)
		fmt.Println()
	}
}
