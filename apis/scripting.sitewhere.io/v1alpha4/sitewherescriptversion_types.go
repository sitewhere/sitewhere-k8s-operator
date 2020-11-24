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

// SiteWhereScriptVersionSpec defines the desired state of SiteWhereScriptVersion
type SiteWhereScriptVersionSpec struct {
	// Comment is the comment on this version
	Comment string `json:"comment,omitempty"`

	// Content is the comment on this version
	Content string `json:"content,omitempty"`
}

// SiteWhereScriptVersionStatus defines the observed state of SiteWhereScriptVersion
type SiteWhereScriptVersionStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=scriptversions,scope=Namespaced,singular=scriptversion,shortName=swscrv,categories=sitewhere-io;core-sitewhere-io

// SiteWhereScriptVersion is the Schema for the sitewherescriptversions API
type SiteWhereScriptVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SiteWhereScriptVersionSpec   `json:"spec,omitempty"`
	Status SiteWhereScriptVersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SiteWhereScriptVersionList contains a list of SiteWhereScriptVersion
type SiteWhereScriptVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SiteWhereScriptVersion `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SiteWhereScriptVersion{}, &SiteWhereScriptVersionList{})
}
