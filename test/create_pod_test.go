package test

import (
	"context"
	"testing"
	"example.com/m/mock_practice"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	fake "k8s.io/client-go/kubernetes/fake"
)

func TestFakePod(t *testing.T) {
	clientset := fake.NewSimpleClientset()

	podName := "test-pod"
	namespace := "default"

	err := mock_practice.CreatePod(clientset, podName, namespace)
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
