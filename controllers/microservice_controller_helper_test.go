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
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	sitewhereiov1alpha4 "github.com/sitewhere/sitewhere-k8s-operator/apis/sitewhere.io/v1alpha4"
)

func TestRenderDeploymentPodSpec(t *testing.T) {
	t.Parallel()
	data := []struct {
		name           string
		swInstance     *sitewhereiov1alpha4.SiteWhereInstance
		swMicroservice *sitewhereiov1alpha4.SiteWhereMicroservice
		result         corev1.PodSpec
	}{
		{
			name: "test-case-01",
			swInstance: &sitewhereiov1alpha4.SiteWhereInstance{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test",
					Namespace: "test",
				},
				Spec: sitewhereiov1alpha4.SiteWhereInstanceSpec{},
			},
			swMicroservice: &sitewhereiov1alpha4.SiteWhereMicroservice{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "ms-test",
					Namespace: "test",
				},
				Spec: sitewhereiov1alpha4.SiteWhereMicroserviceSpec{
					FunctionalArea: "some",
				},
			},
			result: corev1.PodSpec{},
		},
	}
	for _, single := range data {
		t.Run(single.name, func(single struct {
			name           string
			swInstance     *sitewhereiov1alpha4.SiteWhereInstance
			swMicroservice *sitewhereiov1alpha4.SiteWhereMicroservice
			result         corev1.PodSpec
		}) func(t *testing.T) {
			return func(t *testing.T) {
				result := renderDeploymentPodSpec(single.swInstance, single.swMicroservice)

				if len(result.Containers) != len(single.result.Containers) {
					t.Fatalf("expected %d containers, got %d", len(result.Containers), len(single.result.Containers))
				}
			}
		}(single))
	}
}
