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

// ClusterEnforcerProfileInformer provides access to a shared informer and lister for
// ClusterEnforcerProfiles.
type ClusterEnforcerProfileInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ClusterEnforcerProfileLister
}

type clusterEnforcerProfileInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterEnforcerProfileInformer constructs a new informer for ClusterEnforcerProfile type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterEnforcerProfileInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterEnforcerProfileInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterEnforcerProfileInformer constructs a new informer for ClusterEnforcerProfile type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterEnforcerProfileInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkV1().ClusterEnforcerProfiles().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkV1().ClusterEnforcerProfiles().Watch(context.TODO(), options)
			},
		},
		&networkprismacloudiov1.ClusterEnforcerProfile{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterEnforcerProfileInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterEnforcerProfileInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterEnforcerProfileInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&networkprismacloudiov1.ClusterEnforcerProfile{}, f.defaultInformer)
}

func (f *clusterEnforcerProfileInformer) Lister() v1.ClusterEnforcerProfileLister {
	return v1.NewClusterEnforcerProfileLister(f.Informer().GetIndexer())
}