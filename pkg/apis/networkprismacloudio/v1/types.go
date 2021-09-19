

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"go.aporeto.io/gaia"
)


// ClusterNetworkRuleSetPolicy is the Schema for the NetworkRuleSetPolicys API
// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ClusterNetworkRuleSetPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec   gaia.SparseNetworkRuleSetPolicy   `json:"spec,omitempty"`
}
// ClusterNetworkRuleSetPolicyList contains a list of NetworkRuleSetPolicy
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ClusterNetworkRuleSetPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterNetworkRuleSetPolicy `json:"items"`
}


// ClusterExternalNetwork is the Schema for the ExternalNetworks API
// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ClusterExternalNetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec   gaia.SparseExternalNetwork   `json:"spec,omitempty"`
}
// ClusterExternalNetworkList contains a list of ExternalNetwork
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ClusterExternalNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterExternalNetwork `json:"items"`
}


// NetworkRuleSetPolicy is the Schema for the NetworkRuleSetPolicys API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type NetworkRuleSetPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec   gaia.SparseNetworkRuleSetPolicy   `json:"spec,omitempty"`
}
// NetworkRuleSetPolicyList contains a list of NetworkRuleSetPolicy
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type NetworkRuleSetPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkRuleSetPolicy `json:"items"`
}


// ExternalNetwork is the Schema for the ExternalNetworks API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ExternalNetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec   gaia.SparseExternalNetwork   `json:"spec,omitempty"`
}
// ExternalNetworkList contains a list of ExternalNetwork
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ExternalNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExternalNetwork `json:"items"`
}


// ProcessingUnit is the Schema for the ProcessingUnits API
// +genclient
// +genclient:readonly
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ProcessingUnit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec   gaia.SparseProcessingUnit   `json:"spec,omitempty"`
}
// ProcessingUnitList contains a list of ProcessingUnit
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ProcessingUnitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProcessingUnit `json:"items"`
}


// PUTrafficAction is the Schema for the PUTrafficActions API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type PUTrafficAction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec   gaia.SparsePUTrafficAction   `json:"spec,omitempty"`
}
// PUTrafficActionList contains a list of PUTrafficAction
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type PUTrafficActionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PUTrafficAction `json:"items"`
}


// ClusterProcessingUnit is the Schema for the ProcessingUnits API
// +genclient
// +genclient:readonly
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ClusterProcessingUnit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec   gaia.SparseProcessingUnit   `json:"spec,omitempty"`
}
// ClusterProcessingUnitList contains a list of ProcessingUnit
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ClusterProcessingUnitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterProcessingUnit `json:"items"`
}


// ClusterPUTrafficAction is the Schema for the PUTrafficActions API
// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ClusterPUTrafficAction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec   gaia.SparsePUTrafficAction   `json:"spec,omitempty"`
}
// ClusterPUTrafficActionList contains a list of PUTrafficAction
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ClusterPUTrafficActionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterPUTrafficAction `json:"items"`
}


// ClusterEnforcer is the Schema for the Enforcers API
// +genclient
// +genclient:readonly
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ClusterEnforcer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec   gaia.SparseEnforcer   `json:"spec,omitempty"`
}
// ClusterEnforcerList contains a list of Enforcer
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ClusterEnforcerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterEnforcer `json:"items"`
}


// ClusterEnforcerProfile is the Schema for the EnforcerProfiles API
// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ClusterEnforcerProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec   gaia.SparseEnforcerProfile   `json:"spec,omitempty"`
}
// ClusterEnforcerProfileList contains a list of EnforcerProfile
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ClusterEnforcerProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterEnforcerProfile `json:"items"`
}


// add types above this line
