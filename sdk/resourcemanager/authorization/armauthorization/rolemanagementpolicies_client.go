//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armauthorization

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

// RoleManagementPoliciesClient contains the methods for the RoleManagementPolicies group.
// Don't use this type directly, use NewRoleManagementPoliciesClient() instead.
type RoleManagementPoliciesClient struct {
	internal *arm.Client
}

// NewRoleManagementPoliciesClient creates a new instance of RoleManagementPoliciesClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRoleManagementPoliciesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*RoleManagementPoliciesClient, error) {
	cl, err := arm.NewClient(moduleName+".RoleManagementPoliciesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RoleManagementPoliciesClient{
		internal: cl,
	}
	return client, nil
}

// Delete - Delete a role management policy
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-10-01-preview
//   - scope - The scope of the role management policy to upsert.
//   - roleManagementPolicyName - The name (guid) of the role management policy to upsert.
//   - options - RoleManagementPoliciesClientDeleteOptions contains the optional parameters for the RoleManagementPoliciesClient.Delete
//     method.
func (client *RoleManagementPoliciesClient) Delete(ctx context.Context, scope string, roleManagementPolicyName string, options *RoleManagementPoliciesClientDeleteOptions) (RoleManagementPoliciesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, scope, roleManagementPolicyName, options)
	if err != nil {
		return RoleManagementPoliciesClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RoleManagementPoliciesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return RoleManagementPoliciesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return RoleManagementPoliciesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RoleManagementPoliciesClient) deleteCreateRequest(ctx context.Context, scope string, roleManagementPolicyName string, options *RoleManagementPoliciesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleManagementPolicies/{roleManagementPolicyName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleManagementPolicyName == "" {
		return nil, errors.New("parameter roleManagementPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleManagementPolicyName}", url.PathEscape(roleManagementPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the specified role management policy for a resource scope
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-10-01-preview
//   - scope - The scope of the role management policy.
//   - roleManagementPolicyName - The name (guid) of the role management policy to get.
//   - options - RoleManagementPoliciesClientGetOptions contains the optional parameters for the RoleManagementPoliciesClient.Get
//     method.
func (client *RoleManagementPoliciesClient) Get(ctx context.Context, scope string, roleManagementPolicyName string, options *RoleManagementPoliciesClientGetOptions) (RoleManagementPoliciesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, scope, roleManagementPolicyName, options)
	if err != nil {
		return RoleManagementPoliciesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RoleManagementPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleManagementPoliciesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RoleManagementPoliciesClient) getCreateRequest(ctx context.Context, scope string, roleManagementPolicyName string, options *RoleManagementPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleManagementPolicies/{roleManagementPolicyName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleManagementPolicyName == "" {
		return nil, errors.New("parameter roleManagementPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleManagementPolicyName}", url.PathEscape(roleManagementPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RoleManagementPoliciesClient) getHandleResponse(resp *http.Response) (RoleManagementPoliciesClientGetResponse, error) {
	result := RoleManagementPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleManagementPolicy); err != nil {
		return RoleManagementPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// NewListForScopePager - Gets role management policies for a resource scope.
//
// Generated from API version 2020-10-01-preview
//   - scope - The scope of the role management policy.
//   - options - RoleManagementPoliciesClientListForScopeOptions contains the optional parameters for the RoleManagementPoliciesClient.NewListForScopePager
//     method.
func (client *RoleManagementPoliciesClient) NewListForScopePager(scope string, options *RoleManagementPoliciesClientListForScopeOptions) *runtime.Pager[RoleManagementPoliciesClientListForScopeResponse] {
	return runtime.NewPager(runtime.PagingHandler[RoleManagementPoliciesClientListForScopeResponse]{
		More: func(page RoleManagementPoliciesClientListForScopeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RoleManagementPoliciesClientListForScopeResponse) (RoleManagementPoliciesClientListForScopeResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listForScopeCreateRequest(ctx, scope, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RoleManagementPoliciesClientListForScopeResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return RoleManagementPoliciesClientListForScopeResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RoleManagementPoliciesClientListForScopeResponse{}, runtime.NewResponseError(resp)
			}
			return client.listForScopeHandleResponse(resp)
		},
	})
}

// listForScopeCreateRequest creates the ListForScope request.
func (client *RoleManagementPoliciesClient) listForScopeCreateRequest(ctx context.Context, scope string, options *RoleManagementPoliciesClientListForScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleManagementPolicies"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listForScopeHandleResponse handles the ListForScope response.
func (client *RoleManagementPoliciesClient) listForScopeHandleResponse(resp *http.Response) (RoleManagementPoliciesClientListForScopeResponse, error) {
	result := RoleManagementPoliciesClientListForScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleManagementPolicyListResult); err != nil {
		return RoleManagementPoliciesClientListForScopeResponse{}, err
	}
	return result, nil
}

// Update - Update a role management policy
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-10-01-preview
//   - scope - The scope of the role management policy to upsert.
//   - roleManagementPolicyName - The name (guid) of the role management policy to upsert.
//   - parameters - Parameters for the role management policy.
//   - options - RoleManagementPoliciesClientUpdateOptions contains the optional parameters for the RoleManagementPoliciesClient.Update
//     method.
func (client *RoleManagementPoliciesClient) Update(ctx context.Context, scope string, roleManagementPolicyName string, parameters RoleManagementPolicy, options *RoleManagementPoliciesClientUpdateOptions) (RoleManagementPoliciesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, scope, roleManagementPolicyName, parameters, options)
	if err != nil {
		return RoleManagementPoliciesClientUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RoleManagementPoliciesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleManagementPoliciesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *RoleManagementPoliciesClient) updateCreateRequest(ctx context.Context, scope string, roleManagementPolicyName string, parameters RoleManagementPolicy, options *RoleManagementPoliciesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleManagementPolicies/{roleManagementPolicyName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleManagementPolicyName == "" {
		return nil, errors.New("parameter roleManagementPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleManagementPolicyName}", url.PathEscape(roleManagementPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *RoleManagementPoliciesClient) updateHandleResponse(resp *http.Response) (RoleManagementPoliciesClientUpdateResponse, error) {
	result := RoleManagementPoliciesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleManagementPolicy); err != nil {
		return RoleManagementPoliciesClientUpdateResponse{}, err
	}
	return result, nil
}
