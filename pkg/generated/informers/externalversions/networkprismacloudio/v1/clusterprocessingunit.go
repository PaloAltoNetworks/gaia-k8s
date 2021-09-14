/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	networkprismacloudiov1 "go.aporeto.io/gaia-k8s/pkg/apis/networkprismacloudio/v1"
	versioned "go.aporeto.io/gaia-k8s/pkg/generated/clientset/versioned"
	internalinterfaces "go.aporeto.io/gaia-k8s/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "go.aporeto.io/gaia-k8s/pkg/generated/listers/networkprismacloudio/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterProcessingUnitInformer provides access to a shared informer and lister for
// ClusterProcessingUnits.
type ClusterProcessingUnitInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ClusterProcessingUnitLister
}

type clusterProcessingUnitInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterProcessingUnitInformer constructs a new informer for ClusterProcessingUnit type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterProcessingUnitInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterProcessingUnitInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterProcessingUnitInformer constructs a new informer for ClusterProcessingUnit type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterProcessingUnitInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkV1().ClusterProcessingUnits().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkV1().ClusterProcessingUnits().Watch(context.TODO(), options)
			},
		},
		&networkprismacloudiov1.ClusterProcessingUnit{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterProcessingUnitInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterProcessingUnitInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterProcessingUnitInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&networkprismacloudiov1.ClusterProcessingUnit{}, f.defaultInformer)
}

func (f *clusterProcessingUnitInformer) Lister() v1.ClusterProcessingUnitLister {
	return v1.NewClusterProcessingUnitLister(f.Informer().GetIndexer())
}