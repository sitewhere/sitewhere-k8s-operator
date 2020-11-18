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

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	sitewhereiov1alpha4 "github.com/sitewhere/sitewhere-k8s-operator/apis/sitewhere.io/v1alpha4"
	templatesv1alpha4 "github.com/sitewhere/sitewhere-k8s-operator/apis/templates.sitewhere.io/v1alpha4"
)

const (
	//FunctionalAreaAssetManagement is the name of the functional area Asset Management
	FunctionalAreaAssetManagement string = "asset-management"
	//FunctionalAreaBatchOperations is the name of the functional area Batch Operations
	FunctionalAreaBatchOperations string = "batch-operations"
	//FunctionalAreaCommandDelivery is the name of the functional area Command Delivery
	FunctionalAreaCommandDelivery string = "command-delivery"
	//FunctionalAreaDeviceManagement is the name of the functional area Device Management
	FunctionalAreaDeviceManagement string = "device-management"
	//FunctionalAreaDeviceRegistration is the name of the functional area Device Registration
	FunctionalAreaDeviceRegistration string = "device-registration"
	//FunctionalAreaDeviceState is the name of the functional area Device State
	FunctionalAreaDeviceState string = "device-state"
	//FunctionalAreaEventManagement is the name of the functional area Event Management
	FunctionalAreaEventManagement string = "event-management"
	//FunctionalAreaEventSources is the name of the functional area Event Sources
	FunctionalAreaEventSources string = "event-sources"
	//FunctionalAreaInboundProcessing is the name of the functional area Inbound Processing
	FunctionalAreaInboundProcessing string = "inbound-processing"
	//FunctionalAreaInstanceManagement is the name of the functional area Instance Management
	FunctionalAreaInstanceManagement string = "instance-management"
	//FunctionalAreaLabelGeneration is the name of the functional area Label Generation
	FunctionalAreaLabelGeneration string = "label-generation"
	//FunctionalAreaOutboundConnectors is the name of the functional area Outbound Connectors
	FunctionalAreaOutboundConnectors string = "outbound-connectors"
	//FunctionalAreaScheduleManagement is the name of the functional area Schedule Management
	FunctionalAreaScheduleManagement string = "schedule-management"
)

const (
	//DefaultReplica is the default value for replicas
	DefaultReplica int32 = 1
)

