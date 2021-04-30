package appplatform

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AppResourceProvisioningState enumerates the values for app resource provisioning state.
type AppResourceProvisioningState string

const (
	// Creating ...
	Creating AppResourceProvisioningState = "Creating"
	// Failed ...
	Failed AppResourceProvisioningState = "Failed"
	// Succeeded ...
	Succeeded AppResourceProvisioningState = "Succeeded"
	// Updating ...
	Updating AppResourceProvisioningState = "Updating"
)

// PossibleAppResourceProvisioningStateValues returns an array of possible values for the AppResourceProvisioningState const type.
func PossibleAppResourceProvisioningStateValues() []AppResourceProvisioningState {
	return []AppResourceProvisioningState{Creating, Failed, Succeeded, Updating}
}

// ConfigServerState enumerates the values for config server state.
type ConfigServerState string

const (
	// ConfigServerStateDeleted ...
	ConfigServerStateDeleted ConfigServerState = "Deleted"
	// ConfigServerStateFailed ...
	ConfigServerStateFailed ConfigServerState = "Failed"
	// ConfigServerStateNotAvailable ...
	ConfigServerStateNotAvailable ConfigServerState = "NotAvailable"
	// ConfigServerStateSucceeded ...
	ConfigServerStateSucceeded ConfigServerState = "Succeeded"
	// ConfigServerStateUpdating ...
	ConfigServerStateUpdating ConfigServerState = "Updating"
)

// PossibleConfigServerStateValues returns an array of possible values for the ConfigServerState const type.
func PossibleConfigServerStateValues() []ConfigServerState {
	return []ConfigServerState{ConfigServerStateDeleted, ConfigServerStateFailed, ConfigServerStateNotAvailable, ConfigServerStateSucceeded, ConfigServerStateUpdating}
}

// DeploymentResourceProvisioningState enumerates the values for deployment resource provisioning state.
type DeploymentResourceProvisioningState string

const (
	// DeploymentResourceProvisioningStateCreating ...
	DeploymentResourceProvisioningStateCreating DeploymentResourceProvisioningState = "Creating"
	// DeploymentResourceProvisioningStateFailed ...
	DeploymentResourceProvisioningStateFailed DeploymentResourceProvisioningState = "Failed"
	// DeploymentResourceProvisioningStateSucceeded ...
	DeploymentResourceProvisioningStateSucceeded DeploymentResourceProvisioningState = "Succeeded"
	// DeploymentResourceProvisioningStateUpdating ...
	DeploymentResourceProvisioningStateUpdating DeploymentResourceProvisioningState = "Updating"
)

// PossibleDeploymentResourceProvisioningStateValues returns an array of possible values for the DeploymentResourceProvisioningState const type.
func PossibleDeploymentResourceProvisioningStateValues() []DeploymentResourceProvisioningState {
	return []DeploymentResourceProvisioningState{DeploymentResourceProvisioningStateCreating, DeploymentResourceProvisioningStateFailed, DeploymentResourceProvisioningStateSucceeded, DeploymentResourceProvisioningStateUpdating}
}

// DeploymentResourceStatus enumerates the values for deployment resource status.
type DeploymentResourceStatus string

const (
	// DeploymentResourceStatusAllocating ...
	DeploymentResourceStatusAllocating DeploymentResourceStatus = "Allocating"
	// DeploymentResourceStatusCompiling ...
	DeploymentResourceStatusCompiling DeploymentResourceStatus = "Compiling"
	// DeploymentResourceStatusFailed ...
	DeploymentResourceStatusFailed DeploymentResourceStatus = "Failed"
	// DeploymentResourceStatusRunning ...
	DeploymentResourceStatusRunning DeploymentResourceStatus = "Running"
	// DeploymentResourceStatusStopped ...
	DeploymentResourceStatusStopped DeploymentResourceStatus = "Stopped"
	// DeploymentResourceStatusUnknown ...
	DeploymentResourceStatusUnknown DeploymentResourceStatus = "Unknown"
	// DeploymentResourceStatusUpgrading ...
	DeploymentResourceStatusUpgrading DeploymentResourceStatus = "Upgrading"
)

