//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package generated

import "time"

// AzureAppConfigurationClientCheckKeyValueOptions contains the optional parameters for the AzureAppConfigurationClient.CheckKeyValue
// method.
type AzureAppConfigurationClientCheckKeyValueOptions struct {
	// Requests the server to respond with the state of the resource at the specified time.
	AcceptDatetime *string
	// Used to perform an operation only if the targeted resource's etag matches the value provided.
	IfMatch *string
	// Used to perform an operation only if the targeted resource's etag does not match the value provided.
	IfNoneMatch *string
	// The label of the key-value to retrieve.
	Label *string
	// Used to select what fields are present in the returned resource(s).
	Select []Head7ItemsItem
}

// AzureAppConfigurationClientCheckKeyValuesOptions contains the optional parameters for the AzureAppConfigurationClient.CheckKeyValues
// method.
type AzureAppConfigurationClientCheckKeyValuesOptions struct {
	// Requests the server to respond with the state of the resource at the specified time.
	AcceptDatetime *string
	// Instructs the server to return elements that appear after the element referred to by the specified token.
	After *string
	// A filter used to match keys.
	Key *string
	// A filter used to match labels
	Label *string
	// Used to select what fields are present in the returned resource(s).
	Select []Head6ItemsItem
}

// AzureAppConfigurationClientCheckKeysOptions contains the optional parameters for the AzureAppConfigurationClient.CheckKeys
// method.
type AzureAppConfigurationClientCheckKeysOptions struct {
	// Requests the server to respond with the state of the resource at the specified time.
	AcceptDatetime *string
	// Instructs the server to return elements that appear after the element referred to by the specified token.
	After *string
	// A filter for the name of the returned keys.
	Name *string
}

// AzureAppConfigurationClientCheckLabelsOptions contains the optional parameters for the AzureAppConfigurationClient.CheckLabels
// method.
type AzureAppConfigurationClientCheckLabelsOptions struct {
	// Requests the server to respond with the state of the resource at the specified time.
	AcceptDatetime *string
	// Instructs the server to return elements that appear after the element referred to by the specified token.
	After *string
	// A filter for the name of the returned labels.
	Name *string
	// Used to select what fields are present in the returned resource(s).
	Select []Head5ItemsItem
}

// AzureAppConfigurationClientCheckRevisionsOptions contains the optional parameters for the AzureAppConfigurationClient.CheckRevisions
// method.
type AzureAppConfigurationClientCheckRevisionsOptions struct {
	// Requests the server to respond with the state of the resource at the specified time.
	AcceptDatetime *string
	// Instructs the server to return elements that appear after the element referred to by the specified token.
	After *string
	// A filter used to match keys.
	Key *string
	// A filter used to match labels
	Label *string
	// Used to select what fields are present in the returned resource(s).
	Select []Enum7
}

// AzureAppConfigurationClientDeleteKeyValueOptions contains the optional parameters for the AzureAppConfigurationClient.DeleteKeyValue
// method.
type AzureAppConfigurationClientDeleteKeyValueOptions struct {
	// Used to perform an operation only if the targeted resource's etag matches the value provided.
	IfMatch *string
	// The label of the key-value to delete.
	Label *string
}

// AzureAppConfigurationClientDeleteLockOptions contains the optional parameters for the AzureAppConfigurationClient.DeleteLock
// method.
type AzureAppConfigurationClientDeleteLockOptions struct {
	// Used to perform an operation only if the targeted resource's etag matches the value provided.
	IfMatch *string
	// Used to perform an operation only if the targeted resource's etag does not match the value provided.
	IfNoneMatch *string
	// The label, if any, of the key-value to unlock.
	Label *string
}

// AzureAppConfigurationClientGetKeyValueOptions contains the optional parameters for the AzureAppConfigurationClient.GetKeyValue
// method.
type AzureAppConfigurationClientGetKeyValueOptions struct {
	// Requests the server to respond with the state of the resource at the specified time.
	AcceptDatetime *string
	// Used to perform an operation only if the targeted resource's etag matches the value provided.
	IfMatch *string
	// Used to perform an operation only if the targeted resource's etag does not match the value provided.
	IfNoneMatch *string
	// The label of the key-value to retrieve.
	Label *string
	// Used to select what fields are present in the returned resource(s).
	Select []Get7ItemsItem
}

// AzureAppConfigurationClientGetKeyValuesOptions contains the optional parameters for the AzureAppConfigurationClient.GetKeyValues
// method.
type AzureAppConfigurationClientGetKeyValuesOptions struct {
	// Requests the server to respond with the state of the resource at the specified time.
	AcceptDatetime *string
	// Instructs the server to return elements that appear after the element referred to by the specified token.
	After *string
	// A filter used to match keys.
	Key *string
	// A filter used to match labels
	Label *string
	// Used to select what fields are present in the returned resource(s).
	Select []Get6ItemsItem
}

