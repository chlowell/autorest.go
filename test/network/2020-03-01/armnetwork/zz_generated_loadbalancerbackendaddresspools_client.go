// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// LoadBalancerBackendAddressPoolsClient contains the methods for the LoadBalancerBackendAddressPools group.
// Don't use this type directly, use NewLoadBalancerBackendAddressPoolsClient() instead.
type LoadBalancerBackendAddressPoolsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewLoadBalancerBackendAddressPoolsClient creates a new instance of LoadBalancerBackendAddressPoolsClient with the specified values.
func NewLoadBalancerBackendAddressPoolsClient(con *armcore.Connection, subscriptionID string) *LoadBalancerBackendAddressPoolsClient {
	return &LoadBalancerBackendAddressPoolsClient{con: con, subscriptionID: subscriptionID}
}

// Get - Gets load balancer backend address pool.
// If the operation fails it returns the *CloudError error type.
func (client *LoadBalancerBackendAddressPoolsClient) Get(ctx context.Context, resourceGroupName string, loadBalancerName string, backendAddressPoolName string, options *LoadBalancerBackendAddressPoolsGetOptions) (BackendAddressPoolResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, loadBalancerName, backendAddressPoolName, options)
	if err != nil {
		return BackendAddressPoolResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return BackendAddressPoolResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return BackendAddressPoolResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LoadBalancerBackendAddressPoolsClient) getCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, backendAddressPoolName string, options *LoadBalancerBackendAddressPoolsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/backendAddressPools/{backendAddressPoolName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if backendAddressPoolName == "" {
		return nil, errors.New("parameter backendAddressPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backendAddressPoolName}", url.PathEscape(backendAddressPoolName))
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
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LoadBalancerBackendAddressPoolsClient) getHandleResponse(resp *azcore.Response) (BackendAddressPoolResponse, error) {
	var val *BackendAddressPool
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return BackendAddressPoolResponse{}, err
	}
	return BackendAddressPoolResponse{RawResponse: resp.Response, BackendAddressPool: val}, nil
}

// getHandleError handles the Get error response.
func (client *LoadBalancerBackendAddressPoolsClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// List - Gets all the load balancer backed address pools.
// If the operation fails it returns the *CloudError error type.
func (client *LoadBalancerBackendAddressPoolsClient) List(resourceGroupName string, loadBalancerName string, options *LoadBalancerBackendAddressPoolsListOptions) LoadBalancerBackendAddressPoolListResultPager {
	return &loadBalancerBackendAddressPoolListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, loadBalancerName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp LoadBalancerBackendAddressPoolListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.LoadBalancerBackendAddressPoolListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *LoadBalancerBackendAddressPoolsClient) listCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancerBackendAddressPoolsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/backendAddressPools"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
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
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LoadBalancerBackendAddressPoolsClient) listHandleResponse(resp *azcore.Response) (LoadBalancerBackendAddressPoolListResultResponse, error) {
	var val *LoadBalancerBackendAddressPoolListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return LoadBalancerBackendAddressPoolListResultResponse{}, err
	}
	return LoadBalancerBackendAddressPoolListResultResponse{RawResponse: resp.Response, LoadBalancerBackendAddressPoolListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *LoadBalancerBackendAddressPoolsClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
