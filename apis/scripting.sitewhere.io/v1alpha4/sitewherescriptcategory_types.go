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

// SiteWhereScriptCategorySpec defines the desired state of SiteWhereScriptCategory
type SiteWhereScriptCategorySpec struct {
	// TODO: check if metadata.name can be used
	// Name is the name of the dataset
	Name string `json:"name,omitempty"`

	// Description is the description of the tenant dataset template
	Description string `json:"description,omitempty"`
}

// SiteWhereScriptCategoryStatus defines the observed state of SiteWhereScriptCategory
type SiteWhereScriptCategoryStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=scriptcategories,scope=Cluster,singular=scriptcategory,shortName=swscrc,categories=sitewhere-io;core-sitewhere-io

// SiteWhereScriptCategory is the Schema for the scriptcategories API
type SiteWhereScriptCategory struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SiteWhereScriptCategorySpec   `json:"spec,omitempty"`
	Status SiteWhereScriptCategoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SiteWhereScriptCategoryList contains a list of SiteWhereScriptCategory
type SiteWhereScriptCategoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SiteWhereScriptCategory `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SiteWhereScriptCategory{}, &SiteWhereScriptCategoryList{})
}
