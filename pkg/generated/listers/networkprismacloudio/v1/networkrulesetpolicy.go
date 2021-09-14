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

// NetworkRuleSetPolicyLister helps list NetworkRuleSetPolicies.
// All objects returned here must be treated as read-only.
type NetworkRuleSetPolicyLister interface {
	// List lists all NetworkRuleSetPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.NetworkRuleSetPolicy, err error)
	// NetworkRuleSetPolicies returns an object that can list and get NetworkRuleSetPolicies.
	NetworkRuleSetPolicies(namespace string) NetworkRuleSetPolicyNamespaceLister
	NetworkRuleSetPolicyListerExpansion
}

// networkRuleSetPolicyLister implements the NetworkRuleSetPolicyLister interface.
type networkRuleSetPolicyLister struct {
	indexer cache.Indexer
}

// NewNetworkRuleSetPolicyLister returns a new NetworkRuleSetPolicyLister.
func NewNetworkRuleSetPolicyLister(indexer cache.Indexer) NetworkRuleSetPolicyLister {
	return &networkRuleSetPolicyLister{indexer: indexer}
}

// List lists all NetworkRuleSetPolicies in the indexer.
func (s *networkRuleSetPolicyLister) List(selector labels.Selector) (ret []*v1.NetworkRuleSetPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.NetworkRuleSetPolicy))
	})
	return ret, err
}

// NetworkRuleSetPolicies returns an object that can list and get NetworkRuleSetPolicies.
func (s *networkRuleSetPolicyLister) NetworkRuleSetPolicies(namespace string) NetworkRuleSetPolicyNamespaceLister {
	return networkRuleSetPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NetworkRuleSetPolicyNamespaceLister helps list and get NetworkRuleSetPolicies.
// All objects returned here must be treated as read-only.
type NetworkRuleSetPolicyNamespaceLister interface {
	// List lists all NetworkRuleSetPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.NetworkRuleSetPolicy, err error)
	// Get retrieves the NetworkRuleSetPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.NetworkRuleSetPolicy, error)
	NetworkRuleSetPolicyNamespaceListerExpansion
}

// networkRuleSetPolicyNamespaceLister implements the NetworkRuleSetPolicyNamespaceLister
// interface.
type networkRuleSetPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NetworkRuleSetPolicies in the indexer for a given namespace.
func (s networkRuleSetPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1.NetworkRuleSetPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.NetworkRuleSetPolicy))
	})
	return ret, err
}

// Get retrieves the NetworkRuleSetPolicy from the indexer for a given namespace and name.
func (s networkRuleSetPolicyNamespaceLister) Get(name string) (*v1.NetworkRuleSetPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("networkrulesetpolicy"), name)
	}
	return obj.(*v1.NetworkRuleSetPolicy), nil
}