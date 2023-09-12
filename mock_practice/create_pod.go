package mock_practice

import (
    "context"

    v1 "k8s.io/api/core/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes"
)

func CreatePod(clientset kubernetes.Interface, name, namespace string) error {
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
