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
	"time"
)

// VirtualNetworkGatewayConnectionsClient contains the methods for the VirtualNetworkGatewayConnections group.
// Don't use this type directly, use NewVirtualNetworkGatewayConnectionsClient() instead.
type VirtualNetworkGatewayConnectionsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVirtualNetworkGatewayConnectionsClient creates a new instance of VirtualNetworkGatewayConnectionsClient with the specified values.
func NewVirtualNetworkGatewayConnectionsClient(con *armcore.Connection, subscriptionID string) VirtualNetworkGatewayConnectionsClient {
	return VirtualNetworkGatewayConnectionsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client VirtualNetworkGatewayConnectionsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// BeginCreateOrUpdate - Creates or updates a virtual network gateway connection in the specified resource group.
func (client VirtualNetworkGatewayConnectionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters VirtualNetworkGatewayConnection, options *VirtualNetworkGatewayConnectionsCreateOrUpdateOptions) (*VirtualNetworkGatewayConnectionPollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, parameters, options)
	if err != nil {
		return nil, err
	}
	result := &VirtualNetworkGatewayConnectionPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualNetworkGatewayConnectionsClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &virtualNetworkGatewayConnectionPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*VirtualNetworkGatewayConnectionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new VirtualNetworkGatewayConnectionPoller from the specified resume token.
// token - The value must come from a previous call to VirtualNetworkGatewayConnectionPoller.ResumeToken().
func (client VirtualNetworkGatewayConnectionsClient) ResumeCreateOrUpdate(token string) (VirtualNetworkGatewayConnectionPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualNetworkGatewayConnectionsClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &virtualNetworkGatewayConnectionPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates a virtual network gateway connection in the specified resource group.
func (client VirtualNetworkGatewayConnectionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters VirtualNetworkGatewayConnection, options *VirtualNetworkGatewayConnectionsCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client VirtualNetworkGatewayConnectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters VirtualNetworkGatewayConnection, options *VirtualNetworkGatewayConnectionsCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayConnectionName}", url.PathEscape(virtualNetworkGatewayConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client VirtualNetworkGatewayConnectionsClient) createOrUpdateHandleResponse(resp *azcore.Response) (*VirtualNetworkGatewayConnectionResponse, error) {
	result := VirtualNetworkGatewayConnectionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualNetworkGatewayConnection)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client VirtualNetworkGatewayConnectionsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Deletes the specified virtual network Gateway connection.
func (client VirtualNetworkGatewayConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, options *VirtualNetworkGatewayConnectionsDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualNetworkGatewayConnectionsClient.Delete", "location", resp, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client VirtualNetworkGatewayConnectionsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualNetworkGatewayConnectionsClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified virtual network Gateway connection.
func (client VirtualNetworkGatewayConnectionsClient) Delete(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, options *VirtualNetworkGatewayConnectionsDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client VirtualNetworkGatewayConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, options *VirtualNetworkGatewayConnectionsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayConnectionName}", url.PathEscape(virtualNetworkGatewayConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
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

// deleteHandleError handles the Delete error response.
func (client VirtualNetworkGatewayConnectionsClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified virtual network gateway connection by resource group.
func (client VirtualNetworkGatewayConnectionsClient) Get(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, options *VirtualNetworkGatewayConnectionsGetOptions) (*VirtualNetworkGatewayConnectionResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, options)
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
func (client VirtualNetworkGatewayConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, options *VirtualNetworkGatewayConnectionsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayConnectionName}", url.PathEscape(virtualNetworkGatewayConnectionName))
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
func (client VirtualNetworkGatewayConnectionsClient) getHandleResponse(resp *azcore.Response) (*VirtualNetworkGatewayConnectionResponse, error) {
	result := VirtualNetworkGatewayConnectionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualNetworkGatewayConnection)
}

// getHandleError handles the Get error response.
func (client VirtualNetworkGatewayConnectionsClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetSharedKey - The Get VirtualNetworkGatewayConnectionSharedKey operation retrieves information about the specified virtual network gateway connection
// shared key through Network resource provider.
func (client VirtualNetworkGatewayConnectionsClient) GetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, options *VirtualNetworkGatewayConnectionsGetSharedKeyOptions) (*ConnectionSharedKeyResponse, error) {
	req, err := client.getSharedKeyCreateRequest(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getSharedKeyHandleError(resp)
	}
	result, err := client.getSharedKeyHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getSharedKeyCreateRequest creates the GetSharedKey request.
func (client VirtualNetworkGatewayConnectionsClient) getSharedKeyCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, options *VirtualNetworkGatewayConnectionsGetSharedKeyOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/sharedkey"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayConnectionName}", url.PathEscape(virtualNetworkGatewayConnectionName))
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

// getSharedKeyHandleResponse handles the GetSharedKey response.
func (client VirtualNetworkGatewayConnectionsClient) getSharedKeyHandleResponse(resp *azcore.Response) (*ConnectionSharedKeyResponse, error) {
	result := ConnectionSharedKeyResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ConnectionSharedKey)
}

// getSharedKeyHandleError handles the GetSharedKey error response.
func (client VirtualNetworkGatewayConnectionsClient) getSharedKeyHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - The List VirtualNetworkGatewayConnections operation retrieves all the virtual network gateways connections created.
func (client VirtualNetworkGatewayConnectionsClient) List(resourceGroupName string, options *VirtualNetworkGatewayConnectionsListOptions) VirtualNetworkGatewayConnectionListResultPager {
	return &virtualNetworkGatewayConnectionListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp *VirtualNetworkGatewayConnectionListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.VirtualNetworkGatewayConnectionListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client VirtualNetworkGatewayConnectionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *VirtualNetworkGatewayConnectionsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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
func (client VirtualNetworkGatewayConnectionsClient) listHandleResponse(resp *azcore.Response) (*VirtualNetworkGatewayConnectionListResultResponse, error) {
	result := VirtualNetworkGatewayConnectionListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualNetworkGatewayConnectionListResult)
}

// listHandleError handles the List error response.
func (client VirtualNetworkGatewayConnectionsClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginResetSharedKey - The VirtualNetworkGatewayConnectionResetSharedKey operation resets the virtual network gateway connection shared key for passed
// virtual network gateway connection in the specified resource group
// through Network resource provider.
func (client VirtualNetworkGatewayConnectionsClient) BeginResetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionResetSharedKey, options *VirtualNetworkGatewayConnectionsResetSharedKeyOptions) (*ConnectionResetSharedKeyPollerResponse, error) {
	resp, err := client.ResetSharedKey(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, parameters, options)
	if err != nil {
		return nil, err
	}
	result := &ConnectionResetSharedKeyPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualNetworkGatewayConnectionsClient.ResetSharedKey", "location", resp, client.resetSharedKeyHandleError)
	if err != nil {
		return nil, err
	}
	poller := &connectionResetSharedKeyPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ConnectionResetSharedKeyResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeResetSharedKey creates a new ConnectionResetSharedKeyPoller from the specified resume token.
// token - The value must come from a previous call to ConnectionResetSharedKeyPoller.ResumeToken().
func (client VirtualNetworkGatewayConnectionsClient) ResumeResetSharedKey(token string) (ConnectionResetSharedKeyPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualNetworkGatewayConnectionsClient.ResetSharedKey", token, client.resetSharedKeyHandleError)
	if err != nil {
		return nil, err
	}
	return &connectionResetSharedKeyPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// ResetSharedKey - The VirtualNetworkGatewayConnectionResetSharedKey operation resets the virtual network gateway connection shared key for passed virtual
// network gateway connection in the specified resource group
// through Network resource provider.
func (client VirtualNetworkGatewayConnectionsClient) ResetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionResetSharedKey, options *VirtualNetworkGatewayConnectionsResetSharedKeyOptions) (*azcore.Response, error) {
	req, err := client.resetSharedKeyCreateRequest(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.resetSharedKeyHandleError(resp)
	}
	return resp, nil
}

// resetSharedKeyCreateRequest creates the ResetSharedKey request.
func (client VirtualNetworkGatewayConnectionsClient) resetSharedKeyCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionResetSharedKey, options *VirtualNetworkGatewayConnectionsResetSharedKeyOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/sharedkey/reset"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayConnectionName}", url.PathEscape(virtualNetworkGatewayConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// resetSharedKeyHandleResponse handles the ResetSharedKey response.
func (client VirtualNetworkGatewayConnectionsClient) resetSharedKeyHandleResponse(resp *azcore.Response) (*ConnectionResetSharedKeyResponse, error) {
	result := ConnectionResetSharedKeyResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ConnectionResetSharedKey)
}

// resetSharedKeyHandleError handles the ResetSharedKey error response.
func (client VirtualNetworkGatewayConnectionsClient) resetSharedKeyHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginSetSharedKey - The Put VirtualNetworkGatewayConnectionSharedKey operation sets the virtual network gateway connection shared key for passed virtual
// network gateway connection in the specified resource group through
// Network resource provider.
func (client VirtualNetworkGatewayConnectionsClient) BeginSetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionSharedKey, options *VirtualNetworkGatewayConnectionsSetSharedKeyOptions) (*ConnectionSharedKeyPollerResponse, error) {
	resp, err := client.SetSharedKey(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, parameters, options)
	if err != nil {
		return nil, err
	}
	result := &ConnectionSharedKeyPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualNetworkGatewayConnectionsClient.SetSharedKey", "azure-async-operation", resp, client.setSharedKeyHandleError)
	if err != nil {
		return nil, err
	}
	poller := &connectionSharedKeyPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ConnectionSharedKeyResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeSetSharedKey creates a new ConnectionSharedKeyPoller from the specified resume token.
// token - The value must come from a previous call to ConnectionSharedKeyPoller.ResumeToken().
func (client VirtualNetworkGatewayConnectionsClient) ResumeSetSharedKey(token string) (ConnectionSharedKeyPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualNetworkGatewayConnectionsClient.SetSharedKey", token, client.setSharedKeyHandleError)
	if err != nil {
		return nil, err
	}
	return &connectionSharedKeyPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// SetSharedKey - The Put VirtualNetworkGatewayConnectionSharedKey operation sets the virtual network gateway connection shared key for passed virtual network
// gateway connection in the specified resource group through
// Network resource provider.
func (client VirtualNetworkGatewayConnectionsClient) SetSharedKey(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionSharedKey, options *VirtualNetworkGatewayConnectionsSetSharedKeyOptions) (*azcore.Response, error) {
	req, err := client.setSharedKeyCreateRequest(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.setSharedKeyHandleError(resp)
	}
	return resp, nil
}

// setSharedKeyCreateRequest creates the SetSharedKey request.
func (client VirtualNetworkGatewayConnectionsClient) setSharedKeyCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionSharedKey, options *VirtualNetworkGatewayConnectionsSetSharedKeyOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/sharedkey"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayConnectionName}", url.PathEscape(virtualNetworkGatewayConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// setSharedKeyHandleResponse handles the SetSharedKey response.
func (client VirtualNetworkGatewayConnectionsClient) setSharedKeyHandleResponse(resp *azcore.Response) (*ConnectionSharedKeyResponse, error) {
	result := ConnectionSharedKeyResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ConnectionSharedKey)
}

// setSharedKeyHandleError handles the SetSharedKey error response.
func (client VirtualNetworkGatewayConnectionsClient) setSharedKeyHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginStartPacketCapture - Starts packet capture on virtual network gateway connection in the specified resource group.
func (client VirtualNetworkGatewayConnectionsClient) BeginStartPacketCapture(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, options *VirtualNetworkGatewayConnectionsStartPacketCaptureOptions) (*StringPollerResponse, error) {
	resp, err := client.StartPacketCapture(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, options)
	if err != nil {
		return nil, err
	}
	result := &StringPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualNetworkGatewayConnectionsClient.StartPacketCapture", "location", resp, client.startPacketCaptureHandleError)
	if err != nil {
		return nil, err
	}
	poller := &stringPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*StringResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeStartPacketCapture creates a new StringPoller from the specified resume token.
// token - The value must come from a previous call to StringPoller.ResumeToken().
func (client VirtualNetworkGatewayConnectionsClient) ResumeStartPacketCapture(token string) (StringPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualNetworkGatewayConnectionsClient.StartPacketCapture", token, client.startPacketCaptureHandleError)
	if err != nil {
		return nil, err
	}
	return &stringPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// StartPacketCapture - Starts packet capture on virtual network gateway connection in the specified resource group.
func (client VirtualNetworkGatewayConnectionsClient) StartPacketCapture(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, options *VirtualNetworkGatewayConnectionsStartPacketCaptureOptions) (*azcore.Response, error) {
	req, err := client.startPacketCaptureCreateRequest(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.startPacketCaptureHandleError(resp)
	}
	return resp, nil
}

// startPacketCaptureCreateRequest creates the StartPacketCapture request.
func (client VirtualNetworkGatewayConnectionsClient) startPacketCaptureCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, options *VirtualNetworkGatewayConnectionsStartPacketCaptureOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/startPacketCapture"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayConnectionName}", url.PathEscape(virtualNetworkGatewayConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	if options != nil {
		return req, req.MarshalAsJSON(options.Parameters)
	}
	return req, nil
}

// startPacketCaptureHandleResponse handles the StartPacketCapture response.
func (client VirtualNetworkGatewayConnectionsClient) startPacketCaptureHandleResponse(resp *azcore.Response) (*StringResponse, error) {
	result := StringResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// startPacketCaptureHandleError handles the StartPacketCapture error response.
func (client VirtualNetworkGatewayConnectionsClient) startPacketCaptureHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginStopPacketCapture - Stops packet capture on virtual network gateway connection in the specified resource group.
func (client VirtualNetworkGatewayConnectionsClient) BeginStopPacketCapture(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters VpnPacketCaptureStopParameters, options *VirtualNetworkGatewayConnectionsStopPacketCaptureOptions) (*StringPollerResponse, error) {
	resp, err := client.StopPacketCapture(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, parameters, options)
	if err != nil {
		return nil, err
	}
	result := &StringPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualNetworkGatewayConnectionsClient.StopPacketCapture", "location", resp, client.stopPacketCaptureHandleError)
	if err != nil {
		return nil, err
	}
	poller := &stringPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*StringResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeStopPacketCapture creates a new StringPoller from the specified resume token.
// token - The value must come from a previous call to StringPoller.ResumeToken().
func (client VirtualNetworkGatewayConnectionsClient) ResumeStopPacketCapture(token string) (StringPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualNetworkGatewayConnectionsClient.StopPacketCapture", token, client.stopPacketCaptureHandleError)
	if err != nil {
		return nil, err
	}
	return &stringPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// StopPacketCapture - Stops packet capture on virtual network gateway connection in the specified resource group.
func (client VirtualNetworkGatewayConnectionsClient) StopPacketCapture(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters VpnPacketCaptureStopParameters, options *VirtualNetworkGatewayConnectionsStopPacketCaptureOptions) (*azcore.Response, error) {
	req, err := client.stopPacketCaptureCreateRequest(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.stopPacketCaptureHandleError(resp)
	}
	return resp, nil
}

// stopPacketCaptureCreateRequest creates the StopPacketCapture request.
func (client VirtualNetworkGatewayConnectionsClient) stopPacketCaptureCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters VpnPacketCaptureStopParameters, options *VirtualNetworkGatewayConnectionsStopPacketCaptureOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/stopPacketCapture"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayConnectionName}", url.PathEscape(virtualNetworkGatewayConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// stopPacketCaptureHandleResponse handles the StopPacketCapture response.
func (client VirtualNetworkGatewayConnectionsClient) stopPacketCaptureHandleResponse(resp *azcore.Response) (*StringResponse, error) {
	result := StringResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// stopPacketCaptureHandleError handles the StopPacketCapture error response.
func (client VirtualNetworkGatewayConnectionsClient) stopPacketCaptureHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginUpdateTags - Updates a virtual network gateway connection tags.
func (client VirtualNetworkGatewayConnectionsClient) BeginUpdateTags(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters TagsObject, options *VirtualNetworkGatewayConnectionsUpdateTagsOptions) (*VirtualNetworkGatewayConnectionPollerResponse, error) {
	resp, err := client.UpdateTags(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, parameters, options)
	if err != nil {
		return nil, err
	}
	result := &VirtualNetworkGatewayConnectionPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualNetworkGatewayConnectionsClient.UpdateTags", "azure-async-operation", resp, client.updateTagsHandleError)
	if err != nil {
		return nil, err
	}
	poller := &virtualNetworkGatewayConnectionPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*VirtualNetworkGatewayConnectionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeUpdateTags creates a new VirtualNetworkGatewayConnectionPoller from the specified resume token.
// token - The value must come from a previous call to VirtualNetworkGatewayConnectionPoller.ResumeToken().
func (client VirtualNetworkGatewayConnectionsClient) ResumeUpdateTags(token string) (VirtualNetworkGatewayConnectionPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualNetworkGatewayConnectionsClient.UpdateTags", token, client.updateTagsHandleError)
	if err != nil {
		return nil, err
	}
	return &virtualNetworkGatewayConnectionPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// UpdateTags - Updates a virtual network gateway connection tags.
func (client VirtualNetworkGatewayConnectionsClient) UpdateTags(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters TagsObject, options *VirtualNetworkGatewayConnectionsUpdateTagsOptions) (*azcore.Response, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, virtualNetworkGatewayConnectionName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.updateTagsHandleError(resp)
	}
	return resp, nil
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client VirtualNetworkGatewayConnectionsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters TagsObject, options *VirtualNetworkGatewayConnectionsUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkGatewayConnectionName}", url.PathEscape(virtualNetworkGatewayConnectionName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client VirtualNetworkGatewayConnectionsClient) updateTagsHandleResponse(resp *azcore.Response) (*VirtualNetworkGatewayConnectionResponse, error) {
	result := VirtualNetworkGatewayConnectionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VirtualNetworkGatewayConnection)
}

// updateTagsHandleError handles the UpdateTags error response.
func (client VirtualNetworkGatewayConnectionsClient) updateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
