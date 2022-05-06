// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "open-cluster-management.io/api/client/cluster/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ManagedClusters returns a ManagedClusterInformer.
	ManagedClusters() ManagedClusterInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ManagedClusters returns a ManagedClusterInformer.
func (v *version) ManagedClusters() ManagedClusterInformer {
	return &managedClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
