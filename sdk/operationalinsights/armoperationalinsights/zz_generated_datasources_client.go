// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoperationalinsights

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// DataSourcesClient contains the methods for the DataSources group.
// Don't use this type directly, use NewDataSourcesClient() instead.
type DataSourcesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewDataSourcesClient creates a new instance of DataSourcesClient with the specified values.
func NewDataSourcesClient(con *armcore.Connection, subscriptionID string) *DataSourcesClient {
	return &DataSourcesClient{con: con, subscriptionID: subscriptionID}
}

// CreateOrUpdate - Create or update a data source.
// If the operation fails it returns a generic error.
func (client *DataSourcesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceName string, parameters DataSource, options *DataSourcesCreateOrUpdateOptions) (DataSourcesCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, dataSourceName, parameters, options)
	if err != nil {
		return DataSourcesCreateOrUpdateResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DataSourcesCreateOrUpdateResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return DataSourcesCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DataSourcesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceName string, parameters DataSource, options *DataSourcesCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/dataSources/{dataSourceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if dataSourceName == "" {
		return nil, errors.New("parameter dataSourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataSourceName}", url.PathEscape(dataSourceName))
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
	reqQP.Set("api-version", "2020-08-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DataSourcesClient) createOrUpdateHandleResponse(resp *azcore.Response) (DataSourcesCreateOrUpdateResponse, error) {
	result := DataSourcesCreateOrUpdateResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.DataSource); err != nil {
		return DataSourcesCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *DataSourcesClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Delete - Deletes a data source instance.
// If the operation fails it returns a generic error.
func (client *DataSourcesClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceName string, options *DataSourcesDeleteOptions) (DataSourcesDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, dataSourceName, options)
	if err != nil {
		return DataSourcesDeleteResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DataSourcesDeleteResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return DataSourcesDeleteResponse{}, client.deleteHandleError(resp)
	}
	return DataSourcesDeleteResponse{RawResponse: resp.Response}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DataSourcesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceName string, options *DataSourcesDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/dataSources/{dataSourceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if dataSourceName == "" {
		return nil, errors.New("parameter dataSourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataSourceName}", url.PathEscape(dataSourceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *DataSourcesClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Get - Gets a datasource instance.
// If the operation fails it returns a generic error.
func (client *DataSourcesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceName string, options *DataSourcesGetOptions) (DataSourcesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, dataSourceName, options)
	if err != nil {
		return DataSourcesGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DataSourcesGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DataSourcesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DataSourcesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, dataSourceName string, options *DataSourcesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/dataSources/{dataSourceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if dataSourceName == "" {
		return nil, errors.New("parameter dataSourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataSourceName}", url.PathEscape(dataSourceName))
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
	reqQP.Set("api-version", "2020-08-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DataSourcesClient) getHandleResponse(resp *azcore.Response) (DataSourcesGetResponse, error) {
	result := DataSourcesGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.DataSource); err != nil {
		return DataSourcesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *DataSourcesClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListByWorkspace - Gets the first page of data source instances in a workspace with the link to the next page.
// If the operation fails it returns a generic error.
func (client *DataSourcesClient) ListByWorkspace(resourceGroupName string, workspaceName string, filter string, options *DataSourcesListByWorkspaceOptions) DataSourcesListByWorkspacePager {
	return &dataSourcesListByWorkspacePager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, filter, options)
		},
		advancer: func(ctx context.Context, resp DataSourcesListByWorkspaceResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.DataSourceListResult.NextLink)
		},
	}
}

// listByWorkspaceCreateRequest creates the ListByWorkspace request.
func (client *DataSourcesClient) listByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, filter string, options *DataSourcesListByWorkspaceOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/dataSources"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
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
	reqQP.Set("$filter", filter)
	if options != nil && options.Skiptoken != nil {
		reqQP.Set("$skiptoken", *options.Skiptoken)
	}
	reqQP.Set("api-version", "2020-08-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByWorkspaceHandleResponse handles the ListByWorkspace response.
func (client *DataSourcesClient) listByWorkspaceHandleResponse(resp *azcore.Response) (DataSourcesListByWorkspaceResponse, error) {
	result := DataSourcesListByWorkspaceResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.DataSourceListResult); err != nil {
		return DataSourcesListByWorkspaceResponse{}, err
	}
	return result, nil
}

// listByWorkspaceHandleError handles the ListByWorkspace error response.
func (client *DataSourcesClient) listByWorkspaceHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}