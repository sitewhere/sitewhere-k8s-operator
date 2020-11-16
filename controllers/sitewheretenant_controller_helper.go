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
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	sitewhereiov1alpha4 "github.com/sitewhere/sitewhere-k8s-operator/api/v1alpha4"
)

//RenderTenantEngine derives SiteWhereTenantEngine from a SiteWhereMicroservice and a SiteWhereTenant
func RenderTenantEngine(swTenant *sitewhereiov1alpha4.SiteWhereTenant, swMicroservice *sitewhereiov1alpha4.SiteWhereMicroservice) (*sitewhereiov1alpha4.SiteWhereTenantEngine, error) {
	var name = fmt.Sprintf("%s-%s", swTenant.GetName(), swMicroservice.GetName())
	name = name[:63]

	// Look up tenant configuration template for tenant/microservice combination.
	// TenantEngineConfigurationTemplate tecTemplate = getTenantEngineConfigurationTemplate(tenant, microservice);
	// if (tecTemplate == null) {
	//     throw new SiteWhereK8sException(
	// 	    String.format("Unable to resolve default tenant engine configuration for '%s'.", functionalArea));
	// }

	var tenantEngine *sitewhereiov1alpha4.SiteWhereTenantEngine = &sitewhereiov1alpha4.SiteWhereTenantEngine{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: swTenant.ObjectMeta.Namespace,
			Labels: map[string]string{
				sitewhereLabelTenant:         swTenant.ObjectMeta.Name,
				sitewhereLabelMicroservice:   swMicroservice.ObjectMeta.Name,
				sitewhereLabelFunctionalArea: swMicroservice.Spec.FunctionalArea,
			},
		},
	}
	return tenantEngine, nil
}
