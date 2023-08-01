package kubeclient

import (
	"flag"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// kubeClient is a concrete implementation of the Client interface.
type KubeClient struct {
	clientset *kubernetes.Clientset
}

// NewClient creates a new Kubernetes API client and returns it as the Client interface.
func NewClient(kubeconfigPath string) (*KubeClient, error) {
	// config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	// if err != nil {
	// 	return nil, err
	// }

	// // Create the clientset
	// clientset, err := kubernetes.NewForConfig(config)
	// if err != nil {
	// 	return nil, err
	// }

	// return &kubeClient{clientset: clientset}, nil
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "~/.kube/config")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "~/.kube/config")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	return &KubeClient{clientset: clientset}, nil
}
