package kubeclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/Parthiba-Hazra/kubervigil/internal/reporting"
	"github.com/Parthiba-Hazra/kubervigil/internal/shared"
	"gopkg.in/yaml.v3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
)

type DiscoveryClient struct {
	ClientSet       dynamic.Interface
	restConfig      *rest.Config
	DiscoveryClient discovery.DiscoveryInterface
	namespace       string
}

// NewDiscoveryClient returns a new struct with config portions complete.
func NewDiscoveryClient(namespace string, kubeconfigPath string) (*DiscoveryClient, error) {
	dc := &DiscoveryClient{}

	dc.restConfig = GetConfig(kubeconfigPath)

	var err error
	dc.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(dc.restConfig)
	if err != nil {
		return nil, err
	}

	dc.namespace = namespace

	dc.ClientSet, err = dynamic.NewForConfig(dc.restConfig)
	if err != nil {
		return nil, err
	}
	return dc, nil
}

func (dc *DiscoveryClient) GetResourceInfo() ([]*shared.Package, error) {

	resourcelist, err := dc.DiscoveryClient.ServerPreferredResources()
	if err != nil {
		return nil, err
	}

	gvrs := []schema.GroupVersionResource{}
	for _, rl := range resourcelist {
		for i := range rl.APIResources {
			if dc.namespace != "" && !rl.APIResources[i].Namespaced {
				continue
			}
			gv, _ := schema.ParseGroupVersion(rl.GroupVersion)
			ResourceName := rl.APIResources[i].Name
			g := schema.GroupVersionResource{Group: gv.Group, Version: gv.Version, Resource: ResourceName}
			gvrs = append(gvrs, g)
		}
	}

	var results []*shared.Package
	for _, g := range gvrs {
		nri := dc.ClientSet.Resource(g)
		var ri dynamic.ResourceInterface = nri
		if dc.namespace != "" {
			ri = nri.Namespace(dc.namespace)
		}
		log.Printf("fetcing data: %s.%s.%s", g.Resource, g.Version, g.Group)
		rs, err := ri.List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			log.Printf("failed to fetch: %v %v", g, err)
			continue
		}

		if len(rs.Items) == 0 {
			log.Printf("no annotations for Resource-version %s", rs.GetAPIVersion())
			obj := rs.UnstructuredContent()
			data, err := json.Marshal(obj)
			if err != nil {
				log.Printf("failed to marshal data %v", err.Error())
				return nil, err
			}
			// Instead of directly returning, we will append the packages to the results slice.
			packagesFromYaml, err1 := GetPackagesFromYamlManifest(data)
			if err1 != nil {
				packagesFromJSON, err2 := GetPackagesFromJSONManifest(data)
				if err2 != nil {
					return nil, fmt.Errorf("error parsing yaml and json data, %v,  %v", err1, err2)
				}
				results = append(results, packagesFromJSON...)
			} else {
				results = append(results, packagesFromYaml...)
			}
		} else {
			for _, r := range rs.Items {
				if jsonManifest, ok := r.GetAnnotations()["kubectl.kubernetes.io/last-applied-configuration"]; ok {
					var manifest map[string]interface{}

					err := json.Unmarshal([]byte(jsonManifest), &manifest)
					if err != nil {
						fmt.Printf("failed to parse 'last-applied-configuration' annotation of resource %s/%s: %s", r.GetNamespace(), r.GetName(), err.Error())
						continue
					}
					data, err := json.Marshal(manifest)
					if err != nil {
						fmt.Printf("failed to marshal the data %v", err.Error())
						return nil, err
					}
					// Instead of directly returning, we will append the packages to the results slice.
					packagesFromYaml, err1 := GetPackagesFromYamlManifest(data)
					if err1 != nil {
						packagesFromJSON, err2 := GetPackagesFromJSONManifest(data)
						if err2 != nil {
							return nil, fmt.Errorf("error parsing yaml and json data, %v,  %v", err1, err2)
						}
						results = append(results, packagesFromJSON...)
					} else {
						results = append(results, packagesFromYaml...)
					}
				}
			}
		}

	}

	log.Printf("Result from resources: %d", len(results))
	return results, nil
}

// GetPackagesFromYamlManifest extracts packages from the given manifest data.
func GetPackagesFromYamlManifest(data []byte) ([]*shared.Package, error) {
	var packages []*shared.Package
	var tError *yaml.TypeError

	decoder := yaml.NewDecoder(bytes.NewReader(data))
	for {
		pack := &shared.Package{}
		err := decoder.Decode(pack)
		if err != nil {
			if err == io.EOF {
				break
			}
			if errors.As(err, &tError) {
				log.Printf("error decoding package: %v", err)
				continue
			}
			return packages, err

		}
		if len(pack.Items) > 0 {
			log.Printf("found more items in items in current package: %v", len(pack.Items))
			for _, pk := range pack.Items {
				currentPack := pk
				packages = append(packages, &currentPack)
			}
		} else {
			packages = append(packages, pack)
		}
	}

	return packages, nil
}

// // extractPackageFromObject extracts a package from the given object.
// func extractPackageFromObject(obj interface{}) (*shared.Package, error) {
// 	data, err := json.Marshal(obj)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var pk shared.Package
// 	if err := customUnmarshal(data, &pk); err != nil {
// 		return nil, err
// 	}

// 	if pk.Kind == "" || pk.APIVersion == "" {
// 		return nil, fmt.Errorf("invalid package: %+v", pk)
// 	}

// 	return &pk, nil
// }

// GetPackagesFromJSONManifest extracts packages from the given JSON manifest data.
func GetPackagesFromJSONManifest(data []byte) ([]*shared.Package, error) {
	var packages []*shared.Package

	pack := &shared.Package{}
	err := json.Unmarshal(data, pack)
	if err != nil {
		return nil, err
	}
	if len(pack.Items) > 0 {
		log.Printf("found more items in items in current package: %v", len(pack.Items))
		for _, pk := range pack.Items {
			currentPack := pk
			packages = append(packages, &currentPack)
		}
	} else {
		packages = append(packages, pack)
	}

	return packages, nil
}

func PrintAPIinfo(namespace string, kubeconfigPath string) {
	dc, err := NewDiscoveryClient(namespace, kubeconfigPath)
	if err != nil {
		fmt.Printf("error getting discovery client: %v", err)
	}
	info, err := dc.GetResourceInfo()
	if err != nil {
		fmt.Printf("error getting resource info: %v", err)
	}
	reporting.PrintPackages(info)
}
