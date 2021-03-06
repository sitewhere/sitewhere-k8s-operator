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

// Package v1alpha4 contains API Schema definitions for the templates.sitewhere.io v1alpha4 API group
// +kubebuilder:object:generate=true
// +groupName=templates.sitewhere.io
package v1alpha4

import (
	"reflect"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// Package type metadata.
const (
	Group   = "templates.sitewhere.io"
	Version = "v1alpha4"
)

var (
	// GroupVersion is group version used to register these objects
	GroupVersion = schema.GroupVersion{Group: Group, Version: Version}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)

// InstanceConfigurationTemplate type metadata.
var (
	InstanceConfigurationTemplateKind             = reflect.TypeOf(InstanceConfigurationTemplate{}).Name()
	InstanceConfigurationTemplateGroupKind        = schema.GroupKind{Group: Group, Kind: InstanceConfigurationTemplateKind}.String()
	InstanceConfigurationTemplateKindAPIVersion   = InstanceConfigurationTemplateKind + "." + GroupVersion.String()
	InstanceConfigurationTemplateGroupVersionKind = GroupVersion.WithKind(InstanceConfigurationTemplateKind)
)

// TenantConfigurationTemplate type metadata.
var (
	TenantConfigurationTemplateKind             = reflect.TypeOf(TenantConfigurationTemplate{}).Name()
	TenantConfigurationTemplateGroupKind        = schema.GroupKind{Group: Group, Kind: TenantConfigurationTemplateKind}.String()
	TenantConfigurationTemplateKindAPIVersion   = TenantConfigurationTemplateKind + "." + GroupVersion.String()
	TenantConfigurationTemplateGroupVersionKind = GroupVersion.WithKind(TenantConfigurationTemplateKind)
)

// TenantEngineConfigurationTemplate type metadata.
var (
	TenantEngineConfigurationTemplateKind             = reflect.TypeOf(TenantEngineConfigurationTemplate{}).Name()
	TenantEngineConfigurationTemplateGroupKind        = schema.GroupKind{Group: Group, Kind: TenantEngineConfigurationTemplateKind}.String()
	TenantEngineConfigurationTemplateKindAPIVersion   = TenantEngineConfigurationTemplateKind + "." + GroupVersion.String()
	TenantEngineConfigurationTemplateGroupVersionKind = GroupVersion.WithKind(TenantEngineConfigurationTemplateKind)
)