//RenderMicroservices derives SiteWhereMicroservice from a SiteWhereInstance
func RenderMicroservices(swInstance *sitewhereiov1alpha4.SiteWhereInstance, ns *corev1.Namespace) ([]*sitewhereiov1alpha4.SiteWhereMicroservice, error) {

	var result = []*sitewhereiov1alpha4.SiteWhereMicroservice{
		&sitewhereiov1alpha4.SiteWhereMicroservice{
			ObjectMeta: metav1.ObjectMeta{
				Name:      FunctionalAreaAssetManagement,
				Namespace: ns.GetName(),
				Labels: map[string]string{
					sitewhereLabelInstance:       swInstance.GetName(),
					sitewhereLabelFunctionalArea: FunctionalAreaAssetManagement,
				},
			},
			Spec: sitewhereiov1alpha4.SiteWhereMicroserviceSpec{
				FunctionalArea: FunctionalAreaAssetManagement,
				Name:           "Asset Management",
				Description:    "Manages asset model persistence and API operations",
				Icon:           "fa-laptop-house",
				Replicas:       DefaultReplica,
				Multitenant:    true,
			},
		},
		&sitewhereiov1alpha4.SiteWhereMicroservice{
			ObjectMeta: metav1.ObjectMeta{
				Name:      FunctionalAreaBatchOperations,
				Namespace: ns.GetName(),
				Labels: map[string]string{
					sitewhereLabelInstance:       swInstance.GetName(),
					sitewhereLabelFunctionalArea: FunctionalAreaBatchOperations,
				},
			},
			Spec: sitewhereiov1alpha4.SiteWhereMicroserviceSpec{
				FunctionalArea: FunctionalAreaBatchOperations,
				Name:           "Batch Operations",
				Description:    "Manages operations that are applied to a list of target devices",
				Icon:           "fa-th",
				Replicas:       DefaultReplica,
				Multitenant:    true,
			},
		},
		&sitewhereiov1alpha4.SiteWhereMicroservice{
			ObjectMeta: metav1.ObjectMeta{
				Name:      FunctionalAreaCommandDelivery,
				Namespace: ns.GetName(),
				Labels: map[string]string{
					sitewhereLabelInstance:       swInstance.GetName(),
					sitewhereLabelFunctionalArea: FunctionalAreaCommandDelivery,
				},
			},
			Spec: sitewhereiov1alpha4.SiteWhereMicroserviceSpec{
				FunctionalArea: FunctionalAreaCommandDelivery,
				Name:           "Command Delivery",
				Description:    "Delivers command payloads to devices based on preset rules",
				Icon:           "fa-broadcast-tower",
				Replicas:       DefaultReplica,
				Multitenant:    true,
			},
		},
		&sitewhereiov1alpha4.SiteWhereMicroservice{
			ObjectMeta: metav1.ObjectMeta{
				Name:      FunctionalAreaDeviceManagement,
				Namespace: ns.GetName(),
				Labels: map[string]string{
					sitewhereLabelInstance:       swInstance.GetName(),
					sitewhereLabelFunctionalArea: FunctionalAreaDeviceManagement,
				},
			},
			Spec: sitewhereiov1alpha4.SiteWhereMicroserviceSpec{
				FunctionalArea: FunctionalAreaDeviceManagement,
				Name:           "Device Management",
				Description:    "Manages device model persistence and API operations",
				Icon:           "fa-microchip",
				Replicas:       DefaultReplica,
				Multitenant:    true,
			},
		},
		&sitewhereiov1alpha4.SiteWhereMicroservice{
			ObjectMeta: metav1.ObjectMeta{
				Name:      FunctionalAreaDeviceRegistration,
				Namespace: ns.GetName(),
				Labels: map[string]string{
					sitewhereLabelInstance:       swInstance.GetName(),
					sitewhereLabelFunctionalArea: FunctionalAreaDeviceRegistration,
				},
			},
			Spec: sitewhereiov1alpha4.SiteWhereMicroserviceSpec{
				FunctionalArea: FunctionalAreaDeviceRegistration,
				Name:           "Device Registration",
				Description:    "Handles registration of new devices with the system",
				Icon:           "fa-plus-circle",
				Replicas:       DefaultReplica,
				Multitenant:    true,
			},
		},
		&sitewhereiov1alpha4.SiteWhereMicroservice{
			ObjectMeta: metav1.ObjectMeta{
				Name:      FunctionalAreaDeviceState,
				Namespace: ns.GetName(),
				Labels: map[string]string{
					sitewhereLabelInstance:       swInstance.GetName(),
					sitewhereLabelFunctionalArea: FunctionalAreaDeviceState,
				},
			},
			Spec: sitewhereiov1alpha4.SiteWhereMicroserviceSpec{
				FunctionalArea: FunctionalAreaDeviceState,
				Name:           "Device State",
				Description:    "Provides device state management features such as device shadows",
				Icon:           "fa-exclamation-circle",
				Replicas:       DefaultReplica,
				Multitenant:    true,
			},
		},
		&sitewhereiov1alpha4.SiteWhereMicroservice{
			ObjectMeta: metav1.ObjectMeta{
				Name:      FunctionalAreaEventManagement,
				Namespace: ns.GetName(),
				Labels: map[string]string{
					sitewhereLabelInstance:       swInstance.GetName(),
					sitewhereLabelFunctionalArea: FunctionalAreaEventManagement,
				},
			},
			Spec: sitewhereiov1alpha4.SiteWhereMicroserviceSpec{
				FunctionalArea: FunctionalAreaEventManagement,
				Name:           "Event Management",
				Description:    "Provides APIs for persisting and accessing events generated by devices",
				Icon:           "fa-stream",
				Replicas:       DefaultReplica,
				Multitenant:    true,
			},
		},
		&sitewhereiov1alpha4.SiteWhereMicroservice{
			ObjectMeta: metav1.ObjectMeta{
				Name:      FunctionalAreaEventSources,
				Namespace: ns.GetName(),
				Labels: map[string]string{
					sitewhereLabelInstance:       swInstance.GetName(),
					sitewhereLabelFunctionalArea: FunctionalAreaEventSources,
				},
			},
			Spec: sitewhereiov1alpha4.SiteWhereMicroserviceSpec{
				FunctionalArea: FunctionalAreaEventSources,
				Name:           "Event Sources",
				Description:    "Handles inbound device data from various sources, protocols, and formats",
				Icon:           "fa-sign-in-alt",
				Replicas:       DefaultReplica,
				Multitenant:    true,
			},
		},
		&sitewhereiov1alpha4.SiteWhereMicroservice{
			ObjectMeta: metav1.ObjectMeta{
				Name:      FunctionalAreaInboundProcessing,
				Namespace: ns.GetName(),
				Labels: map[string]string{
					sitewhereLabelInstance:       swInstance.GetName(),
					sitewhereLabelFunctionalArea: FunctionalAreaInboundProcessing,
				},
			},
			Spec: sitewhereiov1alpha4.SiteWhereMicroserviceSpec{
				FunctionalArea: FunctionalAreaInboundProcessing,
				Name:           "Inbound Processing",
				Description:    "Common processing logic applied to enrich and direct inbound events",
				Icon:           "fa-cogs",
				Replicas:       DefaultReplica,
				Multitenant:    true,
			},
		},
		&sitewhereiov1alpha4.SiteWhereMicroservice{
			ObjectMeta: metav1.ObjectMeta{
				Name:      FunctionalAreaInstanceManagement,
				Namespace: ns.GetName(),
				Labels: map[string]string{
					sitewhereLabelInstance:       swInstance.GetName(),
					sitewhereLabelFunctionalArea: FunctionalAreaInstanceManagement,
				},
			},
			Spec: sitewhereiov1alpha4.SiteWhereMicroserviceSpec{
				FunctionalArea: FunctionalAreaInstanceManagement,
				Name:           "Instance Management",
				Description:    "Handles APIs for managing global aspects of an instance",
				Icon:           "fa-sitemap",
				Replicas:       DefaultReplica,
				Multitenant:    false,
			},
		},
		&sitewhereiov1alpha4.SiteWhereMicroservice{
			ObjectMeta: metav1.ObjectMeta{
				Name:      FunctionalAreaLabelGeneration,
				Namespace: ns.GetName(),
				Labels: map[string]string{
					sitewhereLabelInstance:       swInstance.GetName(),
					sitewhereLabelFunctionalArea: FunctionalAreaLabelGeneration,
				},
			},
			Spec: sitewhereiov1alpha4.SiteWhereMicroserviceSpec{
				FunctionalArea: FunctionalAreaLabelGeneration,
				Name:           "Label Generation",
				Description:    "Supports generating labels such as bar codes and QR codes for devices",
				Icon:           "fa-qrcode",
				Replicas:       DefaultReplica,
				Multitenant:    true,
			},
		},
		&sitewhereiov1alpha4.SiteWhereMicroservice{
			ObjectMeta: metav1.ObjectMeta{
				Name:      FunctionalAreaOutboundConnectors,
				Namespace: ns.GetName(),
				Labels: map[string]string{
					sitewhereLabelInstance:       swInstance.GetName(),
					sitewhereLabelFunctionalArea: FunctionalAreaOutboundConnectors,
				},
			},
			Spec: sitewhereiov1alpha4.SiteWhereMicroserviceSpec{
				FunctionalArea: FunctionalAreaOutboundConnectors,
				Name:           "Outbound Connectors",
				Description:    "Allows event streams to be delivered to external systems for additional processing",
				Icon:           "fa-sign-out-alt",
				Replicas:       DefaultReplica,
				Multitenant:    true,
			},
		},
		&sitewhereiov1alpha4.SiteWhereMicroservice{
			ObjectMeta: metav1.ObjectMeta{
				Name:      FunctionalAreaScheduleManagement,
				Namespace: ns.GetName(),
				Labels: map[string]string{
					sitewhereLabelInstance:       swInstance.GetName(),
					sitewhereLabelFunctionalArea: FunctionalAreaScheduleManagement,
				},
			},
			Spec: sitewhereiov1alpha4.SiteWhereMicroserviceSpec{
				FunctionalArea: FunctionalAreaScheduleManagement,
				Name:           "Schedule Management",
				Description:    "Supports scheduling of various system operations",
				Icon:           "fa-calendar-alt",
				Replicas:       DefaultReplica,
				Multitenant:    true,
			},
		},
	}

	return result, nil
}

