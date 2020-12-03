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

// SiteWhereTenantEngineSpec defines the desired state of SiteWhereTenantEngine
type SiteWhereTenantEngineSpec struct {
	// Configuration is the configuration for the tenant
	Configuration *runtime.RawExtension `json:"configuration,omitempty"`
}

// SiteWhereTenantEngineStatus defines the observed state of SiteWhereTenantEngine
type SiteWhereTenantEngineStatus struct {
	// BootstrapState is the bootstrap state of the tenant engine
	BootstrapState BootstrapState `json:"bootstrapState,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=tenantengines,scope=Namespaced,singular=tenantengine,shortName=swte,categories=sitewhere-io;core-sitewhere-io

// SiteWhereTenantEngine is the Schema for the sitewheretenantengines API
type SiteWhereTenantEngine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SiteWhereTenantEngineSpec   `json:"spec,omitempty"`
	Status SiteWhereTenantEngineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SiteWhereTenantEngineList contains a list of SiteWhereTenantEngine
type SiteWhereTenantEngineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SiteWhereTenantEngine `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SiteWhereTenantEngine{}, &SiteWhereTenantEngineList{})
}
