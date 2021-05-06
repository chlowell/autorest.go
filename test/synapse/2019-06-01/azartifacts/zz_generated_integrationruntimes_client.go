// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

type integrationRuntimesClient struct {
	con *connection
}

// Get - Get Integration Runtime
// If the operation fails it returns the *ErrorContract error type.
func (client *integrationRuntimesClient) Get(ctx context.Context, integrationRuntimeName string, options *IntegrationRuntimesGetOptions) (IntegrationRuntimeResourceResponse, error) {
	req, err := client.getCreateRequest(ctx, integrationRuntimeName, options)
	if err != nil {
		return IntegrationRuntimeResourceResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return IntegrationRuntimeResourceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return IntegrationRuntimeResourceResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *integrationRuntimesClient) getCreateRequest(ctx context.Context, integrationRuntimeName string, options *IntegrationRuntimesGetOptions) (*azcore.Request, error) {
	urlPath := "/integrationRuntimes/{integrationRuntimeName}"
	if integrationRuntimeName == "" {
		return nil, errors.New("parameter integrationRuntimeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationRuntimeName}", url.PathEscape(integrationRuntimeName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *integrationRuntimesClient) getHandleResponse(resp *azcore.Response) (IntegrationRuntimeResourceResponse, error) {
	var val *IntegrationRuntimeResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return IntegrationRuntimeResourceResponse{}, err
	}
	return IntegrationRuntimeResourceResponse{RawResponse: resp.Response, IntegrationRuntimeResource: val}, nil
}

// getHandleError handles the Get error response.
func (client *integrationRuntimesClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorContract{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// List - List Integration Runtimes
// If the operation fails it returns the *ErrorContract error type.
func (client *integrationRuntimesClient) List(ctx context.Context, options *IntegrationRuntimesListOptions) (IntegrationRuntimeListResponseResponse, error) {
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return IntegrationRuntimeListResponseResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return IntegrationRuntimeListResponseResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return IntegrationRuntimeListResponseResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *integrationRuntimesClient) listCreateRequest(ctx context.Context, options *IntegrationRuntimesListOptions) (*azcore.Request, error) {
	urlPath := "/integrationRuntimes"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *integrationRuntimesClient) listHandleResponse(resp *azcore.Response) (IntegrationRuntimeListResponseResponse, error) {
	var val *IntegrationRuntimeListResponse
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return IntegrationRuntimeListResponseResponse{}, err
	}
	return IntegrationRuntimeListResponseResponse{RawResponse: resp.Response, IntegrationRuntimeListResponse: val}, nil
}

// listHandleError handles the List error response.
func (client *integrationRuntimesClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorContract{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
