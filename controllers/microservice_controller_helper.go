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
	"sigs.k8s.io/controller-runtime/pkg/client"

	sitewhereiov1alpha4 "github.com/sitewhere/sitewhere-k8s-operator/api/v1alpha4"
)

const (
	// ErrLocateInstance is the error while locating parent instance.
	ErrLocateInstance = "cannot locate the parent instance"
)

var (
	deploymentKind       = reflect.TypeOf(appsv1.Deployment{}).Name()
	deploymentAPIVersion = appsv1.SchemeGroupVersion.String()
	serviceKind          = reflect.TypeOf(corev1.Service{}).Name()
	serviceAPIVersion    = corev1.SchemeGroupVersion.String()
)

//RenderMicroservicesDeployment derives apps.Deployment from a SiteWhereMicroservice
func RenderMicroservicesDeployment(swInstance *sitewhereiov1alpha4.SiteWhereInstance, swMicroservice *sitewhereiov1alpha4.SiteWhereMicroservice) (*appsv1.Deployment, error) {

	var labelsSelectorMap = make(map[string]string)

	labelsSelectorMap["app.kubernetes.io/instance"] = swInstance.GetName()
	labelsSelectorMap["app.kubernetes.io/instance"] = swMicroservice.GetName()

	//TODO replace registry, repository and tag for variables of the instance
	var imageName = fmt.Sprintf("docker.io/sitewhere/service-%s:3.0.0.beta1", swMicroservice.GetName())

	return &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       deploymentKind,
			APIVersion: deploymentAPIVersion,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      swMicroservice.Name,
			Namespace: swMicroservice.Namespace,
			Labels: map[string]string{
				"app.kubernetes.io/instance":   swInstance.GetName(),
				"app.kubernetes.io/managed-by": "sitewhere-k8s-operator",
				"app.kubernetes.io/name":       swInstance.GetName(),
				"sitewhere.io/instance":        swInstance.GetName(),
				"sitewhere.io/name":            swMicroservice.GetName(),
				"sitewhere.io/role":            "microservice",
			},
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
					Containers: []corev1.Container{
						corev1.Container{
							Name:  swMicroservice.GetName(),
							Image: imageName,
						},
					},
				},
			},
		},
	}, nil
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
			Name:      acName,
			Namespace: swMicroservice.GetNamespace(),
		}
		if err := client.Get(ctx, nn, eventObj); err != nil {
			return nil, err
		}
		return eventObj, nil
	}
	return nil, errors.Errorf(ErrLocateInstance)
}
