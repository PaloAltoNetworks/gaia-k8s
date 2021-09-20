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

// PUTrafficActionInformer provides access to a shared informer and lister for
// PUTrafficActions.
type PUTrafficActionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.PUTrafficActionLister
}

type pUTrafficActionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPUTrafficActionInformer constructs a new informer for PUTrafficAction type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPUTrafficActionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPUTrafficActionInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPUTrafficActionInformer constructs a new informer for PUTrafficAction type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPUTrafficActionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkV1().PUTrafficActions(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkV1().PUTrafficActions(namespace).Watch(context.TODO(), options)
			},
		},
		&networkprismacloudiov1.PUTrafficAction{},
		resyncPeriod,
		indexers,
	)
}

func (f *pUTrafficActionInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPUTrafficActionInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *pUTrafficActionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&networkprismacloudiov1.PUTrafficAction{}, f.defaultInformer)
}

func (f *pUTrafficActionInformer) Lister() v1.PUTrafficActionLister {
	return v1.NewPUTrafficActionLister(f.Informer().GetIndexer())
}
