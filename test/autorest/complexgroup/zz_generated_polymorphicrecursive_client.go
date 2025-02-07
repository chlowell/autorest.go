//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// PolymorphicrecursiveClient contains the methods for the Polymorphicrecursive group.
// Don't use this type directly, use NewPolymorphicrecursiveClient() instead.
type PolymorphicrecursiveClient struct {
	con *Connection
}

// NewPolymorphicrecursiveClient creates a new instance of PolymorphicrecursiveClient with the specified values.
func NewPolymorphicrecursiveClient(con *Connection) *PolymorphicrecursiveClient {
	return &PolymorphicrecursiveClient{con: con}
}

// GetValid - Get complex types that are polymorphic and have recursive references
// If the operation fails it returns the *Error error type.
func (client *PolymorphicrecursiveClient) GetValid(ctx context.Context, options *PolymorphicrecursiveGetValidOptions) (PolymorphicrecursiveGetValidResponse, error) {
	req, err := client.getValidCreateRequest(ctx, options)
	if err != nil {
		return PolymorphicrecursiveGetValidResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return PolymorphicrecursiveGetValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PolymorphicrecursiveGetValidResponse{}, client.getValidHandleError(resp)
	}
	return client.getValidHandleResponse(resp)
}

// getValidCreateRequest creates the GetValid request.
func (client *PolymorphicrecursiveClient) getValidCreateRequest(ctx context.Context, options *PolymorphicrecursiveGetValidOptions) (*policy.Request, error) {
	urlPath := "/complex/polymorphicrecursive/valid"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getValidHandleResponse handles the GetValid response.
func (client *PolymorphicrecursiveClient) getValidHandleResponse(resp *http.Response) (PolymorphicrecursiveGetValidResponse, error) {
	result := PolymorphicrecursiveGetValidResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result); err != nil {
		return PolymorphicrecursiveGetValidResponse{}, err
	}
	return result, nil
}

// getValidHandleError handles the GetValid error response.
func (client *PolymorphicrecursiveClient) getValidHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// PutValid - Put complex types that are polymorphic and have recursive references
// If the operation fails it returns the *Error error type.
func (client *PolymorphicrecursiveClient) PutValid(ctx context.Context, complexBody FishClassification, options *PolymorphicrecursivePutValidOptions) (PolymorphicrecursivePutValidResponse, error) {
	req, err := client.putValidCreateRequest(ctx, complexBody, options)
	if err != nil {
		return PolymorphicrecursivePutValidResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return PolymorphicrecursivePutValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PolymorphicrecursivePutValidResponse{}, client.putValidHandleError(resp)
	}
	return PolymorphicrecursivePutValidResponse{RawResponse: resp}, nil
}

// putValidCreateRequest creates the PutValid request.
func (client *PolymorphicrecursiveClient) putValidCreateRequest(ctx context.Context, complexBody FishClassification, options *PolymorphicrecursivePutValidOptions) (*policy.Request, error) {
	urlPath := "/complex/polymorphicrecursive/valid"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, complexBody)
}

// putValidHandleError handles the PutValid error response.
func (client *PolymorphicrecursiveClient) putValidHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
