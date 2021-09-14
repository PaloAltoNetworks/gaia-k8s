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
// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "go.aporeto.io/gaia-k8s/pkg/apis/networkprismacloudio/v1"
	"go.aporeto.io/gaia-k8s/pkg/generated/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type NetworkV1Interface interface {
	RESTClient() rest.Interface
	ClusterEnforcersGetter
	ClusterEnforcerProfilesGetter
	ClusterExternalNetworksGetter
	ClusterNetworkRuleSetPoliciesGetter
	ClusterProcessingUnitsGetter
	ExternalNetworksGetter
	NetworkRuleSetPoliciesGetter
	ProcessingUnitsGetter
}

// NetworkV1Client is used to interact with features provided by the network.prismacloud.io group.
type NetworkV1Client struct {
	restClient rest.Interface
}

func (c *NetworkV1Client) ClusterEnforcers() ClusterEnforcerInterface {
	return newClusterEnforcers(c)
}

func (c *NetworkV1Client) ClusterEnforcerProfiles() ClusterEnforcerProfileInterface {
	return newClusterEnforcerProfiles(c)
}

func (c *NetworkV1Client) ClusterExternalNetworks() ClusterExternalNetworkInterface {
	return newClusterExternalNetworks(c)
}

func (c *NetworkV1Client) ClusterNetworkRuleSetPolicies() ClusterNetworkRuleSetPolicyInterface {
	return newClusterNetworkRuleSetPolicies(c)
}

func (c *NetworkV1Client) ClusterProcessingUnits() ClusterProcessingUnitInterface {
	return newClusterProcessingUnits(c)
}

func (c *NetworkV1Client) ExternalNetworks(namespace string) ExternalNetworkInterface {
	return newExternalNetworks(c, namespace)
}

func (c *NetworkV1Client) NetworkRuleSetPolicies(namespace string) NetworkRuleSetPolicyInterface {
	return newNetworkRuleSetPolicies(c, namespace)
}

func (c *NetworkV1Client) ProcessingUnits(namespace string) ProcessingUnitInterface {
	return newProcessingUnits(c, namespace)
}

// NewForConfig creates a new NetworkV1Client for the given config.
func NewForConfig(c *rest.Config) (*NetworkV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &NetworkV1Client{client}, nil
}

// NewForConfigOrDie creates a new NetworkV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *NetworkV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new NetworkV1Client for the given RESTClient.
func New(c rest.Interface) *NetworkV1Client {
	return &NetworkV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *NetworkV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}