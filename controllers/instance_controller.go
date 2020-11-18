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

	"github.com/go-logr/logr"
	core "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	sitewhereiov1alpha4 "github.com/sitewhere/sitewhere-k8s-operator/apis/sitewhere.io/v1alpha4"
)

// SiteWhereInstanceReconciler reconciles a SiteWhereInstance object
type SiteWhereInstanceReconciler struct {
	client.Client
	Log      logr.Logger
	Recorder record.EventRecorder
	Scheme   *runtime.Scheme
}

// +kubebuilder:rbac:groups=sitewhere.io,resources=instances,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=sitewhere.io,resources=instances/status,verbs=get;update;patch

func (r *SiteWhereInstanceReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("sitewhereinstance", req.NamespacedName)
	log.Info("Reconcile SiteWhere Instance")

	var swInstance sitewhereiov1alpha4.SiteWhereInstance
	if err := r.Get(ctx, req.NamespacedName, &swInstance); err != nil {
		if apierrors.IsNotFound(err) {
			log.Info("SiteWhere Instance is deleted")
		}
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	r.Recorder.Event(&swInstance, core.EventTypeNormal, "Updated", "Bootstraping")

	swInstance.Status.TenantManagementBootstrapState = sitewhereiov1alpha4.Bootstrapping
	swInstance.Status.UserManagementBootstrapState = sitewhereiov1alpha4.Bootstrapping

	if err := r.Status().Update(context.Background(), &swInstance); err != nil {
		return ctrl.Result{}, err
	}

	// Reder Namespace base on the instance's name where the operator will be placing all of it's objectes for the instance.
	namespace, err := RenderInstanceNamespace(&swInstance)
	if err != nil {
		log.Error(err, "can not render microservices from instance")
		return ctrl.Result{}, err
	}
	// Set SiteWhereInstace instance as the owner and controller
	ctrl.SetControllerReference(&swInstance, namespace, r.Scheme)
	if err := r.Create(ctx, namespace); err != nil {
		if apierrors.IsAlreadyExists(err) {
			log.Info(fmt.Sprintf("Namespace %s already exists", namespace.GetName()))
		} else {
			log.Error(err, "can not create namespace from instance")
			return ctrl.Result{}, err
		}
	} else {
		var message = fmt.Sprintf("Namespace %s for instance created.", namespace.GetName())
		r.Recorder.Event(&swInstance, core.EventTypeNormal, "Updated", message)
	}

	// Render the service account
	msServiceAccount, err := RenderMicroservicesServiceAccount(&swInstance, namespace)
	if err != nil {
		log.Error(err, "cannot render service account for instace")
		return ctrl.Result{}, err
	}
	// Set ownership
	ctrl.SetControllerReference(&swInstance, msServiceAccount, r.Scheme)
	if err := r.Create(ctx, msServiceAccount); err != nil {
		if apierrors.IsAlreadyExists(err) {
			log.Info(fmt.Sprintf("ServiceAccount %s already exists", msServiceAccount.GetName()))
		} else {
			log.Error(err, "can not create ServiceAccount from instance")
			return ctrl.Result{}, err
		}
	} else {
		var message = fmt.Sprintf("ServiceAccount %s for instance created.", msServiceAccount.GetName())
		r.Recorder.Event(&swInstance, core.EventTypeNormal, "Updated", message)
	}

	// If we don't have configuration, copy from InstanceConfigurationTemplate
	if swInstance.Spec.Configuration == nil {
		instanceConfigurationTemplate, err := FindInstanceConfigurationTemplate(ctx, r.Client, swInstance.Spec.ConfigurationTemplate)
		if err != nil {
			log.Error(err, fmt.Sprintf("can not find instance configuration template %s", swInstance.Spec.ConfigurationTemplate))
			return ctrl.Result{}, err
		}
		swInstance.Spec.Configuration = instanceConfigurationTemplate.Spec.Configuration.DeepCopy()
		if err := r.Update(context.Background(), &swInstance); err != nil {
			log.Error(err, "Failed to update SiteWhereInstance")
			r.Recorder.Event(&swInstance, core.EventTypeWarning, "Configuration", err.Error())
			return ctrl.Result{}, err
		}
		r.Recorder.Event(&swInstance, core.EventTypeNormal, "Configuration", "Updated")
	}

	microservices, err := RenderMicroservices(&swInstance, namespace)
	if err != nil {
		log.Error(err, "can not render microservices from instance")
		return ctrl.Result{}, err
	}
	for _, ms := range microservices {
		// Check if microservice exists
		var swMicroservice sitewhereiov1alpha4.SiteWhereMicroservice
		if err := r.Get(ctx, types.NamespacedName{Namespace: ms.Namespace, Name: ms.Name}, &swMicroservice); err != nil {
			if apierrors.IsNotFound(err) {
				// Set SiteWhereInstace instance as the owner and controller
				ctrl.SetControllerReference(&swInstance, ms, r.Scheme)
				if err := r.Create(ctx, ms); err != nil {
					log.Error(err, "can not create microservices from instance")
					return ctrl.Result{}, err
				}
			}
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager setups up k8s controller.
func (r *SiteWhereInstanceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&sitewhereiov1alpha4.SiteWhereInstance{}).
		Complete(r)
}
