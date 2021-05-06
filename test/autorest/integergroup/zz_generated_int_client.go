// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package integergroup

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"time"
)

// IntClient contains the methods for the Int group.
// Don't use this type directly, use NewIntClient() instead.
type IntClient struct {
	con *Connection
}

// NewIntClient creates a new instance of IntClient with the specified values.
func NewIntClient(con *Connection) *IntClient {
	return &IntClient{con: con}
}

// GetInvalid - Get invalid Int value
// If the operation fails it returns the *Error error type.
func (client *IntClient) GetInvalid(ctx context.Context, options *IntGetInvalidOptions) (Int32Response, error) {
	req, err := client.getInvalidCreateRequest(ctx, options)
	if err != nil {
		return Int32Response{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return Int32Response{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return Int32Response{}, client.getInvalidHandleError(resp)
	}
	return client.getInvalidHandleResponse(resp)
}

// getInvalidCreateRequest creates the GetInvalid request.
func (client *IntClient) getInvalidCreateRequest(ctx context.Context, options *IntGetInvalidOptions) (*azcore.Request, error) {
	urlPath := "/int/invalid"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getInvalidHandleResponse handles the GetInvalid response.
func (client *IntClient) getInvalidHandleResponse(resp *azcore.Response) (Int32Response, error) {
	var val *int32
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return Int32Response{}, err
	}
	return Int32Response{RawResponse: resp.Response, Value: val}, nil
}

// getInvalidHandleError handles the GetInvalid error response.
func (client *IntClient) getInvalidHandleError(resp *azcore.Response) error {
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

// GetInvalidUnixTime - Get invalid Unix time value
// If the operation fails it returns the *Error error type.
func (client *IntClient) GetInvalidUnixTime(ctx context.Context, options *IntGetInvalidUnixTimeOptions) (TimeResponse, error) {
	req, err := client.getInvalidUnixTimeCreateRequest(ctx, options)
	if err != nil {
		return TimeResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TimeResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return TimeResponse{}, client.getInvalidUnixTimeHandleError(resp)
	}
	return client.getInvalidUnixTimeHandleResponse(resp)
}

// getInvalidUnixTimeCreateRequest creates the GetInvalidUnixTime request.
func (client *IntClient) getInvalidUnixTimeCreateRequest(ctx context.Context, options *IntGetInvalidUnixTimeOptions) (*azcore.Request, error) {
	urlPath := "/int/invalidunixtime"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getInvalidUnixTimeHandleResponse handles the GetInvalidUnixTime response.
func (client *IntClient) getInvalidUnixTimeHandleResponse(resp *azcore.Response) (TimeResponse, error) {
	var aux *timeUnix
	if err := resp.UnmarshalAsJSON(&aux); err != nil {
		return TimeResponse{}, err
	}
	return TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, nil
}

// getInvalidUnixTimeHandleError handles the GetInvalidUnixTime error response.
func (client *IntClient) getInvalidUnixTimeHandleError(resp *azcore.Response) error {
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

// GetNull - Get null Int value
// If the operation fails it returns the *Error error type.
func (client *IntClient) GetNull(ctx context.Context, options *IntGetNullOptions) (Int32Response, error) {
	req, err := client.getNullCreateRequest(ctx, options)
	if err != nil {
		return Int32Response{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return Int32Response{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return Int32Response{}, client.getNullHandleError(resp)
	}
	return client.getNullHandleResponse(resp)
}

// getNullCreateRequest creates the GetNull request.
func (client *IntClient) getNullCreateRequest(ctx context.Context, options *IntGetNullOptions) (*azcore.Request, error) {
	urlPath := "/int/null"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getNullHandleResponse handles the GetNull response.
func (client *IntClient) getNullHandleResponse(resp *azcore.Response) (Int32Response, error) {
	var val *int32
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return Int32Response{}, err
	}
	return Int32Response{RawResponse: resp.Response, Value: val}, nil
}

// getNullHandleError handles the GetNull error response.
func (client *IntClient) getNullHandleError(resp *azcore.Response) error {
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

// GetNullUnixTime - Get null Unix time value
// If the operation fails it returns the *Error error type.
func (client *IntClient) GetNullUnixTime(ctx context.Context, options *IntGetNullUnixTimeOptions) (TimeResponse, error) {
	req, err := client.getNullUnixTimeCreateRequest(ctx, options)
	if err != nil {
		return TimeResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TimeResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return TimeResponse{}, client.getNullUnixTimeHandleError(resp)
	}
	return client.getNullUnixTimeHandleResponse(resp)
}

// getNullUnixTimeCreateRequest creates the GetNullUnixTime request.
func (client *IntClient) getNullUnixTimeCreateRequest(ctx context.Context, options *IntGetNullUnixTimeOptions) (*azcore.Request, error) {
	urlPath := "/int/nullunixtime"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getNullUnixTimeHandleResponse handles the GetNullUnixTime response.
func (client *IntClient) getNullUnixTimeHandleResponse(resp *azcore.Response) (TimeResponse, error) {
	var aux *timeUnix
	if err := resp.UnmarshalAsJSON(&aux); err != nil {
		return TimeResponse{}, err
	}
	return TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, nil
}

// getNullUnixTimeHandleError handles the GetNullUnixTime error response.
func (client *IntClient) getNullUnixTimeHandleError(resp *azcore.Response) error {
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

// GetOverflowInt32 - Get overflow Int32 value
// If the operation fails it returns the *Error error type.
func (client *IntClient) GetOverflowInt32(ctx context.Context, options *IntGetOverflowInt32Options) (Int32Response, error) {
	req, err := client.getOverflowInt32CreateRequest(ctx, options)
	if err != nil {
		return Int32Response{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return Int32Response{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return Int32Response{}, client.getOverflowInt32HandleError(resp)
	}
	return client.getOverflowInt32HandleResponse(resp)
}

// getOverflowInt32CreateRequest creates the GetOverflowInt32 request.
func (client *IntClient) getOverflowInt32CreateRequest(ctx context.Context, options *IntGetOverflowInt32Options) (*azcore.Request, error) {
	urlPath := "/int/overflowint32"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getOverflowInt32HandleResponse handles the GetOverflowInt32 response.
func (client *IntClient) getOverflowInt32HandleResponse(resp *azcore.Response) (Int32Response, error) {
	var val *int32
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return Int32Response{}, err
	}
	return Int32Response{RawResponse: resp.Response, Value: val}, nil
}

// getOverflowInt32HandleError handles the GetOverflowInt32 error response.
func (client *IntClient) getOverflowInt32HandleError(resp *azcore.Response) error {
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

// GetOverflowInt64 - Get overflow Int64 value
// If the operation fails it returns the *Error error type.
func (client *IntClient) GetOverflowInt64(ctx context.Context, options *IntGetOverflowInt64Options) (Int64Response, error) {
	req, err := client.getOverflowInt64CreateRequest(ctx, options)
	if err != nil {
		return Int64Response{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return Int64Response{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return Int64Response{}, client.getOverflowInt64HandleError(resp)
	}
	return client.getOverflowInt64HandleResponse(resp)
}

// getOverflowInt64CreateRequest creates the GetOverflowInt64 request.
func (client *IntClient) getOverflowInt64CreateRequest(ctx context.Context, options *IntGetOverflowInt64Options) (*azcore.Request, error) {
	urlPath := "/int/overflowint64"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getOverflowInt64HandleResponse handles the GetOverflowInt64 response.
func (client *IntClient) getOverflowInt64HandleResponse(resp *azcore.Response) (Int64Response, error) {
	var val *int64
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return Int64Response{}, err
	}
	return Int64Response{RawResponse: resp.Response, Value: val}, nil
}

// getOverflowInt64HandleError handles the GetOverflowInt64 error response.
func (client *IntClient) getOverflowInt64HandleError(resp *azcore.Response) error {
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

// GetUnderflowInt32 - Get underflow Int32 value
// If the operation fails it returns the *Error error type.
func (client *IntClient) GetUnderflowInt32(ctx context.Context, options *IntGetUnderflowInt32Options) (Int32Response, error) {
	req, err := client.getUnderflowInt32CreateRequest(ctx, options)
	if err != nil {
		return Int32Response{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return Int32Response{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return Int32Response{}, client.getUnderflowInt32HandleError(resp)
	}
	return client.getUnderflowInt32HandleResponse(resp)
}

// getUnderflowInt32CreateRequest creates the GetUnderflowInt32 request.
func (client *IntClient) getUnderflowInt32CreateRequest(ctx context.Context, options *IntGetUnderflowInt32Options) (*azcore.Request, error) {
	urlPath := "/int/underflowint32"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getUnderflowInt32HandleResponse handles the GetUnderflowInt32 response.
func (client *IntClient) getUnderflowInt32HandleResponse(resp *azcore.Response) (Int32Response, error) {
	var val *int32
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return Int32Response{}, err
	}
	return Int32Response{RawResponse: resp.Response, Value: val}, nil
}

// getUnderflowInt32HandleError handles the GetUnderflowInt32 error response.
func (client *IntClient) getUnderflowInt32HandleError(resp *azcore.Response) error {
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

// GetUnderflowInt64 - Get underflow Int64 value
// If the operation fails it returns the *Error error type.
func (client *IntClient) GetUnderflowInt64(ctx context.Context, options *IntGetUnderflowInt64Options) (Int64Response, error) {
	req, err := client.getUnderflowInt64CreateRequest(ctx, options)
	if err != nil {
		return Int64Response{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return Int64Response{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return Int64Response{}, client.getUnderflowInt64HandleError(resp)
	}
	return client.getUnderflowInt64HandleResponse(resp)
}

// getUnderflowInt64CreateRequest creates the GetUnderflowInt64 request.
func (client *IntClient) getUnderflowInt64CreateRequest(ctx context.Context, options *IntGetUnderflowInt64Options) (*azcore.Request, error) {
	urlPath := "/int/underflowint64"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getUnderflowInt64HandleResponse handles the GetUnderflowInt64 response.
func (client *IntClient) getUnderflowInt64HandleResponse(resp *azcore.Response) (Int64Response, error) {
	var val *int64
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return Int64Response{}, err
	}
	return Int64Response{RawResponse: resp.Response, Value: val}, nil
}

// getUnderflowInt64HandleError handles the GetUnderflowInt64 error response.
func (client *IntClient) getUnderflowInt64HandleError(resp *azcore.Response) error {
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

// GetUnixTime - Get datetime encoded as Unix time value
// If the operation fails it returns the *Error error type.
func (client *IntClient) GetUnixTime(ctx context.Context, options *IntGetUnixTimeOptions) (TimeResponse, error) {
	req, err := client.getUnixTimeCreateRequest(ctx, options)
	if err != nil {
		return TimeResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TimeResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return TimeResponse{}, client.getUnixTimeHandleError(resp)
	}
	return client.getUnixTimeHandleResponse(resp)
}

// getUnixTimeCreateRequest creates the GetUnixTime request.
func (client *IntClient) getUnixTimeCreateRequest(ctx context.Context, options *IntGetUnixTimeOptions) (*azcore.Request, error) {
	urlPath := "/int/unixtime"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getUnixTimeHandleResponse handles the GetUnixTime response.
func (client *IntClient) getUnixTimeHandleResponse(resp *azcore.Response) (TimeResponse, error) {
	var aux *timeUnix
	if err := resp.UnmarshalAsJSON(&aux); err != nil {
		return TimeResponse{}, err
	}
	return TimeResponse{RawResponse: resp.Response, Value: (*time.Time)(aux)}, nil
}

// getUnixTimeHandleError handles the GetUnixTime error response.
func (client *IntClient) getUnixTimeHandleError(resp *azcore.Response) error {
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

// PutMax32 - Put max int32 value
// If the operation fails it returns the *Error error type.
func (client *IntClient) PutMax32(ctx context.Context, intBody int32, options *IntPutMax32Options) (*http.Response, error) {
	req, err := client.putMax32CreateRequest(ctx, intBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putMax32HandleError(resp)
	}
	return resp.Response, nil
}

// putMax32CreateRequest creates the PutMax32 request.
func (client *IntClient) putMax32CreateRequest(ctx context.Context, intBody int32, options *IntPutMax32Options) (*azcore.Request, error) {
	urlPath := "/int/max/32"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(intBody)
}

// putMax32HandleError handles the PutMax32 error response.
func (client *IntClient) putMax32HandleError(resp *azcore.Response) error {
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

// PutMax64 - Put max int64 value
// If the operation fails it returns the *Error error type.
func (client *IntClient) PutMax64(ctx context.Context, intBody int64, options *IntPutMax64Options) (*http.Response, error) {
	req, err := client.putMax64CreateRequest(ctx, intBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putMax64HandleError(resp)
	}
	return resp.Response, nil
}

// putMax64CreateRequest creates the PutMax64 request.
func (client *IntClient) putMax64CreateRequest(ctx context.Context, intBody int64, options *IntPutMax64Options) (*azcore.Request, error) {
	urlPath := "/int/max/64"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(intBody)
}

// putMax64HandleError handles the PutMax64 error response.
func (client *IntClient) putMax64HandleError(resp *azcore.Response) error {
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

// PutMin32 - Put min int32 value
// If the operation fails it returns the *Error error type.
func (client *IntClient) PutMin32(ctx context.Context, intBody int32, options *IntPutMin32Options) (*http.Response, error) {
	req, err := client.putMin32CreateRequest(ctx, intBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putMin32HandleError(resp)
	}
	return resp.Response, nil
}

// putMin32CreateRequest creates the PutMin32 request.
func (client *IntClient) putMin32CreateRequest(ctx context.Context, intBody int32, options *IntPutMin32Options) (*azcore.Request, error) {
	urlPath := "/int/min/32"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(intBody)
}

// putMin32HandleError handles the PutMin32 error response.
func (client *IntClient) putMin32HandleError(resp *azcore.Response) error {
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

// PutMin64 - Put min int64 value
// If the operation fails it returns the *Error error type.
func (client *IntClient) PutMin64(ctx context.Context, intBody int64, options *IntPutMin64Options) (*http.Response, error) {
	req, err := client.putMin64CreateRequest(ctx, intBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putMin64HandleError(resp)
	}
	return resp.Response, nil
}

// putMin64CreateRequest creates the PutMin64 request.
func (client *IntClient) putMin64CreateRequest(ctx context.Context, intBody int64, options *IntPutMin64Options) (*azcore.Request, error) {
	urlPath := "/int/min/64"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(intBody)
}

// putMin64HandleError handles the PutMin64 error response.
func (client *IntClient) putMin64HandleError(resp *azcore.Response) error {
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

// PutUnixTimeDate - Put datetime encoded as Unix time
// If the operation fails it returns the *Error error type.
func (client *IntClient) PutUnixTimeDate(ctx context.Context, intBody time.Time, options *IntPutUnixTimeDateOptions) (*http.Response, error) {
	req, err := client.putUnixTimeDateCreateRequest(ctx, intBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putUnixTimeDateHandleError(resp)
	}
	return resp.Response, nil
}

// putUnixTimeDateCreateRequest creates the PutUnixTimeDate request.
func (client *IntClient) putUnixTimeDateCreateRequest(ctx context.Context, intBody time.Time, options *IntPutUnixTimeDateOptions) (*azcore.Request, error) {
	urlPath := "/int/unixtime"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	aux := timeUnix(intBody)
	return req, req.MarshalAsJSON(aux)
}

// putUnixTimeDateHandleError handles the PutUnixTimeDate error response.
func (client *IntClient) putUnixTimeDateHandleError(resp *azcore.Response) error {
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
