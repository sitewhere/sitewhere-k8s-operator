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
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var sitewhereinstancelog = logf.Log.WithName("sitewhereinstance-resource")

const (
	defaultRegistry   string = "docker.io"
	defaultRepository string = "sitewhere"
	defaultTag        string = "3.0.0.beta1"
)

func (r *SiteWhereInstance) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-sitewhere-io-v1alpha4-sitewhereinstance,mutating=true,failurePolicy=fail,groups=sitewhere.io,resources=instances,verbs=create;update,versions=v1alpha4,name=msitewhereinstance.kb.io

var _ webhook.Defaulter = &SiteWhereInstance{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *SiteWhereInstance) Default() {
	sitewhereinstancelog.Info("default", "name", r.Name)
	if r.Spec.DockerSpec == nil {
		sitewhereinstancelog.Info("Creating", "DockerSpec", r.Name)
		r.Spec.DockerSpec = &DockerSpec{
			Registry:   defaultRegistry,
			Repository: defaultRepository,
			Tag:        defaultTag,
		}
	} else {
		if r.Spec.DockerSpec.Registry == "" {
			sitewhereinstancelog.Info("Updateing", "DockerSpec.Registry", r.Name)
			r.Spec.DockerSpec.Registry = defaultRegistry
		}
		if r.Spec.DockerSpec.Repository == "" {
			sitewhereinstancelog.Info("Updateing", "DockerSpec.Repository", r.Name)
			r.Spec.DockerSpec.Repository = defaultRepository
		}
		if r.Spec.DockerSpec.Tag == "" {
			sitewhereinstancelog.Info("Updateing", "DockerSpec.Tag", r.Name)
			r.Spec.DockerSpec.Tag = defaultTag
		}
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update,path=/validate-sitewhere-io-v1alpha4-sitewhereinstance,mutating=false,failurePolicy=fail,groups=sitewhere.io,resources=instances,versions=v1alpha4,name=vsitewhereinstance.kb.io

var _ webhook.Validator = &SiteWhereInstance{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *SiteWhereInstance) ValidateCreate() error {
	sitewhereinstancelog.Info("validate create", "name", r.Name)
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *SiteWhereInstance) ValidateUpdate(old runtime.Object) error {
	sitewhereinstancelog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *SiteWhereInstance) ValidateDelete() error {
	sitewhereinstancelog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
