//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataboxedge

import (
	"context"
	"errors"
	"fmt"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DiagnosticSettingsClient contains the methods for the DiagnosticSettings group.
// Don't use this type directly, use NewDiagnosticSettingsClient() instead.
type DiagnosticSettingsClient struct {
	con            *connection
	subscriptionID string
}

// NewDiagnosticSettingsClient creates a new instance of DiagnosticSettingsClient with the specified values.
func NewDiagnosticSettingsClient(con *connection, subscriptionID string) *DiagnosticSettingsClient {
	return &DiagnosticSettingsClient{con: con, subscriptionID: subscriptionID}
}

// GetDiagnosticProactiveLogCollectionSettings - Gets the proactive log collection settings of the specified Data Box Edge/Data Box Gateway device.
// If the operation fails it returns the *CloudError error type.
func (client *DiagnosticSettingsClient) GetDiagnosticProactiveLogCollectionSettings(ctx context.Context, deviceName string, resourceGroupName string, options *DiagnosticSettingsGetDiagnosticProactiveLogCollectionSettingsOptions) (DiagnosticSettingsGetDiagnosticProactiveLogCollectionSettingsResponse, error) {
	req, err := client.getDiagnosticProactiveLogCollectionSettingsCreateRequest(ctx, deviceName, resourceGroupName, options)
	if err != nil {
		return DiagnosticSettingsGetDiagnosticProactiveLogCollectionSettingsResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DiagnosticSettingsGetDiagnosticProactiveLogCollectionSettingsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DiagnosticSettingsGetDiagnosticProactiveLogCollectionSettingsResponse{}, client.getDiagnosticProactiveLogCollectionSettingsHandleError(resp)
	}
	return client.getDiagnosticProactiveLogCollectionSettingsHandleResponse(resp)
}

// getDiagnosticProactiveLogCollectionSettingsCreateRequest creates the GetDiagnosticProactiveLogCollectionSettings request.
func (client *DiagnosticSettingsClient) getDiagnosticProactiveLogCollectionSettingsCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, options *DiagnosticSettingsGetDiagnosticProactiveLogCollectionSettingsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/diagnosticProactiveLogCollectionSettings/default"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getDiagnosticProactiveLogCollectionSettingsHandleResponse handles the GetDiagnosticProactiveLogCollectionSettings response.
func (client *DiagnosticSettingsClient) getDiagnosticProactiveLogCollectionSettingsHandleResponse(resp *http.Response) (DiagnosticSettingsGetDiagnosticProactiveLogCollectionSettingsResponse, error) {
	result := DiagnosticSettingsGetDiagnosticProactiveLogCollectionSettingsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiagnosticProactiveLogCollectionSettings); err != nil {
		return DiagnosticSettingsGetDiagnosticProactiveLogCollectionSettingsResponse{}, err
	}
	return result, nil
}