// PossibleDeploymentResourceStatusValues returns an array of possible values for the DeploymentResourceStatus const type.
func PossibleDeploymentResourceStatusValues() []DeploymentResourceStatus {
	return []DeploymentResourceStatus{DeploymentResourceStatusAllocating, DeploymentResourceStatusCompiling, DeploymentResourceStatusFailed, DeploymentResourceStatusRunning, DeploymentResourceStatusStopped, DeploymentResourceStatusUnknown, DeploymentResourceStatusUpgrading}
}

// ManagedIdentityType enumerates the values for managed identity type.
type ManagedIdentityType string

const (
	// None ...
	None ManagedIdentityType = "None"
	// SystemAssigned ...
	SystemAssigned ManagedIdentityType = "SystemAssigned"
	// SystemAssignedUserAssigned ...
	SystemAssignedUserAssigned ManagedIdentityType = "SystemAssigned,UserAssigned"
	// UserAssigned ...
	UserAssigned ManagedIdentityType = "UserAssigned"
)

// PossibleManagedIdentityTypeValues returns an array of possible values for the ManagedIdentityType const type.
func PossibleManagedIdentityTypeValues() []ManagedIdentityType {
	return []ManagedIdentityType{None, SystemAssigned, SystemAssignedUserAssigned, UserAssigned}
}

// MonitoringSettingState enumerates the values for monitoring setting state.
type MonitoringSettingState string

const (
	// MonitoringSettingStateFailed ...
	MonitoringSettingStateFailed MonitoringSettingState = "Failed"
	// MonitoringSettingStateNotAvailable ...
	MonitoringSettingStateNotAvailable MonitoringSettingState = "NotAvailable"
	// MonitoringSettingStateSucceeded ...
	MonitoringSettingStateSucceeded MonitoringSettingState = "Succeeded"
	// MonitoringSettingStateUpdating ...
	MonitoringSettingStateUpdating MonitoringSettingState = "Updating"
)

// PossibleMonitoringSettingStateValues returns an array of possible values for the MonitoringSettingState const type.
func PossibleMonitoringSettingStateValues() []MonitoringSettingState {
	return []MonitoringSettingState{MonitoringSettingStateFailed, MonitoringSettingStateNotAvailable, MonitoringSettingStateSucceeded, MonitoringSettingStateUpdating}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// ProvisioningStateCreating ...
	ProvisioningStateCreating ProvisioningState = "Creating"
	// ProvisioningStateDeleted ...
	ProvisioningStateDeleted ProvisioningState = "Deleted"
	// ProvisioningStateDeleting ...
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed ...
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateMoved ...
	ProvisioningStateMoved ProvisioningState = "Moved"
	// ProvisioningStateMoveFailed ...
	ProvisioningStateMoveFailed ProvisioningState = "MoveFailed"
	// ProvisioningStateMoving ...
	ProvisioningStateMoving ProvisioningState = "Moving"
	// ProvisioningStateSucceeded ...
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUpdating ...
	ProvisioningStateUpdating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{ProvisioningStateCreating, ProvisioningStateDeleted, ProvisioningStateDeleting, ProvisioningStateFailed, ProvisioningStateMoved, ProvisioningStateMoveFailed, ProvisioningStateMoving, ProvisioningStateSucceeded, ProvisioningStateUpdating}
}

// ResourceSkuRestrictionsReasonCode enumerates the values for resource sku restrictions reason code.
type ResourceSkuRestrictionsReasonCode string

const (
	// NotAvailableForSubscription ...
	NotAvailableForSubscription ResourceSkuRestrictionsReasonCode = "NotAvailableForSubscription"
	// QuotaID ...
	QuotaID ResourceSkuRestrictionsReasonCode = "QuotaId"
)

// PossibleResourceSkuRestrictionsReasonCodeValues returns an array of possible values for the ResourceSkuRestrictionsReasonCode const type.
func PossibleResourceSkuRestrictionsReasonCodeValues() []ResourceSkuRestrictionsReasonCode {
	return []ResourceSkuRestrictionsReasonCode{NotAvailableForSubscription, QuotaID}
}

// ResourceSkuRestrictionsType enumerates the values for resource sku restrictions type.
type ResourceSkuRestrictionsType string

const (
	// Location ...
	Location ResourceSkuRestrictionsType = "Location"
	// Zone ...
	Zone ResourceSkuRestrictionsType = "Zone"
)

