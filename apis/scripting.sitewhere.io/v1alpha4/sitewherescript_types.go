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

// SiteWhereScriptSpec defines the desired state of SiteWhereScript
type SiteWhereScriptSpec struct {
	// ID is the ID of the Script
	ID string `json:"scriptId,omitempty"`

	// TODO: check if metadata.name can be used
	// Name is the name of the dataset
	Name string `json:"name,omitempty"`

	// Description is the description of the tenant dataset template
	Description string `json:"description,omitempty"`

	// InterpreterType interpreter type
	InterpreterType string `json:"interpreterType,omitempty"`

	// ActiveVersion is the active version of the script
	ActiveVersion string `json:"activeVersion,omitempty"`
}

// SiteWhereScriptStatus defines the observed state of SiteWhereScript
type SiteWhereScriptStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=scripts,scope=Namespaced,singular=script,shortName=swscr,categories=sitewhere-io;core-sitewhere-io

// SiteWhereScript is the Schema for the scripts API
type SiteWhereScript struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SiteWhereScriptSpec   `json:"spec,omitempty"`
	Status SiteWhereScriptStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SiteWhereScriptList contains a list of SiteWhereScript
type SiteWhereScriptList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SiteWhereScript `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SiteWhereScript{}, &SiteWhereScriptList{})
}
