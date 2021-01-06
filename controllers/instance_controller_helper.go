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
	b64 "encoding/base64"
	"encoding/json"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/sitewhere/sitewhere-k8s-operator/pkg/funcarea"
	"github.com/sitewhere/sitewhere-k8s-operator/pkg/rand"

	sitewhereiov1alpha4 "github.com/sitewhere/sitewhere-k8s-operator/apis/sitewhere.io/v1alpha4"
	templatesv1alpha4 "github.com/sitewhere/sitewhere-k8s-operator/apis/templates.sitewhere.io/v1alpha4"
)

const (
	//DefaultReplica is the default value for replicas
	DefaultReplica int32 = 1
)

const (
	// SiteWhere Instance Cluster Role
	swInstanceClusterRoleName = "sitewhere:instance"

	// SiteWhere Instance Role Name
	swInstanceRoleName = "sitewhere-system-reader"
)

const (
	// Client Secret key
	clientSecretKey = "client-secret"
)

//RenderMicroservices derives SiteWhereMicroservice from a SiteWhereInstance
func RenderMicroservices(swInstance *sitewhereiov1alpha4.SiteWhereInstance, ns *corev1.Namespace) ([]*sitewhereiov1alpha4.SiteWhereMicroservice, error) {

	var result []*sitewhereiov1alpha4.SiteWhereMicroservice

	for _, fa := range swInstance.Spec.FunctionalAreas {
		microservice := RenderMicroseviceFromFunctionalArea(fa, swInstance, ns)
		result = append(result, microservice)
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

// RenderMicroservicesClusterRole derices a ClusterRole for the Deployments of SW Instace
func RenderMicroservicesClusterRole(swInstance *sitewhereiov1alpha4.SiteWhereInstance) (*rbacv1.ClusterRole, error) {
	return &rbacv1.ClusterRole{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ClusterRole",
			APIVersion: "rbac.authorization.k8s.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: swInstanceClusterRoleName,
			Labels: map[string]string{
				"app": "sitewhere",
			},
		},
		Rules: []rbacv1.PolicyRule{
			{
				APIGroups: []string{
					"sitewhere.io",
				},
				Resources: []string{
					"instances",
					"instances/status",
					"microservices",
					"tenants",
					"tenantengines",
					"tenantengines/status",
				},
				Verbs: []string{
					"*",
				},
			}, {
				APIGroups: []string{
					"templates.sitewhere.io",
				},
				Resources: []string{
					"instanceconfigurations",
					"instancedatasets",
					"tenantconfigurations",
					"tenantengineconfigurations",
					"tenantdatasets",
					"tenantenginedatasets",
				},
				Verbs: []string{
					"*",
				},
			}, {
				APIGroups: []string{
					"scripting.sitewhere.io",
				},
				Resources: []string{
					"scriptcategories",
					"scripttemplates",
					"scripts",
					"scriptversions",
				},
				Verbs: []string{
					"*",
				},
			}, {
				APIGroups: []string{
					"apiextensions.k8s.io",
				},
				Resources: []string{
					"customresourcedefinitions",
				},
				Verbs: []string{
					"*",
				},
			},
		},
	}, nil
}

// RenderMicroservicesClusterRoleBinding derices a ClusterRoleBinding for the Deployments of SW Instace
func RenderMicroservicesClusterRoleBinding(swInstance *sitewhereiov1alpha4.SiteWhereInstance,
	namespace *corev1.Namespace,
	sa *corev1.ServiceAccount,
	cr *rbacv1.ClusterRole) (*rbacv1.ClusterRoleBinding, error) {
	roleBindingName := fmt.Sprintf("sitewhere:instance:%s", swInstance.GetName())
	return &rbacv1.ClusterRoleBinding{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ClusterRoleBinding",
			APIVersion: "rbac.authorization.k8s.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: roleBindingName,
			Labels: map[string]string{
				"app": "sitewhere",
			},
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Namespace: namespace.ObjectMeta.Name,
				Name:      sa.ObjectMeta.Name,
			},
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     cr.ObjectMeta.Name,
		},
	}, nil
}

// RenderMicroservicesRoleBinding derices a RoleBinding for the Deployments of SW Instace
func RenderMicroservicesRoleBinding(swInstance *sitewhereiov1alpha4.SiteWhereInstance,
	namespace *corev1.Namespace,
	sa *corev1.ServiceAccount) (*rbacv1.RoleBinding, error) {
	roleBindingName := fmt.Sprintf("sitewhere:instance:%s", swInstance.GetName())
	return &rbacv1.RoleBinding{
		TypeMeta: metav1.TypeMeta{
			Kind:       "RoleBinding",
			APIVersion: "rbac.authorization.k8s.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      roleBindingName,
			Namespace: "sitewhere-system",
			Labels: map[string]string{
				"app": "sitewhere",
			},
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Namespace: namespace.ObjectMeta.Name,
				Name:      sa.ObjectMeta.Name,
			},
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     swInstanceRoleName,
		},
	}, nil
}