// getDiagnosticProactiveLogCollectionSettingsHandleError handles the GetDiagnosticProactiveLogCollectionSettings error response.
func (client *DiagnosticSettingsClient) getDiagnosticProactiveLogCollectionSettingsHandleError(resp *http.Response) error {
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

// GetDiagnosticRemoteSupportSettings - Gets the diagnostic remote support settings of the specified Data Box Edge/Data Box Gateway device.
// If the operation fails it returns the *CloudError error type.
func (client *DiagnosticSettingsClient) GetDiagnosticRemoteSupportSettings(ctx context.Context, deviceName string, resourceGroupName string, options *DiagnosticSettingsGetDiagnosticRemoteSupportSettingsOptions) (DiagnosticSettingsGetDiagnosticRemoteSupportSettingsResponse, error) {
	req, err := client.getDiagnosticRemoteSupportSettingsCreateRequest(ctx, deviceName, resourceGroupName, options)
	if err != nil {
		return DiagnosticSettingsGetDiagnosticRemoteSupportSettingsResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DiagnosticSettingsGetDiagnosticRemoteSupportSettingsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DiagnosticSettingsGetDiagnosticRemoteSupportSettingsResponse{}, client.getDiagnosticRemoteSupportSettingsHandleError(resp)
	}
	return client.getDiagnosticRemoteSupportSettingsHandleResponse(resp)
}

// getDiagnosticRemoteSupportSettingsCreateRequest creates the GetDiagnosticRemoteSupportSettings request.
func (client *DiagnosticSettingsClient) getDiagnosticRemoteSupportSettingsCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, options *DiagnosticSettingsGetDiagnosticRemoteSupportSettingsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/diagnosticRemoteSupportSettings/default"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getDiagnosticRemoteSupportSettingsHandleResponse handles the GetDiagnosticRemoteSupportSettings response.
func (client *DiagnosticSettingsClient) getDiagnosticRemoteSupportSettingsHandleResponse(resp *http.Response) (DiagnosticSettingsGetDiagnosticRemoteSupportSettingsResponse, error) {
	result := DiagnosticSettingsGetDiagnosticRemoteSupportSettingsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiagnosticRemoteSupportSettings); err != nil {
		return DiagnosticSettingsGetDiagnosticRemoteSupportSettingsResponse{}, err
	}
	return result, nil
}

// getDiagnosticRemoteSupportSettingsHandleError handles the GetDiagnosticRemoteSupportSettings error response.
func (client *DiagnosticSettingsClient) getDiagnosticRemoteSupportSettingsHandleError(resp *http.Response) error {
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

// BeginUpdateDiagnosticProactiveLogCollectionSettings - Updates the proactive log collection settings on a Data Box Edge/Data Box Gateway device.
// If the operation fails it returns the *CloudError error type.
func (client *DiagnosticSettingsClient) BeginUpdateDiagnosticProactiveLogCollectionSettings(ctx context.Context, deviceName string, resourceGroupName string, proactiveLogCollectionSettings DiagnosticProactiveLogCollectionSettings, options *DiagnosticSettingsBeginUpdateDiagnosticProactiveLogCollectionSettingsOptions) (DiagnosticSettingsUpdateDiagnosticProactiveLogCollectionSettingsPollerResponse, error) {
	resp, err := client.updateDiagnosticProactiveLogCollectionSettings(ctx, deviceName, resourceGroupName, proactiveLogCollectionSettings, options)
	if err != nil {
		return DiagnosticSettingsUpdateDiagnosticProactiveLogCollectionSettingsPollerResponse{}, err
	}
	result := DiagnosticSettingsUpdateDiagnosticProactiveLogCollectionSettingsPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DiagnosticSettingsClient.UpdateDiagnosticProactiveLogCollectionSettings", "", resp, client.con.Pipeline(), client.updateDiagnosticProactiveLogCollectionSettingsHandleError)
	if err != nil {
		return DiagnosticSettingsUpdateDiagnosticProactiveLogCollectionSettingsPollerResponse{}, err
	}
	result.Poller = &DiagnosticSettingsUpdateDiagnosticProactiveLogCollectionSettingsPoller{
		pt: pt,
	}
	return result, nil
}

// UpdateDiagnosticProactiveLogCollectionSettings - Updates the proactive log collection settings on a Data Box Edge/Data Box Gateway device.
// If the operation fails it returns the *CloudError error type.
func (client *DiagnosticSettingsClient) updateDiagnosticProactiveLogCollectionSettings(ctx context.Context, deviceName string, resourceGroupName string, proactiveLogCollectionSettings DiagnosticProactiveLogCollectionSettings, options *DiagnosticSettingsBeginUpdateDiagnosticProactiveLogCollectionSettingsOptions) (*http.Response, error) {
	req, err := client.updateDiagnosticProactiveLogCollectionSettingsCreateRequest(ctx, deviceName, resourceGroupName, proactiveLogCollectionSettings, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.updateDiagnosticProactiveLogCollectionSettingsHandleError(resp)
	}
	return resp, nil
}

// updateDiagnosticProactiveLogCollectionSettingsCreateRequest creates the UpdateDiagnosticProactiveLogCollectionSettings request.
func (client *DiagnosticSettingsClient) updateDiagnosticProactiveLogCollectionSettingsCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, proactiveLogCollectionSettings DiagnosticProactiveLogCollectionSettings, options *DiagnosticSettingsBeginUpdateDiagnosticProactiveLogCollectionSettingsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/diagnosticProactiveLogCollectionSettings/default"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, proactiveLogCollectionSettings)
}

