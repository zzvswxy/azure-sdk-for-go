package netapp

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ServiceLevel enumerates the values for service level.
type ServiceLevel string

const (
	// Premium Premium service level
	Premium ServiceLevel = "Premium"
	// Standard Standard service level
	Standard ServiceLevel = "Standard"
	// Ultra Ultra service level
	Ultra ServiceLevel = "Ultra"
)

// PossibleServiceLevelValues returns an array of possible values for the ServiceLevel const type.
func PossibleServiceLevelValues() []ServiceLevel {
	return []ServiceLevel{Premium, Standard, Ultra}
}