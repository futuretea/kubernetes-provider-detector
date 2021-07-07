package providers

import (
	"context"

	"k8s.io/client-go/kubernetes"
)

// Provider is the interface all providers need to implement
type Provider interface {
	GetName() string
	Detect(ctx context.Context, k8sClient kubernetes.Interface) (bool, error)
}
