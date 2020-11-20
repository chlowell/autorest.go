// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package httpinfrastructuregroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// HTTPServerFailureClient contains the methods for the HTTPServerFailure group.
// Don't use this type directly, use NewHTTPServerFailureClient() instead.
type HTTPServerFailureClient struct {
	con *Connection
}

// NewHTTPServerFailureClient creates a new instance of HTTPServerFailureClient with the specified values.
func NewHTTPServerFailureClient(con *Connection) HTTPServerFailureClient {
	return HTTPServerFailureClient{con: con}
}

// Pipeline returns the pipeline associated with this client.
func (client HTTPServerFailureClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// Delete505 - Return 505 status code - should be represented in the client as an error
func (client HTTPServerFailureClient) Delete505(ctx context.Context, options *HTTPServerFailureDelete505Options) (*http.Response, error) {
	req, err := client.delete505CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode() {
		return nil, client.delete505HandleError(resp)
	}
	return resp.Response, nil
}

// delete505CreateRequest creates the Delete505 request.
func (client HTTPServerFailureClient) delete505CreateRequest(ctx context.Context, options *HTTPServerFailureDelete505Options) (*azcore.Request, error) {
	urlPath := "/http/failure/server/505"
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// delete505HandleError handles the Delete505 error response.
func (client HTTPServerFailureClient) delete505HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get501 - Return 501 status code - should be represented in the client as an error
func (client HTTPServerFailureClient) Get501(ctx context.Context, options *HTTPServerFailureGet501Options) (*http.Response, error) {
	req, err := client.get501CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode() {
		return nil, client.get501HandleError(resp)
	}
	return resp.Response, nil
}

// get501CreateRequest creates the Get501 request.
func (client HTTPServerFailureClient) get501CreateRequest(ctx context.Context, options *HTTPServerFailureGet501Options) (*azcore.Request, error) {
	urlPath := "/http/failure/server/501"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// get501HandleError handles the Get501 error response.
func (client HTTPServerFailureClient) get501HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Head501 - Return 501 status code - should be represented in the client as an error
func (client HTTPServerFailureClient) Head501(ctx context.Context, options *HTTPServerFailureHead501Options) (*http.Response, error) {
	req, err := client.head501CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode() {
		return nil, client.head501HandleError(resp)
	}
	return resp.Response, nil
}

// head501CreateRequest creates the Head501 request.
func (client HTTPServerFailureClient) head501CreateRequest(ctx context.Context, options *HTTPServerFailureHead501Options) (*azcore.Request, error) {
	urlPath := "/http/failure/server/501"
	req, err := azcore.NewRequest(ctx, http.MethodHead, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// head501HandleError handles the Head501 error response.
func (client HTTPServerFailureClient) head501HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Post505 - Return 505 status code - should be represented in the client as an error
func (client HTTPServerFailureClient) Post505(ctx context.Context, options *HTTPServerFailurePost505Options) (*http.Response, error) {
	req, err := client.post505CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode() {
		return nil, client.post505HandleError(resp)
	}
	return resp.Response, nil
}

// post505CreateRequest creates the Post505 request.
func (client HTTPServerFailureClient) post505CreateRequest(ctx context.Context, options *HTTPServerFailurePost505Options) (*azcore.Request, error) {
	urlPath := "/http/failure/server/505"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// post505HandleError handles the Post505 error response.
func (client HTTPServerFailureClient) post505HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
