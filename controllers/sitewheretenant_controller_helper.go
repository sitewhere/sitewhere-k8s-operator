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

package controllers

import (
	"context"
	"fmt"

	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	sitewhereiov1alpha4 "github.com/sitewhere/sitewhere-k8s-operator/apis/sitewhere.io/v1alpha4"
	templatesv1alpha4 "github.com/sitewhere/sitewhere-k8s-operator/apis/templates.sitewhere.io/v1alpha4"
)

const (
	// ErrLocateTenantEngineConfigurationTemplate is the error while locating tenant engine template
	ErrLocateTenantEngineConfigurationTemplate = "cannot locate the tenant engine template"
)

//RenderTenantEngine derives SiteWhereTenantEngine from a SiteWhereMicroservice and a SiteWhereTenant
func RenderTenantEngine(ctx context.Context, client client.Client, swTenant *sitewhereiov1alpha4.SiteWhereTenant, swMicroservice *sitewhereiov1alpha4.SiteWhereMicroservice) (*sitewhereiov1alpha4.SiteWhereTenantEngine, error) {
	var name = fmt.Sprintf("%s-%s", swTenant.GetName(), swMicroservice.GetName())
	//name = name[:63]

	tecTemplate, err := FindTenantEngineConfigurationTemplate(ctx, client, swTenant, swMicroservice)

	if err != nil {
		return nil, errors.Errorf(ErrLocateTenantEngineConfigurationTemplate)
	}

	var tenantEngine *sitewhereiov1alpha4.SiteWhereTenantEngine = &sitewhereiov1alpha4.SiteWhereTenantEngine{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: swTenant.ObjectMeta.Namespace,
			Labels: map[string]string{
				sitewhereLabelTenant:         swTenant.ObjectMeta.Name,
				sitewhereLabelMicroservice:   swMicroservice.ObjectMeta.Name,
				sitewhereLabelFunctionalArea: string(swMicroservice.Spec.FunctionalArea),
			},
		},
		Spec: sitewhereiov1alpha4.SiteWhereTenantEngineSpec{
			Configuration: tecTemplate.Spec.Configuration,
		},
	}
	return tenantEngine, nil
}

// FindTenantEngineConfigurationTemplate finds a tenant engine configuration template
func FindTenantEngineConfigurationTemplate(ctx context.Context,
	client client.Client,
	swTenant *sitewhereiov1alpha4.SiteWhereTenant,
	swMicroservice *sitewhereiov1alpha4.SiteWhereMicroservice) (*templatesv1alpha4.TenantEngineConfigurationTemplate, error) {
	var tenantTemplate = &templatesv1alpha4.TenantConfigurationTemplate{}
	if err := client.Get(ctx, types.NamespacedName{Name: swTenant.Spec.ConfigurationTemplate}, tenantTemplate); err != nil {
		return nil, err
	}

	var key = strcase.ToLowerCamel(string(swMicroservice.Spec.FunctionalArea))
	var name = tenantTemplate.Spec.TenantEngineTemplates[key]

	var tenantEngineTemplate = &templatesv1alpha4.TenantEngineConfigurationTemplate{}
	if err := client.Get(ctx, types.NamespacedName{Name: name}, tenantEngineTemplate); err != nil {
		return nil, err
	}
	return tenantEngineTemplate, nil
}
