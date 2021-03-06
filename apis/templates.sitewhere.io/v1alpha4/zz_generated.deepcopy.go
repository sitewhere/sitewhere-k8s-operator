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
func (in *InstanceConfigurationTemplate) DeepCopyInto(out *InstanceConfigurationTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceConfigurationTemplate.
func (in *InstanceConfigurationTemplate) DeepCopy() *InstanceConfigurationTemplate {
	if in == nil {
		return nil
	}
	out := new(InstanceConfigurationTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstanceConfigurationTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceConfigurationTemplateList) DeepCopyInto(out *InstanceConfigurationTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InstanceConfigurationTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceConfigurationTemplateList.
func (in *InstanceConfigurationTemplateList) DeepCopy() *InstanceConfigurationTemplateList {
	if in == nil {
		return nil
	}
	out := new(InstanceConfigurationTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstanceConfigurationTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceConfigurationTemplateSpec) DeepCopyInto(out *InstanceConfigurationTemplateSpec) {
	*out = *in
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceConfigurationTemplateSpec.
func (in *InstanceConfigurationTemplateSpec) DeepCopy() *InstanceConfigurationTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(InstanceConfigurationTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceConfigurationTemplateStatus) DeepCopyInto(out *InstanceConfigurationTemplateStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceConfigurationTemplateStatus.
func (in *InstanceConfigurationTemplateStatus) DeepCopy() *InstanceConfigurationTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(InstanceConfigurationTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceDatasetTemplate) DeepCopyInto(out *InstanceDatasetTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceDatasetTemplate.
func (in *InstanceDatasetTemplate) DeepCopy() *InstanceDatasetTemplate {
	if in == nil {
		return nil
	}
	out := new(InstanceDatasetTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstanceDatasetTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceDatasetTemplateList) DeepCopyInto(out *InstanceDatasetTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InstanceDatasetTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceDatasetTemplateList.
func (in *InstanceDatasetTemplateList) DeepCopy() *InstanceDatasetTemplateList {
	if in == nil {
		return nil
	}
	out := new(InstanceDatasetTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstanceDatasetTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceDatasetTemplateSpec) DeepCopyInto(out *InstanceDatasetTemplateSpec) {
	*out = *in
	if in.Datasets != nil {
		in, out := &in.Datasets, &out.Datasets
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceDatasetTemplateSpec.
func (in *InstanceDatasetTemplateSpec) DeepCopy() *InstanceDatasetTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(InstanceDatasetTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceDatasetTemplateStatus) DeepCopyInto(out *InstanceDatasetTemplateStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceDatasetTemplateStatus.
func (in *InstanceDatasetTemplateStatus) DeepCopy() *InstanceDatasetTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(InstanceDatasetTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantConfigurationTemplate) DeepCopyInto(out *TenantConfigurationTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantConfigurationTemplate.
func (in *TenantConfigurationTemplate) DeepCopy() *TenantConfigurationTemplate {
	if in == nil {
		return nil
	}
	out := new(TenantConfigurationTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TenantConfigurationTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantConfigurationTemplateList) DeepCopyInto(out *TenantConfigurationTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TenantConfigurationTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantConfigurationTemplateList.
func (in *TenantConfigurationTemplateList) DeepCopy() *TenantConfigurationTemplateList {
	if in == nil {
		return nil
	}
	out := new(TenantConfigurationTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TenantConfigurationTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantConfigurationTemplateSpec) DeepCopyInto(out *TenantConfigurationTemplateSpec) {
	*out = *in
	if in.TenantEngineTemplates != nil {
		in, out := &in.TenantEngineTemplates, &out.TenantEngineTemplates
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantConfigurationTemplateSpec.
func (in *TenantConfigurationTemplateSpec) DeepCopy() *TenantConfigurationTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(TenantConfigurationTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantConfigurationTemplateStatus) DeepCopyInto(out *TenantConfigurationTemplateStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantConfigurationTemplateStatus.
func (in *TenantConfigurationTemplateStatus) DeepCopy() *TenantConfigurationTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(TenantConfigurationTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantDatasetTemplate) DeepCopyInto(out *TenantDatasetTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantDatasetTemplate.
func (in *TenantDatasetTemplate) DeepCopy() *TenantDatasetTemplate {
	if in == nil {
		return nil
	}
	out := new(TenantDatasetTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TenantDatasetTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantDatasetTemplateList) DeepCopyInto(out *TenantDatasetTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TenantDatasetTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantDatasetTemplateList.
func (in *TenantDatasetTemplateList) DeepCopy() *TenantDatasetTemplateList {
	if in == nil {
		return nil
	}
	out := new(TenantDatasetTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TenantDatasetTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantDatasetTemplateSpec) DeepCopyInto(out *TenantDatasetTemplateSpec) {
	*out = *in
	if in.TenantEngineTemplates != nil {
		in, out := &in.TenantEngineTemplates, &out.TenantEngineTemplates
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantDatasetTemplateSpec.
func (in *TenantDatasetTemplateSpec) DeepCopy() *TenantDatasetTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(TenantDatasetTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantDatasetTemplateStatus) DeepCopyInto(out *TenantDatasetTemplateStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantDatasetTemplateStatus.
func (in *TenantDatasetTemplateStatus) DeepCopy() *TenantDatasetTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(TenantDatasetTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantEngineConfigurationTemplate) DeepCopyInto(out *TenantEngineConfigurationTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantEngineConfigurationTemplate.
func (in *TenantEngineConfigurationTemplate) DeepCopy() *TenantEngineConfigurationTemplate {
	if in == nil {
		return nil
	}
	out := new(TenantEngineConfigurationTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TenantEngineConfigurationTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantEngineConfigurationTemplateList) DeepCopyInto(out *TenantEngineConfigurationTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TenantEngineConfigurationTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantEngineConfigurationTemplateList.
func (in *TenantEngineConfigurationTemplateList) DeepCopy() *TenantEngineConfigurationTemplateList {
	if in == nil {
		return nil
	}
	out := new(TenantEngineConfigurationTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TenantEngineConfigurationTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantEngineConfigurationTemplateSpec) DeepCopyInto(out *TenantEngineConfigurationTemplateSpec) {
	*out = *in
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantEngineConfigurationTemplateSpec.
func (in *TenantEngineConfigurationTemplateSpec) DeepCopy() *TenantEngineConfigurationTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(TenantEngineConfigurationTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantEngineConfigurationTemplateStatus) DeepCopyInto(out *TenantEngineConfigurationTemplateStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantEngineConfigurationTemplateStatus.
func (in *TenantEngineConfigurationTemplateStatus) DeepCopy() *TenantEngineConfigurationTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(TenantEngineConfigurationTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantEngineDatasetTemplate) DeepCopyInto(out *TenantEngineDatasetTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantEngineDatasetTemplate.
func (in *TenantEngineDatasetTemplate) DeepCopy() *TenantEngineDatasetTemplate {
	if in == nil {
		return nil
	}
	out := new(TenantEngineDatasetTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TenantEngineDatasetTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantEngineDatasetTemplateList) DeepCopyInto(out *TenantEngineDatasetTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TenantEngineDatasetTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantEngineDatasetTemplateList.
func (in *TenantEngineDatasetTemplateList) DeepCopy() *TenantEngineDatasetTemplateList {
	if in == nil {
		return nil
	}
	out := new(TenantEngineDatasetTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TenantEngineDatasetTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantEngineDatasetTemplateSpec) DeepCopyInto(out *TenantEngineDatasetTemplateSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantEngineDatasetTemplateSpec.
func (in *TenantEngineDatasetTemplateSpec) DeepCopy() *TenantEngineDatasetTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(TenantEngineDatasetTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantEngineDatasetTemplateStatus) DeepCopyInto(out *TenantEngineDatasetTemplateStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantEngineDatasetTemplateStatus.
func (in *TenantEngineDatasetTemplateStatus) DeepCopy() *TenantEngineDatasetTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(TenantEngineDatasetTemplateStatus)
	in.DeepCopyInto(out)
	return out
}
