package main

import (
	"context"
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	fake "k8s.io/client-go/kubernetes/fake"
)

func createPod(clientset kubernetes.Interface, name, namespace string) error {
	pod := &v1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:            "nginx",
					Image:           "nginx",
					ImagePullPolicy: "Always",
				},
			},
		},
	}

	_, err := clientset.CoreV1().Pods(namespace).Create(context.Background(), pod, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}

func TestCreatePod(t *testing.T) {
	clientset := fake.NewSimpleClientset()

	podName := "test-pod"
	namespace := "default"

	err := createPod(clientset, podName, namespace)
	if err != nil {
		t.Errorf("Error creating pod: %v", err)
	}

	podList, err := clientset.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		t.Errorf("Error listing pods: %v", err)
	}

	if len(podList.Items) != 1 {
		t.Errorf("Expected 1 pod to be created, but got %d", len(podList.Items))
	}

	createdPod := podList.Items[0]
	if createdPod.Name != podName {
		t.Errorf("Expected pod name to be '%s', but got '%s'", podName, createdPod.Name)
	}
}