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
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	sitewhereiov1alpha4 "github.com/sitewhere/sitewhere-k8s-operator/apis/sitewhere.io/v1alpha4"
)

// SiteWhereMicroserviceReconciler reconciles a SiteWhereMicroservice object
type SiteWhereMicroserviceReconciler struct {
	client.Client
	Log      logr.Logger
	Recorder record.EventRecorder
	Scheme   *runtime.Scheme
}

// +kubebuilder:rbac:groups=sitewhere.io,resources=sitewheremicroservices,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=sitewhere.io,resources=sitewheremicroservices/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delet

func (r *SiteWhereMicroserviceReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("sitewheremicroservice", req.NamespacedName)
	log.Info("Reconcile SiteWhere Microservice")

	var swMicroservice sitewhereiov1alpha4.SiteWhereMicroservice
	if err := r.Get(ctx, req.NamespacedName, &swMicroservice); err != nil {
		if apierrors.IsNotFound(err) {
			log.Info("SiteWhere Microservice is deleted")
		}
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	// Render the deployment
	msParentInstance, err := LocateParentSiteWhereInstance(ctx, r.Client, &swMicroservice)
	if err != nil {
		log.Error(err, "cannot locate instance for microservice")
		return ctrl.Result{}, err
	}
	msDeployment, err := RenderMicroservicesDeployment(msParentInstance, &swMicroservice)
	if err != nil {
		log.Error(err, "cannot render deployment for microservice")
		return ctrl.Result{}, err
	}
	// Set ownership
	ctrl.SetControllerReference(&swMicroservice, msDeployment, r.Scheme)
	// server side apply, only the fields we set are touched
	applyOpts := []client.PatchOption{client.ForceOwnership, client.FieldOwner(swMicroservice.GetUID())}
	if err := r.Patch(ctx, msDeployment, client.Apply, applyOpts...); err != nil {
		log.Error(err, "Failed to apply to a deployment")
		r.Recorder.Event(&swMicroservice, core.EventTypeWarning, "Deployment", err.Error())
		return ctrl.Result{}, err
	}
	r.Recorder.Event(&swMicroservice, core.EventTypeNormal, "Deployment", "Created")
	// Render the service
	msServices, err := RenderMicroservicesService(msParentInstance, &swMicroservice, msDeployment)
	if err != nil {
		log.Error(err, "cannot render services for microservice")
		return ctrl.Result{}, err
	}

	for _, msService := range msServices {
		// Set ownership
		ctrl.SetControllerReference(&swMicroservice, msService, r.Scheme)
		// server side apply, only the fields we set are touched
		if err := r.Patch(ctx, msService, client.Apply, applyOpts...); err != nil {
			log.Error(err, "Failed to apply to a service")
			r.Recorder.Event(&swMicroservice, core.EventTypeWarning, "Service", err.Error())
			return ctrl.Result{}, err
		}
		var message = fmt.Sprintf("%s Created", msService.GetName())
		r.Recorder.Event(&swMicroservice, core.EventTypeNormal, "Service", message)
	}
	swMicroservice.Status.Deployment = msDeployment.GetName()
	swMicroservice.Status.Services = nil
	for _, msService := range msServices {
		swMicroservice.Status.Services = append(swMicroservice.Status.Services, msService.GetName())
	}
	if err := r.Status().Update(ctx, &swMicroservice); err != nil {
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

func (r *SiteWhereMicroserviceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&sitewhereiov1alpha4.SiteWhereMicroservice{}).
		Complete(r)
}
