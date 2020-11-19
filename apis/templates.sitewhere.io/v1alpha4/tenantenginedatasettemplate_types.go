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

// TenantEngineDatasetTemplateSpec defines the desired state of TenantEngineDatasetTemplate
type TenantEngineDatasetTemplateSpec struct {
	// +nullable

	// Configuration is the configuration for the tenant
	Configuration *runtime.RawExtension `json:"configuration,omitempty"`
}

// TenantEngineDatasetTemplateStatus defines the observed state of TenantEngineDatasetTemplate
type TenantEngineDatasetTemplateStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=tenantenginedatasets,scope=Cluster,singular=tenantenginedataset,shortName=tedt,categories=sitewhere-io;core-sitewhere-io

// TenantEngineDatasetTemplate is the Schema for the tenantenginedatasettemplates API
type TenantEngineDatasetTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TenantEngineDatasetTemplateSpec   `json:"spec,omitempty"`
	Status TenantEngineDatasetTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TenantEngineDatasetTemplateList contains a list of TenantEngineDatasetTemplate
type TenantEngineDatasetTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TenantEngineDatasetTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TenantEngineDatasetTemplate{}, &TenantEngineDatasetTemplateList{})
}
