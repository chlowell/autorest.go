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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// PeerExpressRouteCircuitConnectionsClient contains the methods for the PeerExpressRouteCircuitConnections group.
// Don't use this type directly, use NewPeerExpressRouteCircuitConnectionsClient() instead.
type PeerExpressRouteCircuitConnectionsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewPeerExpressRouteCircuitConnectionsClient creates a new instance of PeerExpressRouteCircuitConnectionsClient with the specified values.
func NewPeerExpressRouteCircuitConnectionsClient(con *arm.Connection, subscriptionID string) *PeerExpressRouteCircuitConnectionsClient {
	return &PeerExpressRouteCircuitConnectionsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Get - Gets the specified Peer Express Route Circuit Connection from the specified express route circuit.
// If the operation fails it returns the *CloudError error type.
func (client *PeerExpressRouteCircuitConnectionsClient) Get(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, connectionName string, options *PeerExpressRouteCircuitConnectionsGetOptions) (PeerExpressRouteCircuitConnectionsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, circuitName, peeringName, connectionName, options)
	if err != nil {
		return PeerExpressRouteCircuitConnectionsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PeerExpressRouteCircuitConnectionsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PeerExpressRouteCircuitConnectionsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PeerExpressRouteCircuitConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, connectionName string, options *PeerExpressRouteCircuitConnectionsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/peerConnections/{connectionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if circuitName == "" {
		return nil, errors.New("parameter circuitName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PeerExpressRouteCircuitConnectionsClient) getHandleResponse(resp *http.Response) (PeerExpressRouteCircuitConnectionsGetResponse, error) {
	result := PeerExpressRouteCircuitConnectionsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PeerExpressRouteCircuitConnection); err != nil {
		return PeerExpressRouteCircuitConnectionsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *PeerExpressRouteCircuitConnectionsClient) getHandleError(resp *http.Response) error {
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

// List - Gets all global reach peer connections associated with a private peering in an express route circuit.
// If the operation fails it returns the *CloudError error type.
func (client *PeerExpressRouteCircuitConnectionsClient) List(resourceGroupName string, circuitName string, peeringName string, options *PeerExpressRouteCircuitConnectionsListOptions) *PeerExpressRouteCircuitConnectionsListPager {
	return &PeerExpressRouteCircuitConnectionsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, circuitName, peeringName, options)
		},
		advancer: func(ctx context.Context, resp PeerExpressRouteCircuitConnectionsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PeerExpressRouteCircuitConnectionListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *PeerExpressRouteCircuitConnectionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, options *PeerExpressRouteCircuitConnectionsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/peerConnections"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if circuitName == "" {
		return nil, errors.New("parameter circuitName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
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
func (client *PeerExpressRouteCircuitConnectionsClient) listHandleResponse(resp *http.Response) (PeerExpressRouteCircuitConnectionsListResponse, error) {
	result := PeerExpressRouteCircuitConnectionsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PeerExpressRouteCircuitConnectionListResult); err != nil {
		return PeerExpressRouteCircuitConnectionsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *PeerExpressRouteCircuitConnectionsClient) listHandleError(resp *http.Response) error {
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
