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

// SiteWhereScriptTemplateSpec defines the desired state of SiteWhereScriptTemplate
type SiteWhereScriptTemplateSpec struct {
	// TODO: check if metadata.name can be used
	// Name is the name of the dataset
	Name string `json:"name,omitempty"`

	// Description is the description of the tenant dataset template
	Description string `json:"description,omitempty"`

	// InterpreterType interpreter type
	InterpreterType string `json:"interpreterType,omitempty"`

	// Script script to execute
	Script string `json:"script,omitempty"`
}

// SiteWhereScriptTemplateStatus defines the observed state of SiteWhereScriptTemplate
type SiteWhereScriptTemplateStatus struct {
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=scripttemplates,scope=Cluster,singular=scripttemplate,shortName=swscrt,categories=sitewhere-io;core-sitewhere-io

// SiteWhereScriptTemplate is the Schema for the scripttemplates API
type SiteWhereScriptTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SiteWhereScriptTemplateSpec   `json:"spec,omitempty"`
	Status SiteWhereScriptTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SiteWhereScriptTemplateList contains a list of SiteWhereScriptTemplate
type SiteWhereScriptTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SiteWhereScriptTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SiteWhereScriptTemplate{}, &SiteWhereScriptTemplateList{})
}
