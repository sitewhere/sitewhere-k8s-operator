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

package sitewhereio

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

// +kubebuilder:rbac:groups=sitewhere.io,resources=microservices,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=sitewhere.io,resources=microservices/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=sitewhere.io,resources=microservices/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the SiteWhereMicroservice object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.7.0/pkg/reconcile
func (r *SiteWhereMicroserviceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("sitewheremicroservice", req.NamespacedName)
	log.Info("Reconcile SiteWhere Microservice")

	var swMicroservice sitewhereiov1alpha4.SiteWhereMicroservice
	if err := r.Get(ctx, req.NamespacedName, &swMicroservice); err != nil {
		if apierrors.IsNotFound(err) {
			log.Info("SiteWhere Microservice is deleted")
		}
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	// Look up parent instance
	msParentInstance, err := LocateParentSiteWhereInstance(ctx, r.Client, &swMicroservice)
	if err != nil {
		log.Error(err, "cannot locate instance for microservice")
		return ctrl.Result{}, err
	}

	// TODO Add SA, Role, RoleBindings for deployment
	// Render the deployment
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

// SetupWithManager sets up the controller with the Manager.
func (r *SiteWhereMicroserviceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&sitewhereiov1alpha4.SiteWhereMicroservice{}).
		Complete(r)
}
