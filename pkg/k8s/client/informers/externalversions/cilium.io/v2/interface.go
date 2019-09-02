// Copyright 2017-2019 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v2

import (
	internalinterfaces "github.com/cilium/cilium/pkg/k8s/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// CiliumEndpoints returns a CiliumEndpointInformer.
	CiliumEndpoints() CiliumEndpointInformer
	// CiliumIdentities returns a CiliumIdentityInformer.
	CiliumIdentities() CiliumIdentityInformer
	// CiliumNetworkPolicies returns a CiliumNetworkPolicyInformer.
	CiliumNetworkPolicies() CiliumNetworkPolicyInformer
	// CiliumNodes returns a CiliumNodeInformer.
	CiliumNodes() CiliumNodeInformer
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

// CiliumEndpoints returns a CiliumEndpointInformer.
func (v *version) CiliumEndpoints() CiliumEndpointInformer {
	return &ciliumEndpointInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// CiliumIdentities returns a CiliumIdentityInformer.
func (v *version) CiliumIdentities() CiliumIdentityInformer {
	return &ciliumIdentityInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// CiliumNetworkPolicies returns a CiliumNetworkPolicyInformer.
func (v *version) CiliumNetworkPolicies() CiliumNetworkPolicyInformer {
	return &ciliumNetworkPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// CiliumNodes returns a CiliumNodeInformer.
func (v *version) CiliumNodes() CiliumNodeInformer {
	return &ciliumNodeInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
