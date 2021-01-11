/*
Copyright (c) SiteWhere, LLC. All rights reserved. http://www.sitewhere.com

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

package v1alpha4

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// InstanceConfigurationTemplateSpec defines the desired state of InstanceConfigurationTemplate
type InstanceConfigurationTemplateSpec struct {
	// +nullable
	// +kubebuilder:pruning:PreserveUnknownFields

	// Configuration is the configuration for the tenant
	Configuration *runtime.RawExtension `json:"configuration,omitempty"`
}

// InstanceConfigurationTemplateStatus defines the observed state of InstanceConfigurationTemplate
type InstanceConfigurationTemplateStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=instanceconfigurations,scope=Cluster,singular=instanceconfiguration,shortName=ict,categories=sitewhere-io;core-sitewhere-io

// InstanceConfigurationTemplate is the Schema for the instanceconfigurations API
type InstanceConfigurationTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InstanceConfigurationTemplateSpec   `json:"spec,omitempty"`
	Status InstanceConfigurationTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceConfigurationTemplateList contains a list of InstanceConfigurationTemplate
type InstanceConfigurationTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceConfigurationTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&InstanceConfigurationTemplate{}, &InstanceConfigurationTemplateList{})
}
