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
// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "go.aporeto.io/gaia-k8s/pkg/apis/networkprismacloudio/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterExternalNetworkLister helps list ClusterExternalNetworks.
// All objects returned here must be treated as read-only.
type ClusterExternalNetworkLister interface {
	// List lists all ClusterExternalNetworks in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ClusterExternalNetwork, err error)
	// Get retrieves the ClusterExternalNetwork from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ClusterExternalNetwork, error)
	ClusterExternalNetworkListerExpansion
}

// clusterExternalNetworkLister implements the ClusterExternalNetworkLister interface.
type clusterExternalNetworkLister struct {
	indexer cache.Indexer
}

// NewClusterExternalNetworkLister returns a new ClusterExternalNetworkLister.
func NewClusterExternalNetworkLister(indexer cache.Indexer) ClusterExternalNetworkLister {
	return &clusterExternalNetworkLister{indexer: indexer}
}

// List lists all ClusterExternalNetworks in the indexer.
func (s *clusterExternalNetworkLister) List(selector labels.Selector) (ret []*v1.ClusterExternalNetwork, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClusterExternalNetwork))
	})
	return ret, err
}

// Get retrieves the ClusterExternalNetwork from the index for a given name.
func (s *clusterExternalNetworkLister) Get(name string) (*v1.ClusterExternalNetwork, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("clusterexternalnetwork"), name)
	}
	return obj.(*v1.ClusterExternalNetwork), nil
}
