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
)

// InstanceDatasetTemplateSpec defines the desired state of InstanceDatasetTemplate
type InstanceDatasetTemplateSpec struct {
	// Datasets is map of datasets
	Datasets map[string]string `json:"datasets,omitempty"`
}

// InstanceDatasetTemplateStatus defines the observed state of InstanceDatasetTemplate
type InstanceDatasetTemplateStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=instancedatasets,scope=Cluster,singular=instancedataset,shortName=idt,categories=sitewhere-io;core-sitewhere-io

// InstanceDatasetTemplate is the Schema for the instancedatasettemplates API
type InstanceDatasetTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InstanceDatasetTemplateSpec   `json:"spec,omitempty"`
	Status InstanceDatasetTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceDatasetTemplateList contains a list of InstanceDatasetTemplate
type InstanceDatasetTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceDatasetTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&InstanceDatasetTemplate{}, &InstanceDatasetTemplateList{})
}
