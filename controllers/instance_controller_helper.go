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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	sitewhereiov1alpha4 "github.com/sitewhere/sitewhere-k8s-operator/api/v1alpha4"
)

const (
	//FunctionalAreaAssetManagement is the name of the functional area Asset Management
	FunctionalAreaAssetManagement string = "asset-management"
	//FunctionalAreaBatchOperations is the name of the functional area Batch Operations
	FunctionalAreaBatchOperations string = "batch-operations"
)

const (
	//DefaultReplica is the default value for replicas
	DefaultReplica int32 = 1
)

//RenderMicroservices derives SiteWhereMicroservice from a SiteWhereInstance
func RenderMicroservices(swInstance *sitewhereiov1alpha4.SiteWhereInstance) ([]*sitewhereiov1alpha4.SiteWhereMicroservice, error) {

	var result = []*sitewhereiov1alpha4.SiteWhereMicroservice{
		&sitewhereiov1alpha4.SiteWhereMicroservice{
			ObjectMeta: metav1.ObjectMeta{
				Name:      FunctionalAreaAssetManagement,
				Namespace: swInstance.ObjectMeta.Namespace,
			},
			Spec: sitewhereiov1alpha4.SiteWhereMicroserviceSpec{
				FunctionalArea: FunctionalAreaAssetManagement,
				Name:           "Asset Management",
				Description:    "Manages asset model persistence and API operations",
				Icon:           "fa-laptop-house",
				Replicas:       DefaultReplica,
			},
		},
		&sitewhereiov1alpha4.SiteWhereMicroservice{
			ObjectMeta: metav1.ObjectMeta{
				Name:      FunctionalAreaBatchOperations,
				Namespace: swInstance.ObjectMeta.Namespace,
			},
			Spec: sitewhereiov1alpha4.SiteWhereMicroserviceSpec{
				FunctionalArea: FunctionalAreaBatchOperations,
				Name:           "Batch Operations",
				Description:    "Manages operations that are applied to a list of target devices",
				Icon:           "fa-th",
				Replicas:       DefaultReplica,
			},
		},
	}

	return result, nil
}
