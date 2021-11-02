//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmediaservices

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
	"strconv"
	"strings"
)

// AssetsClient contains the methods for the Assets group.
// Don't use this type directly, use NewAssetsClient() instead.
type AssetsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewAssetsClient creates a new instance of AssetsClient with the specified values.
func NewAssetsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *AssetsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &AssetsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CreateOrUpdate - Creates or updates an Asset in the Media Services account
// If the operation fails it returns the *ErrorResponse error type.
func (client *AssetsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, assetName string, parameters Asset, options *AssetsCreateOrUpdateOptions) (AssetsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, assetName, parameters, options)
	if err != nil {
		return AssetsCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AssetsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return AssetsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AssetsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, assetName string, parameters Asset, options *AssetsCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if assetName == "" {
		return nil, errors.New("parameter assetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assetName}", url.PathEscape(assetName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *AssetsClient) createOrUpdateHandleResponse(resp *http.Response) (AssetsCreateOrUpdateResponse, error) {
	result := AssetsCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Asset); err != nil {
		return AssetsCreateOrUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *AssetsClient) createOrUpdateHandleError(resp *http.Response) error {
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

// Delete - Deletes an Asset in the Media Services account
// If the operation fails it returns the *ErrorResponse error type.
func (client *AssetsClient) Delete(ctx context.Context, resourceGroupName string, accountName string, assetName string, options *AssetsDeleteOptions) (AssetsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, assetName, options)
	if err != nil {
		return AssetsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AssetsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return AssetsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return AssetsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AssetsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, assetName string, options *AssetsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if assetName == "" {
		return nil, errors.New("parameter assetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assetName}", url.PathEscape(assetName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *AssetsClient) deleteHandleError(resp *http.Response) error {
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

// Get - Get the details of an Asset in the Media Services account
// If the operation fails it returns the *ErrorResponse error type.
func (client *AssetsClient) Get(ctx context.Context, resourceGroupName string, accountName string, assetName string, options *AssetsGetOptions) (AssetsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, assetName, options)
	if err != nil {
		return AssetsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AssetsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AssetsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AssetsClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, assetName string, options *AssetsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if assetName == "" {
		return nil, errors.New("parameter assetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assetName}", url.PathEscape(assetName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AssetsClient) getHandleResponse(resp *http.Response) (AssetsGetResponse, error) {
	result := AssetsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Asset); err != nil {
		return AssetsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *AssetsClient) getHandleError(resp *http.Response) error {
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

// GetEncryptionKey - Gets the Asset storage encryption keys used to decrypt content created by version 2 of the Media Services API
// If the operation fails it returns the *ErrorResponse error type.
func (client *AssetsClient) GetEncryptionKey(ctx context.Context, resourceGroupName string, accountName string, assetName string, options *AssetsGetEncryptionKeyOptions) (AssetsGetEncryptionKeyResponse, error) {
	req, err := client.getEncryptionKeyCreateRequest(ctx, resourceGroupName, accountName, assetName, options)
	if err != nil {
		return AssetsGetEncryptionKeyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AssetsGetEncryptionKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AssetsGetEncryptionKeyResponse{}, client.getEncryptionKeyHandleError(resp)
	}
	return client.getEncryptionKeyHandleResponse(resp)
}

// getEncryptionKeyCreateRequest creates the GetEncryptionKey request.
func (client *AssetsClient) getEncryptionKeyCreateRequest(ctx context.Context, resourceGroupName string, accountName string, assetName string, options *AssetsGetEncryptionKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/getEncryptionKey"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if assetName == "" {
		return nil, errors.New("parameter assetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assetName}", url.PathEscape(assetName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getEncryptionKeyHandleResponse handles the GetEncryptionKey response.
func (client *AssetsClient) getEncryptionKeyHandleResponse(resp *http.Response) (AssetsGetEncryptionKeyResponse, error) {
	result := AssetsGetEncryptionKeyResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageEncryptedAssetDecryptionData); err != nil {
		return AssetsGetEncryptionKeyResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getEncryptionKeyHandleError handles the GetEncryptionKey error response.
func (client *AssetsClient) getEncryptionKeyHandleError(resp *http.Response) error {
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

// List - List Assets in the Media Services account with optional filtering and ordering
// If the operation fails it returns the *ErrorResponse error type.
func (client *AssetsClient) List(resourceGroupName string, accountName string, options *AssetsListOptions) *AssetsListPager {
	return &AssetsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, accountName, options)
		},
		advancer: func(ctx context.Context, resp AssetsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AssetCollection.ODataNextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *AssetsClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *AssetsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AssetsClient) listHandleResponse(resp *http.Response) (AssetsListResponse, error) {
	result := AssetsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssetCollection); err != nil {
		return AssetsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *AssetsClient) listHandleError(resp *http.Response) error {
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

// ListContainerSas - Lists storage container URLs with shared access signatures (SAS) for uploading and downloading Asset content. The signatures are derived
// from the storage account keys.
// If the operation fails it returns the *ErrorResponse error type.
func (client *AssetsClient) ListContainerSas(ctx context.Context, resourceGroupName string, accountName string, assetName string, parameters ListContainerSasInput, options *AssetsListContainerSasOptions) (AssetsListContainerSasResponse, error) {
	req, err := client.listContainerSasCreateRequest(ctx, resourceGroupName, accountName, assetName, parameters, options)
	if err != nil {
		return AssetsListContainerSasResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AssetsListContainerSasResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AssetsListContainerSasResponse{}, client.listContainerSasHandleError(resp)
	}
	return client.listContainerSasHandleResponse(resp)
}

// listContainerSasCreateRequest creates the ListContainerSas request.
func (client *AssetsClient) listContainerSasCreateRequest(ctx context.Context, resourceGroupName string, accountName string, assetName string, parameters ListContainerSasInput, options *AssetsListContainerSasOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/listContainerSas"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if assetName == "" {
		return nil, errors.New("parameter assetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assetName}", url.PathEscape(assetName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// listContainerSasHandleResponse handles the ListContainerSas response.
func (client *AssetsClient) listContainerSasHandleResponse(resp *http.Response) (AssetsListContainerSasResponse, error) {
	result := AssetsListContainerSasResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssetContainerSas); err != nil {
		return AssetsListContainerSasResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listContainerSasHandleError handles the ListContainerSas error response.
func (client *AssetsClient) listContainerSasHandleError(resp *http.Response) error {
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

// ListStreamingLocators - Lists Streaming Locators which are associated with this asset.
// If the operation fails it returns the *ErrorResponse error type.
func (client *AssetsClient) ListStreamingLocators(ctx context.Context, resourceGroupName string, accountName string, assetName string, options *AssetsListStreamingLocatorsOptions) (AssetsListStreamingLocatorsResponse, error) {
	req, err := client.listStreamingLocatorsCreateRequest(ctx, resourceGroupName, accountName, assetName, options)
	if err != nil {
		return AssetsListStreamingLocatorsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AssetsListStreamingLocatorsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AssetsListStreamingLocatorsResponse{}, client.listStreamingLocatorsHandleError(resp)
	}
	return client.listStreamingLocatorsHandleResponse(resp)
}

// listStreamingLocatorsCreateRequest creates the ListStreamingLocators request.
func (client *AssetsClient) listStreamingLocatorsCreateRequest(ctx context.Context, resourceGroupName string, accountName string, assetName string, options *AssetsListStreamingLocatorsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/listStreamingLocators"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if assetName == "" {
		return nil, errors.New("parameter assetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assetName}", url.PathEscape(assetName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listStreamingLocatorsHandleResponse handles the ListStreamingLocators response.
func (client *AssetsClient) listStreamingLocatorsHandleResponse(resp *http.Response) (AssetsListStreamingLocatorsResponse, error) {
	result := AssetsListStreamingLocatorsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListStreamingLocatorsResponse); err != nil {
		return AssetsListStreamingLocatorsResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listStreamingLocatorsHandleError handles the ListStreamingLocators error response.
func (client *AssetsClient) listStreamingLocatorsHandleError(resp *http.Response) error {
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

// Update - Updates an existing Asset in the Media Services account
// If the operation fails it returns the *ErrorResponse error type.
func (client *AssetsClient) Update(ctx context.Context, resourceGroupName string, accountName string, assetName string, parameters Asset, options *AssetsUpdateOptions) (AssetsUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, assetName, parameters, options)
	if err != nil {
		return AssetsUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AssetsUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AssetsUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *AssetsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, assetName string, parameters Asset, options *AssetsUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if assetName == "" {
		return nil, errors.New("parameter assetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assetName}", url.PathEscape(assetName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *AssetsClient) updateHandleResponse(resp *http.Response) (AssetsUpdateResponse, error) {
	result := AssetsUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Asset); err != nil {
		return AssetsUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *AssetsClient) updateHandleError(resp *http.Response) error {
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