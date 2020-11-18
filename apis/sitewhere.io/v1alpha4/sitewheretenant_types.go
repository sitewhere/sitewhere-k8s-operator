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

// TenantBrandingSpecification defines the branding of a tenant
type TenantBrandingSpecification struct {
	// BackgroundColor is the background color of the branding
	BackgroundColor string `json:"backgroundColor,omitempty"`
	// ForegroundColor is the foreground color of the branding
	ForegroundColor string `json:"foregroundColor,omitempty"`
	// BorderColor is the border color of the branding
	BorderColor string `json:"borderColor,omitempty"`
	// Icon is the icon of the branding
	Icon string `json:"icon,omitempty"`
	// ImageURL is the image URL of the branding
	ImageURL string `json:"imageUrl,omitempty"`
}

// SiteWhereTenantSpec defines the desired state of SiteWhereTenant
type SiteWhereTenantSpec struct {
	// Name is the name of the tenant (ObjectMeta.name)
	Name string `json:"name,omitempty"`
	// AuthenticationToken is the token used for authenticating the tenant
	AuthenticationToken string `json:"authenticationToken,omitempty"`
	// Authorized are the IDs of the users that are authorized to use the tenant
	AuthorizedUserIds []string `json:"authorizedUserIds,omitempty"`
	// ConfigurationTemplate is the configuration template used for the tenant
	ConfigurationTemplate string `json:"configurationTemplate,omitempty"`
	// DatasetTemplate is the dataset template used for the tenant
	DatasetTemplate string `json:"datasetTemplate,omitempty"`
	// Branding is the branding information for the tenant
	Branding *TenantBrandingSpecification `json:"branding,omitempty"`
	// Metadata is the metadata of the tenant
	Metadata map[string]string `json:"metadata,omitempty"`
}

// SiteWhereTenantStatus defines the observed state of SiteWhereTenant
type SiteWhereTenantStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=tenants,scope=Namespaced,singular=tenant,shortName=swt,categories=sitewhere-io;core-sitewhere-io

// SiteWhereTenant is the Schema for the sitewheretenants API
type SiteWhereTenant struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SiteWhereTenantSpec   `json:"spec,omitempty"`
	Status SiteWhereTenantStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SiteWhereTenantList contains a list of SiteWhereTenant
type SiteWhereTenantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SiteWhereTenant `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SiteWhereTenant{}, &SiteWhereTenantList{})
}
