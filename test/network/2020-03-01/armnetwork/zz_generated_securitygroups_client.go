//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

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

// SecurityGroupsClient contains the methods for the NetworkSecurityGroups group.
// Don't use this type directly, use NewSecurityGroupsClient() instead.
type SecurityGroupsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSecurityGroupsClient creates a new instance of SecurityGroupsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSecurityGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SecurityGroupsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	client := &SecurityGroupsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Host),
		pl:             armruntime.NewPipeline(module, version, credential, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates a network security group in the specified resource group.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// networkSecurityGroupName - The name of the network security group.
// parameters - Parameters supplied to the create or update network security group operation.
// options - NetworkSecurityGroupsBeginCreateOrUpdateOptions contains the optional parameters for the SecurityGroupsClient.BeginCreateOrUpdate
// method.
func (client *SecurityGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, parameters SecurityGroup, options *NetworkSecurityGroupsBeginCreateOrUpdateOptions) (NetworkSecurityGroupsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, networkSecurityGroupName, parameters, options)
	if err != nil {
		return NetworkSecurityGroupsCreateOrUpdatePollerResponse{}, err
	}
	result := NetworkSecurityGroupsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SecurityGroupsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return NetworkSecurityGroupsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &NetworkSecurityGroupsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a network security group in the specified resource group.
// If the operation fails it returns the *CloudError error type.
func (client *SecurityGroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, parameters SecurityGroup, options *NetworkSecurityGroupsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SecurityGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, parameters SecurityGroup, options *NetworkSecurityGroupsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityGroupName == "" {
		return nil, errors.New("parameter networkSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *SecurityGroupsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes the specified network security group.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// networkSecurityGroupName - The name of the network security group.
// options - NetworkSecurityGroupsBeginDeleteOptions contains the optional parameters for the SecurityGroupsClient.BeginDelete
// method.
func (client *SecurityGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, options *NetworkSecurityGroupsBeginDeleteOptions) (NetworkSecurityGroupsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, networkSecurityGroupName, options)
	if err != nil {
		return NetworkSecurityGroupsDeletePollerResponse{}, err
	}
	result := NetworkSecurityGroupsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SecurityGroupsClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return NetworkSecurityGroupsDeletePollerResponse{}, err
	}
	result.Poller = &NetworkSecurityGroupsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified network security group.
// If the operation fails it returns the *CloudError error type.
func (client *SecurityGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, options *NetworkSecurityGroupsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SecurityGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, options *NetworkSecurityGroupsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityGroupName == "" {
		return nil, errors.New("parameter networkSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *SecurityGroupsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets the specified network security group.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// networkSecurityGroupName - The name of the network security group.
// options - NetworkSecurityGroupsGetOptions contains the optional parameters for the SecurityGroupsClient.Get method.
func (client *SecurityGroupsClient) Get(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, options *NetworkSecurityGroupsGetOptions) (NetworkSecurityGroupsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, options)
	if err != nil {
		return NetworkSecurityGroupsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NetworkSecurityGroupsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NetworkSecurityGroupsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SecurityGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, options *NetworkSecurityGroupsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityGroupName == "" {
		return nil, errors.New("parameter networkSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SecurityGroupsClient) getHandleResponse(resp *http.Response) (NetworkSecurityGroupsGetResponse, error) {
	result := NetworkSecurityGroupsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityGroup); err != nil {
		return NetworkSecurityGroupsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *SecurityGroupsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Gets all network security groups in a resource group.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// options - NetworkSecurityGroupsListOptions contains the optional parameters for the SecurityGroupsClient.List method.
func (client *SecurityGroupsClient) List(resourceGroupName string, options *NetworkSecurityGroupsListOptions) *NetworkSecurityGroupsListPager {
	return &NetworkSecurityGroupsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp NetworkSecurityGroupsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SecurityGroupListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *SecurityGroupsClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *NetworkSecurityGroupsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SecurityGroupsClient) listHandleResponse(resp *http.Response) (NetworkSecurityGroupsListResponse, error) {
	result := NetworkSecurityGroupsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityGroupListResult); err != nil {
		return NetworkSecurityGroupsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *SecurityGroupsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListAll - Gets all network security groups in a subscription.
// If the operation fails it returns the *CloudError error type.
// options - NetworkSecurityGroupsListAllOptions contains the optional parameters for the SecurityGroupsClient.ListAll method.
func (client *SecurityGroupsClient) ListAll(options *NetworkSecurityGroupsListAllOptions) *NetworkSecurityGroupsListAllPager {
	return &NetworkSecurityGroupsListAllPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listAllCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp NetworkSecurityGroupsListAllResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SecurityGroupListResult.NextLink)
		},
	}
}

// listAllCreateRequest creates the ListAll request.
func (client *SecurityGroupsClient) listAllCreateRequest(ctx context.Context, options *NetworkSecurityGroupsListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkSecurityGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *SecurityGroupsClient) listAllHandleResponse(resp *http.Response) (NetworkSecurityGroupsListAllResponse, error) {
	result := NetworkSecurityGroupsListAllResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityGroupListResult); err != nil {
		return NetworkSecurityGroupsListAllResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listAllHandleError handles the ListAll error response.
func (client *SecurityGroupsClient) listAllHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// UpdateTags - Updates a network security group tags.
// If the operation fails it returns the *CloudError error type.
// resourceGroupName - The name of the resource group.
// networkSecurityGroupName - The name of the network security group.
// parameters - Parameters supplied to update network security group tags.
// options - NetworkSecurityGroupsUpdateTagsOptions contains the optional parameters for the SecurityGroupsClient.UpdateTags
// method.
func (client *SecurityGroupsClient) UpdateTags(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, parameters TagsObject, options *NetworkSecurityGroupsUpdateTagsOptions) (NetworkSecurityGroupsUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, parameters, options)
	if err != nil {
		return NetworkSecurityGroupsUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NetworkSecurityGroupsUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NetworkSecurityGroupsUpdateTagsResponse{}, client.updateTagsHandleError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *SecurityGroupsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, parameters TagsObject, options *NetworkSecurityGroupsUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityGroupName == "" {
		return nil, errors.New("parameter networkSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *SecurityGroupsClient) updateTagsHandleResponse(resp *http.Response) (NetworkSecurityGroupsUpdateTagsResponse, error) {
	result := NetworkSecurityGroupsUpdateTagsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SecurityGroup); err != nil {
		return NetworkSecurityGroupsUpdateTagsResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *SecurityGroupsClient) updateTagsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}