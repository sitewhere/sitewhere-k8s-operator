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

// TenantEngineConfigurationTemplateSpec defines the desired state of TenantEngineConfigurationTemplate
type TenantEngineConfigurationTemplateSpec struct {
	// Configuration is the configuration for the tenant
	Configuration *runtime.RawExtension `json:"configuration,omitempty"`
}

// TenantEngineConfigurationTemplateStatus defines the observed state of TenantEngineConfigurationTemplate
type TenantEngineConfigurationTemplateStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=tenantengineconfigurations,scope=Cluster,singular=tenantengineconfiguration,shortName=tect,categories=sitewhere-io;core-sitewhere-io

// TenantEngineConfigurationTemplate is the Schema for the tenantengineconfigurationtemplates API
type TenantEngineConfigurationTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TenantEngineConfigurationTemplateSpec   `json:"spec,omitempty"`
	Status TenantEngineConfigurationTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TenantEngineConfigurationTemplateList contains a list of TenantEngineConfigurationTemplate
type TenantEngineConfigurationTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TenantEngineConfigurationTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TenantEngineConfigurationTemplate{}, &TenantEngineConfigurationTemplateList{})
}
