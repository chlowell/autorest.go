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

// BasicClient contains the methods for the Basic group.
// Don't use this type directly, use NewBasicClient() instead.
type BasicClient struct {
	con *Connection
}

// NewBasicClient creates a new instance of BasicClient with the specified values.
func NewBasicClient(con *Connection) *BasicClient {
	return &BasicClient{con: con}
}

// GetEmpty - Get a basic complex type that is empty
// If the operation fails it returns the *Error error type.
func (client *BasicClient) GetEmpty(ctx context.Context, options *BasicGetEmptyOptions) (BasicGetEmptyResponse, error) {
	req, err := client.getEmptyCreateRequest(ctx, options)
	if err != nil {
		return BasicGetEmptyResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return BasicGetEmptyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BasicGetEmptyResponse{}, client.getEmptyHandleError(resp)
	}
	return client.getEmptyHandleResponse(resp)
}

// getEmptyCreateRequest creates the GetEmpty request.
func (client *BasicClient) getEmptyCreateRequest(ctx context.Context, options *BasicGetEmptyOptions) (*policy.Request, error) {
	urlPath := "/complex/basic/empty"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getEmptyHandleResponse handles the GetEmpty response.
func (client *BasicClient) getEmptyHandleResponse(resp *http.Response) (BasicGetEmptyResponse, error) {
	result := BasicGetEmptyResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Basic); err != nil {
		return BasicGetEmptyResponse{}, err
	}
	return result, nil
}

// getEmptyHandleError handles the GetEmpty error response.
func (client *BasicClient) getEmptyHandleError(resp *http.Response) error {
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

// GetInvalid - Get a basic complex type that is invalid for the local strong type
// If the operation fails it returns the *Error error type.
func (client *BasicClient) GetInvalid(ctx context.Context, options *BasicGetInvalidOptions) (BasicGetInvalidResponse, error) {
	req, err := client.getInvalidCreateRequest(ctx, options)
	if err != nil {
		return BasicGetInvalidResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return BasicGetInvalidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BasicGetInvalidResponse{}, client.getInvalidHandleError(resp)
	}
	return client.getInvalidHandleResponse(resp)
}

// getInvalidCreateRequest creates the GetInvalid request.
func (client *BasicClient) getInvalidCreateRequest(ctx context.Context, options *BasicGetInvalidOptions) (*policy.Request, error) {
	urlPath := "/complex/basic/invalid"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getInvalidHandleResponse handles the GetInvalid response.
func (client *BasicClient) getInvalidHandleResponse(resp *http.Response) (BasicGetInvalidResponse, error) {
	result := BasicGetInvalidResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Basic); err != nil {
		return BasicGetInvalidResponse{}, err
	}
	return result, nil
}

// getInvalidHandleError handles the GetInvalid error response.
func (client *BasicClient) getInvalidHandleError(resp *http.Response) error {
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

// GetNotProvided - Get a basic complex type while the server doesn't provide a response payload
// If the operation fails it returns the *Error error type.
func (client *BasicClient) GetNotProvided(ctx context.Context, options *BasicGetNotProvidedOptions) (BasicGetNotProvidedResponse, error) {
	req, err := client.getNotProvidedCreateRequest(ctx, options)
	if err != nil {
		return BasicGetNotProvidedResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return BasicGetNotProvidedResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BasicGetNotProvidedResponse{}, client.getNotProvidedHandleError(resp)
	}
	return client.getNotProvidedHandleResponse(resp)
}

// getNotProvidedCreateRequest creates the GetNotProvided request.
func (client *BasicClient) getNotProvidedCreateRequest(ctx context.Context, options *BasicGetNotProvidedOptions) (*policy.Request, error) {
	urlPath := "/complex/basic/notprovided"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getNotProvidedHandleResponse handles the GetNotProvided response.
func (client *BasicClient) getNotProvidedHandleResponse(resp *http.Response) (BasicGetNotProvidedResponse, error) {
	result := BasicGetNotProvidedResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Basic); err != nil {
		return BasicGetNotProvidedResponse{}, err
	}
	return result, nil
}

// getNotProvidedHandleError handles the GetNotProvided error response.
func (client *BasicClient) getNotProvidedHandleError(resp *http.Response) error {
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

// GetNull - Get a basic complex type whose properties are null
// If the operation fails it returns the *Error error type.
func (client *BasicClient) GetNull(ctx context.Context, options *BasicGetNullOptions) (BasicGetNullResponse, error) {
	req, err := client.getNullCreateRequest(ctx, options)
	if err != nil {
		return BasicGetNullResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return BasicGetNullResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BasicGetNullResponse{}, client.getNullHandleError(resp)
	}
	return client.getNullHandleResponse(resp)
}

// getNullCreateRequest creates the GetNull request.
func (client *BasicClient) getNullCreateRequest(ctx context.Context, options *BasicGetNullOptions) (*policy.Request, error) {
	urlPath := "/complex/basic/null"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getNullHandleResponse handles the GetNull response.
func (client *BasicClient) getNullHandleResponse(resp *http.Response) (BasicGetNullResponse, error) {
	result := BasicGetNullResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Basic); err != nil {
		return BasicGetNullResponse{}, err
	}
	return result, nil
}

// getNullHandleError handles the GetNull error response.
func (client *BasicClient) getNullHandleError(resp *http.Response) error {
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

// GetValid - Get complex type {id: 2, name: 'abc', color: 'YELLOW'}
// If the operation fails it returns the *Error error type.
func (client *BasicClient) GetValid(ctx context.Context, options *BasicGetValidOptions) (BasicGetValidResponse, error) {
	req, err := client.getValidCreateRequest(ctx, options)
	if err != nil {
		return BasicGetValidResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return BasicGetValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BasicGetValidResponse{}, client.getValidHandleError(resp)
	}
	return client.getValidHandleResponse(resp)
}

// getValidCreateRequest creates the GetValid request.
func (client *BasicClient) getValidCreateRequest(ctx context.Context, options *BasicGetValidOptions) (*policy.Request, error) {
	urlPath := "/complex/basic/valid"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getValidHandleResponse handles the GetValid response.
func (client *BasicClient) getValidHandleResponse(resp *http.Response) (BasicGetValidResponse, error) {
	result := BasicGetValidResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Basic); err != nil {
		return BasicGetValidResponse{}, err
	}
	return result, nil
}

// getValidHandleError handles the GetValid error response.
func (client *BasicClient) getValidHandleError(resp *http.Response) error {
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

// PutValid - Please put {id: 2, name: 'abc', color: 'Magenta'}
// If the operation fails it returns the *Error error type.
func (client *BasicClient) PutValid(ctx context.Context, complexBody Basic, options *BasicPutValidOptions) (BasicPutValidResponse, error) {
	req, err := client.putValidCreateRequest(ctx, complexBody, options)
	if err != nil {
		return BasicPutValidResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return BasicPutValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BasicPutValidResponse{}, client.putValidHandleError(resp)
	}
	return BasicPutValidResponse{RawResponse: resp}, nil
}

// putValidCreateRequest creates the PutValid request.
func (client *BasicClient) putValidCreateRequest(ctx context.Context, complexBody Basic, options *BasicPutValidOptions) (*policy.Request, error) {
	urlPath := "/complex/basic/valid"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2016-02-29")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, complexBody)
}

// putValidHandleError handles the PutValid error response.
func (client *BasicClient) putValidHandleError(resp *http.Response) error {
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
