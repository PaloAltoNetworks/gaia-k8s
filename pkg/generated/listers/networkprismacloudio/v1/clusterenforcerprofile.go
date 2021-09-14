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

// ClusterEnforcerProfileLister helps list ClusterEnforcerProfiles.
// All objects returned here must be treated as read-only.
type ClusterEnforcerProfileLister interface {
	// List lists all ClusterEnforcerProfiles in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ClusterEnforcerProfile, err error)
	// Get retrieves the ClusterEnforcerProfile from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ClusterEnforcerProfile, error)
	ClusterEnforcerProfileListerExpansion
}

// clusterEnforcerProfileLister implements the ClusterEnforcerProfileLister interface.
type clusterEnforcerProfileLister struct {
	indexer cache.Indexer
}

// NewClusterEnforcerProfileLister returns a new ClusterEnforcerProfileLister.
func NewClusterEnforcerProfileLister(indexer cache.Indexer) ClusterEnforcerProfileLister {
	return &clusterEnforcerProfileLister{indexer: indexer}
}

// List lists all ClusterEnforcerProfiles in the indexer.
func (s *clusterEnforcerProfileLister) List(selector labels.Selector) (ret []*v1.ClusterEnforcerProfile, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClusterEnforcerProfile))
	})
	return ret, err
}

// Get retrieves the ClusterEnforcerProfile from the index for a given name.
func (s *clusterEnforcerProfileLister) Get(name string) (*v1.ClusterEnforcerProfile, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("clusterenforcerprofile"), name)
	}
	return obj.(*v1.ClusterEnforcerProfile), nil
}
