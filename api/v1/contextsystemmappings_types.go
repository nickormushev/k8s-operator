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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ContextSystemMappingsSpec defines the desired state of ContextSystemMappings
type ContextSystemMappingsSpec struct {
	ID        int `json:"id"`
	ContextID int `json:"context_id"`
	SystemID  int `json:"system_id"`
}

// ContextSystemMappingsStatus defines the observed state of ContextSystemMappings
type ContextSystemMappingsStatus struct {
	SecretStatus string `json:"secret_status"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ContextSystemMappings is the Schema for the contextsystemmappings API
type ContextSystemMappings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContextSystemMappingsSpec   `json:"spec,omitempty"`
	Status ContextSystemMappingsStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ContextSystemMappingsList contains a list of ContextSystemMappings
type ContextSystemMappingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContextSystemMappings `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ContextSystemMappings{}, &ContextSystemMappingsList{})
}
