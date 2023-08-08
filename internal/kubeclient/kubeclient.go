package kubeclient

import (
	"flag"
	"log"
	"path/filepath"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// kubeClient is a concrete implementation of the Client interface.
type KubeClient struct {
	Clientset     *kubernetes.Clientset
	DynamicClient dynamic.Interface
	Namespace     string
}

var kubeconfig *string

func GetConfig(kubeconfigPath string) *rest.Config {
	if len(kubeconfigPath) != 0 {
		kubeconfig = &kubeconfigPath
	} else {
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "~/.kube/config")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "~/.kube/config")
		}
		flag.Parse()
	}

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Printf("config error: %v", err)
		panic(err.Error())
	}
	return config
}

// NewClient creates a new Kubernetes API client and returns it as the Client interface.
func NewClient(kubeconfigPath string) (*KubeClient, error) {

	config := GetConfig(kubeconfigPath)

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Printf("client error: %v", err)
		panic(err.Error())
	}

	// fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	return &KubeClient{Clientset: clientset}, nil
}
