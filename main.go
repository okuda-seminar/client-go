// package main

// import (
// 	"context"
// 	"flag"
// 	"fmt"
// 	"path/filepath"

// 	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
// 	"k8s.io/client-go/kubernetes"
// 	"k8s.io/client-go/tools/clientcmd"
// 	"k8s.io/client-go/util/homedir"
// )
// func main() {
// 	var kubeconfig *string
// 	if home := homedir.HomeDir(); home != "" {
// 		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
// 	} else {
// 		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
// 	}
// 	flag.Parse()

// 	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
// 	if err != nil {
// 		panic(err)
// 	}
// 	clientset, err := kubernetes.NewForConfig(config)
// 	if err != nil {
// 		panic(err)
// 	}

// 	podsClient := clientset.CoreV1().Pods("default")
// 	podList, err := podsClient.List(context.TODO(), metav1.ListOptions{})
// 	if err != nil {
// 		panic(err)
// 	}

// 	for _, pod := range podList.Items {
// 		fmt.Printf("Pod: %s\n", pod.GetName())
// 	}
// }

// package main

// import (
// 	"flag"
// 	"fmt"
// 	"path/filepath"

// 	"k8s.io/client-go/kubernetes"
// 	"k8s.io/client-go/tools/clientcmd"
// 	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
// 	v1 "k8s.io/api/core/v1"
// 	"context"
// )

// func main() {
// 	var kubeconfig *string
// 	home := homedir()
// 	if home != "" {
// 		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
// 	} else {
// 		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
// 	}
// 	flag.Parse()

// 	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
// 	if err != nil {
// 		panic(err)
// 	}
// 	clientset, err := kubernetes.NewForConfig(config)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Get information for each resource
// 	podsList, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
// 	if err != nil {
// 		panic(err)
// 	}

// 	getAndPrintPods(podsList)
// }

// func getAndPrintPods(podsList *v1.PodList) {
// 	fmt.Println("Pods:")
// 	for _, pod := range podsList.Items {
// 		fmt.Printf("- %s\n", pod.Name)
// 	}
// }

// func homedir() string {
// 	home := ""
// 	// Implement the logic to determine the home directory path here
// 	return home
// }

package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"os/user"
)

func main() {
	var kubeconfig *string
	home := homedir()
	if home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// Get information for each resource
	// make function
	resourceGet(clientset)
}

func homedir() string {
	usr, err := user.Current()
	if err != nil {
		// handle the error if needed
		return ""
	}
	return usr.HomeDir
}

func resourceGet(clientset *kubernetes.Clientset) {
	resources, err := clientset.Discovery().ServerPreferredResources() // this line cannot control
	if err != nil {
		panic(err)
	}

	for _, resourceList := range resources {
		groupVersion, err := schema.ParseGroupVersion(resourceList.GroupVersion)
		if err != nil {
			panic(err)
		}
		for _, resource := range resourceList.APIResources {
			fmt.Printf("Resource: %s\n", resource.Name)
			fmt.Printf("  Group: %s\n", groupVersion.Group)
			fmt.Printf("  Version: %s\n", groupVersion.Version)
			fmt.Printf("  Kind: %s\n", resource.Kind)
			fmt.Printf("  Namespaced: %v\n", resource.Namespaced)
			fmt.Println()
		}
	}
}


// unit test
// do mock for test
// 