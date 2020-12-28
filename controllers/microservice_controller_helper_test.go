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
			result: corev1.PodSpec{
				Containers: []corev1.Container{
					corev1.Container{}},
			},
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

func TestRenderContainerImageName(t *testing.T) {
	t.Parallel()
	data := []struct {
		name           string
		swInstance     *sitewhereiov1alpha4.SiteWhereInstance
		swMicroservice *sitewhereiov1alpha4.SiteWhereMicroservice
		result         string
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
					Name:      "some",
					Namespace: "test",
				},
				Spec: sitewhereiov1alpha4.SiteWhereMicroserviceSpec{
					FunctionalArea: "some",
				},
			},
			result: fmt.Sprintf("%s/%s/service-%s:%s", sitewhereiov1alpha4.DefaultDockerSpec.Registry, sitewhereiov1alpha4.DefaultDockerSpec.Repository, "some", sitewhereiov1alpha4.DefaultDockerSpec.Tag),
		},
		{
			name: "test-case-02",
			swInstance: &sitewhereiov1alpha4.SiteWhereInstance{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test",
					Namespace: "test",
				},
				Spec: sitewhereiov1alpha4.SiteWhereInstanceSpec{},
			},
			swMicroservice: &sitewhereiov1alpha4.SiteWhereMicroservice{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "some",
					Namespace: "test",
				},
				Spec: sitewhereiov1alpha4.SiteWhereMicroserviceSpec{
					FunctionalArea: "some",
					PodSpec: &sitewhereiov1alpha4.MicroservicePodSpecification{
						DockerSpec: &sitewhereiov1alpha4.DockerSpec{
							Registry:   "some-reg",
							Repository: "some-repo",
							Tag:        "some-tag",
						},
					},
				},
			},
			result: fmt.Sprintf("%s/%s/service-%s:%s", "some-reg", "some-repo", "some", "some-tag"),
		},
	}
	for _, single := range data {
		t.Run(single.name, func(single struct {
			name           string
			swInstance     *sitewhereiov1alpha4.SiteWhereInstance
			swMicroservice *sitewhereiov1alpha4.SiteWhereMicroservice
			result         string
		}) func(t *testing.T) {
			return func(t *testing.T) {
				result := renderContainerImageName(single.swInstance, single.swMicroservice)

				if result != single.result {
					t.Fatalf("expected %s containers, got %s", result, single.result)
				}
			}
		}(single))
	}
}

func TestRenderContainerImagePullPolicy(t *testing.T) {
	t.Parallel()
	data := []struct {
		name           string
		swInstance     *sitewhereiov1alpha4.SiteWhereInstance
		swMicroservice *sitewhereiov1alpha4.SiteWhereMicroservice
		result         corev1.PullPolicy
	}{
		{
			name: "when_no_podspec_return_default",
			swInstance: &sitewhereiov1alpha4.SiteWhereInstance{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test",
					Namespace: "test",
				},
				Spec: sitewhereiov1alpha4.SiteWhereInstanceSpec{},
			},
			swMicroservice: &sitewhereiov1alpha4.SiteWhereMicroservice{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "some",
					Namespace: "test",
				},
				Spec: sitewhereiov1alpha4.SiteWhereMicroserviceSpec{
					FunctionalArea: "some",
				},
			},
			result: DefaultImagePullPolicy,
		},
		{
			name: "when_podspec_override_return_override",
			swInstance: &sitewhereiov1alpha4.SiteWhereInstance{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test",
					Namespace: "test",
				},
				Spec: sitewhereiov1alpha4.SiteWhereInstanceSpec{},
			},
			swMicroservice: &sitewhereiov1alpha4.SiteWhereMicroservice{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "some",
					Namespace: "test",
				},
				Spec: sitewhereiov1alpha4.SiteWhereMicroserviceSpec{
					FunctionalArea: "some",
					PodSpec: &sitewhereiov1alpha4.MicroservicePodSpecification{
						ImagePullPolicy: corev1.PullNever,
					},
				},
			},
			result: corev1.PullNever,
		},
	}
	for _, single := range data {
		t.Run(single.name, func(single struct {
			name           string
			swInstance     *sitewhereiov1alpha4.SiteWhereInstance
			swMicroservice *sitewhereiov1alpha4.SiteWhereMicroservice
			result         corev1.PullPolicy
		}) func(t *testing.T) {
			return func(t *testing.T) {
				result := renderContainerImagePullPolicy(single.swInstance, single.swMicroservice)

				if result != single.result {
					t.Fatalf("expected %s containers, got %s", result, single.result)
				}
			}
		}(single))
	}
}

func TestRenderDeploymentPodSpecContainerPorts(t *testing.T) {
	t.Parallel()
	data := []struct {
		name           string
		swInstance     *sitewhereiov1alpha4.SiteWhereInstance
		swMicroservice *sitewhereiov1alpha4.SiteWhereMicroservice
		result         []corev1.ContainerPort
	}{
		{
			name: "when_no_podspec_return_default",
			swInstance: &sitewhereiov1alpha4.SiteWhereInstance{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test",
					Namespace: "test",
				},
				Spec: sitewhereiov1alpha4.SiteWhereInstanceSpec{},
			},
			swMicroservice: &sitewhereiov1alpha4.SiteWhereMicroservice{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "some",
					Namespace: "test",
				},
				Spec: sitewhereiov1alpha4.SiteWhereMicroserviceSpec{
					FunctionalArea: "some",
				},
			},
			result: DefaultContainerPorts,
		},
		{
			name: "when_podspec_override_return_override",
			swInstance: &sitewhereiov1alpha4.SiteWhereInstance{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test",
					Namespace: "test",
				},
				Spec: sitewhereiov1alpha4.SiteWhereInstanceSpec{},
			},
			swMicroservice: &sitewhereiov1alpha4.SiteWhereMicroservice{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "some",
					Namespace: "test",
				},
				Spec: sitewhereiov1alpha4.SiteWhereMicroserviceSpec{
					FunctionalArea: "some",
					PodSpec: &sitewhereiov1alpha4.MicroservicePodSpecification{
						Ports: []corev1.ContainerPort{
							corev1.ContainerPort{
								ContainerPort: 8080,
								Protocol:      corev1.ProtocolTCP,
							},
						},
					},
				},
			},
			result: []corev1.ContainerPort{
				corev1.ContainerPort{
					ContainerPort: 8080,
					Protocol:      corev1.ProtocolTCP,
				},
			},
		},
	}
	for _, single := range data {
		t.Run(single.name, func(single struct {
			name           string
			swInstance     *sitewhereiov1alpha4.SiteWhereInstance
			swMicroservice *sitewhereiov1alpha4.SiteWhereMicroservice
			result         []corev1.ContainerPort
		}) func(t *testing.T) {
			return func(t *testing.T) {
				result := renderDeploymentPodSpecContainerPorts(single.swInstance, single.swMicroservice)

				if !compareContainerPorts(result, single.result) {
					t.Fatalf("expected %v containers, got %v", result, single.result)
				}
			}
		}(single))
	}
}

func compareContainerPorts(lhs []corev1.ContainerPort, rhs []corev1.ContainerPort) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for index := 0; index < len(lhs); index++ {
		l := lhs[index]
		r := rhs[index]
		if !compareContainerPort(l, r) {
			return false
		}
	}
	return true
}

func compareContainerPort(lhs corev1.ContainerPort, rhs corev1.ContainerPort) bool {
	return (lhs.ContainerPort == rhs.ContainerPort) && (lhs.Protocol == rhs.Protocol)
}
