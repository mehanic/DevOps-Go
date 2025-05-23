package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// initialize the Kubernetes cluster with kubeadm
	initCmd := exec.Command("sudo", "kubeadm", "init")
	initOutput, err := initCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(initOutput))
	// configure kubectl to connect to the Kubernetes API server
	configCmd := exec.Command("sudo", "cp", "/etc/kubernetes/admin.conf", "$HOME/")
	configOutput, err := configCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(configOutput))
	// apply a network plugin for the cluster
	netCmd := exec.Command("kubectl", "apply", "-f", "https://docs.projectcalico.org/v3.19/manifests/calico.yaml")
	netOutput, err := netCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(netOutput))
}

