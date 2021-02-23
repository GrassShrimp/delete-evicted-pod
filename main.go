package main

import (
	"context"
	"fmt"
	"sync"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{
		FieldSelector: "status.phase=Failed",
	})

	if err != nil {
		panic(err.Error())
	}

	wg := new(sync.WaitGroup)

	for _, pod := range pods.Items {
		wg.Add(1)
		go (func(_pod v1.Pod, _wg *sync.WaitGroup) {
			defer _wg.Done()

			err := clientset.CoreV1().Pods(_pod.Namespace).Delete(context.TODO(), _pod.Name, metav1.DeleteOptions{})

			if err != nil {
				panic(err.Error())
			}

			fmt.Printf("Deleted Evicted Pod %s in namespace %s\n", _pod.Name, _pod.Namespace)
		})(pod, wg)
	}

	wg.Wait()

	fmt.Println("Deleted all Evicted Pods")
}
