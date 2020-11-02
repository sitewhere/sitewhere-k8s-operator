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

// SiteWhereMicroserviceSpec defines the desired state of SiteWhereMicroservice
type SiteWhereMicroserviceSpec struct {
	// Functional Area
	FunctionalArea string `json:"functionalArea,omitempty"`

	// Name is the name displayed for microservice
	Name string `json:"name,omitempty"`

	// Description is the description of the microservice
	Description string `json:"description,omitempty"`

	// Icon displayed for microservice
	Icon string `json:"icon,omitempty"`

	// Replicas is the number of desired replicas of the microservice
	Replicas int32 `json:"replicas,omitempty"`
}

// SiteWhereMicroserviceStatus defines the observed state of SiteWhereMicroservice
type SiteWhereMicroserviceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=microservices,scope=Namespaced,singular=microservice,shortName=swm,categories=sitewhere-io;core-sitewhere-io

// SiteWhereMicroservice is the Schema for the sitewheremicroservices API
type SiteWhereMicroservice struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SiteWhereMicroserviceSpec   `json:"spec,omitempty"`
	Status SiteWhereMicroserviceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SiteWhereMicroserviceList contains a list of SiteWhereMicroservice
type SiteWhereMicroserviceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SiteWhereMicroservice `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SiteWhereMicroservice{}, &SiteWhereMicroserviceList{})
}
