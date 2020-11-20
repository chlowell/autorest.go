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

// ExpressRouteCircuitsClient contains the methods for the ExpressRouteCircuits group.
// Don't use this type directly, use NewExpressRouteCircuitsClient() instead.
type ExpressRouteCircuitsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewExpressRouteCircuitsClient creates a new instance of ExpressRouteCircuitsClient with the specified values.
func NewExpressRouteCircuitsClient(con *armcore.Connection, subscriptionID string) ExpressRouteCircuitsClient {
	return ExpressRouteCircuitsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client ExpressRouteCircuitsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// BeginCreateOrUpdate - Creates or updates an express route circuit.
func (client ExpressRouteCircuitsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, parameters ExpressRouteCircuit, options *ExpressRouteCircuitsCreateOrUpdateOptions) (*ExpressRouteCircuitPollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, circuitName, parameters, options)
	if err != nil {
		return nil, err
	}
	result := &ExpressRouteCircuitPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRouteCircuitsClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &expressRouteCircuitPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ExpressRouteCircuitResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new ExpressRouteCircuitPoller from the specified resume token.
// token - The value must come from a previous call to ExpressRouteCircuitPoller.ResumeToken().
func (client ExpressRouteCircuitsClient) ResumeCreateOrUpdate(token string) (ExpressRouteCircuitPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteCircuitsClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &expressRouteCircuitPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates an express route circuit.
func (client ExpressRouteCircuitsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, circuitName string, parameters ExpressRouteCircuit, options *ExpressRouteCircuitsCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, circuitName, parameters, options)
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
func (client ExpressRouteCircuitsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, parameters ExpressRouteCircuit, options *ExpressRouteCircuitsCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
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
func (client ExpressRouteCircuitsClient) createOrUpdateHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitResponse, error) {
	result := ExpressRouteCircuitResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCircuit)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client ExpressRouteCircuitsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Deletes the specified express route circuit.
func (client ExpressRouteCircuitsClient) BeginDelete(ctx context.Context, resourceGroupName string, circuitName string, options *ExpressRouteCircuitsDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, circuitName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRouteCircuitsClient.Delete", "location", resp, client.deleteHandleError)
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
func (client ExpressRouteCircuitsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteCircuitsClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified express route circuit.
func (client ExpressRouteCircuitsClient) Delete(ctx context.Context, resourceGroupName string, circuitName string, options *ExpressRouteCircuitsDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, circuitName, options)
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
func (client ExpressRouteCircuitsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, options *ExpressRouteCircuitsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
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
func (client ExpressRouteCircuitsClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets information about the specified express route circuit.
func (client ExpressRouteCircuitsClient) Get(ctx context.Context, resourceGroupName string, circuitName string, options *ExpressRouteCircuitsGetOptions) (*ExpressRouteCircuitResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, circuitName, options)
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
func (client ExpressRouteCircuitsClient) getCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, options *ExpressRouteCircuitsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
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
func (client ExpressRouteCircuitsClient) getHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitResponse, error) {
	result := ExpressRouteCircuitResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCircuit)
}

// getHandleError handles the Get error response.
func (client ExpressRouteCircuitsClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetPeeringStats - Gets all stats from an express route circuit in a resource group.
func (client ExpressRouteCircuitsClient) GetPeeringStats(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, options *ExpressRouteCircuitsGetPeeringStatsOptions) (*ExpressRouteCircuitStatsResponse, error) {
	req, err := client.getPeeringStatsCreateRequest(ctx, resourceGroupName, circuitName, peeringName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getPeeringStatsHandleError(resp)
	}
	result, err := client.getPeeringStatsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getPeeringStatsCreateRequest creates the GetPeeringStats request.
func (client ExpressRouteCircuitsClient) getPeeringStatsCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, options *ExpressRouteCircuitsGetPeeringStatsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/stats"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
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

// getPeeringStatsHandleResponse handles the GetPeeringStats response.
func (client ExpressRouteCircuitsClient) getPeeringStatsHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitStatsResponse, error) {
	result := ExpressRouteCircuitStatsResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCircuitStats)
}

// getPeeringStatsHandleError handles the GetPeeringStats error response.
func (client ExpressRouteCircuitsClient) getPeeringStatsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetStats - Gets all the stats from an express route circuit in a resource group.
func (client ExpressRouteCircuitsClient) GetStats(ctx context.Context, resourceGroupName string, circuitName string, options *ExpressRouteCircuitsGetStatsOptions) (*ExpressRouteCircuitStatsResponse, error) {
	req, err := client.getStatsCreateRequest(ctx, resourceGroupName, circuitName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getStatsHandleError(resp)
	}
	result, err := client.getStatsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getStatsCreateRequest creates the GetStats request.
func (client ExpressRouteCircuitsClient) getStatsCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, options *ExpressRouteCircuitsGetStatsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/stats"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
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

// getStatsHandleResponse handles the GetStats response.
func (client ExpressRouteCircuitsClient) getStatsHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitStatsResponse, error) {
	result := ExpressRouteCircuitStatsResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCircuitStats)
}

// getStatsHandleError handles the GetStats error response.
func (client ExpressRouteCircuitsClient) getStatsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all the express route circuits in a resource group.
func (client ExpressRouteCircuitsClient) List(resourceGroupName string, options *ExpressRouteCircuitsListOptions) ExpressRouteCircuitListResultPager {
	return &expressRouteCircuitListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp *ExpressRouteCircuitListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ExpressRouteCircuitListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client ExpressRouteCircuitsClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *ExpressRouteCircuitsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits"
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
func (client ExpressRouteCircuitsClient) listHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitListResultResponse, error) {
	result := ExpressRouteCircuitListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCircuitListResult)
}

// listHandleError handles the List error response.
func (client ExpressRouteCircuitsClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListAll - Gets all the express route circuits in a subscription.
func (client ExpressRouteCircuitsClient) ListAll(options *ExpressRouteCircuitsListAllOptions) ExpressRouteCircuitListResultPager {
	return &expressRouteCircuitListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listAllCreateRequest(ctx, options)
		},
		responder: client.listAllHandleResponse,
		errorer:   client.listAllHandleError,
		advancer: func(ctx context.Context, resp *ExpressRouteCircuitListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ExpressRouteCircuitListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listAllCreateRequest creates the ListAll request.
func (client ExpressRouteCircuitsClient) listAllCreateRequest(ctx context.Context, options *ExpressRouteCircuitsListAllOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteCircuits"
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

// listAllHandleResponse handles the ListAll response.
func (client ExpressRouteCircuitsClient) listAllHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitListResultResponse, error) {
	result := ExpressRouteCircuitListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCircuitListResult)
}

// listAllHandleError handles the ListAll error response.
func (client ExpressRouteCircuitsClient) listAllHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginListArpTable - Gets the currently advertised ARP table associated with the express route circuit in a resource group.
func (client ExpressRouteCircuitsClient) BeginListArpTable(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string, options *ExpressRouteCircuitsListArpTableOptions) (*ExpressRouteCircuitsArpTableListResultPollerResponse, error) {
	resp, err := client.ListArpTable(ctx, resourceGroupName, circuitName, peeringName, devicePath, options)
	if err != nil {
		return nil, err
	}
	result := &ExpressRouteCircuitsArpTableListResultPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRouteCircuitsClient.ListArpTable", "location", resp, client.listArpTableHandleError)
	if err != nil {
		return nil, err
	}
	poller := &expressRouteCircuitsArpTableListResultPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ExpressRouteCircuitsArpTableListResultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeListArpTable creates a new ExpressRouteCircuitsArpTableListResultPoller from the specified resume token.
// token - The value must come from a previous call to ExpressRouteCircuitsArpTableListResultPoller.ResumeToken().
func (client ExpressRouteCircuitsClient) ResumeListArpTable(token string) (ExpressRouteCircuitsArpTableListResultPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteCircuitsClient.ListArpTable", token, client.listArpTableHandleError)
	if err != nil {
		return nil, err
	}
	return &expressRouteCircuitsArpTableListResultPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// ListArpTable - Gets the currently advertised ARP table associated with the express route circuit in a resource group.
func (client ExpressRouteCircuitsClient) ListArpTable(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string, options *ExpressRouteCircuitsListArpTableOptions) (*azcore.Response, error) {
	req, err := client.listArpTableCreateRequest(ctx, resourceGroupName, circuitName, peeringName, devicePath, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.listArpTableHandleError(resp)
	}
	return resp, nil
}

// listArpTableCreateRequest creates the ListArpTable request.
func (client ExpressRouteCircuitsClient) listArpTableCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string, options *ExpressRouteCircuitsListArpTableOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/arpTables/{devicePath}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	urlPath = strings.ReplaceAll(urlPath, "{devicePath}", url.PathEscape(devicePath))
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
	return req, nil
}

// listArpTableHandleResponse handles the ListArpTable response.
func (client ExpressRouteCircuitsClient) listArpTableHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitsArpTableListResultResponse, error) {
	result := ExpressRouteCircuitsArpTableListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCircuitsArpTableListResult)
}

// listArpTableHandleError handles the ListArpTable error response.
func (client ExpressRouteCircuitsClient) listArpTableHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginListRoutesTable - Gets the currently advertised routes table associated with the express route circuit in a resource group.
func (client ExpressRouteCircuitsClient) BeginListRoutesTable(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string, options *ExpressRouteCircuitsListRoutesTableOptions) (*ExpressRouteCircuitsRoutesTableListResultPollerResponse, error) {
	resp, err := client.ListRoutesTable(ctx, resourceGroupName, circuitName, peeringName, devicePath, options)
	if err != nil {
		return nil, err
	}
	result := &ExpressRouteCircuitsRoutesTableListResultPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRouteCircuitsClient.ListRoutesTable", "location", resp, client.listRoutesTableHandleError)
	if err != nil {
		return nil, err
	}
	poller := &expressRouteCircuitsRoutesTableListResultPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ExpressRouteCircuitsRoutesTableListResultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeListRoutesTable creates a new ExpressRouteCircuitsRoutesTableListResultPoller from the specified resume token.
// token - The value must come from a previous call to ExpressRouteCircuitsRoutesTableListResultPoller.ResumeToken().
func (client ExpressRouteCircuitsClient) ResumeListRoutesTable(token string) (ExpressRouteCircuitsRoutesTableListResultPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteCircuitsClient.ListRoutesTable", token, client.listRoutesTableHandleError)
	if err != nil {
		return nil, err
	}
	return &expressRouteCircuitsRoutesTableListResultPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// ListRoutesTable - Gets the currently advertised routes table associated with the express route circuit in a resource group.
func (client ExpressRouteCircuitsClient) ListRoutesTable(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string, options *ExpressRouteCircuitsListRoutesTableOptions) (*azcore.Response, error) {
	req, err := client.listRoutesTableCreateRequest(ctx, resourceGroupName, circuitName, peeringName, devicePath, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.listRoutesTableHandleError(resp)
	}
	return resp, nil
}

// listRoutesTableCreateRequest creates the ListRoutesTable request.
func (client ExpressRouteCircuitsClient) listRoutesTableCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string, options *ExpressRouteCircuitsListRoutesTableOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/routeTables/{devicePath}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	urlPath = strings.ReplaceAll(urlPath, "{devicePath}", url.PathEscape(devicePath))
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
	return req, nil
}

// listRoutesTableHandleResponse handles the ListRoutesTable response.
func (client ExpressRouteCircuitsClient) listRoutesTableHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitsRoutesTableListResultResponse, error) {
	result := ExpressRouteCircuitsRoutesTableListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCircuitsRoutesTableListResult)
}

// listRoutesTableHandleError handles the ListRoutesTable error response.
func (client ExpressRouteCircuitsClient) listRoutesTableHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginListRoutesTableSummary - Gets the currently advertised routes table summary associated with the express route circuit in a resource group.
func (client ExpressRouteCircuitsClient) BeginListRoutesTableSummary(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string, options *ExpressRouteCircuitsListRoutesTableSummaryOptions) (*ExpressRouteCircuitsRoutesTableSummaryListResultPollerResponse, error) {
	resp, err := client.ListRoutesTableSummary(ctx, resourceGroupName, circuitName, peeringName, devicePath, options)
	if err != nil {
		return nil, err
	}
	result := &ExpressRouteCircuitsRoutesTableSummaryListResultPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("ExpressRouteCircuitsClient.ListRoutesTableSummary", "location", resp, client.listRoutesTableSummaryHandleError)
	if err != nil {
		return nil, err
	}
	poller := &expressRouteCircuitsRoutesTableSummaryListResultPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ExpressRouteCircuitsRoutesTableSummaryListResultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeListRoutesTableSummary creates a new ExpressRouteCircuitsRoutesTableSummaryListResultPoller from the specified resume token.
// token - The value must come from a previous call to ExpressRouteCircuitsRoutesTableSummaryListResultPoller.ResumeToken().
func (client ExpressRouteCircuitsClient) ResumeListRoutesTableSummary(token string) (ExpressRouteCircuitsRoutesTableSummaryListResultPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("ExpressRouteCircuitsClient.ListRoutesTableSummary", token, client.listRoutesTableSummaryHandleError)
	if err != nil {
		return nil, err
	}
	return &expressRouteCircuitsRoutesTableSummaryListResultPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// ListRoutesTableSummary - Gets the currently advertised routes table summary associated with the express route circuit in a resource group.
func (client ExpressRouteCircuitsClient) ListRoutesTableSummary(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string, options *ExpressRouteCircuitsListRoutesTableSummaryOptions) (*azcore.Response, error) {
	req, err := client.listRoutesTableSummaryCreateRequest(ctx, resourceGroupName, circuitName, peeringName, devicePath, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.listRoutesTableSummaryHandleError(resp)
	}
	return resp, nil
}

// listRoutesTableSummaryCreateRequest creates the ListRoutesTableSummary request.
func (client ExpressRouteCircuitsClient) listRoutesTableSummaryCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, peeringName string, devicePath string, options *ExpressRouteCircuitsListRoutesTableSummaryOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/peerings/{peeringName}/routeTablesSummary/{devicePath}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
	urlPath = strings.ReplaceAll(urlPath, "{devicePath}", url.PathEscape(devicePath))
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
	return req, nil
}

// listRoutesTableSummaryHandleResponse handles the ListRoutesTableSummary response.
func (client ExpressRouteCircuitsClient) listRoutesTableSummaryHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitsRoutesTableSummaryListResultResponse, error) {
	result := ExpressRouteCircuitsRoutesTableSummaryListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCircuitsRoutesTableSummaryListResult)
}

// listRoutesTableSummaryHandleError handles the ListRoutesTableSummary error response.
func (client ExpressRouteCircuitsClient) listRoutesTableSummaryHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// UpdateTags - Updates an express route circuit tags.
func (client ExpressRouteCircuitsClient) UpdateTags(ctx context.Context, resourceGroupName string, circuitName string, parameters TagsObject, options *ExpressRouteCircuitsUpdateTagsOptions) (*ExpressRouteCircuitResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, circuitName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.updateTagsHandleError(resp)
	}
	result, err := client.updateTagsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client ExpressRouteCircuitsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, circuitName string, parameters TagsObject, options *ExpressRouteCircuitsUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{circuitName}", url.PathEscape(circuitName))
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
func (client ExpressRouteCircuitsClient) updateTagsHandleResponse(resp *azcore.Response) (*ExpressRouteCircuitResponse, error) {
	result := ExpressRouteCircuitResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRouteCircuit)
}

// updateTagsHandleError handles the UpdateTags error response.
func (client ExpressRouteCircuitsClient) updateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
