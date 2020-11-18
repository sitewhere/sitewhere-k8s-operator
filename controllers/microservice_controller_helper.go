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
	"reflect"

	"github.com/pkg/errors"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/client"

	sitewhereiov1alpha4 "github.com/sitewhere/sitewhere-k8s-operator/apis/sitewhere.io/v1alpha4"
)

const (
	// ErrLocateInstance is the error while locating parent instance.
	ErrLocateInstance = "cannot locate the parent instance"
)

const (
	defaultLabelInstance = "app.kubernetes.io/instance"
	defaultLabelManageBy = "app.kubernetes.io/managed-by"
	defaultLabelName     = "app.kubernetes.io/name"

	sitewhereLabelInstance       = "sitewhere.io/instance"
	sitewhereLabelName           = "sitewhere.io/name"
	sitewhereLabelRole           = "sitewhere.io/role"
	sitewhereLabelTenant         = "sitewhere.io/tenant"
	sitewhereLabelMicroservice   = "sitewhere.io/microservice"
	sitewhereLabelFunctionalArea = "sitewhere.io/functional-area"

	labelManagedBySiteWhere = "sitewhere-k8s-operator"
	labelRoleMicroservice   = "microservice"
)

var (
	deploymentKind       = reflect.TypeOf(appsv1.Deployment{}).Name()
	deploymentAPIVersion = appsv1.SchemeGroupVersion.String()
	serviceKind          = reflect.TypeOf(corev1.Service{}).Name()
	serviceAPIVersion    = corev1.SchemeGroupVersion.String()
)

//RenderMicroservicesDeployment derives apps.Deployment from a SiteWhereMicroservice
func RenderMicroservicesDeployment(swInstance *sitewhereiov1alpha4.SiteWhereInstance, swMicroservice *sitewhereiov1alpha4.SiteWhereMicroservice) (*appsv1.Deployment, error) {

	var labelsSelectorMap = buildLabelsSelectors(swInstance, swMicroservice)

	var dockerSpec = sitewhereiov1alpha4.DefaultDockerSpec
	if swInstance.Spec.DockerSpec != nil {
		dockerSpec = swInstance.Spec.DockerSpec
	}

	var imageName = fmt.Sprintf("%s/%s/service-%s:%s",
		dockerSpec.Registry,
		dockerSpec.Repository,
		swMicroservice.GetName(),
		dockerSpec.Tag)

	var deployment = &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       deploymentKind,
			APIVersion: deploymentAPIVersion,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      swMicroservice.Name,
			Namespace: swMicroservice.Namespace,
			Labels:    buildObjectMetaLabels(swInstance, swMicroservice),
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &swMicroservice.Spec.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: labelsSelectorMap,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labelsSelectorMap,
					Annotations: map[string]string{
						"prometheus.io/port":   "9090",
						"prometheus.io/scheme": "http",
						"prometheus.io/scrape": "true",
					},
				},
				Spec: corev1.PodSpec{
					ServiceAccountName: swInstance.GetName(),
					Containers: []corev1.Container{
						corev1.Container{
							Name:  swMicroservice.GetName(),
							Image: imageName,
							Ports: []corev1.ContainerPort{
								corev1.ContainerPort{
									ContainerPort: 9000,
									Protocol:      corev1.ProtocolTCP,
								},
								corev1.ContainerPort{
									ContainerPort: 9090,
									Protocol:      corev1.ProtocolTCP,
								},
							},
							Env: []corev1.EnvVar{
								corev1.EnvVar{
									Name: "sitewhere.config.k8s.name",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											APIVersion: "v1",
											FieldPath:  "metadata.name",
										},
									},
								},
								corev1.EnvVar{
									Name: "sitewhere.config.k8s.namespace",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											APIVersion: "v1",
											FieldPath:  "metadata.namespace",
										},
									},
								},
								corev1.EnvVar{
									Name: "sitewhere.config.k8s.pod.ip",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											APIVersion: "v1",
											FieldPath:  "status.podIP",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	// Handle Instance Management special case
	if swMicroservice.GetName() == FunctionalAreaInstanceManagement {
		deployment.Spec.Template.Spec.Containers[0].Ports = append(deployment.Spec.Template.Spec.Containers[0].Ports,
			corev1.ContainerPort{
				ContainerPort: 8080,
				Protocol:      corev1.ProtocolTCP,
			})
	}

	return deployment, nil
}

