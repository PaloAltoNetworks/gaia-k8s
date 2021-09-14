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

package fake

import (
	v1 "go.aporeto.io/gaia-k8s/pkg/generated/clientset/versioned/typed/networkprismacloudio/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeNetworkV1 struct {
	*testing.Fake
}

func (c *FakeNetworkV1) ClusterEnforcers() v1.ClusterEnforcerInterface {
	return &FakeClusterEnforcers{c}
}

func (c *FakeNetworkV1) ClusterEnforcerProfiles() v1.ClusterEnforcerProfileInterface {
	return &FakeClusterEnforcerProfiles{c}
}

func (c *FakeNetworkV1) ClusterExternalNetworks() v1.ClusterExternalNetworkInterface {
	return &FakeClusterExternalNetworks{c}
}

func (c *FakeNetworkV1) ClusterNetworkRuleSetPolicies() v1.ClusterNetworkRuleSetPolicyInterface {
	return &FakeClusterNetworkRuleSetPolicies{c}
}

func (c *FakeNetworkV1) ClusterProcessingUnits() v1.ClusterProcessingUnitInterface {
	return &FakeClusterProcessingUnits{c}
}

func (c *FakeNetworkV1) ExternalNetworks(namespace string) v1.ExternalNetworkInterface {
	return &FakeExternalNetworks{c, namespace}
}

func (c *FakeNetworkV1) NetworkRuleSetPolicies(namespace string) v1.NetworkRuleSetPolicyInterface {
	return &FakeNetworkRuleSetPolicies{c, namespace}
}

func (c *FakeNetworkV1) ProcessingUnits(namespace string) v1.ProcessingUnitInterface {
	return &FakeProcessingUnits{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeNetworkV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
