package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	kwild "kwild/pkg/client"

	flag "github.com/spf13/pflag"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {

	clientset, err := kwild.GetKubeClient()
	if err != nil {
		fmt.Printf("Error creating Kubernetes client: %v\n", err)
		return
	}
	object := os.Args[1]
	searchString := flag.StringP("search", "s", "test", "String to search for in pod names")
	flag.Parse()

	if *searchString == "" {
		fmt.Println("Please provide a search string using the -search flag.")
		return
	}
	var podList *corev1.PodList
	var deploymentList *appsv1.DeploymentList
	var serviceList *corev1.ServiceList
	var statefulSetList *appsv1.StatefulSetList
	var namespaceList *corev1.NamespaceList

	if strings.ToLower(object) == "pod" || strings.ToLower(object) == "pods" {
		podList, err = clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		// podList = podList.(*corev1.PodList)
	} else if strings.ToLower(object) == "service" || strings.ToLower(object) == "services" {
		serviceList, err = clientset.CoreV1().Services("").List(context.TODO(), metav1.ListOptions{})
	} else if strings.ToLower(object) == "deployment" || strings.ToLower(object) == "deployments" || strings.ToLower(object) == "deploy" || strings.ToLower(object) == "deploys" {
		deploymentList, err = clientset.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})
	} else if strings.ToLower(object) == "statefulset" || strings.ToLower(object) == "statefulsets" || strings.ToLower(object) == "sts" {
		statefulSetList, err = clientset.AppsV1().StatefulSets("").List(context.TODO(), metav1.ListOptions{})
	} else if strings.ToLower(object) == "namespace" || strings.ToLower(object) == "namespaces" || strings.ToLower(object) == "ns" {
		namespaceList, err = clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})

	} else {
		fmt.Printf("Unsupported object type: %s. Please specify 'pod', 'service', 'deployment', or 'statefulset'.\n", object)
		return
	}
	// if strings.ToLower(object) == "pod" || strings.ToLower(object) == "pods" {
	// 	podList, err = clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	// 	// podList = podList.(*corev1.PodList)
	// }
	if podList != nil {
		for _, pod := range podList.Items {
			if !strings.Contains(pod.Name, *searchString) {
				continue
			}
			fmt.Printf("%s\t%s\n", pod.Namespace, pod.Name)
		}
	}
	if serviceList != nil {
		for _, service := range serviceList.Items {
			if !strings.Contains(service.Name, *searchString) {
				continue
			}
			fmt.Printf("%s\t%s\n", service.Namespace, service.Name)
		}
	}

	if deploymentList != nil {
		for _, deployment := range deploymentList.Items {
			if !strings.Contains(deployment.Name, *searchString) {
				continue
			}
			fmt.Printf("%s\t%s\n", deployment.Namespace, deployment.Name)
		}
	}

	if statefulSetList != nil {
		for _, statefulSet := range statefulSetList.Items {
			if !strings.Contains(statefulSet.Name, *searchString) {
				continue
			}
			fmt.Printf("%s\t%s\n", statefulSet.Namespace, statefulSet.Name)
		}
	}
	if namespaceList != nil {
		for _, namespace := range namespaceList.Items {
			if !strings.Contains(namespace.Name, *searchString) {
				continue
			}
			fmt.Printf("%s\n", namespace.Name)
		}
	}
}
