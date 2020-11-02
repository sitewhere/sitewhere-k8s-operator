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
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	sitewhereiov1alpha4 "github.com/sitewhere/sitewhere-k8s-operator/api/v1alpha4"
)

// SiteWhereInstanceReconciler reconciles a SiteWhereInstance object
type SiteWhereInstanceReconciler struct {
	client.Client
	Log      logr.Logger
	Recorder record.EventRecorder
	Scheme   *runtime.Scheme
}

// +kubebuilder:rbac:groups=sitewhere.io,resources=sitewhereinstances,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=sitewhere.io,resources=sitewhereinstances/status,verbs=get;update;patch

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

	r.Recorder.Event(&swInstance, "Warning", "ProcessingError", "some error")

	return ctrl.Result{}, nil
}

// SetupWithManager setups up k8s controller.
func (r *SiteWhereInstanceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&sitewhereiov1alpha4.SiteWhereInstance{}).
		Complete(r)
}
