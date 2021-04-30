// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresources

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// TagsClient contains the methods for the Tags group.
// Don't use this type directly, use NewTagsClient() instead.
type TagsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewTagsClient creates a new instance of TagsClient with the specified values.
func NewTagsClient(con *armcore.Connection, subscriptionID string) *TagsClient {
	return &TagsClient{con: con, subscriptionID: subscriptionID}
}

// CreateOrUpdate - This operation allows adding a name to the list of predefined tag names for the given subscription. A tag name can have a maximum of
// 512 characters and is case-insensitive. Tag names cannot have the
// following prefixes which are reserved for Azure use: 'microsoft', 'azure', 'windows'.
func (client *TagsClient) CreateOrUpdate(ctx context.Context, tagName string, options *TagsCreateOrUpdateOptions) (TagDetailsResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, tagName, options)
	if err != nil {
		return TagDetailsResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TagDetailsResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return TagDetailsResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *TagsClient) createOrUpdateCreateRequest(ctx context.Context, tagName string, options *TagsCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/tagNames/{tagName}"
	urlPath = strings.ReplaceAll(urlPath, "{tagName}", url.PathEscape(tagName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *TagsClient) createOrUpdateHandleResponse(resp *azcore.Response) (TagDetailsResponse, error) {
	var val *TagDetails
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return TagDetailsResponse{}, err
	}
	return TagDetailsResponse{RawResponse: resp.Response, TagDetails: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *TagsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// CreateOrUpdateAtScope - This operation allows adding or replacing the entire set of tags on the specified resource or subscription. The specified entity
// can have a maximum of 50 tags.
func (client *TagsClient) CreateOrUpdateAtScope(ctx context.Context, scope string, parameters TagsResource, options *TagsCreateOrUpdateAtScopeOptions) (TagsResourceResponse, error) {
	req, err := client.createOrUpdateAtScopeCreateRequest(ctx, scope, parameters, options)
	if err != nil {
		return TagsResourceResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TagsResourceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return TagsResourceResponse{}, client.createOrUpdateAtScopeHandleError(resp)
	}
	return client.createOrUpdateAtScopeHandleResponse(resp)
}

// createOrUpdateAtScopeCreateRequest creates the CreateOrUpdateAtScope request.
func (client *TagsClient) createOrUpdateAtScopeCreateRequest(ctx context.Context, scope string, parameters TagsResource, options *TagsCreateOrUpdateAtScopeOptions) (*azcore.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Resources/tags/default"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateAtScopeHandleResponse handles the CreateOrUpdateAtScope response.
func (client *TagsClient) createOrUpdateAtScopeHandleResponse(resp *azcore.Response) (TagsResourceResponse, error) {
	var val *TagsResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return TagsResourceResponse{}, err
	}
	return TagsResourceResponse{RawResponse: resp.Response, TagsResource: val}, nil
}

// createOrUpdateAtScopeHandleError handles the CreateOrUpdateAtScope error response.
func (client *TagsClient) createOrUpdateAtScopeHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// CreateOrUpdateValue - This operation allows adding a value to the list of predefined values for an existing predefined tag name. A tag value can have
// a maximum of 256 characters.
func (client *TagsClient) CreateOrUpdateValue(ctx context.Context, tagName string, tagValue string, options *TagsCreateOrUpdateValueOptions) (TagValueResponse, error) {
	req, err := client.createOrUpdateValueCreateRequest(ctx, tagName, tagValue, options)
	if err != nil {
		return TagValueResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TagValueResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return TagValueResponse{}, client.createOrUpdateValueHandleError(resp)
	}
	return client.createOrUpdateValueHandleResponse(resp)
}

// createOrUpdateValueCreateRequest creates the CreateOrUpdateValue request.
func (client *TagsClient) createOrUpdateValueCreateRequest(ctx context.Context, tagName string, tagValue string, options *TagsCreateOrUpdateValueOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/tagNames/{tagName}/tagValues/{tagValue}"
	urlPath = strings.ReplaceAll(urlPath, "{tagName}", url.PathEscape(tagName))
	urlPath = strings.ReplaceAll(urlPath, "{tagValue}", url.PathEscape(tagValue))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// createOrUpdateValueHandleResponse handles the CreateOrUpdateValue response.
func (client *TagsClient) createOrUpdateValueHandleResponse(resp *azcore.Response) (TagValueResponse, error) {
	var val *TagValue
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return TagValueResponse{}, err
	}
	return TagValueResponse{RawResponse: resp.Response, TagValue: val}, nil
}

// createOrUpdateValueHandleError handles the CreateOrUpdateValue error response.
func (client *TagsClient) createOrUpdateValueHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Delete - This operation allows deleting a name from the list of predefined tag names for the given subscription. The name being deleted must not be in
// use as a tag name for any resource. All predefined values
// for the given name must have already been deleted.
func (client *TagsClient) Delete(ctx context.Context, tagName string, options *TagsDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, tagName, options)
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
func (client *TagsClient) deleteCreateRequest(ctx context.Context, tagName string, options *TagsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/tagNames/{tagName}"
	urlPath = strings.ReplaceAll(urlPath, "{tagName}", url.PathEscape(tagName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *TagsClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// DeleteAtScope - Deletes the entire set of tags on a resource or subscription.
func (client *TagsClient) DeleteAtScope(ctx context.Context, scope string, options *TagsDeleteAtScopeOptions) (*http.Response, error) {
	req, err := client.deleteAtScopeCreateRequest(ctx, scope, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.deleteAtScopeHandleError(resp)
	}
	return resp.Response, nil
}

// deleteAtScopeCreateRequest creates the DeleteAtScope request.
func (client *TagsClient) deleteAtScopeCreateRequest(ctx context.Context, scope string, options *TagsDeleteAtScopeOptions) (*azcore.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Resources/tags/default"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteAtScopeHandleError handles the DeleteAtScope error response.
func (client *TagsClient) deleteAtScopeHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// DeleteValue - This operation allows deleting a value from the list of predefined values for an existing predefined tag name. The value being deleted
// must not be in use as a tag value for the given tag name for any
// resource.
func (client *TagsClient) DeleteValue(ctx context.Context, tagName string, tagValue string, options *TagsDeleteValueOptions) (*http.Response, error) {
	req, err := client.deleteValueCreateRequest(ctx, tagName, tagValue, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return nil, client.deleteValueHandleError(resp)
	}
	return resp.Response, nil
}

// deleteValueCreateRequest creates the DeleteValue request.
func (client *TagsClient) deleteValueCreateRequest(ctx context.Context, tagName string, tagValue string, options *TagsDeleteValueOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/tagNames/{tagName}/tagValues/{tagValue}"
	urlPath = strings.ReplaceAll(urlPath, "{tagName}", url.PathEscape(tagName))
	urlPath = strings.ReplaceAll(urlPath, "{tagValue}", url.PathEscape(tagValue))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteValueHandleError handles the DeleteValue error response.
func (client *TagsClient) deleteValueHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetAtScope - Gets the entire set of tags on a resource or subscription.
func (client *TagsClient) GetAtScope(ctx context.Context, scope string, options *TagsGetAtScopeOptions) (TagsResourceResponse, error) {
	req, err := client.getAtScopeCreateRequest(ctx, scope, options)
	if err != nil {
		return TagsResourceResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TagsResourceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return TagsResourceResponse{}, client.getAtScopeHandleError(resp)
	}
	return client.getAtScopeHandleResponse(resp)
}

// getAtScopeCreateRequest creates the GetAtScope request.
func (client *TagsClient) getAtScopeCreateRequest(ctx context.Context, scope string, options *TagsGetAtScopeOptions) (*azcore.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Resources/tags/default"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getAtScopeHandleResponse handles the GetAtScope response.
func (client *TagsClient) getAtScopeHandleResponse(resp *azcore.Response) (TagsResourceResponse, error) {
	var val *TagsResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return TagsResourceResponse{}, err
	}
	return TagsResourceResponse{RawResponse: resp.Response, TagsResource: val}, nil
}

// getAtScopeHandleError handles the GetAtScope error response.
func (client *TagsClient) getAtScopeHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - This operation performs a union of predefined tags, resource tags, resource group tags and subscription tags, and returns a summary of usage for
// each tag name and value under the given subscription.
// In case of a large number of tags, this operation may return a previously cached result.
func (client *TagsClient) List(options *TagsListOptions) TagsListResultPager {
	return &tagsListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp TagsListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.TagsListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *TagsClient) listCreateRequest(ctx context.Context, options *TagsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/tagNames"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TagsClient) listHandleResponse(resp *azcore.Response) (TagsListResultResponse, error) {
	var val *TagsListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return TagsListResultResponse{}, err
	}
	return TagsListResultResponse{RawResponse: resp.Response, TagsListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *TagsClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// UpdateAtScope - This operation allows replacing, merging or selectively deleting tags on the specified resource or subscription. The specified entity
// can have a maximum of 50 tags at the end of the operation. The
// 'replace' option replaces the entire set of existing tags with a new set. The 'merge' option allows adding tags with new names and updating the values
// of tags with existing names. The 'delete' option
// allows selectively deleting tags based on given names or name/value pairs.
func (client *TagsClient) UpdateAtScope(ctx context.Context, scope string, parameters TagsPatchResource, options *TagsUpdateAtScopeOptions) (TagsResourceResponse, error) {
	req, err := client.updateAtScopeCreateRequest(ctx, scope, parameters, options)
	if err != nil {
		return TagsResourceResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TagsResourceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return TagsResourceResponse{}, client.updateAtScopeHandleError(resp)
	}
	return client.updateAtScopeHandleResponse(resp)
}

// updateAtScopeCreateRequest creates the UpdateAtScope request.
func (client *TagsClient) updateAtScopeCreateRequest(ctx context.Context, scope string, parameters TagsPatchResource, options *TagsUpdateAtScopeOptions) (*azcore.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Resources/tags/default"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// updateAtScopeHandleResponse handles the UpdateAtScope response.
func (client *TagsClient) updateAtScopeHandleResponse(resp *azcore.Response) (TagsResourceResponse, error) {
	var val *TagsResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return TagsResourceResponse{}, err
	}
	return TagsResourceResponse{RawResponse: resp.Response, TagsResource: val}, nil
}

// updateAtScopeHandleError handles the UpdateAtScope error response.
func (client *TagsClient) updateAtScopeHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}