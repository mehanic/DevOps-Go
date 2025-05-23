package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// Set up Kubernetes client
	kubeconfig := flag.String("kubeconfig", getDefaultKubeconfigPath(), "Path to kubeconfig file")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		// Try in-cluster config if kubeconfig not available
		config, err = rest.InClusterConfig()
		if err != nil {
			fmt.Printf("Error creating Kubernetes config: %v\n", err)
			os.Exit(1)
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error creating Kubernetes client: %v\n", err)
		os.Exit(1)
	}

	// Watch Pods
	go watchPods(clientset)

	// Handle graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	fmt.Println("Shutting down controller...")
}

// watchPods monitors pod events across all namespaces
func watchPods(clientset *kubernetes.Clientset) {
	watcher, err := clientset.CoreV1().Pods("").Watch(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Failed to start pod watcher: %v\n", err)
		return
	}
	defer watcher.Stop()

	fmt.Println("Watching pod events...")

	for event := range watcher.ResultChan() {
		pod, ok := event.Object.(*corev1.Pod)
		if !ok {
			fmt.Printf("Unexpected type\n")
			continue
		}

		switch event.Type {
		case watch.Added:
			fmt.Printf("🟢 Pod ADDED: %s/%s\n", pod.Namespace, pod.Name)
		case watch.Modified:
			fmt.Printf("🟡 Pod MODIFIED: %s/%s\n", pod.Namespace, pod.Name)
		case watch.Deleted:
			fmt.Printf("🔴 Pod DELETED: %s/%s\n", pod.Namespace, pod.Name)
		default:
			fmt.Printf("❓ Unknown event type: %s\n", event.Type)
		}
	}
}

// getDefaultKubeconfigPath returns the default kubeconfig path
func getDefaultKubeconfigPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return home + "/.kube/config"
}

