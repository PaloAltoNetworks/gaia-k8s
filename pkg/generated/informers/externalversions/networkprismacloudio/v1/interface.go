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
	internalinterfaces "go.aporeto.io/gaia-k8s/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ClusterEnforcers returns a ClusterEnforcerInformer.
	ClusterEnforcers() ClusterEnforcerInformer
	// ClusterEnforcerProfiles returns a ClusterEnforcerProfileInformer.
	ClusterEnforcerProfiles() ClusterEnforcerProfileInformer
	// ClusterExternalNetworks returns a ClusterExternalNetworkInformer.
	ClusterExternalNetworks() ClusterExternalNetworkInformer
	// ClusterNetworkRuleSetPolicies returns a ClusterNetworkRuleSetPolicyInformer.
	ClusterNetworkRuleSetPolicies() ClusterNetworkRuleSetPolicyInformer
	// ClusterProcessingUnits returns a ClusterProcessingUnitInformer.
	ClusterProcessingUnits() ClusterProcessingUnitInformer
	// ExternalNetworks returns a ExternalNetworkInformer.
	ExternalNetworks() ExternalNetworkInformer
	// NetworkRuleSetPolicies returns a NetworkRuleSetPolicyInformer.
	NetworkRuleSetPolicies() NetworkRuleSetPolicyInformer
	// ProcessingUnits returns a ProcessingUnitInformer.
	ProcessingUnits() ProcessingUnitInformer
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

// ClusterEnforcers returns a ClusterEnforcerInformer.
func (v *version) ClusterEnforcers() ClusterEnforcerInformer {
	return &clusterEnforcerInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterEnforcerProfiles returns a ClusterEnforcerProfileInformer.
func (v *version) ClusterEnforcerProfiles() ClusterEnforcerProfileInformer {
	return &clusterEnforcerProfileInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterExternalNetworks returns a ClusterExternalNetworkInformer.
func (v *version) ClusterExternalNetworks() ClusterExternalNetworkInformer {
	return &clusterExternalNetworkInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterNetworkRuleSetPolicies returns a ClusterNetworkRuleSetPolicyInformer.
func (v *version) ClusterNetworkRuleSetPolicies() ClusterNetworkRuleSetPolicyInformer {
	return &clusterNetworkRuleSetPolicyInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterProcessingUnits returns a ClusterProcessingUnitInformer.
func (v *version) ClusterProcessingUnits() ClusterProcessingUnitInformer {
	return &clusterProcessingUnitInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ExternalNetworks returns a ExternalNetworkInformer.
func (v *version) ExternalNetworks() ExternalNetworkInformer {
	return &externalNetworkInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// NetworkRuleSetPolicies returns a NetworkRuleSetPolicyInformer.
func (v *version) NetworkRuleSetPolicies() NetworkRuleSetPolicyInformer {
	return &networkRuleSetPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ProcessingUnits returns a ProcessingUnitInformer.
func (v *version) ProcessingUnits() ProcessingUnitInformer {
	return &processingUnitInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
