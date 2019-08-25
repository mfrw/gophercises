package main

import (
	"flag"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kc := flag.String("kubeconfg", "/Users/mfrw/.kube/kind-config-kind", "kubeconfig file")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kc)
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	pod, err := clientset.CoreV1().Pods("default").List(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range pod.Items {
		fmt.Println(p.GetName())
	}
}
