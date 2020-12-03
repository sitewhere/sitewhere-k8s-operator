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

// BootstrapState State values for bootstrapping a component.
type BootstrapState string

const (
	// NotBootstrapped Component not bootstrapped
	NotBootstrapped BootstrapState = "NotBootstrapped"

	// Bootstrapping Component in process of bootstrapping
	Bootstrapping = "Bootstrapping"

	// Bootstrapped Component bootstrapped
	Bootstrapped = "Bootstrapped"

	// BootstrapFailed Bootstrap attempted but failed
	BootstrapFailed = "BootstrapFailed"
)

const (
	defaultRegistry   string = "docker.io"
	defaultRepository string = "sitewhere"
	defaultTag        string = "3.0.0.beta3"
)

// DefaultDockerSpec is th default docker spec
var DefaultDockerSpec = &DockerSpec{
	Registry:   defaultRegistry,
	Repository: defaultRepository,
	Tag:        defaultTag,
}

// DockerSpec are the spec to access sitewhere docker images
type DockerSpec struct {
	// Registry is the Docker registry
	Registry string `json:"registry,omitempty"`
	// Repository is the Repository for docker images
	Repository string `json:"repository,omitempty"`
	// Tag is the tag for docker images
	Tag string `json:"tag,omitempty"`
}

// SiteWhereInstanceSpec defines the desired state of SiteWhereInstance
type SiteWhereInstanceSpec struct {
	// ConfigurationTemplate is the name of the configuration template of the instance
	ConfigurationTemplate string `json:"configurationTemplate,omitempty"`
	// DatasetTemplate is the name of the dataset template of the instance
	DatasetTemplate string `json:"datasetTemplate,omitempty"`
	// DockerSpec is the
	DockerSpec *DockerSpec `json:"dockerSpec,omitempty"`
	// Configuration is the configuration of the instance
	Configuration *runtime.RawExtension `json:"configuration,omitempty"`
}

// SiteWhereInstanceStatus defines the observed state of SiteWhereInstance
type SiteWhereInstanceStatus struct {
	// TenantManagementBootstrapState Bootstrap state of Tenant Management
	TenantManagementBootstrapState BootstrapState `json:"tenantManagementBootstrapState,omitempty"`
	// UserManagementBootstrapState Bootstrap state of User Management
	UserManagementBootstrapState BootstrapState `json:"userManagementBootstrapState,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=instances,scope=Cluster,singular=instance,shortName=swi,categories=sitewhere-io;core-sitewhere-io
// +kubebuilder:printcolumn:name="Config",type=string,JSONPath=`.spec.configurationTemplate`
// +kubebuilder:printcolumn:name="Dataset",type=string,JSONPath=`.spec.datasetTemplate`
// +kubebuilder:printcolumn:name="Tenant Mng",type=string,JSONPath=`.status.tenantManagementBootstrapState`
// +kubebuilder:printcolumn:name="User Mng",type=string,JSONPath=`.status.userManagementBootstrapState`

// SiteWhereInstance is the Schema for the Sitewhere Instances API
type SiteWhereInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SiteWhereInstanceSpec   `json:"spec,omitempty"`
	Status SiteWhereInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SiteWhereInstanceList contains a list of SiteWhereInstance
type SiteWhereInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SiteWhereInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SiteWhereInstance{}, &SiteWhereInstanceList{})
}
