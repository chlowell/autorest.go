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

// ResourceNavigationLinksClient contains the methods for the ResourceNavigationLinks group.
// Don't use this type directly, use NewResourceNavigationLinksClient() instead.
type ResourceNavigationLinksClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewResourceNavigationLinksClient creates a new instance of ResourceNavigationLinksClient with the specified values.
func NewResourceNavigationLinksClient(con *armcore.Connection, subscriptionID string) *ResourceNavigationLinksClient {
	return &ResourceNavigationLinksClient{con: con, subscriptionID: subscriptionID}
}

// List - Gets a list of resource navigation links for a subnet.
// If the operation fails it returns the *CloudError error type.
func (client *ResourceNavigationLinksClient) List(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, options *ResourceNavigationLinksListOptions) (ResourceNavigationLinksListResultResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, virtualNetworkName, subnetName, options)
	if err != nil {
		return ResourceNavigationLinksListResultResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ResourceNavigationLinksListResultResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ResourceNavigationLinksListResultResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ResourceNavigationLinksClient) listCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, options *ResourceNavigationLinksListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}/ResourceNavigationLinks"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	if subnetName == "" {
		return nil, errors.New("parameter subnetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subnetName}", url.PathEscape(subnetName))
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
func (client *ResourceNavigationLinksClient) listHandleResponse(resp *azcore.Response) (ResourceNavigationLinksListResultResponse, error) {
	var val *ResourceNavigationLinksListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ResourceNavigationLinksListResultResponse{}, err
	}
	return ResourceNavigationLinksListResultResponse{RawResponse: resp.Response, ResourceNavigationLinksListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *ResourceNavigationLinksClient) listHandleError(resp *azcore.Response) error {
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
