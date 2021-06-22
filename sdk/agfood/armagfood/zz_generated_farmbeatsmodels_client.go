// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armagfood

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// FarmBeatsModelsClient contains the methods for the FarmBeatsModels group.
// Don't use this type directly, use NewFarmBeatsModelsClient() instead.
type FarmBeatsModelsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewFarmBeatsModelsClient creates a new instance of FarmBeatsModelsClient with the specified values.
func NewFarmBeatsModelsClient(con *armcore.Connection, subscriptionID string) *FarmBeatsModelsClient {
	return &FarmBeatsModelsClient{con: con, subscriptionID: subscriptionID}
}

// CreateOrUpdate - Create or update FarmBeats resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *FarmBeatsModelsClient) CreateOrUpdate(ctx context.Context, farmBeatsResourceName string, resourceGroupName string, body FarmBeats, options *FarmBeatsModelsCreateOrUpdateOptions) (FarmBeatsResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, farmBeatsResourceName, resourceGroupName, body, options)
	if err != nil {
		return FarmBeatsResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return FarmBeatsResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return FarmBeatsResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *FarmBeatsModelsClient) createOrUpdateCreateRequest(ctx context.Context, farmBeatsResourceName string, resourceGroupName string, body FarmBeats, options *FarmBeatsModelsCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AgFoodPlatform/farmBeats/{farmBeatsResourceName}"
	if farmBeatsResourceName == "" {
		return nil, errors.New("parameter farmBeatsResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{farmBeatsResourceName}", url.PathEscape(farmBeatsResourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-05-12-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(body)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *FarmBeatsModelsClient) createOrUpdateHandleResponse(resp *azcore.Response) (FarmBeatsResponse, error) {
	var val *FarmBeats
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return FarmBeatsResponse{}, err
	}
	return FarmBeatsResponse{RawResponse: resp.Response, FarmBeats: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *FarmBeatsModelsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Delete - Delete a FarmBeats resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *FarmBeatsModelsClient) Delete(ctx context.Context, resourceGroupName string, farmBeatsResourceName string, options *FarmBeatsModelsDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, farmBeatsResourceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp.Response, nil
}

// deleteCreateRequest creates the Delete request.
func (client *FarmBeatsModelsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, farmBeatsResourceName string, options *FarmBeatsModelsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AgFoodPlatform/farmBeats/{farmBeatsResourceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if farmBeatsResourceName == "" {
		return nil, errors.New("parameter farmBeatsResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{farmBeatsResourceName}", url.PathEscape(farmBeatsResourceName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-05-12-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *FarmBeatsModelsClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Get FarmBeats resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *FarmBeatsModelsClient) Get(ctx context.Context, resourceGroupName string, farmBeatsResourceName string, options *FarmBeatsModelsGetOptions) (FarmBeatsResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, farmBeatsResourceName, options)
	if err != nil {
		return FarmBeatsResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return FarmBeatsResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return FarmBeatsResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *FarmBeatsModelsClient) getCreateRequest(ctx context.Context, resourceGroupName string, farmBeatsResourceName string, options *FarmBeatsModelsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AgFoodPlatform/farmBeats/{farmBeatsResourceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if farmBeatsResourceName == "" {
		return nil, errors.New("parameter farmBeatsResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{farmBeatsResourceName}", url.PathEscape(farmBeatsResourceName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-05-12-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FarmBeatsModelsClient) getHandleResponse(resp *azcore.Response) (FarmBeatsResponse, error) {
	var val *FarmBeats
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return FarmBeatsResponse{}, err
	}
	return FarmBeatsResponse{RawResponse: resp.Response, FarmBeats: val}, nil
}

// getHandleError handles the Get error response.
func (client *FarmBeatsModelsClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListByResourceGroup - Lists the FarmBeats instances for a resource group.
// If the operation fails it returns the *ErrorResponse error type.
func (client *FarmBeatsModelsClient) ListByResourceGroup(resourceGroupName string, options *FarmBeatsModelsListByResourceGroupOptions) FarmBeatsListResponsePager {
	return &farmBeatsListResponsePager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listByResourceGroupHandleResponse,
		errorer:   client.listByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp FarmBeatsListResponseResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.FarmBeatsListResponse.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *FarmBeatsModelsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *FarmBeatsModelsListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AgFoodPlatform/farmBeats"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.MaxPageSize != nil {
		reqQP.Set("$maxPageSize", strconv.FormatInt(int64(*options.MaxPageSize), 10))
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	reqQP.Set("api-version", "2020-05-12-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *FarmBeatsModelsClient) listByResourceGroupHandleResponse(resp *azcore.Response) (FarmBeatsListResponseResponse, error) {
	var val *FarmBeatsListResponse
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return FarmBeatsListResponseResponse{}, err
	}
	return FarmBeatsListResponseResponse{RawResponse: resp.Response, FarmBeatsListResponse: val}, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *FarmBeatsModelsClient) listByResourceGroupHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListBySubscription - Lists the FarmBeats instances for a subscription.
// If the operation fails it returns the *ErrorResponse error type.
func (client *FarmBeatsModelsClient) ListBySubscription(options *FarmBeatsModelsListBySubscriptionOptions) FarmBeatsListResponsePager {
	return &farmBeatsListResponsePager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		responder: client.listBySubscriptionHandleResponse,
		errorer:   client.listBySubscriptionHandleError,
		advancer: func(ctx context.Context, resp FarmBeatsListResponseResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.FarmBeatsListResponse.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *FarmBeatsModelsClient) listBySubscriptionCreateRequest(ctx context.Context, options *FarmBeatsModelsListBySubscriptionOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AgFoodPlatform/farmBeats"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.MaxPageSize != nil {
		reqQP.Set("$maxPageSize", strconv.FormatInt(int64(*options.MaxPageSize), 10))
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	reqQP.Set("api-version", "2020-05-12-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *FarmBeatsModelsClient) listBySubscriptionHandleResponse(resp *azcore.Response) (FarmBeatsListResponseResponse, error) {
	var val *FarmBeatsListResponse
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return FarmBeatsListResponseResponse{}, err
	}
	return FarmBeatsListResponseResponse{RawResponse: resp.Response, FarmBeatsListResponse: val}, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *FarmBeatsModelsClient) listBySubscriptionHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Update - Update a FarmBeats resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *FarmBeatsModelsClient) Update(ctx context.Context, farmBeatsResourceName string, resourceGroupName string, body FarmBeatsUpdateRequestModel, options *FarmBeatsModelsUpdateOptions) (FarmBeatsResponse, error) {
	req, err := client.updateCreateRequest(ctx, farmBeatsResourceName, resourceGroupName, body, options)
	if err != nil {
		return FarmBeatsResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return FarmBeatsResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return FarmBeatsResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *FarmBeatsModelsClient) updateCreateRequest(ctx context.Context, farmBeatsResourceName string, resourceGroupName string, body FarmBeatsUpdateRequestModel, options *FarmBeatsModelsUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AgFoodPlatform/farmBeats/{farmBeatsResourceName}"
	if farmBeatsResourceName == "" {
		return nil, errors.New("parameter farmBeatsResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{farmBeatsResourceName}", url.PathEscape(farmBeatsResourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-05-12-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(body)
}

// updateHandleResponse handles the Update response.
func (client *FarmBeatsModelsClient) updateHandleResponse(resp *azcore.Response) (FarmBeatsResponse, error) {
	var val *FarmBeats
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return FarmBeatsResponse{}, err
	}
	return FarmBeatsResponse{RawResponse: resp.Response, FarmBeats: val}, nil
}

// updateHandleError handles the Update error response.
func (client *FarmBeatsModelsClient) updateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}