//RenderMicroservicesService derives corev1.Service from a SiteWhereMicroservice
func RenderMicroservicesService(swInstance *sitewhereiov1alpha4.SiteWhereInstance,
	swMicroservice *sitewhereiov1alpha4.SiteWhereMicroservice,
	deploy *appsv1.Deployment) ([]*corev1.Service, error) {

	var services = []*corev1.Service{
		&corev1.Service{
			TypeMeta: metav1.TypeMeta{
				Kind:       serviceKind,
				APIVersion: serviceAPIVersion,
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      swMicroservice.Name,
				Namespace: swMicroservice.Namespace,
				Labels:    buildObjectMetaLabels(swInstance, swMicroservice),
			},
			Spec: corev1.ServiceSpec{
				Selector: deploy.Spec.Selector.MatchLabels,
				Type:     corev1.ServiceTypeClusterIP,
				Ports: []corev1.ServicePort{
					corev1.ServicePort{
						Name:       "grpc-api",
						Port:       9000,
						Protocol:   corev1.ProtocolTCP,
						TargetPort: intstr.IntOrString{IntVal: 9000},
					},
					corev1.ServicePort{
						Name:       "http-metrics",
						Port:       9090,
						Protocol:   corev1.ProtocolTCP,
						TargetPort: intstr.IntOrString{IntVal: 9090},
					},
				},
			},
		},
	}

	// Handle Instance Management special case
	if swMicroservice.GetName() == FunctionalAreaInstanceManagement {
		var service = &corev1.Service{
			TypeMeta: metav1.TypeMeta{
				Kind:       serviceKind,
				APIVersion: serviceAPIVersion,
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("%s-rest", swMicroservice.Name),
				Namespace: swMicroservice.Namespace,
				Labels:    buildObjectMetaLabels(swInstance, swMicroservice),
			},
			Spec: corev1.ServiceSpec{
				Selector: deploy.Spec.Selector.MatchLabels,
				Type:     corev1.ServiceTypeLoadBalancer,
				Ports: []corev1.ServicePort{
					corev1.ServicePort{
						Name:       "http-rest",
						Port:       8080,
						Protocol:   corev1.ProtocolTCP,
						TargetPort: intstr.IntOrString{IntVal: 8080},
					},
				},
			},
		}
		services = append(services, service)
	}

	return services, nil
}

//LocateParentSiteWhereInstance locates the parent SiteWhereInstance of a Microservice
func LocateParentSiteWhereInstance(ctx context.Context, client client.Client, swMicroservice *sitewhereiov1alpha4.SiteWhereMicroservice) (*sitewhereiov1alpha4.SiteWhereInstance, error) {
	var acName string
	var eventObj = &sitewhereiov1alpha4.SiteWhereInstance{}

	for _, o := range swMicroservice.GetOwnerReferences() {
		if o.Kind == sitewhereiov1alpha4.SiteWhereInstanceKind {
			acName = o.Name
			break
		}
	}
	if len(acName) > 0 {
		nn := types.NamespacedName{
			Name: acName,
		}
		if err := client.Get(ctx, nn, eventObj); err != nil {
			return nil, err
		}
		return eventObj, nil
	}
	return nil, errors.Errorf(ErrLocateInstance)
}

// buildObjectMetaLabels buils the map of labels for object metadata
func buildObjectMetaLabels(
	swInstance *sitewhereiov1alpha4.SiteWhereInstance,
	swMicroservice *sitewhereiov1alpha4.SiteWhereMicroservice) map[string]string {
	return map[string]string{
		defaultLabelInstance:   swInstance.GetName(),
		defaultLabelManageBy:   labelManagedBySiteWhere,
		defaultLabelName:       swMicroservice.GetName(),
		sitewhereLabelInstance: swInstance.GetName(),
		sitewhereLabelName:     swMicroservice.GetName(),
		sitewhereLabelRole:     labelRoleMicroservice,
	}
}

// buildLabelsSelectors builds the map for labeles for deployment selector
func buildLabelsSelectors(
	swInstance *sitewhereiov1alpha4.SiteWhereInstance,
	swMicroservice *sitewhereiov1alpha4.SiteWhereMicroservice) map[string]string {
	return map[string]string{
		defaultLabelName:     swMicroservice.GetName(),
		defaultLabelInstance: swInstance.GetName(),
	}
}
