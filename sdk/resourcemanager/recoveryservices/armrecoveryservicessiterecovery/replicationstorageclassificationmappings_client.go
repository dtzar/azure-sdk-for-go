//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicessiterecovery

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ReplicationStorageClassificationMappingsClient contains the methods for the ReplicationStorageClassificationMappings group.
// Don't use this type directly, use NewReplicationStorageClassificationMappingsClient() instead.
type ReplicationStorageClassificationMappingsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewReplicationStorageClassificationMappingsClient creates a new instance of ReplicationStorageClassificationMappingsClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewReplicationStorageClassificationMappingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ReplicationStorageClassificationMappingsClient, error) {
	cl, err := arm.NewClient(moduleName+".ReplicationStorageClassificationMappingsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ReplicationStorageClassificationMappingsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - The operation to create a storage classification mapping.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric name.
//   - storageClassificationName - Storage classification name.
//   - storageClassificationMappingName - Storage classification mapping name.
//   - pairingInput - Pairing input.
//   - options - ReplicationStorageClassificationMappingsClientBeginCreateOptions contains the optional parameters for the ReplicationStorageClassificationMappingsClient.BeginCreate
//     method.
func (client *ReplicationStorageClassificationMappingsClient) BeginCreate(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, storageClassificationName string, storageClassificationMappingName string, pairingInput StorageClassificationMappingInput, options *ReplicationStorageClassificationMappingsClientBeginCreateOptions) (*runtime.Poller[ReplicationStorageClassificationMappingsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceName, resourceGroupName, fabricName, storageClassificationName, storageClassificationMappingName, pairingInput, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[ReplicationStorageClassificationMappingsClientCreateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ReplicationStorageClassificationMappingsClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - The operation to create a storage classification mapping.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
func (client *ReplicationStorageClassificationMappingsClient) create(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, storageClassificationName string, storageClassificationMappingName string, pairingInput StorageClassificationMappingInput, options *ReplicationStorageClassificationMappingsClientBeginCreateOptions) (*http.Response, error) {
	var err error
	req, err := client.createCreateRequest(ctx, resourceName, resourceGroupName, fabricName, storageClassificationName, storageClassificationMappingName, pairingInput, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCreateRequest creates the Create request.
func (client *ReplicationStorageClassificationMappingsClient) createCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, storageClassificationName string, storageClassificationMappingName string, pairingInput StorageClassificationMappingInput, options *ReplicationStorageClassificationMappingsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationStorageClassifications/{storageClassificationName}/replicationStorageClassificationMappings/{storageClassificationMappingName}"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if storageClassificationName == "" {
		return nil, errors.New("parameter storageClassificationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageClassificationName}", url.PathEscape(storageClassificationName))
	if storageClassificationMappingName == "" {
		return nil, errors.New("parameter storageClassificationMappingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageClassificationMappingName}", url.PathEscape(storageClassificationMappingName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, pairingInput); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - The operation to delete a storage classification mapping.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric name.
//   - storageClassificationName - Storage classification name.
//   - storageClassificationMappingName - Storage classification mapping name.
//   - options - ReplicationStorageClassificationMappingsClientBeginDeleteOptions contains the optional parameters for the ReplicationStorageClassificationMappingsClient.BeginDelete
//     method.
func (client *ReplicationStorageClassificationMappingsClient) BeginDelete(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, storageClassificationName string, storageClassificationMappingName string, options *ReplicationStorageClassificationMappingsClientBeginDeleteOptions) (*runtime.Poller[ReplicationStorageClassificationMappingsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceName, resourceGroupName, fabricName, storageClassificationName, storageClassificationMappingName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[ReplicationStorageClassificationMappingsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ReplicationStorageClassificationMappingsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - The operation to delete a storage classification mapping.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
func (client *ReplicationStorageClassificationMappingsClient) deleteOperation(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, storageClassificationName string, storageClassificationMappingName string, options *ReplicationStorageClassificationMappingsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceName, resourceGroupName, fabricName, storageClassificationName, storageClassificationMappingName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ReplicationStorageClassificationMappingsClient) deleteCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, storageClassificationName string, storageClassificationMappingName string, options *ReplicationStorageClassificationMappingsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationStorageClassifications/{storageClassificationName}/replicationStorageClassificationMappings/{storageClassificationMappingName}"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if storageClassificationName == "" {
		return nil, errors.New("parameter storageClassificationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageClassificationName}", url.PathEscape(storageClassificationName))
	if storageClassificationMappingName == "" {
		return nil, errors.New("parameter storageClassificationMappingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageClassificationMappingName}", url.PathEscape(storageClassificationMappingName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets the details of the specified storage classification mapping.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric name.
//   - storageClassificationName - Storage classification name.
//   - storageClassificationMappingName - Storage classification mapping name.
//   - options - ReplicationStorageClassificationMappingsClientGetOptions contains the optional parameters for the ReplicationStorageClassificationMappingsClient.Get
//     method.
func (client *ReplicationStorageClassificationMappingsClient) Get(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, storageClassificationName string, storageClassificationMappingName string, options *ReplicationStorageClassificationMappingsClientGetOptions) (ReplicationStorageClassificationMappingsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceName, resourceGroupName, fabricName, storageClassificationName, storageClassificationMappingName, options)
	if err != nil {
		return ReplicationStorageClassificationMappingsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ReplicationStorageClassificationMappingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ReplicationStorageClassificationMappingsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ReplicationStorageClassificationMappingsClient) getCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, storageClassificationName string, storageClassificationMappingName string, options *ReplicationStorageClassificationMappingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationStorageClassifications/{storageClassificationName}/replicationStorageClassificationMappings/{storageClassificationMappingName}"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if storageClassificationName == "" {
		return nil, errors.New("parameter storageClassificationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageClassificationName}", url.PathEscape(storageClassificationName))
	if storageClassificationMappingName == "" {
		return nil, errors.New("parameter storageClassificationMappingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageClassificationMappingName}", url.PathEscape(storageClassificationMappingName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ReplicationStorageClassificationMappingsClient) getHandleResponse(resp *http.Response) (ReplicationStorageClassificationMappingsClientGetResponse, error) {
	result := ReplicationStorageClassificationMappingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageClassificationMapping); err != nil {
		return ReplicationStorageClassificationMappingsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the storage classification mappings in the vault.
//
// Generated from API version 2023-06-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - options - ReplicationStorageClassificationMappingsClientListOptions contains the optional parameters for the ReplicationStorageClassificationMappingsClient.NewListPager
//     method.
func (client *ReplicationStorageClassificationMappingsClient) NewListPager(resourceName string, resourceGroupName string, options *ReplicationStorageClassificationMappingsClientListOptions) *runtime.Pager[ReplicationStorageClassificationMappingsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReplicationStorageClassificationMappingsClientListResponse]{
		More: func(page ReplicationStorageClassificationMappingsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReplicationStorageClassificationMappingsClientListResponse) (ReplicationStorageClassificationMappingsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceName, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReplicationStorageClassificationMappingsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ReplicationStorageClassificationMappingsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReplicationStorageClassificationMappingsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ReplicationStorageClassificationMappingsClient) listCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, options *ReplicationStorageClassificationMappingsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationStorageClassificationMappings"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ReplicationStorageClassificationMappingsClient) listHandleResponse(resp *http.Response) (ReplicationStorageClassificationMappingsClientListResponse, error) {
	result := ReplicationStorageClassificationMappingsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageClassificationMappingCollection); err != nil {
		return ReplicationStorageClassificationMappingsClientListResponse{}, err
	}
	return result, nil
}

// NewListByReplicationStorageClassificationsPager - Lists the storage classification mappings for the fabric.
//
// Generated from API version 2023-06-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric name.
//   - storageClassificationName - Storage classification name.
//   - options - ReplicationStorageClassificationMappingsClientListByReplicationStorageClassificationsOptions contains the optional
//     parameters for the ReplicationStorageClassificationMappingsClient.NewListByReplicationStorageClassificationsPager method.
func (client *ReplicationStorageClassificationMappingsClient) NewListByReplicationStorageClassificationsPager(resourceName string, resourceGroupName string, fabricName string, storageClassificationName string, options *ReplicationStorageClassificationMappingsClientListByReplicationStorageClassificationsOptions) *runtime.Pager[ReplicationStorageClassificationMappingsClientListByReplicationStorageClassificationsResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReplicationStorageClassificationMappingsClientListByReplicationStorageClassificationsResponse]{
		More: func(page ReplicationStorageClassificationMappingsClientListByReplicationStorageClassificationsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReplicationStorageClassificationMappingsClientListByReplicationStorageClassificationsResponse) (ReplicationStorageClassificationMappingsClientListByReplicationStorageClassificationsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByReplicationStorageClassificationsCreateRequest(ctx, resourceName, resourceGroupName, fabricName, storageClassificationName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReplicationStorageClassificationMappingsClientListByReplicationStorageClassificationsResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ReplicationStorageClassificationMappingsClientListByReplicationStorageClassificationsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReplicationStorageClassificationMappingsClientListByReplicationStorageClassificationsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByReplicationStorageClassificationsHandleResponse(resp)
		},
	})
}

// listByReplicationStorageClassificationsCreateRequest creates the ListByReplicationStorageClassifications request.
func (client *ReplicationStorageClassificationMappingsClient) listByReplicationStorageClassificationsCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, storageClassificationName string, options *ReplicationStorageClassificationMappingsClientListByReplicationStorageClassificationsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationStorageClassifications/{storageClassificationName}/replicationStorageClassificationMappings"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if storageClassificationName == "" {
		return nil, errors.New("parameter storageClassificationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageClassificationName}", url.PathEscape(storageClassificationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByReplicationStorageClassificationsHandleResponse handles the ListByReplicationStorageClassifications response.
func (client *ReplicationStorageClassificationMappingsClient) listByReplicationStorageClassificationsHandleResponse(resp *http.Response) (ReplicationStorageClassificationMappingsClientListByReplicationStorageClassificationsResponse, error) {
	result := ReplicationStorageClassificationMappingsClientListByReplicationStorageClassificationsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageClassificationMappingCollection); err != nil {
		return ReplicationStorageClassificationMappingsClientListByReplicationStorageClassificationsResponse{}, err
	}
	return result, nil
}
