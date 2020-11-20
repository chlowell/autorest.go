// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// PolymorphicrecursiveClient contains the methods for the Polymorphicrecursive group.
// Don't use this type directly, use NewPolymorphicrecursiveClient() instead.
type PolymorphicrecursiveClient struct {
	con *Connection
}

// NewPolymorphicrecursiveClient creates a new instance of PolymorphicrecursiveClient with the specified values.
func NewPolymorphicrecursiveClient(con *Connection) PolymorphicrecursiveClient {
	return PolymorphicrecursiveClient{con: con}
}

// Pipeline returns the pipeline associated with this client.
func (client PolymorphicrecursiveClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// GetValid - Get complex types that are polymorphic and have recursive references
func (client PolymorphicrecursiveClient) GetValid(ctx context.Context, options *PolymorphicrecursiveGetValidOptions) (*FishResponse, error) {
	req, err := client.getValidCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getValidHandleError(resp)
	}
	result, err := client.getValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getValidCreateRequest creates the GetValid request.
func (client PolymorphicrecursiveClient) getValidCreateRequest(ctx context.Context, options *PolymorphicrecursiveGetValidOptions) (*azcore.Request, error) {
	urlPath := "/complex/polymorphicrecursive/valid"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getValidHandleResponse handles the GetValid response.
func (client PolymorphicrecursiveClient) getValidHandleResponse(resp *azcore.Response) (*FishResponse, error) {
	result := FishResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result)
}

// getValidHandleError handles the GetValid error response.
func (client PolymorphicrecursiveClient) getValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutValid - Put complex types that are polymorphic and have recursive references
func (client PolymorphicrecursiveClient) PutValid(ctx context.Context, complexBody FishClassification, options *PolymorphicrecursivePutValidOptions) (*http.Response, error) {
	req, err := client.putValidCreateRequest(ctx, complexBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putValidHandleError(resp)
	}
	return resp.Response, nil
}

// putValidCreateRequest creates the PutValid request.
func (client PolymorphicrecursiveClient) putValidCreateRequest(ctx context.Context, complexBody FishClassification, options *PolymorphicrecursivePutValidOptions) (*azcore.Request, error) {
	urlPath := "/complex/polymorphicrecursive/valid"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// putValidHandleError handles the PutValid error response.
func (client PolymorphicrecursiveClient) putValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
