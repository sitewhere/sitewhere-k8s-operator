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

// Package v1alpha4 contains API Schema definitions for the sitewhere.io v1alpha4 API group
// +kubebuilder:object:generate=true
// +groupName=sitewhere.io
package v1alpha4

import (
	"reflect"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// Package type metadata.
const (
	Group   = "sitewhere.io"
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

// SiteWhereInstance type metadata.
var (
	SiteWhereInstanceKind             = reflect.TypeOf(SiteWhereInstance{}).Name()
	SiteWhereInstanceGroupKind        = schema.GroupKind{Group: Group, Kind: SiteWhereInstanceKind}.String()
	SiteWhereInstanceKindAPIVersion   = SiteWhereInstanceKind + "." + GroupVersion.String()
	SiteWhereInstanceGroupVersionKind = GroupVersion.WithKind(SiteWhereInstanceKind)
)

// SiteWhereMicroservice type metadata.
var (
	SiteWhereMicroserviceKind             = reflect.TypeOf(SiteWhereMicroservice{}).Name()
	SiteWhereMicroserviceGroupKind        = schema.GroupKind{Group: Group, Kind: SiteWhereMicroserviceKind}.String()
	SiteWhereMicroserviceKindAPIVersion   = SiteWhereMicroserviceKind + "." + GroupVersion.String()
	SiteWhereMicroserviceGroupVersionKind = GroupVersion.WithKind(SiteWhereMicroserviceKind)
)

// SiteWhereTenant type metadata.
var (
	SiteWhereTenantKind             = reflect.TypeOf(SiteWhereTenant{}).Name()
	SiteWhereTenantGroupKind        = schema.GroupKind{Group: Group, Kind: SiteWhereTenantKind}.String()
	SiteWhereTenantKindAPIVersion   = SiteWhereTenantKind + "." + GroupVersion.String()
	SiteWhereTenantGroupVersionKind = GroupVersion.WithKind(SiteWhereTenantKind)
)

// SiteWhereTenantEngine type metadata.
var (
	SiteWhereTenantEngineKind             = reflect.TypeOf(SiteWhereTenantEngine{}).Name()
	SiteWhereTenantEngineGroupKind        = schema.GroupKind{Group: Group, Kind: SiteWhereTenantEngineKind}.String()
	SiteWhereTenantEngineKindAPIVersion   = SiteWhereTenantEngineKind + "." + GroupVersion.String()
	SiteWhereTenantEngineGroupVersionKind = GroupVersion.WithKind(SiteWhereTenantEngineKind)
)
