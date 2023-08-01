// reporting/reporting.go

package reporting

import (
	"fmt"
)

func PrintAPIVersionInfo(versionInfo map[string]string) {
	fmt.Println("API Version Info:")
	fmt.Println("=================")

	for kind, apiVersion := range versionInfo {
		fmt.Printf("Resource Kind: %s\n", kind)
		fmt.Printf("Preferred API Version: %s\n", apiVersion)
		fmt.Println("=================")
	}
}

// // Example usage:
// func Print(namespace string, configPath string) error {

// 	versionInfo, err := analyzer.CompareVersions(configPath, namespace)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return err
// 	}

// 	PrintVersionInfo(versionInfo)
// 	return nil
// }
