

// Package v1 contains API Schema definitions for the networkprismacloudio v1 API group
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// GroupName holds the API group name.
const GroupName = "network.prismacloud.io"

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{Group: GroupName , Version: "v1"}

var (
	// SchemeBuilder registers the known types
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	// AddToScheme applies all the stored functions to the scheme
	AddToScheme   = SchemeBuilder.AddToScheme
)

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// Adds the list of known types to the given scheme.
// TODO: we need to add all the API types in the schema over here.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
&ClusterNetworkRuleSetPolicy{},
&ClusterNetworkRuleSetPolicyList{},
&ClusterExternalNetwork{},
&ClusterExternalNetworkList{},
&NetworkRuleSetPolicy{},
&NetworkRuleSetPolicyList{},
&ExternalNetwork{},
&ExternalNetworkList{},
&ProcessingUnit{},
&ProcessingUnitList{},
&PUTrafficAction{},
&PUTrafficActionList{},
&ClusterProcessingUnit{},
&ClusterProcessingUnitList{},
&ClusterPUTrafficAction{},
&ClusterPUTrafficActionList{},
&ClusterEnforcer{},
&ClusterEnforcerList{},
&ClusterEnforcerProfile{},
&ClusterEnforcerProfileList{},
// add types above this line
	)
	
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
