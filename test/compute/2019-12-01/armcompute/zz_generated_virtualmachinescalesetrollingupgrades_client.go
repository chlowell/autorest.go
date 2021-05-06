// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// VirtualMachineScaleSetRollingUpgradesClient contains the methods for the VirtualMachineScaleSetRollingUpgrades group.
// Don't use this type directly, use NewVirtualMachineScaleSetRollingUpgradesClient() instead.
type VirtualMachineScaleSetRollingUpgradesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVirtualMachineScaleSetRollingUpgradesClient creates a new instance of VirtualMachineScaleSetRollingUpgradesClient with the specified values.
func NewVirtualMachineScaleSetRollingUpgradesClient(con *armcore.Connection, subscriptionID string) *VirtualMachineScaleSetRollingUpgradesClient {
	return &VirtualMachineScaleSetRollingUpgradesClient{con: con, subscriptionID: subscriptionID}
}

// BeginCancel - Cancels the current virtual machine scale set rolling upgrade.
// If the operation fails it returns a generic error.
func (client *VirtualMachineScaleSetRollingUpgradesClient) BeginCancel(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetRollingUpgradesBeginCancelOptions) (HTTPPollerResponse, error) {
	resp, err := client.cancel(ctx, resourceGroupName, vmScaleSetName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualMachineScaleSetRollingUpgradesClient.Cancel", "", resp, client.cancelHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCancel creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *VirtualMachineScaleSetRollingUpgradesClient) ResumeCancel(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualMachineScaleSetRollingUpgradesClient.Cancel", token, client.cancelHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Cancel - Cancels the current virtual machine scale set rolling upgrade.
// If the operation fails it returns a generic error.
func (client *VirtualMachineScaleSetRollingUpgradesClient) cancel(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetRollingUpgradesBeginCancelOptions) (*azcore.Response, error) {
	req, err := client.cancelCreateRequest(ctx, resourceGroupName, vmScaleSetName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.cancelHandleError(resp)
	}
	return resp, nil
}

// cancelCreateRequest creates the Cancel request.
func (client *VirtualMachineScaleSetRollingUpgradesClient) cancelCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetRollingUpgradesBeginCancelOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/rollingUpgrades/cancel"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.URL.RawQuery = reqQP.Encode()
	return req, nil
}

// cancelHandleError handles the Cancel error response.
func (client *VirtualMachineScaleSetRollingUpgradesClient) cancelHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// GetLatest - Gets the status of the latest virtual machine scale set rolling upgrade.
// If the operation fails it returns a generic error.
func (client *VirtualMachineScaleSetRollingUpgradesClient) GetLatest(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetRollingUpgradesGetLatestOptions) (RollingUpgradeStatusInfoResponse, error) {
	req, err := client.getLatestCreateRequest(ctx, resourceGroupName, vmScaleSetName, options)
	if err != nil {
		return RollingUpgradeStatusInfoResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return RollingUpgradeStatusInfoResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return RollingUpgradeStatusInfoResponse{}, client.getLatestHandleError(resp)
	}
	return client.getLatestHandleResponse(resp)
}

// getLatestCreateRequest creates the GetLatest request.
func (client *VirtualMachineScaleSetRollingUpgradesClient) getLatestCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetRollingUpgradesGetLatestOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/rollingUpgrades/latest"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
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
	reqQP.Set("api-version", "2019-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getLatestHandleResponse handles the GetLatest response.
func (client *VirtualMachineScaleSetRollingUpgradesClient) getLatestHandleResponse(resp *azcore.Response) (RollingUpgradeStatusInfoResponse, error) {
	var val *RollingUpgradeStatusInfo
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RollingUpgradeStatusInfoResponse{}, err
	}
	return RollingUpgradeStatusInfoResponse{RawResponse: resp.Response, RollingUpgradeStatusInfo: val}, nil
}

// getLatestHandleError handles the GetLatest error response.
func (client *VirtualMachineScaleSetRollingUpgradesClient) getLatestHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// BeginStartExtensionUpgrade - Starts a rolling upgrade to move all extensions for all virtual machine scale set instances to the latest available extension
// version. Instances which are already running the latest extension versions
// are not affected.
// If the operation fails it returns a generic error.
func (client *VirtualMachineScaleSetRollingUpgradesClient) BeginStartExtensionUpgrade(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetRollingUpgradesBeginStartExtensionUpgradeOptions) (HTTPPollerResponse, error) {
	resp, err := client.startExtensionUpgrade(ctx, resourceGroupName, vmScaleSetName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualMachineScaleSetRollingUpgradesClient.StartExtensionUpgrade", "", resp, client.startExtensionUpgradeHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeStartExtensionUpgrade creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *VirtualMachineScaleSetRollingUpgradesClient) ResumeStartExtensionUpgrade(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualMachineScaleSetRollingUpgradesClient.StartExtensionUpgrade", token, client.startExtensionUpgradeHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// StartExtensionUpgrade - Starts a rolling upgrade to move all extensions for all virtual machine scale set instances to the latest available extension
// version. Instances which are already running the latest extension versions
// are not affected.
// If the operation fails it returns a generic error.
func (client *VirtualMachineScaleSetRollingUpgradesClient) startExtensionUpgrade(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetRollingUpgradesBeginStartExtensionUpgradeOptions) (*azcore.Response, error) {
	req, err := client.startExtensionUpgradeCreateRequest(ctx, resourceGroupName, vmScaleSetName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.startExtensionUpgradeHandleError(resp)
	}
	return resp, nil
}

// startExtensionUpgradeCreateRequest creates the StartExtensionUpgrade request.
func (client *VirtualMachineScaleSetRollingUpgradesClient) startExtensionUpgradeCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetRollingUpgradesBeginStartExtensionUpgradeOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/extensionRollingUpgrade"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.URL.RawQuery = reqQP.Encode()
	return req, nil
}

// startExtensionUpgradeHandleError handles the StartExtensionUpgrade error response.
func (client *VirtualMachineScaleSetRollingUpgradesClient) startExtensionUpgradeHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// BeginStartOSUpgrade - Starts a rolling upgrade to move all virtual machine scale set instances to the latest available Platform Image OS version. Instances
// which are already running the latest available OS version are not
// affected.
// If the operation fails it returns a generic error.
func (client *VirtualMachineScaleSetRollingUpgradesClient) BeginStartOSUpgrade(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetRollingUpgradesBeginStartOSUpgradeOptions) (HTTPPollerResponse, error) {
	resp, err := client.startOSUpgrade(ctx, resourceGroupName, vmScaleSetName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VirtualMachineScaleSetRollingUpgradesClient.StartOSUpgrade", "", resp, client.startOSUpgradeHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeStartOSUpgrade creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *VirtualMachineScaleSetRollingUpgradesClient) ResumeStartOSUpgrade(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("VirtualMachineScaleSetRollingUpgradesClient.StartOSUpgrade", token, client.startOSUpgradeHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// StartOSUpgrade - Starts a rolling upgrade to move all virtual machine scale set instances to the latest available Platform Image OS version. Instances
// which are already running the latest available OS version are not
// affected.
// If the operation fails it returns a generic error.
func (client *VirtualMachineScaleSetRollingUpgradesClient) startOSUpgrade(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetRollingUpgradesBeginStartOSUpgradeOptions) (*azcore.Response, error) {
	req, err := client.startOSUpgradeCreateRequest(ctx, resourceGroupName, vmScaleSetName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.startOSUpgradeHandleError(resp)
	}
	return resp, nil
}

// startOSUpgradeCreateRequest creates the StartOSUpgrade request.
func (client *VirtualMachineScaleSetRollingUpgradesClient) startOSUpgradeCreateRequest(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *VirtualMachineScaleSetRollingUpgradesBeginStartOSUpgradeOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}/osRollingUpgrade"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmScaleSetName == "" {
		return nil, errors.New("parameter vmScaleSetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmScaleSetName}", url.PathEscape(vmScaleSetName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.URL.RawQuery = reqQP.Encode()
	return req, nil
}

// startOSUpgradeHandleError handles the StartOSUpgrade error response.
func (client *VirtualMachineScaleSetRollingUpgradesClient) startOSUpgradeHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
