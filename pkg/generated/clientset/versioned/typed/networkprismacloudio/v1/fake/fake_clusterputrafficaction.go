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
	"context"

	networkprismacloudiov1 "go.aporeto.io/gaia-k8s/pkg/apis/networkprismacloudio/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterPUTrafficActions implements ClusterPUTrafficActionInterface
type FakeClusterPUTrafficActions struct {
	Fake *FakeNetworkV1
}

var clusterputrafficactionsResource = schema.GroupVersionResource{Group: "network.prismacloud.io", Version: "v1", Resource: "clusterputrafficactions"}

var clusterputrafficactionsKind = schema.GroupVersionKind{Group: "network.prismacloud.io", Version: "v1", Kind: "ClusterPUTrafficAction"}

// Get takes name of the clusterPUTrafficAction, and returns the corresponding clusterPUTrafficAction object, and an error if there is any.
func (c *FakeClusterPUTrafficActions) Get(ctx context.Context, name string, options v1.GetOptions) (result *networkprismacloudiov1.ClusterPUTrafficAction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterputrafficactionsResource, name), &networkprismacloudiov1.ClusterPUTrafficAction{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkprismacloudiov1.ClusterPUTrafficAction), err
}

// List takes label and field selectors, and returns the list of ClusterPUTrafficActions that match those selectors.
func (c *FakeClusterPUTrafficActions) List(ctx context.Context, opts v1.ListOptions) (result *networkprismacloudiov1.ClusterPUTrafficActionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterputrafficactionsResource, clusterputrafficactionsKind, opts), &networkprismacloudiov1.ClusterPUTrafficActionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &networkprismacloudiov1.ClusterPUTrafficActionList{ListMeta: obj.(*networkprismacloudiov1.ClusterPUTrafficActionList).ListMeta}
	for _, item := range obj.(*networkprismacloudiov1.ClusterPUTrafficActionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterPUTrafficActions.
func (c *FakeClusterPUTrafficActions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterputrafficactionsResource, opts))
}

// Create takes the representation of a clusterPUTrafficAction and creates it.  Returns the server's representation of the clusterPUTrafficAction, and an error, if there is any.
func (c *FakeClusterPUTrafficActions) Create(ctx context.Context, clusterPUTrafficAction *networkprismacloudiov1.ClusterPUTrafficAction, opts v1.CreateOptions) (result *networkprismacloudiov1.ClusterPUTrafficAction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterputrafficactionsResource, clusterPUTrafficAction), &networkprismacloudiov1.ClusterPUTrafficAction{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkprismacloudiov1.ClusterPUTrafficAction), err
}

// Update takes the representation of a clusterPUTrafficAction and updates it. Returns the server's representation of the clusterPUTrafficAction, and an error, if there is any.
func (c *FakeClusterPUTrafficActions) Update(ctx context.Context, clusterPUTrafficAction *networkprismacloudiov1.ClusterPUTrafficAction, opts v1.UpdateOptions) (result *networkprismacloudiov1.ClusterPUTrafficAction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterputrafficactionsResource, clusterPUTrafficAction), &networkprismacloudiov1.ClusterPUTrafficAction{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkprismacloudiov1.ClusterPUTrafficAction), err
}

// Delete takes name of the clusterPUTrafficAction and deletes it. Returns an error if one occurs.
func (c *FakeClusterPUTrafficActions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(clusterputrafficactionsResource, name), &networkprismacloudiov1.ClusterPUTrafficAction{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterPUTrafficActions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterputrafficactionsResource, listOpts)

	_, err := c.Fake.Invokes(action, &networkprismacloudiov1.ClusterPUTrafficActionList{})
	return err
}

// Patch applies the patch and returns the patched clusterPUTrafficAction.
func (c *FakeClusterPUTrafficActions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *networkprismacloudiov1.ClusterPUTrafficAction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterputrafficactionsResource, name, pt, data, subresources...), &networkprismacloudiov1.ClusterPUTrafficAction{})
	if obj == nil {
		return nil, err
	}
	return obj.(*networkprismacloudiov1.ClusterPUTrafficAction), err
}