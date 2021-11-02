//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SoftwareUpdateConfigurationsClient contains the methods for the SoftwareUpdateConfigurations group.
// Don't use this type directly, use NewSoftwareUpdateConfigurationsClient() instead.
type SoftwareUpdateConfigurationsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewSoftwareUpdateConfigurationsClient creates a new instance of SoftwareUpdateConfigurationsClient with the specified values.
func NewSoftwareUpdateConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SoftwareUpdateConfigurationsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &SoftwareUpdateConfigurationsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Create - Create a new software update configuration with the name given in the URI.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SoftwareUpdateConfigurationsClient) Create(ctx context.Context, resourceGroupName string, automationAccountName string, softwareUpdateConfigurationName string, parameters SoftwareUpdateConfiguration, options *SoftwareUpdateConfigurationsCreateOptions) (SoftwareUpdateConfigurationsCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, automationAccountName, softwareUpdateConfigurationName, parameters, options)
	if err != nil {
		return SoftwareUpdateConfigurationsCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SoftwareUpdateConfigurationsCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return SoftwareUpdateConfigurationsCreateResponse{}, client.createHandleError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *SoftwareUpdateConfigurationsClient) createCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, softwareUpdateConfigurationName string, parameters SoftwareUpdateConfiguration, options *SoftwareUpdateConfigurationsCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/softwareUpdateConfigurations/{softwareUpdateConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if softwareUpdateConfigurationName == "" {
		return nil, errors.New("parameter softwareUpdateConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{softwareUpdateConfigurationName}", url.PathEscape(softwareUpdateConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.ClientRequestID != nil {
		req.Raw().Header.Set("clientRequestId", *options.ClientRequestID)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createHandleResponse handles the Create response.
func (client *SoftwareUpdateConfigurationsClient) createHandleResponse(resp *http.Response) (SoftwareUpdateConfigurationsCreateResponse, error) {
	result := SoftwareUpdateConfigurationsCreateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SoftwareUpdateConfiguration); err != nil {
		return SoftwareUpdateConfigurationsCreateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createHandleError handles the Create error response.
func (client *SoftwareUpdateConfigurationsClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - delete a specific software update configuration.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SoftwareUpdateConfigurationsClient) Delete(ctx context.Context, resourceGroupName string, automationAccountName string, softwareUpdateConfigurationName string, options *SoftwareUpdateConfigurationsDeleteOptions) (SoftwareUpdateConfigurationsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, automationAccountName, softwareUpdateConfigurationName, options)
	if err != nil {
		return SoftwareUpdateConfigurationsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SoftwareUpdateConfigurationsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return SoftwareUpdateConfigurationsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return SoftwareUpdateConfigurationsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SoftwareUpdateConfigurationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, softwareUpdateConfigurationName string, options *SoftwareUpdateConfigurationsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/softwareUpdateConfigurations/{softwareUpdateConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if softwareUpdateConfigurationName == "" {
		return nil, errors.New("parameter softwareUpdateConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{softwareUpdateConfigurationName}", url.PathEscape(softwareUpdateConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.ClientRequestID != nil {
		req.Raw().Header.Set("clientRequestId", *options.ClientRequestID)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *SoftwareUpdateConfigurationsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetByName - Get a single software update configuration by name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SoftwareUpdateConfigurationsClient) GetByName(ctx context.Context, resourceGroupName string, automationAccountName string, softwareUpdateConfigurationName string, options *SoftwareUpdateConfigurationsGetByNameOptions) (SoftwareUpdateConfigurationsGetByNameResponse, error) {
	req, err := client.getByNameCreateRequest(ctx, resourceGroupName, automationAccountName, softwareUpdateConfigurationName, options)
	if err != nil {
		return SoftwareUpdateConfigurationsGetByNameResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SoftwareUpdateConfigurationsGetByNameResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SoftwareUpdateConfigurationsGetByNameResponse{}, client.getByNameHandleError(resp)
	}
	return client.getByNameHandleResponse(resp)
}

// getByNameCreateRequest creates the GetByName request.
func (client *SoftwareUpdateConfigurationsClient) getByNameCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, softwareUpdateConfigurationName string, options *SoftwareUpdateConfigurationsGetByNameOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/softwareUpdateConfigurations/{softwareUpdateConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if softwareUpdateConfigurationName == "" {
		return nil, errors.New("parameter softwareUpdateConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{softwareUpdateConfigurationName}", url.PathEscape(softwareUpdateConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.ClientRequestID != nil {
		req.Raw().Header.Set("clientRequestId", *options.ClientRequestID)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getByNameHandleResponse handles the GetByName response.
func (client *SoftwareUpdateConfigurationsClient) getByNameHandleResponse(resp *http.Response) (SoftwareUpdateConfigurationsGetByNameResponse, error) {
	result := SoftwareUpdateConfigurationsGetByNameResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SoftwareUpdateConfiguration); err != nil {
		return SoftwareUpdateConfigurationsGetByNameResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getByNameHandleError handles the GetByName error response.
func (client *SoftwareUpdateConfigurationsClient) getByNameHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Get all software update configurations for the account.
// If the operation fails it returns the *ErrorResponse error type.
func (client *SoftwareUpdateConfigurationsClient) List(ctx context.Context, resourceGroupName string, automationAccountName string, options *SoftwareUpdateConfigurationsListOptions) (SoftwareUpdateConfigurationsListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, automationAccountName, options)
	if err != nil {
		return SoftwareUpdateConfigurationsListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SoftwareUpdateConfigurationsListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SoftwareUpdateConfigurationsListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *SoftwareUpdateConfigurationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *SoftwareUpdateConfigurationsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/softwareUpdateConfigurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.ClientRequestID != nil {
		req.Raw().Header.Set("clientRequestId", *options.ClientRequestID)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SoftwareUpdateConfigurationsClient) listHandleResponse(resp *http.Response) (SoftwareUpdateConfigurationsListResponse, error) {
	result := SoftwareUpdateConfigurationsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SoftwareUpdateConfigurationListResult); err != nil {
		return SoftwareUpdateConfigurationsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *SoftwareUpdateConfigurationsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}