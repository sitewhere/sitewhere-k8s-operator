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

	"github.com/go-logr/logr"
	core "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	sitewhereiov1alpha4 "github.com/sitewhere/sitewhere-k8s-operator/api/v1alpha4"
)

// SiteWhereTenantReconciler reconciles a SiteWhereTenant object
type SiteWhereTenantReconciler struct {
	client.Client
	Log      logr.Logger
	Recorder record.EventRecorder
	Scheme   *runtime.Scheme
}

// +kubebuilder:rbac:groups=sitewhere.io,resources=sitewheretenants,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=sitewhere.io,resources=sitewheretenants/status,verbs=get;update;patch

func (r *SiteWhereTenantReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("sitewheretenant", req.NamespacedName)
	log.Info("Reconcile SiteWhere Tenant")

	var swTenant sitewhereiov1alpha4.SiteWhereTenant
	if err := r.Get(ctx, req.NamespacedName, &swTenant); err != nil {
		if apierrors.IsNotFound(err) {
			log.Info("SiteWhere Tenant is deleted")
		}
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	r.Recorder.Event(&swTenant, core.EventTypeNormal, "Updated", "Bootstraping")

	var msList sitewhereiov1alpha4.SiteWhereMicroserviceList
	if err := r.Get(ctx, types.NamespacedName{Namespace: req.NamespacedName.Namespace}, &msList); err != nil {
		if apierrors.IsNotFound(err) {
			log.Info("No Microservices found in namespace")
		}
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	for _, swMicroservice := range msList.Items {
		// Render Tenant Engine from SiteWhereTenant/SiteWhereMicroservice
		if swMicroservice.Spec.Multitenant {
			tenantEngine, err := RenderTenantEngine(&swTenant, &swMicroservice)
			if err != nil {
				log.Error(err, "can not render tenant engine from tenant and microservice")
				return ctrl.Result{}, err
			}
			// Set SiteWhereTenant as the owner and controller
			ctrl.SetControllerReference(&swTenant, tenantEngine, r.Scheme)
			if err := r.Create(ctx, tenantEngine); err != nil {
				log.Error(err, "can not create tenant engine from tenant and microservice")
				return ctrl.Result{}, err
			}
		}
	}

	return ctrl.Result{}, nil
}

func (r *SiteWhereTenantReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&sitewhereiov1alpha4.SiteWhereTenant{}).
		Complete(r)
}
