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

// TenantConfigurationTemplateSpec defines the desired state of TenantConfigurationTemplate
type TenantConfigurationTemplateSpec struct {
	// Name is the name of the tenant configuration template
	Name string `json:"name,omitempty"`
	// Description is the name of the tenant configuration template
	Description string `json:"description,omitempty"`
	// TenantEngineTemplates is the name of the tenant configuration template
	TenantEngineTemplates map[string]string `json:"tenantEngineTemplates,omitempty"`
}

// TenantConfigurationTemplateStatus defines the observed state of TenantConfigurationTemplate
type TenantConfigurationTemplateStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=tenantconfigurations,scope=Cluster,singular=tenantconfiguration,shortName=tct,categories=sitewhere-io;core-sitewhere-io

// TenantConfigurationTemplate is the Schema for the tenantconfigurations API
type TenantConfigurationTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TenantConfigurationTemplateSpec   `json:"spec,omitempty"`
	Status TenantConfigurationTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TenantConfigurationTemplateList contains a list of TenantConfigurationTemplate
type TenantConfigurationTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TenantConfigurationTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TenantConfigurationTemplate{}, &TenantConfigurationTemplateList{})
}
