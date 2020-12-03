// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha4

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerSpec) DeepCopyInto(out *DockerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerSpec.
func (in *DockerSpec) DeepCopy() *DockerSpec {
	if in == nil {
		return nil
	}
	out := new(DockerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceMangementConfiguration) DeepCopyInto(out *InstanceMangementConfiguration) {
	*out = *in
	if in.UserManagementConfiguration != nil {
		in, out := &in.UserManagementConfiguration, &out.UserManagementConfiguration
		*out = new(UserManagementConfiguration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceMangementConfiguration.
func (in *InstanceMangementConfiguration) DeepCopy() *InstanceMangementConfiguration {
	if in == nil {
		return nil
	}
	out := new(InstanceMangementConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiteWhereInstance) DeepCopyInto(out *SiteWhereInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiteWhereInstance.
func (in *SiteWhereInstance) DeepCopy() *SiteWhereInstance {
	if in == nil {
		return nil
	}
	out := new(SiteWhereInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SiteWhereInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiteWhereInstanceList) DeepCopyInto(out *SiteWhereInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SiteWhereInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiteWhereInstanceList.
func (in *SiteWhereInstanceList) DeepCopy() *SiteWhereInstanceList {
	if in == nil {
		return nil
	}
	out := new(SiteWhereInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SiteWhereInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiteWhereInstanceSpec) DeepCopyInto(out *SiteWhereInstanceSpec) {
	*out = *in
	if in.DockerSpec != nil {
		in, out := &in.DockerSpec, &out.DockerSpec
		*out = new(DockerSpec)
		**out = **in
	}
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiteWhereInstanceSpec.
func (in *SiteWhereInstanceSpec) DeepCopy() *SiteWhereInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(SiteWhereInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiteWhereInstanceStatus) DeepCopyInto(out *SiteWhereInstanceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiteWhereInstanceStatus.
func (in *SiteWhereInstanceStatus) DeepCopy() *SiteWhereInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(SiteWhereInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiteWhereMicroservice) DeepCopyInto(out *SiteWhereMicroservice) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiteWhereMicroservice.
func (in *SiteWhereMicroservice) DeepCopy() *SiteWhereMicroservice {
	if in == nil {
		return nil
	}
	out := new(SiteWhereMicroservice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SiteWhereMicroservice) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiteWhereMicroserviceList) DeepCopyInto(out *SiteWhereMicroserviceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SiteWhereMicroservice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiteWhereMicroserviceList.
func (in *SiteWhereMicroserviceList) DeepCopy() *SiteWhereMicroserviceList {
	if in == nil {
		return nil
	}
	out := new(SiteWhereMicroserviceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SiteWhereMicroserviceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiteWhereMicroserviceSpec) DeepCopyInto(out *SiteWhereMicroserviceSpec) {
	*out = *in
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiteWhereMicroserviceSpec.
func (in *SiteWhereMicroserviceSpec) DeepCopy() *SiteWhereMicroserviceSpec {
	if in == nil {
		return nil
	}
	out := new(SiteWhereMicroserviceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiteWhereMicroserviceStatus) DeepCopyInto(out *SiteWhereMicroserviceStatus) {
	*out = *in
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiteWhereMicroserviceStatus.
func (in *SiteWhereMicroserviceStatus) DeepCopy() *SiteWhereMicroserviceStatus {
	if in == nil {
		return nil
	}
	out := new(SiteWhereMicroserviceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiteWhereTenant) DeepCopyInto(out *SiteWhereTenant) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiteWhereTenant.
func (in *SiteWhereTenant) DeepCopy() *SiteWhereTenant {
	if in == nil {
		return nil
	}
	out := new(SiteWhereTenant)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SiteWhereTenant) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiteWhereTenantEngine) DeepCopyInto(out *SiteWhereTenantEngine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiteWhereTenantEngine.
func (in *SiteWhereTenantEngine) DeepCopy() *SiteWhereTenantEngine {
	if in == nil {
		return nil
	}
	out := new(SiteWhereTenantEngine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SiteWhereTenantEngine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiteWhereTenantEngineList) DeepCopyInto(out *SiteWhereTenantEngineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SiteWhereTenantEngine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiteWhereTenantEngineList.
func (in *SiteWhereTenantEngineList) DeepCopy() *SiteWhereTenantEngineList {
	if in == nil {
		return nil
	}
	out := new(SiteWhereTenantEngineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SiteWhereTenantEngineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiteWhereTenantEngineSpec) DeepCopyInto(out *SiteWhereTenantEngineSpec) {
	*out = *in
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiteWhereTenantEngineSpec.
func (in *SiteWhereTenantEngineSpec) DeepCopy() *SiteWhereTenantEngineSpec {
	if in == nil {
		return nil
	}
	out := new(SiteWhereTenantEngineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiteWhereTenantEngineStatus) DeepCopyInto(out *SiteWhereTenantEngineStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiteWhereTenantEngineStatus.
func (in *SiteWhereTenantEngineStatus) DeepCopy() *SiteWhereTenantEngineStatus {
	if in == nil {
		return nil
	}
	out := new(SiteWhereTenantEngineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiteWhereTenantList) DeepCopyInto(out *SiteWhereTenantList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SiteWhereTenant, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiteWhereTenantList.
func (in *SiteWhereTenantList) DeepCopy() *SiteWhereTenantList {
	if in == nil {
		return nil
	}
	out := new(SiteWhereTenantList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SiteWhereTenantList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiteWhereTenantSpec) DeepCopyInto(out *SiteWhereTenantSpec) {
	*out = *in
	if in.AuthorizedUserIds != nil {
		in, out := &in.AuthorizedUserIds, &out.AuthorizedUserIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Branding != nil {
		in, out := &in.Branding, &out.Branding
		*out = new(TenantBrandingSpecification)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiteWhereTenantSpec.
func (in *SiteWhereTenantSpec) DeepCopy() *SiteWhereTenantSpec {
	if in == nil {
		return nil
	}
	out := new(SiteWhereTenantSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SiteWhereTenantStatus) DeepCopyInto(out *SiteWhereTenantStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SiteWhereTenantStatus.
func (in *SiteWhereTenantStatus) DeepCopy() *SiteWhereTenantStatus {
	if in == nil {
		return nil
	}
	out := new(SiteWhereTenantStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantBrandingSpecification) DeepCopyInto(out *TenantBrandingSpecification) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantBrandingSpecification.
func (in *TenantBrandingSpecification) DeepCopy() *TenantBrandingSpecification {
	if in == nil {
		return nil
	}
	out := new(TenantBrandingSpecification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserManagementConfiguration) DeepCopyInto(out *UserManagementConfiguration) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserManagementConfiguration.
func (in *UserManagementConfiguration) DeepCopy() *UserManagementConfiguration {
	if in == nil {
		return nil
	}
	out := new(UserManagementConfiguration)
	in.DeepCopyInto(out)
	return out
}
