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
	"time"
)

// SecurityRulesClient contains the methods for the SecurityRules group.
// Don't use this type directly, use NewSecurityRulesClient() instead.
type SecurityRulesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewSecurityRulesClient creates a new instance of SecurityRulesClient with the specified values.
func NewSecurityRulesClient(con *armcore.Connection, subscriptionID string) *SecurityRulesClient {
	return &SecurityRulesClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a security rule in the specified network security group.
// If the operation fails it returns the *CloudError error type.
func (client *SecurityRulesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, securityRuleParameters SecurityRule, options *SecurityRulesBeginCreateOrUpdateOptions) (SecurityRulePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, networkSecurityGroupName, securityRuleName, securityRuleParameters, options)
	if err != nil {
		return SecurityRulePollerResponse{}, err
	}
	result := SecurityRulePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("SecurityRulesClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return SecurityRulePollerResponse{}, err
	}
	poller := &securityRulePoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (SecurityRuleResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new SecurityRulePoller from the specified resume token.
// token - The value must come from a previous call to SecurityRulePoller.ResumeToken().
func (client *SecurityRulesClient) ResumeCreateOrUpdate(ctx context.Context, token string) (SecurityRulePollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("SecurityRulesClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return SecurityRulePollerResponse{}, err
	}
	poller := &securityRulePoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return SecurityRulePollerResponse{}, err
	}
	result := SecurityRulePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (SecurityRuleResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a security rule in the specified network security group.
// If the operation fails it returns the *CloudError error type.
func (client *SecurityRulesClient) createOrUpdate(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, securityRuleParameters SecurityRule, options *SecurityRulesBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, securityRuleName, securityRuleParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SecurityRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, securityRuleParameters SecurityRule, options *SecurityRulesBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules/{securityRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityGroupName == "" {
		return nil, errors.New("parameter networkSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
	if securityRuleName == "" {
		return nil, errors.New("parameter securityRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityRuleName}", url.PathEscape(securityRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(securityRuleParameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *SecurityRulesClient) createOrUpdateHandleResponse(resp *azcore.Response) (SecurityRuleResponse, error) {
	var val *SecurityRule
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return SecurityRuleResponse{}, err
	}
	return SecurityRuleResponse{RawResponse: resp.Response, SecurityRule: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *SecurityRulesClient) createOrUpdateHandleError(resp *azcore.Response) error {
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

// BeginDelete - Deletes the specified network security rule.
// If the operation fails it returns the *CloudError error type.
func (client *SecurityRulesClient) BeginDelete(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, options *SecurityRulesBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, networkSecurityGroupName, securityRuleName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("SecurityRulesClient.Delete", "location", resp, client.deleteHandleError)
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

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *SecurityRulesClient) ResumeDelete(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("SecurityRulesClient.Delete", token, client.deleteHandleError)
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

// Delete - Deletes the specified network security rule.
// If the operation fails it returns the *CloudError error type.
func (client *SecurityRulesClient) deleteOperation(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, options *SecurityRulesBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, securityRuleName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SecurityRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, options *SecurityRulesBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules/{securityRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityGroupName == "" {
		return nil, errors.New("parameter networkSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
	if securityRuleName == "" {
		return nil, errors.New("parameter securityRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityRuleName}", url.PathEscape(securityRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
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

// deleteHandleError handles the Delete error response.
func (client *SecurityRulesClient) deleteHandleError(resp *azcore.Response) error {
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

// Get - Get the specified network security rule.
// If the operation fails it returns the *CloudError error type.
func (client *SecurityRulesClient) Get(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, options *SecurityRulesGetOptions) (SecurityRuleResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, securityRuleName, options)
	if err != nil {
		return SecurityRuleResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SecurityRuleResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SecurityRuleResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SecurityRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, securityRuleName string, options *SecurityRulesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules/{securityRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkSecurityGroupName == "" {
		return nil, errors.New("parameter networkSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkSecurityGroupName}", url.PathEscape(networkSecurityGroupName))
	if securityRuleName == "" {
		return nil, errors.New("parameter securityRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityRuleName}", url.PathEscape(securityRuleName))
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
func (client *SecurityRulesClient) getHandleResponse(resp *azcore.Response) (SecurityRuleResponse, error) {
	var val *SecurityRule
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return SecurityRuleResponse{}, err
	}
	return SecurityRuleResponse{RawResponse: resp.Response, SecurityRule: val}, nil
}

// getHandleError handles the Get error response.
func (client *SecurityRulesClient) getHandleError(resp *azcore.Response) error {
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

// List - Gets all security rules in a network security group.
// If the operation fails it returns the *CloudError error type.
func (client *SecurityRulesClient) List(resourceGroupName string, networkSecurityGroupName string, options *SecurityRulesListOptions) SecurityRuleListResultPager {
	return &securityRuleListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, networkSecurityGroupName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp SecurityRuleListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.SecurityRuleListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *SecurityRulesClient) listCreateRequest(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, options *SecurityRulesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules"
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
func (client *SecurityRulesClient) listHandleResponse(resp *azcore.Response) (SecurityRuleListResultResponse, error) {
	var val *SecurityRuleListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return SecurityRuleListResultResponse{}, err
	}
	return SecurityRuleListResultResponse{RawResponse: resp.Response, SecurityRuleListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *SecurityRulesClient) listHandleError(resp *azcore.Response) error {
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
