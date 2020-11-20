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

// VpnLinkConnectionsClient contains the methods for the VpnLinkConnections group.
// Don't use this type directly, use NewVpnLinkConnectionsClient() instead.
type VpnLinkConnectionsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVpnLinkConnectionsClient creates a new instance of VpnLinkConnectionsClient with the specified values.
func NewVpnLinkConnectionsClient(con *armcore.Connection, subscriptionID string) VpnLinkConnectionsClient {
	return VpnLinkConnectionsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client VpnLinkConnectionsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// ListByVpnConnection - Retrieves all vpn site link connections for a particular virtual wan vpn gateway vpn connection.
func (client VpnLinkConnectionsClient) ListByVpnConnection(resourceGroupName string, gatewayName string, connectionName string, options *VpnLinkConnectionsListByVpnConnectionOptions) ListVpnSiteLinkConnectionsResultPager {
	return &listVpnSiteLinkConnectionsResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByVpnConnectionCreateRequest(ctx, resourceGroupName, gatewayName, connectionName, options)
		},
		responder: client.listByVpnConnectionHandleResponse,
		errorer:   client.listByVpnConnectionHandleError,
		advancer: func(ctx context.Context, resp *ListVpnSiteLinkConnectionsResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ListVpnSiteLinkConnectionsResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listByVpnConnectionCreateRequest creates the ListByVpnConnection request.
func (client VpnLinkConnectionsClient) listByVpnConnectionCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *VpnLinkConnectionsListByVpnConnectionOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}/vpnLinkConnections"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
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

// listByVpnConnectionHandleResponse handles the ListByVpnConnection response.
func (client VpnLinkConnectionsClient) listByVpnConnectionHandleResponse(resp *azcore.Response) (*ListVpnSiteLinkConnectionsResultResponse, error) {
	result := ListVpnSiteLinkConnectionsResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ListVpnSiteLinkConnectionsResult)
}

// listByVpnConnectionHandleError handles the ListByVpnConnection error response.
func (client VpnLinkConnectionsClient) listByVpnConnectionHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
