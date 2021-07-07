package providers

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const Harvester = "harvester"

var _ Provider = &HarvesterProvider{}

type HarvesterProvider struct{}

func (p *HarvesterProvider) GetName() string {
	return Harvester
}

func (p *HarvesterProvider) Detect(ctx context.Context, k8sClient kubernetes.Interface) (bool, error) {
	// Look for nodes that have an Harvester specific label
	listOpts := metav1.ListOptions{
		LabelSelector: "harvesterhci.io/managed",
		// Only need one
		Limit: 1,
	}

	nodes, err := k8sClient.CoreV1().Nodes().List(ctx, listOpts)
	if err != nil {
		return false, err
	}
	if len(nodes.Items) > 0 {
		return true, nil
	}
	return false, nil
}