// PossibleResourceSkuRestrictionsTypeValues returns an array of possible values for the ResourceSkuRestrictionsType const type.
func PossibleResourceSkuRestrictionsTypeValues() []ResourceSkuRestrictionsType {
	return []ResourceSkuRestrictionsType{Location, Zone}
}

// RuntimeVersion enumerates the values for runtime version.
type RuntimeVersion string

const (
	// Java11 ...
	Java11 RuntimeVersion = "Java_11"
	// Java8 ...
	Java8 RuntimeVersion = "Java_8"
	// NetCore31 ...
	NetCore31 RuntimeVersion = "NetCore_31"
)

// PossibleRuntimeVersionValues returns an array of possible values for the RuntimeVersion const type.
func PossibleRuntimeVersionValues() []RuntimeVersion {
	return []RuntimeVersion{Java11, Java8, NetCore31}
}

// SkuScaleType enumerates the values for sku scale type.
type SkuScaleType string

const (
	// SkuScaleTypeAutomatic ...
	SkuScaleTypeAutomatic SkuScaleType = "Automatic"
	// SkuScaleTypeManual ...
	SkuScaleTypeManual SkuScaleType = "Manual"
	// SkuScaleTypeNone ...
	SkuScaleTypeNone SkuScaleType = "None"
)

// PossibleSkuScaleTypeValues returns an array of possible values for the SkuScaleType const type.
func PossibleSkuScaleTypeValues() []SkuScaleType {
	return []SkuScaleType{SkuScaleTypeAutomatic, SkuScaleTypeManual, SkuScaleTypeNone}
}

// SupportedRuntimePlatform enumerates the values for supported runtime platform.
type SupportedRuntimePlatform string

const (
	// Java ...
	Java SupportedRuntimePlatform = "Java"
	// NETCore ...
	NETCore SupportedRuntimePlatform = ".NET Core"
)

// PossibleSupportedRuntimePlatformValues returns an array of possible values for the SupportedRuntimePlatform const type.
func PossibleSupportedRuntimePlatformValues() []SupportedRuntimePlatform {
	return []SupportedRuntimePlatform{Java, NETCore}
}

// SupportedRuntimeValue enumerates the values for supported runtime value.
type SupportedRuntimeValue string

const (
	// SupportedRuntimeValueJava11 ...
	SupportedRuntimeValueJava11 SupportedRuntimeValue = "Java_11"
	// SupportedRuntimeValueJava8 ...
	SupportedRuntimeValueJava8 SupportedRuntimeValue = "Java_8"
	// SupportedRuntimeValueNetCore31 ...
	SupportedRuntimeValueNetCore31 SupportedRuntimeValue = "NetCore_31"
)

// PossibleSupportedRuntimeValueValues returns an array of possible values for the SupportedRuntimeValue const type.
func PossibleSupportedRuntimeValueValues() []SupportedRuntimeValue {
	return []SupportedRuntimeValue{SupportedRuntimeValueJava11, SupportedRuntimeValueJava8, SupportedRuntimeValueNetCore31}
}

// TestKeyType enumerates the values for test key type.
type TestKeyType string

const (
	// Primary ...
	Primary TestKeyType = "Primary"
	// Secondary ...
	Secondary TestKeyType = "Secondary"
)

// PossibleTestKeyTypeValues returns an array of possible values for the TestKeyType const type.
func PossibleTestKeyTypeValues() []TestKeyType {
	return []TestKeyType{Primary, Secondary}
}

// TrafficDirection enumerates the values for traffic direction.
type TrafficDirection string

const (
	// Inbound ...
	Inbound TrafficDirection = "Inbound"
	// Outbound ...
	Outbound TrafficDirection = "Outbound"
)

// PossibleTrafficDirectionValues returns an array of possible values for the TrafficDirection const type.
func PossibleTrafficDirectionValues() []TrafficDirection {
	return []TrafficDirection{Inbound, Outbound}
}

// UserSourceType enumerates the values for user source type.
type UserSourceType string

const (
	// Jar ...
	Jar UserSourceType = "Jar"
	// NetCoreZip ...
	NetCoreZip UserSourceType = "NetCoreZip"
	// Source ...
	Source UserSourceType = "Source"
)

// PossibleUserSourceTypeValues returns an array of possible values for the UserSourceType const type.
func PossibleUserSourceTypeValues() []UserSourceType {
	return []UserSourceType{Jar, NetCoreZip, Source}
}