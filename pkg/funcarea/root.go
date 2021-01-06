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

package funcarea

// FunctionalArea are the SiteWhere Functional Areas
type FunctionalArea string

const (
	//FunctionalAreaAssetManagement is the name of the functional area Asset Management
	FunctionalAreaAssetManagement FunctionalArea = "asset-management"
	//FunctionalAreaBatchOperations is the name of the functional area Batch Operations
	FunctionalAreaBatchOperations FunctionalArea = "batch-operations"
	//FunctionalAreaCommandDelivery is the name of the functional area Command Delivery
	FunctionalAreaCommandDelivery FunctionalArea = "command-delivery"
	//FunctionalAreaDeviceManagement is the name of the functional area Device Management
	FunctionalAreaDeviceManagement FunctionalArea = "device-management"
	//FunctionalAreaDeviceRegistration is the name of the functional area Device Registration
	FunctionalAreaDeviceRegistration FunctionalArea = "device-registration"
	//FunctionalAreaDeviceState is the name of the functional area Device State
	FunctionalAreaDeviceState FunctionalArea = "device-state"
	//FunctionalAreaEventManagement is the name of the functional area Event Management
	FunctionalAreaEventManagement FunctionalArea = "event-management"
	//FunctionalAreaEventSources is the name of the functional area Event Sources
	FunctionalAreaEventSources FunctionalArea = "event-sources"
	//FunctionalAreaInboundProcessing is the name of the functional area Inbound Processing
	FunctionalAreaInboundProcessing FunctionalArea = "inbound-processing"
	//FunctionalAreaInstanceManagement is the name of the functional area Instance Management
	FunctionalAreaInstanceManagement FunctionalArea = "instance-management"
	//FunctionalAreaLabelGeneration is the name of the functional area Label Generation
	FunctionalAreaLabelGeneration FunctionalArea = "label-generation"
	//FunctionalAreaOutboundConnectors is the name of the functional area Outbound Connectors
	FunctionalAreaOutboundConnectors FunctionalArea = "outbound-connectors"
	//FunctionalAreaScheduleManagement is the name of the functional area Schedule Management
	FunctionalAreaScheduleManagement FunctionalArea = "schedule-management"
)

var (
	// DefaultFunctionalAreas are the default Functional Areas of SiteWhere
	DefaultFunctionalAreas = []FunctionalArea{
		FunctionalAreaAssetManagement,
		FunctionalAreaBatchOperations,
		FunctionalAreaCommandDelivery,
		FunctionalAreaDeviceManagement,
		FunctionalAreaDeviceRegistration,
		FunctionalAreaDeviceState,
		FunctionalAreaEventManagement,
		FunctionalAreaEventSources,
		FunctionalAreaInboundProcessing,
		FunctionalAreaInstanceManagement,
		FunctionalAreaLabelGeneration,
		FunctionalAreaOutboundConnectors,
		FunctionalAreaScheduleManagement,
	}
)

// HasFunctionalArea checks is an array of functional areas has a functional area
func HasFunctionalArea(items []FunctionalArea, area FunctionalArea) bool {
	for _, item := range items {
		if item == area {
			return true
		}
	}
	return false
}
