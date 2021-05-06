// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azurespecialsgroup

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// HeaderClient contains the methods for the Header group.
// Don't use this type directly, use NewHeaderClient() instead.
type HeaderClient struct {
	con *Connection
}

// NewHeaderClient creates a new instance of HeaderClient with the specified values.
func NewHeaderClient(con *Connection) *HeaderClient {
	return &HeaderClient{con: con}
}

// CustomNamedRequestID - Send foo-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 in the header of the request
// If the operation fails it returns the *Error error type.
func (client *HeaderClient) CustomNamedRequestID(ctx context.Context, fooClientRequestID string, options *HeaderCustomNamedRequestIDOptions) (HeaderCustomNamedRequestIDResponse, error) {
	req, err := client.customNamedRequestIDCreateRequest(ctx, fooClientRequestID, options)
	if err != nil {
		return HeaderCustomNamedRequestIDResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return HeaderCustomNamedRequestIDResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return HeaderCustomNamedRequestIDResponse{}, client.customNamedRequestIDHandleError(resp)
	}
	return client.customNamedRequestIDHandleResponse(resp)
}

// customNamedRequestIDCreateRequest creates the CustomNamedRequestID request.
func (client *HeaderClient) customNamedRequestIDCreateRequest(ctx context.Context, fooClientRequestID string, options *HeaderCustomNamedRequestIDOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/customNamedRequestId"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("foo-client-request-id", fooClientRequestID)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// customNamedRequestIDHandleResponse handles the CustomNamedRequestID response.
func (client *HeaderClient) customNamedRequestIDHandleResponse(resp *azcore.Response) (HeaderCustomNamedRequestIDResponse, error) {
	result := HeaderCustomNamedRequestIDResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("foo-request-id"); val != "" {
		result.FooRequestID = &val
	}
	return result, nil
}

// customNamedRequestIDHandleError handles the CustomNamedRequestID error response.
func (client *HeaderClient) customNamedRequestIDHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := Error{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// CustomNamedRequestIDHead - Send foo-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 in the header of the request
// If the operation fails it returns the *Error error type.
func (client *HeaderClient) CustomNamedRequestIDHead(ctx context.Context, fooClientRequestID string, options *HeaderCustomNamedRequestIDHeadOptions) (HeaderCustomNamedRequestIDHeadResponse, error) {
	req, err := client.customNamedRequestIDHeadCreateRequest(ctx, fooClientRequestID, options)
	if err != nil {
		return HeaderCustomNamedRequestIDHeadResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return HeaderCustomNamedRequestIDHeadResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNotFound) {
		return HeaderCustomNamedRequestIDHeadResponse{}, client.customNamedRequestIDHeadHandleError(resp)
	}
	return client.customNamedRequestIDHeadHandleResponse(resp)
}

// customNamedRequestIDHeadCreateRequest creates the CustomNamedRequestIDHead request.
func (client *HeaderClient) customNamedRequestIDHeadCreateRequest(ctx context.Context, fooClientRequestID string, options *HeaderCustomNamedRequestIDHeadOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/customNamedRequestIdHead"
	req, err := azcore.NewRequest(ctx, http.MethodHead, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("foo-client-request-id", fooClientRequestID)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// customNamedRequestIDHeadHandleResponse handles the CustomNamedRequestIDHead response.
func (client *HeaderClient) customNamedRequestIDHeadHandleResponse(resp *azcore.Response) (HeaderCustomNamedRequestIDHeadResponse, error) {
	result := HeaderCustomNamedRequestIDHeadResponse{RawResponse: resp.Response}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	if val := resp.Header.Get("foo-request-id"); val != "" {
		result.FooRequestID = &val
	}
	return result, nil
}

// customNamedRequestIDHeadHandleError handles the CustomNamedRequestIDHead error response.
func (client *HeaderClient) customNamedRequestIDHeadHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := Error{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// CustomNamedRequestIDParamGrouping - Send foo-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 in the header of the request, via a parameter group
// If the operation fails it returns the *Error error type.
func (client *HeaderClient) CustomNamedRequestIDParamGrouping(ctx context.Context, headerCustomNamedRequestIDParamGroupingParameters HeaderCustomNamedRequestIDParamGroupingParameters) (HeaderCustomNamedRequestIDParamGroupingResponse, error) {
	req, err := client.customNamedRequestIDParamGroupingCreateRequest(ctx, headerCustomNamedRequestIDParamGroupingParameters)
	if err != nil {
		return HeaderCustomNamedRequestIDParamGroupingResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return HeaderCustomNamedRequestIDParamGroupingResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return HeaderCustomNamedRequestIDParamGroupingResponse{}, client.customNamedRequestIDParamGroupingHandleError(resp)
	}
	return client.customNamedRequestIDParamGroupingHandleResponse(resp)
}

// customNamedRequestIDParamGroupingCreateRequest creates the CustomNamedRequestIDParamGrouping request.
func (client *HeaderClient) customNamedRequestIDParamGroupingCreateRequest(ctx context.Context, headerCustomNamedRequestIDParamGroupingParameters HeaderCustomNamedRequestIDParamGroupingParameters) (*azcore.Request, error) {
	urlPath := "/azurespecials/customNamedRequestIdParamGrouping"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("foo-client-request-id", headerCustomNamedRequestIDParamGroupingParameters.FooClientRequestID)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// customNamedRequestIDParamGroupingHandleResponse handles the CustomNamedRequestIDParamGrouping response.
func (client *HeaderClient) customNamedRequestIDParamGroupingHandleResponse(resp *azcore.Response) (HeaderCustomNamedRequestIDParamGroupingResponse, error) {
	result := HeaderCustomNamedRequestIDParamGroupingResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("foo-request-id"); val != "" {
		result.FooRequestID = &val
	}
	return result, nil
}

// customNamedRequestIDParamGroupingHandleError handles the CustomNamedRequestIDParamGrouping error response.
func (client *HeaderClient) customNamedRequestIDParamGroupingHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := Error{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