//RenderInstanceNamespace derives a Namespace from the Instance
func RenderInstanceNamespace(swInstance *sitewhereiov1alpha4.SiteWhereInstance) (*corev1.Namespace, error) {
	return &corev1.Namespace{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Namespace",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: swInstance.GetName(),
		},
	}, nil
}

// RenderMicroservicesServiceAccount derices a Service Account for the Deployments of SW Instace
func RenderMicroservicesServiceAccount(swInstance *sitewhereiov1alpha4.SiteWhereInstance, namespace *corev1.Namespace) (*corev1.ServiceAccount, error) {
	return &corev1.ServiceAccount{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ServiceAccount",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      swInstance.GetName(),
			Namespace: namespace.GetName(),
		},
	}, nil
}

// FindInstanceConfigurationTemplate retrieves a InstanceConfigurationTemplate
func FindInstanceConfigurationTemplate(ctx context.Context,
	client client.Client, name string) (*templatesv1alpha4.InstanceConfigurationTemplate, error) {
	var intanceConfigTemplate = &templatesv1alpha4.InstanceConfigurationTemplate{}
	if err := client.Get(ctx, types.NamespacedName{Name: name}, intanceConfigTemplate); err != nil {
		return nil, err
	}
	return intanceConfigTemplate, nil
}
