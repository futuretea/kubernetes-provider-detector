package detector

import (
	"context"
	"errors"

	"k8s.io/client-go/kubernetes"

	"github.com/rancher/kubernetes-provider-detector/providers"
)

var (
	allProviders       []providers.Provider
	ErrUnknownProvider = errors.New("unknown provider")
)

func addProvider(provider providers.Provider) {
	allProviders = append(allProviders, provider)
}

func init() {
	addProvider(&providers.HarvesterProvider{})
	addProvider(&providers.AKSProvider{})
	addProvider(&providers.DockerProvider{})
	addProvider(&providers.EKSProvider{})
	addProvider(&providers.GKEProvider{})
	addProvider(&providers.K3sProvider{})
	addProvider(&providers.MinikubeProvider{})
	addProvider(&providers.RKEProvider{})
	addProvider(&providers.RKEWindowsProvider{})
	addProvider(&providers.RKE2Provider{})
}

// DetectProvider accepts a k8s interface and checks all registered providers for a match
func DetectProvider(ctx context.Context, k8sClient kubernetes.Interface) (string, error) {
	for _, p := range allProviders {
		// Check the context before calling the provider
		if err := ctx.Err(); err != nil {
			return "", err
		}

		if ok, err := p.Detect(ctx, k8sClient); err != nil {
			return "", err
		} else if ok {
			return p.GetName(), nil
		}
	}
	return "", ErrUnknownProvider
}

// ListRegisteredProviders returns a list of the names of all providers
func ListRegisteredProviders() []string {
	providerNames := make([]string, len(allProviders))
	for _, provider := range allProviders {
		providerNames = append(providerNames, provider.GetName())
	}
	return providerNames
}
