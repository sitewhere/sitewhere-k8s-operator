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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// UserManagementConfiguration is the configuration for User Management
type UserManagementConfiguration struct {
	// SyncopeHost is the hostname of Syncope API
	SyncopeHost string `json:"syncopeHost,omitempty"`

	// SyncopePort is the port of Syncope API
	SyncopePort int `json:"syncopePort,omitempty"`

	// JWTExpirationInMinutes is the expiration in minutes of the JWT.
	JWTExpirationInMinutes int `json:"jwtExpirationInMinutes,omitempty"`
}

// MicroservicePodSpecification is the specificacion of the microservice pod
type MicroservicePodSpecification struct {
	// Annotations is an unstructured key value map stored with a resource that may be
	// set by external tools to store and retrieve arbitrary metadata. They are not
	// queryable and should be preserved when modifying objects.
	// More info: http://kubernetes.io/docs/user-guide/annotations
	// +optional
	Annotations map[string]string `json:"annotations,omitempty" protobuf:"bytes,12,rep,name=annotations"`

	// Name must be unique within a namespace. Is required when creating resources, although
	// some resources may allow a client to request the generation of an appropriate name
	// automatically. Name is primarily intended for creation idempotence and configuration
	// definition.
	// Cannot be updated.
	// More info: http://kubernetes.io/docs/user-guide/identifiers#names
	// +optional
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`

	// DockerSpec is the Docker specification of the microservice.
	// If this value is set a the microservice level, it will override
	// the value set at the instance level.
	// +optional
	DockerSpec *DockerSpec `json:"dockerSpec,omitempty"`

	// Image pull policy.
	// One of Always, Never, IfNotPresent.
	// Defaults to Always if :latest tag is specified, or IfNotPresent otherwise.
	// Cannot be updated.
	// More info: https://kubernetes.io/docs/concepts/containers/images#updating-images
	// +optional
	ImagePullPolicy corev1.PullPolicy `json:"imagePullPolicy,omitempty" protobuf:"bytes,14,opt,name=imagePullPolicy,casttype=PullPolicy"`

	// List of ports to expose from the container. Exposing a port here gives
	// the system additional information about the network connections a
	// container uses, but is primarily informational. Not specifying a port here
	// DOES NOT prevent that port from being exposed. Any port which is
	// listening on the default "0.0.0.0" address inside a container will be
	// accessible from the network.
	// Cannot be updated.
	// +optional
	// +patchMergeKey=containerPort
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=containerPort
	// +listMapKey=protocol
	Ports []corev1.ContainerPort `json:"ports,omitempty" patchStrategy:"merge" patchMergeKey:"containerPort" protobuf:"bytes,6,rep,name=ports"`

	// List of environment variables to set in the container.
	// Cannot be updated.
	// +optional
	// +patchMergeKey=name
	// +patchStrategy=merge
	Env []corev1.EnvVar `json:"env,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,7,rep,name=env"`

	// Compute Resources required by this container.
	// Cannot be updated.
	// More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
	// +optional
	Resources *corev1.ResourceRequirements `json:"resources,omitempty" protobuf:"bytes,8,opt,name=resources"`

	// Periodic probe of container liveness.
	// Container will be restarted if the probe fails.
	// Cannot be updated.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// +optional
	LivenessProbe *corev1.Probe `json:"livenessProbe,omitempty" protobuf:"bytes,10,opt,name=livenessProbe"`
	// Periodic probe of container service readiness.
	// Container will be removed from service endpoints if the probe fails.
	// Cannot be updated.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// +optional
	ReadinessProbe *corev1.Probe `json:"readinessProbe,omitempty" protobuf:"bytes,11,opt,name=readinessProbe"`
}

// MicroserviceServiceSpecification is the service specificacion of the microservice
type MicroserviceServiceSpecification struct {
	// The list of ports that are exposed by this service.
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	// +patchMergeKey=port
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=port
	// +listMapKey=protocol
	Ports []corev1.ServicePort `json:"ports,omitempty" patchStrategy:"merge" patchMergeKey:"port" protobuf:"bytes,1,rep,name=ports"`

	// type determines how the Service is exposed. Defaults to ClusterIP. Valid
	// options are ExternalName, ClusterIP, NodePort, and LoadBalancer.
	// "ExternalName" maps to the specified externalName.
	// "ClusterIP" allocates a cluster-internal IP address for load-balancing to
	// endpoints. Endpoints are determined by the selector or if that is not
	// specified, by manual construction of an Endpoints object. If clusterIP is
	// "None", no virtual IP is allocated and the endpoints are published as a
	// set of endpoints rather than a stable IP.
	// "NodePort" builds on ClusterIP and allocates a port on every node which
	// routes to the clusterIP.
	// "LoadBalancer" builds on NodePort and creates an
	// external load-balancer (if supported in the current cloud) which routes
	// to the clusterIP.
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types
	// +optional
	Type *corev1.ServiceType `json:"type,omitempty" protobuf:"bytes,4,opt,name=type,casttype=ServiceType"`
}

// MicroserviceLoggingEntry is the logging level
type MicroserviceLoggingEntry struct {
	Logger string `json:"logger,omitempty"`

	Level string `json:"level,omitempty"`
}

// MicroserviceDebugSpecification is the debug specificacion of the microservice
type MicroserviceDebugSpecification struct {
	Enabled bool `json:"enabled,omitempty"`

	JDWPPort int `json:"jdwpPort,omitempty"`

	JMXPort int `json:"jmxPort,omitempty"`
}

// MicroserviceLoggingSpecification is the logging specification of the microservice
type MicroserviceLoggingSpecification struct {
	Overrides []MicroserviceLoggingEntry `json:"overrides,omitempty"`
}

// InstanceMangementConfiguration is the configuration of instance management
type InstanceMangementConfiguration struct {
	UserManagementConfiguration *UserManagementConfiguration `json:"userManagement,omitempty"`
}

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

	// Multitenant indicates whether microservice has tenant engines
	Multitenant bool `json:"multitenant,omitempty"`

	// PodSpec is the microservice pod specificacion
	// +optional
	PodSpec *MicroservicePodSpecification `json:"podSpec,omitempty"`

	// ServiceSpec is the Service specification
	// +optional
	SerivceSpec *MicroserviceServiceSpecification `json:"serviceSpec,omitempty"`

	// Debug is the Debug specification
	// +optional
	Debug *MicroserviceDebugSpecification `json:"debug,omitempty"`

	// Logging is the Logging specificacion
	// +optional
	Logging *MicroserviceLoggingSpecification `json:"logging,omitempty"`

	// Configuration is the configuration of the microservice
	// +optional
	Configuration *runtime.RawExtension `json:"configuration,omitempty"`
}

// SiteWhereMicroserviceStatus defines the observed state of SiteWhereMicroservice
type SiteWhereMicroserviceStatus struct {
	// Deployment name of the deployment
	Deployment string `json:"deployment,omitempty"`
	// Services name of services
	Services []string `json:"services,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=microservices,scope=Namespaced,singular=microservice,shortName=swm,categories=sitewhere-io;core-sitewhere-io
// +kubebuilder:printcolumn:name="Area",type=string,JSONPath=`.spec.functionalArea`

// SiteWhereMicroservice is the Schema for the microservices API
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
