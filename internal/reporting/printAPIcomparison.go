package reporting

import (
	"fmt"

	"github.com/Parthiba-Hazra/kubervigil/internal/shared"
)

func PrintComparisonResults(results []*shared.CompareVersion) {
	if len(results) == 0 {
		fmt.Println("No resources found.")
		return
	}

	fmt.Println("Comparison Results:")
	fmt.Println("===================")

	for _, result := range results {
		message := "Your API version is good."
		if !result.PreferredEquals {
			message = fmt.Sprintf("Your API version for this resource of kind '%s' is not the same as the preferred version. It might be deprecated or removed.", result.ResourceName)
		}

		fmt.Println("Resource:", result.ResourceName)
		fmt.Println("Resource kind:", result.ResourceKind)
		fmt.Println("Resource namespace:", result.Namespace)
		fmt.Println("Preferred API Version:", result.PreferredAPI)
		fmt.Println("Current API Version:", result.CurrentAPI)
		fmt.Println("Message:", message)
		fmt.Println("===================")
	}
}
