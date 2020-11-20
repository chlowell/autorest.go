// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// NetworkInterfaceIPConfigurationsClient contains the methods for the NetworkInterfaceIPConfigurations group.
// Don't use this type directly, use NewNetworkInterfaceIPConfigurationsClient() instead.
type NetworkInterfaceIPConfigurationsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewNetworkInterfaceIPConfigurationsClient creates a new instance of NetworkInterfaceIPConfigurationsClient with the specified values.
func NewNetworkInterfaceIPConfigurationsClient(con *armcore.Connection, subscriptionID string) NetworkInterfaceIPConfigurationsClient {
	return NetworkInterfaceIPConfigurationsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client NetworkInterfaceIPConfigurationsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// Get - Gets the specified network interface ip configuration.
func (client NetworkInterfaceIPConfigurationsClient) Get(ctx context.Context, resourceGroupName string, networkInterfaceName string, ipConfigurationName string, options *NetworkInterfaceIPConfigurationsGetOptions) (*NetworkInterfaceIPConfigurationResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkInterfaceName, ipConfigurationName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client NetworkInterfaceIPConfigurationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkInterfaceName string, ipConfigurationName string, options *NetworkInterfaceIPConfigurationsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/ipConfigurations/{ipConfigurationName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkInterfaceName}", url.PathEscape(networkInterfaceName))
	urlPath = strings.ReplaceAll(urlPath, "{ipConfigurationName}", url.PathEscape(ipConfigurationName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client NetworkInterfaceIPConfigurationsClient) getHandleResponse(resp *azcore.Response) (*NetworkInterfaceIPConfigurationResponse, error) {
	result := NetworkInterfaceIPConfigurationResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NetworkInterfaceIPConfiguration)
}

// getHandleError handles the Get error response.
func (client NetworkInterfaceIPConfigurationsClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Get all ip configurations in a network interface.
func (client NetworkInterfaceIPConfigurationsClient) List(resourceGroupName string, networkInterfaceName string, options *NetworkInterfaceIPConfigurationsListOptions) NetworkInterfaceIPConfigurationListResultPager {
	return &networkInterfaceIPConfigurationListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, networkInterfaceName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp *NetworkInterfaceIPConfigurationListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.NetworkInterfaceIPConfigurationListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client NetworkInterfaceIPConfigurationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, networkInterfaceName string, options *NetworkInterfaceIPConfigurationsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/ipConfigurations"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkInterfaceName}", url.PathEscape(networkInterfaceName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client NetworkInterfaceIPConfigurationsClient) listHandleResponse(resp *azcore.Response) (*NetworkInterfaceIPConfigurationListResultResponse, error) {
	result := NetworkInterfaceIPConfigurationListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.NetworkInterfaceIPConfigurationListResult)
}

// listHandleError handles the List error response.
func (client NetworkInterfaceIPConfigurationsClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
