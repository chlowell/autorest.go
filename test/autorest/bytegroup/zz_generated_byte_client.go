//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package bytegroup

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// ByteClient contains the methods for the Byte group.
// Don't use this type directly, use NewByteClient() instead.
type ByteClient struct {
	con *Connection
}

// NewByteClient creates a new instance of ByteClient with the specified values.
func NewByteClient(con *Connection) *ByteClient {
	return &ByteClient{con: con}
}

// GetEmpty - Get empty byte value ''
// If the operation fails it returns the *Error error type.
func (client *ByteClient) GetEmpty(ctx context.Context, options *ByteGetEmptyOptions) (ByteGetEmptyResponse, error) {
	req, err := client.getEmptyCreateRequest(ctx, options)
	if err != nil {
		return ByteGetEmptyResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ByteGetEmptyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ByteGetEmptyResponse{}, client.getEmptyHandleError(resp)
	}
	return client.getEmptyHandleResponse(resp)
}

// getEmptyCreateRequest creates the GetEmpty request.
func (client *ByteClient) getEmptyCreateRequest(ctx context.Context, options *ByteGetEmptyOptions) (*policy.Request, error) {
	urlPath := "/byte/empty"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getEmptyHandleResponse handles the GetEmpty response.
func (client *ByteClient) getEmptyHandleResponse(resp *http.Response) (ByteGetEmptyResponse, error) {
	result := ByteGetEmptyResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsByteArray(resp, &result.Value, runtime.Base64StdFormat); err != nil {
		return ByteGetEmptyResponse{}, err
	}
	return result, nil
}

// getEmptyHandleError handles the GetEmpty error response.
func (client *ByteClient) getEmptyHandleError(resp *http.Response) error {
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

// GetInvalid - Get invalid byte value ':::SWAGGER::::'
// If the operation fails it returns the *Error error type.
func (client *ByteClient) GetInvalid(ctx context.Context, options *ByteGetInvalidOptions) (ByteGetInvalidResponse, error) {
	req, err := client.getInvalidCreateRequest(ctx, options)
	if err != nil {
		return ByteGetInvalidResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ByteGetInvalidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ByteGetInvalidResponse{}, client.getInvalidHandleError(resp)
	}
	return client.getInvalidHandleResponse(resp)
}

// getInvalidCreateRequest creates the GetInvalid request.
func (client *ByteClient) getInvalidCreateRequest(ctx context.Context, options *ByteGetInvalidOptions) (*policy.Request, error) {
	urlPath := "/byte/invalid"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getInvalidHandleResponse handles the GetInvalid response.
func (client *ByteClient) getInvalidHandleResponse(resp *http.Response) (ByteGetInvalidResponse, error) {
	result := ByteGetInvalidResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsByteArray(resp, &result.Value, runtime.Base64StdFormat); err != nil {
		return ByteGetInvalidResponse{}, err
	}
	return result, nil
}

// getInvalidHandleError handles the GetInvalid error response.
func (client *ByteClient) getInvalidHandleError(resp *http.Response) error {
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

// GetNonASCII - Get non-ascii byte string hex(FF FE FD FC FB FA F9 F8 F7 F6)
// If the operation fails it returns the *Error error type.
func (client *ByteClient) GetNonASCII(ctx context.Context, options *ByteGetNonASCIIOptions) (ByteGetNonASCIIResponse, error) {
	req, err := client.getNonASCIICreateRequest(ctx, options)
	if err != nil {
		return ByteGetNonASCIIResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ByteGetNonASCIIResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ByteGetNonASCIIResponse{}, client.getNonASCIIHandleError(resp)
	}
	return client.getNonASCIIHandleResponse(resp)
}

// getNonASCIICreateRequest creates the GetNonASCII request.
func (client *ByteClient) getNonASCIICreateRequest(ctx context.Context, options *ByteGetNonASCIIOptions) (*policy.Request, error) {
	urlPath := "/byte/nonAscii"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getNonASCIIHandleResponse handles the GetNonASCII response.
func (client *ByteClient) getNonASCIIHandleResponse(resp *http.Response) (ByteGetNonASCIIResponse, error) {
	result := ByteGetNonASCIIResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsByteArray(resp, &result.Value, runtime.Base64StdFormat); err != nil {
		return ByteGetNonASCIIResponse{}, err
	}
	return result, nil
}

// getNonASCIIHandleError handles the GetNonASCII error response.
func (client *ByteClient) getNonASCIIHandleError(resp *http.Response) error {
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

// GetNull - Get null byte value
// If the operation fails it returns the *Error error type.
func (client *ByteClient) GetNull(ctx context.Context, options *ByteGetNullOptions) (ByteGetNullResponse, error) {
	req, err := client.getNullCreateRequest(ctx, options)
	if err != nil {
		return ByteGetNullResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ByteGetNullResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ByteGetNullResponse{}, client.getNullHandleError(resp)
	}
	return client.getNullHandleResponse(resp)
}

// getNullCreateRequest creates the GetNull request.
func (client *ByteClient) getNullCreateRequest(ctx context.Context, options *ByteGetNullOptions) (*policy.Request, error) {
	urlPath := "/byte/null"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getNullHandleResponse handles the GetNull response.
func (client *ByteClient) getNullHandleResponse(resp *http.Response) (ByteGetNullResponse, error) {
	result := ByteGetNullResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsByteArray(resp, &result.Value, runtime.Base64StdFormat); err != nil {
		return ByteGetNullResponse{}, err
	}
	return result, nil
}

// getNullHandleError handles the GetNull error response.
func (client *ByteClient) getNullHandleError(resp *http.Response) error {
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

// PutNonASCII - Put non-ascii byte string hex(FF FE FD FC FB FA F9 F8 F7 F6)
// If the operation fails it returns the *Error error type.
func (client *ByteClient) PutNonASCII(ctx context.Context, byteBody []byte, options *BytePutNonASCIIOptions) (BytePutNonASCIIResponse, error) {
	req, err := client.putNonASCIICreateRequest(ctx, byteBody, options)
	if err != nil {
		return BytePutNonASCIIResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return BytePutNonASCIIResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BytePutNonASCIIResponse{}, client.putNonASCIIHandleError(resp)
	}
	return BytePutNonASCIIResponse{RawResponse: resp}, nil
}

// putNonASCIICreateRequest creates the PutNonASCII request.
func (client *ByteClient) putNonASCIICreateRequest(ctx context.Context, byteBody []byte, options *BytePutNonASCIIOptions) (*policy.Request, error) {
	urlPath := "/byte/nonAscii"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsByteArray(req, byteBody, runtime.Base64StdFormat)
}

// putNonASCIIHandleError handles the PutNonASCII error response.
func (client *ByteClient) putNonASCIIHandleError(resp *http.Response) error {
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