// updateDiagnosticProactiveLogCollectionSettingsHandleError handles the UpdateDiagnosticProactiveLogCollectionSettings error response.
func (client *DiagnosticSettingsClient) updateDiagnosticProactiveLogCollectionSettingsHandleError(resp *http.Response) error {
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

// BeginUpdateDiagnosticRemoteSupportSettings - Updates the diagnostic remote support settings on a Data Box Edge/Data Box Gateway device.
// If the operation fails it returns the *CloudError error type.
func (client *DiagnosticSettingsClient) BeginUpdateDiagnosticRemoteSupportSettings(ctx context.Context, deviceName string, resourceGroupName string, diagnosticRemoteSupportSettings DiagnosticRemoteSupportSettings, options *DiagnosticSettingsBeginUpdateDiagnosticRemoteSupportSettingsOptions) (DiagnosticSettingsUpdateDiagnosticRemoteSupportSettingsPollerResponse, error) {
	resp, err := client.updateDiagnosticRemoteSupportSettings(ctx, deviceName, resourceGroupName, diagnosticRemoteSupportSettings, options)
	if err != nil {
		return DiagnosticSettingsUpdateDiagnosticRemoteSupportSettingsPollerResponse{}, err
	}
	result := DiagnosticSettingsUpdateDiagnosticRemoteSupportSettingsPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DiagnosticSettingsClient.UpdateDiagnosticRemoteSupportSettings", "", resp, client.con.Pipeline(), client.updateDiagnosticRemoteSupportSettingsHandleError)
	if err != nil {
		return DiagnosticSettingsUpdateDiagnosticRemoteSupportSettingsPollerResponse{}, err
	}
	result.Poller = &DiagnosticSettingsUpdateDiagnosticRemoteSupportSettingsPoller{
		pt: pt,
	}
	return result, nil
}

// UpdateDiagnosticRemoteSupportSettings - Updates the diagnostic remote support settings on a Data Box Edge/Data Box Gateway device.
// If the operation fails it returns the *CloudError error type.
func (client *DiagnosticSettingsClient) updateDiagnosticRemoteSupportSettings(ctx context.Context, deviceName string, resourceGroupName string, diagnosticRemoteSupportSettings DiagnosticRemoteSupportSettings, options *DiagnosticSettingsBeginUpdateDiagnosticRemoteSupportSettingsOptions) (*http.Response, error) {
	req, err := client.updateDiagnosticRemoteSupportSettingsCreateRequest(ctx, deviceName, resourceGroupName, diagnosticRemoteSupportSettings, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.updateDiagnosticRemoteSupportSettingsHandleError(resp)
	}
	return resp, nil
}

// updateDiagnosticRemoteSupportSettingsCreateRequest creates the UpdateDiagnosticRemoteSupportSettings request.
func (client *DiagnosticSettingsClient) updateDiagnosticRemoteSupportSettingsCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, diagnosticRemoteSupportSettings DiagnosticRemoteSupportSettings, options *DiagnosticSettingsBeginUpdateDiagnosticRemoteSupportSettingsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/diagnosticRemoteSupportSettings/default"
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, diagnosticRemoteSupportSettings)
}

// updateDiagnosticRemoteSupportSettingsHandleError handles the UpdateDiagnosticRemoteSupportSettings error response.
func (client *DiagnosticSettingsClient) updateDiagnosticRemoteSupportSettingsHandleError(resp *http.Response) error {
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
