package main

import (
	"context"
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	fake "k8s.io/client-go/kubernetes/fake"
)

func TestCreatePod(t *testing.T) {
	// Create a fake clientset
	clientset := fake.NewSimpleClientset()

	pod := &v1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-pod",
			Namespace: "default",
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

	// Call the Create function of the clientset to create the pod
	_, err := clientset.CoreV1().Pods("default").Create(context.Background(), pod, metav1.CreateOptions{})
	if err != nil {
		t.Errorf("Error creating pod: %v", err)
	}

	// Verify that the pod was created in the fake clientset
	podList, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		t.Errorf("Error listing pods: %v", err)
	}

	if len(podList.Items) != 1 {
		t.Errorf("Expected 1 pod to be created, but got %d", len(podList.Items))
	}

	createdPod := podList.Items[0]
	if createdPod.Name != "test-pod" {
		t.Errorf("Expected pod name to be 'test-pod', but got '%s'", createdPod.Name)
	}
}