// RenderInstanceSecret renders a Secret for the instace
func RenderInstanceSecret(swInstance *sitewhereiov1alpha4.SiteWhereInstance,
	namespace *corev1.Namespace) (*corev1.Secret, error) {

	var randomSecret = rand.String(16)
	encodedClientSecret := b64.StdEncoding.EncodeToString([]byte(randomSecret))

	return &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      swInstance.GetName(),
			Namespace: namespace.GetName(),
			Labels: map[string]string{
				"app": "sitewhere",
			},
		},
		Type: corev1.SecretTypeOpaque,
		Data: map[string][]byte{
			clientSecretKey: []byte(encodedClientSecret),
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

// RenderMicroseviceFromFunctionalArea renders a SiteWhere Microservice for a Functional Area
func RenderMicroseviceFromFunctionalArea(fa funcarea.FunctionalArea, swInstance *sitewhereiov1alpha4.SiteWhereInstance, ns *corev1.Namespace) *sitewhereiov1alpha4.SiteWhereMicroservice {
	switch fa {
	case funcarea.FunctionalAreaAssetManagement:
		return &sitewhereiov1alpha4.SiteWhereMicroservice{
			ObjectMeta: metav1.ObjectMeta{
				Name:      string(funcarea.FunctionalAreaAssetManagement),
				Namespace: ns.GetName(),
				Labels: map[string]string{
					sitewhereLabelInstance:       swInstance.GetName(),
					sitewhereLabelFunctionalArea: string(funcarea.FunctionalAreaAssetManagement),
				},
			},
			Spec: sitewhereiov1alpha4.SiteWhereMicroserviceSpec{
				FunctionalArea: funcarea.FunctionalAreaAssetManagement,
				Name:           "Asset Management",
				Description:    "Manages asset model persistence and API operations",
				Icon:           "fa-laptop-house",
				Replicas:       DefaultReplica,
				Multitenant:    true,
			},
		}
	case funcarea.FunctionalAreaBatchOperations:
		return &sitewhereiov1alpha4.SiteWhereMicroservice{
			ObjectMeta: metav1.ObjectMeta{
				Name:      string(funcarea.FunctionalAreaBatchOperations),
				Namespace: ns.GetName(),
				Labels: map[string]string{
					sitewhereLabelInstance:       swInstance.GetName(),
					sitewhereLabelFunctionalArea: string(funcarea.FunctionalAreaBatchOperations),
				},
			},
			Spec: sitewhereiov1alpha4.SiteWhereMicroserviceSpec{
				FunctionalArea: funcarea.FunctionalAreaBatchOperations,
				Name:           "Batch Operations",
				Description:    "Manages operations that are applied to a list of target devices",
				Icon:           "fa-th",
				Replicas:       DefaultReplica,
				Multitenant:    true,
			},
		}
	case funcarea.FunctionalAreaCommandDelivery:
		return &sitewhereiov1alpha4.SiteWhereMicroservice{
			ObjectMeta: metav1.ObjectMeta{
				Name:      string(funcarea.FunctionalAreaCommandDelivery),
				Namespace: ns.GetName(),
				Labels: map[string]string{
					sitewhereLabelInstance:       swInstance.GetName(),
					sitewhereLabelFunctionalArea: string(funcarea.FunctionalAreaCommandDelivery),
				},
			},
			Spec: sitewhereiov1alpha4.SiteWhereMicroserviceSpec{
				FunctionalArea: funcarea.FunctionalAreaCommandDelivery,
				Name:           "Command Delivery",
				Description:    "Delivers command payloads to devices based on preset rules",
				Icon:           "fa-broadcast-tower",
				Replicas:       DefaultReplica,
				Multitenant:    true,
			},
		}
	case funcarea.FunctionalAreaDeviceManagement:
		return &sitewhereiov1alpha4.SiteWhereMicroservice{
			ObjectMeta: metav1.ObjectMeta{
				Name:      string(funcarea.FunctionalAreaDeviceManagement),
				Namespace: ns.GetName(),
				Labels: map[string]string{
					sitewhereLabelInstance:       swInstance.GetName(),
					sitewhereLabelFunctionalArea: string(funcarea.FunctionalAreaDeviceManagement),
				},
			},
			Spec: sitewhereiov1alpha4.SiteWhereMicroserviceSpec{
				FunctionalArea: funcarea.FunctionalAreaDeviceManagement,
				Name:           "Device Management",
				Description:    "Manages device model persistence and API operations",
				Icon:           "fa-microchip",
				Replicas:       DefaultReplica,
				Multitenant:    true,
			},
		}
	case funcarea.FunctionalAreaDeviceRegistration:
		return &sitewhereiov1alpha4.SiteWhereMicroservice{
			ObjectMeta: metav1.ObjectMeta{
				Name:      string(funcarea.FunctionalAreaDeviceRegistration),
				Namespace: ns.GetName(),
				Labels: map[string]string{
					sitewhereLabelInstance:       swInstance.GetName(),
					sitewhereLabelFunctionalArea: string(funcarea.FunctionalAreaDeviceRegistration),
				},
			},
			Spec: sitewhereiov1alpha4.SiteWhereMicroserviceSpec{
				FunctionalArea: funcarea.FunctionalAreaDeviceRegistration,
				Name:           "Device Registration",
				Description:    "Handles registration of new devices with the system",
				Icon:           "fa-plus-circle",
				Replicas:       DefaultReplica,
				Multitenant:    true,
			},
		}
	case funcarea.FunctionalAreaDeviceState:
		return &sitewhereiov1alpha4.SiteWhereMicroservice{
			ObjectMeta: metav1.ObjectMeta{
				Name:      string(funcarea.FunctionalAreaDeviceState),
				Namespace: ns.GetName(),
				Labels: map[string]string{
					sitewhereLabelInstance:       swInstance.GetName(),
					sitewhereLabelFunctionalArea: string(funcarea.FunctionalAreaDeviceState),
				},
			},
			Spec: sitewhereiov1alpha4.SiteWhereMicroserviceSpec{
				FunctionalArea: funcarea.FunctionalAreaDeviceState,
				Name:           "Device State",
				Description:    "Provides device state management features such as device shadows",
				Icon:           "fa-exclamation-circle",
				Replicas:       DefaultReplica,
				Multitenant:    true,
			},
		}
	case funcarea.FunctionalAreaEventManagement:
		return &sitewhereiov1alpha4.SiteWhereMicroservice{
			ObjectMeta: metav1.ObjectMeta{
				Name:      string(funcarea.FunctionalAreaEventManagement),
				Namespace: ns.GetName(),
				Labels: map[string]string{
					sitewhereLabelInstance:       swInstance.GetName(),
					sitewhereLabelFunctionalArea: string(funcarea.FunctionalAreaEventManagement),
				},
			},
			Spec: sitewhereiov1alpha4.SiteWhereMicroserviceSpec{
				FunctionalArea: funcarea.FunctionalAreaEventManagement,
				Name:           "Event Management",
				Description:    "Provides APIs for persisting and accessing events generated by devices",
				Icon:           "fa-stream",
				Replicas:       DefaultReplica,
				Multitenant:    true,
			},
		}
	case funcarea.FunctionalAreaEventSources:
		return &sitewhereiov1alpha4.SiteWhereMicroservice{
			ObjectMeta: metav1.ObjectMeta{
				Name:      string(funcarea.FunctionalAreaEventSources),
				Namespace: ns.GetName(),
				Labels: map[string]string{
					sitewhereLabelInstance:       swInstance.GetName(),
					sitewhereLabelFunctionalArea: string(funcarea.FunctionalAreaEventSources),
				},
			},
			Spec: sitewhereiov1alpha4.SiteWhereMicroserviceSpec{
				FunctionalArea: funcarea.FunctionalAreaEventSources,
				Name:           "Event Sources",
				Description:    "Handles inbound device data from various sources, protocols, and formats",
				Icon:           "fa-sign-in-alt",
				Replicas:       DefaultReplica,
				Multitenant:    true,
			},
		}
	case funcarea.FunctionalAreaInboundProcessing:
		return &sitewhereiov1alpha4.SiteWhereMicroservice{
			ObjectMeta: metav1.ObjectMeta{
				Name:      string(funcarea.FunctionalAreaInboundProcessing),
				Namespace: ns.GetName(),
				Labels: map[string]string{
					sitewhereLabelInstance:       swInstance.GetName(),
					sitewhereLabelFunctionalArea: string(funcarea.FunctionalAreaInboundProcessing),
				},
			},
			Spec: sitewhereiov1alpha4.SiteWhereMicroserviceSpec{
				FunctionalArea: funcarea.FunctionalAreaInboundProcessing,
				Name:           "Inbound Processing",
				Description:    "Common processing logic applied to enrich and direct inbound events",
				Icon:           "fa-cogs",
				Replicas:       DefaultReplica,
				Multitenant:    true,
			},
		}
	case funcarea.FunctionalAreaInstanceManagement:
		var imConfiguration = &sitewhereiov1alpha4.InstanceMangementConfiguration{
			UserManagementConfiguration: &sitewhereiov1alpha4.UserManagementConfiguration{
				SyncopeHost:            "sitewhere-syncope.sitewhere-system.cluster.local",
				SyncopePort:            8080,
				JWTExpirationInMinutes: 60,
			},
		}
		marshalledBytes, err := json.Marshal(imConfiguration)
		if err != nil {
			return nil
		}
		var instanceManagementConfiguration = &runtime.RawExtension{}
		err = instanceManagementConfiguration.UnmarshalJSON(marshalledBytes)
		if err != nil {
			return nil
		}
		return &sitewhereiov1alpha4.SiteWhereMicroservice{
			ObjectMeta: metav1.ObjectMeta{
				Name:      string(funcarea.FunctionalAreaInstanceManagement),
				Namespace: ns.GetName(),
				Labels: map[string]string{
					sitewhereLabelInstance:       swInstance.GetName(),
					sitewhereLabelFunctionalArea: string(funcarea.FunctionalAreaInstanceManagement),
				},
			},
			Spec: sitewhereiov1alpha4.SiteWhereMicroserviceSpec{
				FunctionalArea: funcarea.FunctionalAreaInstanceManagement,
				Name:           "Instance Management",
				Description:    "Handles APIs for managing global aspects of an instance",
				Icon:           "fa-sitemap",
				Replicas:       DefaultReplica,
				Multitenant:    false,
				Configuration:  instanceManagementConfiguration,
			},
		}
	case funcarea.FunctionalAreaLabelGeneration:
		return &sitewhereiov1alpha4.SiteWhereMicroservice{
			ObjectMeta: metav1.ObjectMeta{
				Name:      string(funcarea.FunctionalAreaLabelGeneration),
				Namespace: ns.GetName(),
				Labels: map[string]string{
					sitewhereLabelInstance:       swInstance.GetName(),
					sitewhereLabelFunctionalArea: string(funcarea.FunctionalAreaLabelGeneration),
				},
			},
			Spec: sitewhereiov1alpha4.SiteWhereMicroserviceSpec{
				FunctionalArea: funcarea.FunctionalAreaLabelGeneration,
				Name:           "Label Generation",
				Description:    "Supports generating labels such as bar codes and QR codes for devices",
				Icon:           "fa-qrcode",
				Replicas:       DefaultReplica,
				Multitenant:    true,
			},
		}
	case funcarea.FunctionalAreaOutboundConnectors:
		return &sitewhereiov1alpha4.SiteWhereMicroservice{
			ObjectMeta: metav1.ObjectMeta{
				Name:      string(funcarea.FunctionalAreaOutboundConnectors),
				Namespace: ns.GetName(),
				Labels: map[string]string{
					sitewhereLabelInstance:       swInstance.GetName(),
					sitewhereLabelFunctionalArea: string(funcarea.FunctionalAreaOutboundConnectors),
				},
			},
			Spec: sitewhereiov1alpha4.SiteWhereMicroserviceSpec{
				FunctionalArea: funcarea.FunctionalAreaOutboundConnectors,
				Name:           "Outbound Connectors",
				Description:    "Allows event streams to be delivered to external systems for additional processing",
				Icon:           "fa-sign-out-alt",
				Replicas:       DefaultReplica,
				Multitenant:    true,
			},
		}
	case funcarea.FunctionalAreaScheduleManagement:
		return &sitewhereiov1alpha4.SiteWhereMicroservice{
			ObjectMeta: metav1.ObjectMeta{
				Name:      string(funcarea.FunctionalAreaScheduleManagement),
				Namespace: ns.GetName(),
				Labels: map[string]string{
					sitewhereLabelInstance:       swInstance.GetName(),
					sitewhereLabelFunctionalArea: string(funcarea.FunctionalAreaScheduleManagement),
				},
			},
			Spec: sitewhereiov1alpha4.SiteWhereMicroserviceSpec{
				FunctionalArea: funcarea.FunctionalAreaScheduleManagement,
				Name:           "Schedule Management",
				Description:    "Supports scheduling of various system operations",
				Icon:           "fa-calendar-alt",
				Replicas:       DefaultReplica,
				Multitenant:    true,
			},
		}
	}
	return nil
}
