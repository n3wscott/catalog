/*
Copyright 2019 The Knative Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"knative.dev/pkg/apis"
	duckv1beta1 "knative.dev/pkg/apis/duck/v1beta1"
	"knative.dev/pkg/kmeta"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Broker is a Knative abstraction that encapsulates the interface by which Knative
// components express a desire to have a particular image cached.
type Broker struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec holds the desired state of the Broker (from the client).
	// +optional
	Spec BrokerSpec `json:"spec,omitempty"`

	// Status communicates the observed state of the Broker (from the controller).
	// +optional
	Status BrokerStatus `json:"status,omitempty"`
}

// Check that Broker can be validated and defaulted.
var _ apis.Validatable = (*Broker)(nil)
var _ apis.Defaultable = (*Broker)(nil)
var _ kmeta.OwnerRefable = (*Broker)(nil)

// BrokerSpec holds the desired state of the Broker (from the client).
type BrokerSpec struct {
	// URL is the address used to communicate with the Open Service Broker.
	URL *apis.URL `json:"url,omitempty"`

	// ServiceName holds the name of the Kubernetes Service to expose as an "addressable".
	ServiceName string `json:"serviceName"`
}

const (
	// AddressableServiceConditionReady is set when the revision is starting to materialize
	// runtime resources, and becomes true when those resources are ready.
	AddressableServiceConditionReady = apis.ConditionReady
)

// BrokerStatus communicates the observed state of the Broker (from the controller).
type BrokerStatus struct {
	duckv1beta1.Status `json:",inline"`

	// LastSyncTime is the last time the service catalog has been listed.
	// +optional
	LastSyncTime apis.VolatileTime `json:"lastSyncTime,omitempty" description:"Last time the service catalog has been listed."`

	// Address holds the information needed to connect this Addressable up to receive events.
	// +optional
	Address *duckv1beta1.Addressable `json:"address,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AddressableServiceList is a list of Broker resources
type AddressableServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Broker `json:"items"`
}