// AzureAppConfigurationClientGetKeysOptions contains the optional parameters for the AzureAppConfigurationClient.GetKeys
// method.
type AzureAppConfigurationClientGetKeysOptions struct {
	// Requests the server to respond with the state of the resource at the specified time.
	AcceptDatetime *string
	// Instructs the server to return elements that appear after the element referred to by the specified token.
	After *string
	// A filter for the name of the returned keys.
	Name *string
}

// AzureAppConfigurationClientGetLabelsOptions contains the optional parameters for the AzureAppConfigurationClient.GetLabels
// method.
type AzureAppConfigurationClientGetLabelsOptions struct {
	// Requests the server to respond with the state of the resource at the specified time.
	AcceptDatetime *string
	// Instructs the server to return elements that appear after the element referred to by the specified token.
	After *string
	// A filter for the name of the returned labels.
	Name *string
	// Used to select what fields are present in the returned resource(s).
	Select []Get5ItemsItem
}

// AzureAppConfigurationClientGetRevisionsOptions contains the optional parameters for the AzureAppConfigurationClient.GetRevisions
// method.
type AzureAppConfigurationClientGetRevisionsOptions struct {
	// Requests the server to respond with the state of the resource at the specified time.
	AcceptDatetime *string
	// Instructs the server to return elements that appear after the element referred to by the specified token.
	After *string
	// A filter used to match keys.
	Key *string
	// A filter used to match labels
	Label *string
	// Used to select what fields are present in the returned resource(s).
	Select []Enum6
}

// AzureAppConfigurationClientPutKeyValueOptions contains the optional parameters for the AzureAppConfigurationClient.PutKeyValue
// method.
type AzureAppConfigurationClientPutKeyValueOptions struct {
	// The key-value to create.
	Entity *KeyValue
	// Used to perform an operation only if the targeted resource's etag matches the value provided.
	IfMatch *string
	// Used to perform an operation only if the targeted resource's etag does not match the value provided.
	IfNoneMatch *string
	// The label of the key-value to create.
	Label *string
}

// AzureAppConfigurationClientPutLockOptions contains the optional parameters for the AzureAppConfigurationClient.PutLock
// method.
type AzureAppConfigurationClientPutLockOptions struct {
	// Used to perform an operation only if the targeted resource's etag matches the value provided.
	IfMatch *string
	// Used to perform an operation only if the targeted resource's etag does not match the value provided.
	IfNoneMatch *string
	// The label, if any, of the key-value to lock.
	Label *string
}

// Error - Azure App Configuration error object.
type Error struct {
	// A detailed description of the error.
	Detail *string `json:"detail,omitempty"`

	// The name of the parameter that resulted in the error.
	Name *string `json:"name,omitempty"`

	// The HTTP status code that the error maps to.
	Status *int32 `json:"status,omitempty"`

	// A brief summary of the error.
	Title *string `json:"title,omitempty"`

	// The type of the error.
	Type *string `json:"type,omitempty"`
}

type Key struct {
	// READ-ONLY
	Name *string `json:"name,omitempty" azure:"ro"`
}

// KeyListResult - The result of a list request.
type KeyListResult struct {
	// The collection value.
	Items []*Key `json:"items,omitempty"`

	// The URI that can be used to request the next set of paged results.
	NextLink *string `json:"@nextLink,omitempty"`
}

type KeyValue struct {
	ContentType  *string    `json:"content_type,omitempty"`
	Etag         *string    `json:"etag,omitempty"`
	Key          *string    `json:"key,omitempty"`
	Label        *string    `json:"label,omitempty"`
	LastModified *time.Time `json:"last_modified,omitempty"`
	Locked       *bool      `json:"locked,omitempty"`

	// Dictionary of
	Tags  map[string]*string `json:"tags,omitempty"`
	Value *string            `json:"value,omitempty"`
}

// KeyValueListResult - The result of a list request.
type KeyValueListResult struct {
	// The collection value.
	Items []*KeyValue `json:"items,omitempty"`

	// The URI that can be used to request the next set of paged results.
	NextLink *string `json:"@nextLink,omitempty"`
}

type Label struct {
	// READ-ONLY
	Name *string `json:"name,omitempty" azure:"ro"`
}

// LabelListResult - The result of a list request.
type LabelListResult struct {
	// The collection value.
	Items []*Label `json:"items,omitempty"`

	// The URI that can be used to request the next set of paged results.
	NextLink *string `json:"@nextLink,omitempty"`
}
