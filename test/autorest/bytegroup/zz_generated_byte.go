// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package bytegroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
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
func (client *ByteClient) GetEmpty(ctx context.Context, options *ByteGetEmptyOptions) (ByteArrayResponse, error) {
	req, err := client.getEmptyCreateRequest(ctx, options)
	if err != nil {
		return ByteArrayResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ByteArrayResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ByteArrayResponse{}, client.getEmptyHandleError(resp)
	}
	return client.getEmptyHandleResponse(resp)
}

// getEmptyCreateRequest creates the GetEmpty request.
func (client *ByteClient) getEmptyCreateRequest(ctx context.Context, options *ByteGetEmptyOptions) (*azcore.Request, error) {
	urlPath := "/byte/empty"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getEmptyHandleResponse handles the GetEmpty response.
func (client *ByteClient) getEmptyHandleResponse(resp *azcore.Response) (ByteArrayResponse, error) {
	var val *[]byte
	if err := resp.UnmarshalAsByteArray(&val, azcore.Base64StdFormat); err != nil {
		return ByteArrayResponse{}, err
	}
	return ByteArrayResponse{RawResponse: resp.Response, Value: val}, nil
}

// getEmptyHandleError handles the GetEmpty error response.
func (client *ByteClient) getEmptyHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetInvalid - Get invalid byte value ':::SWAGGER::::'
func (client *ByteClient) GetInvalid(ctx context.Context, options *ByteGetInvalidOptions) (ByteArrayResponse, error) {
	req, err := client.getInvalidCreateRequest(ctx, options)
	if err != nil {
		return ByteArrayResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ByteArrayResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ByteArrayResponse{}, client.getInvalidHandleError(resp)
	}
	return client.getInvalidHandleResponse(resp)
}

// getInvalidCreateRequest creates the GetInvalid request.
func (client *ByteClient) getInvalidCreateRequest(ctx context.Context, options *ByteGetInvalidOptions) (*azcore.Request, error) {
	urlPath := "/byte/invalid"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getInvalidHandleResponse handles the GetInvalid response.
func (client *ByteClient) getInvalidHandleResponse(resp *azcore.Response) (ByteArrayResponse, error) {
	var val *[]byte
	if err := resp.UnmarshalAsByteArray(&val, azcore.Base64StdFormat); err != nil {
		return ByteArrayResponse{}, err
	}
	return ByteArrayResponse{RawResponse: resp.Response, Value: val}, nil
}

// getInvalidHandleError handles the GetInvalid error response.
func (client *ByteClient) getInvalidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetNonASCII - Get non-ascii byte string hex(FF FE FD FC FB FA F9 F8 F7 F6)
func (client *ByteClient) GetNonASCII(ctx context.Context, options *ByteGetNonASCIIOptions) (ByteArrayResponse, error) {
	req, err := client.getNonASCIICreateRequest(ctx, options)
	if err != nil {
		return ByteArrayResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ByteArrayResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ByteArrayResponse{}, client.getNonASCIIHandleError(resp)
	}
	return client.getNonASCIIHandleResponse(resp)
}

// getNonASCIICreateRequest creates the GetNonASCII request.
func (client *ByteClient) getNonASCIICreateRequest(ctx context.Context, options *ByteGetNonASCIIOptions) (*azcore.Request, error) {
	urlPath := "/byte/nonAscii"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getNonASCIIHandleResponse handles the GetNonASCII response.
func (client *ByteClient) getNonASCIIHandleResponse(resp *azcore.Response) (ByteArrayResponse, error) {
	var val *[]byte
	if err := resp.UnmarshalAsByteArray(&val, azcore.Base64StdFormat); err != nil {
		return ByteArrayResponse{}, err
	}
	return ByteArrayResponse{RawResponse: resp.Response, Value: val}, nil
}

// getNonASCIIHandleError handles the GetNonASCII error response.
func (client *ByteClient) getNonASCIIHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetNull - Get null byte value
func (client *ByteClient) GetNull(ctx context.Context, options *ByteGetNullOptions) (ByteArrayResponse, error) {
	req, err := client.getNullCreateRequest(ctx, options)
	if err != nil {
		return ByteArrayResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ByteArrayResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ByteArrayResponse{}, client.getNullHandleError(resp)
	}
	return client.getNullHandleResponse(resp)
}

// getNullCreateRequest creates the GetNull request.
func (client *ByteClient) getNullCreateRequest(ctx context.Context, options *ByteGetNullOptions) (*azcore.Request, error) {
	urlPath := "/byte/null"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getNullHandleResponse handles the GetNull response.
func (client *ByteClient) getNullHandleResponse(resp *azcore.Response) (ByteArrayResponse, error) {
	var val *[]byte
	if err := resp.UnmarshalAsByteArray(&val, azcore.Base64StdFormat); err != nil {
		return ByteArrayResponse{}, err
	}
	return ByteArrayResponse{RawResponse: resp.Response, Value: val}, nil
}

// getNullHandleError handles the GetNull error response.
func (client *ByteClient) getNullHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutNonASCII - Put non-ascii byte string hex(FF FE FD FC FB FA F9 F8 F7 F6)
func (client *ByteClient) PutNonASCII(ctx context.Context, byteBody []byte, options *BytePutNonASCIIOptions) (*http.Response, error) {
	req, err := client.putNonASCIICreateRequest(ctx, byteBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putNonASCIIHandleError(resp)
	}
	return resp.Response, nil
}

// putNonASCIICreateRequest creates the PutNonASCII request.
func (client *ByteClient) putNonASCIICreateRequest(ctx context.Context, byteBody []byte, options *BytePutNonASCIIOptions) (*azcore.Request, error) {
	urlPath := "/byte/nonAscii"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsByteArray(byteBody, azcore.Base64StdFormat)
}

// putNonASCIIHandleError handles the PutNonASCII error response.
func (client *ByteClient) putNonASCIIHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}