//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package dategroup

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"time"
)

// DateClient contains the methods for the Date group.
// Don't use this type directly, use NewDateClient() instead.
type DateClient struct {
	con *Connection
}

// NewDateClient creates a new instance of DateClient with the specified values.
func NewDateClient(con *Connection) *DateClient {
	return &DateClient{con: con}
}

// GetInvalidDate - Get invalid date value
// If the operation fails it returns the *Error error type.
func (client *DateClient) GetInvalidDate(ctx context.Context, options *DateGetInvalidDateOptions) (DateGetInvalidDateResponse, error) {
	req, err := client.getInvalidDateCreateRequest(ctx, options)
	if err != nil {
		return DateGetInvalidDateResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DateGetInvalidDateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DateGetInvalidDateResponse{}, client.getInvalidDateHandleError(resp)
	}
	return client.getInvalidDateHandleResponse(resp)
}

// getInvalidDateCreateRequest creates the GetInvalidDate request.
func (client *DateClient) getInvalidDateCreateRequest(ctx context.Context, options *DateGetInvalidDateOptions) (*policy.Request, error) {
	urlPath := "/date/invaliddate"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getInvalidDateHandleResponse handles the GetInvalidDate response.
func (client *DateClient) getInvalidDateHandleResponse(resp *http.Response) (DateGetInvalidDateResponse, error) {
	result := DateGetInvalidDateResponse{RawResponse: resp}
	var aux *dateType
	if err := runtime.UnmarshalAsJSON(resp, &aux); err != nil {
		return DateGetInvalidDateResponse{}, err
	}
	result.Value = (*time.Time)(aux)
	return result, nil
}

// getInvalidDateHandleError handles the GetInvalidDate error response.
func (client *DateClient) getInvalidDateHandleError(resp *http.Response) error {
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

// GetMaxDate - Get max date value 9999-12-31
// If the operation fails it returns the *Error error type.
func (client *DateClient) GetMaxDate(ctx context.Context, options *DateGetMaxDateOptions) (DateGetMaxDateResponse, error) {
	req, err := client.getMaxDateCreateRequest(ctx, options)
	if err != nil {
		return DateGetMaxDateResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DateGetMaxDateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DateGetMaxDateResponse{}, client.getMaxDateHandleError(resp)
	}
	return client.getMaxDateHandleResponse(resp)
}

// getMaxDateCreateRequest creates the GetMaxDate request.
func (client *DateClient) getMaxDateCreateRequest(ctx context.Context, options *DateGetMaxDateOptions) (*policy.Request, error) {
	urlPath := "/date/max"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getMaxDateHandleResponse handles the GetMaxDate response.
func (client *DateClient) getMaxDateHandleResponse(resp *http.Response) (DateGetMaxDateResponse, error) {
	result := DateGetMaxDateResponse{RawResponse: resp}
	var aux *dateType
	if err := runtime.UnmarshalAsJSON(resp, &aux); err != nil {
		return DateGetMaxDateResponse{}, err
	}
	result.Value = (*time.Time)(aux)
	return result, nil
}

// getMaxDateHandleError handles the GetMaxDate error response.
func (client *DateClient) getMaxDateHandleError(resp *http.Response) error {
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

// GetMinDate - Get min date value 0000-01-01
// If the operation fails it returns the *Error error type.
func (client *DateClient) GetMinDate(ctx context.Context, options *DateGetMinDateOptions) (DateGetMinDateResponse, error) {
	req, err := client.getMinDateCreateRequest(ctx, options)
	if err != nil {
		return DateGetMinDateResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DateGetMinDateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DateGetMinDateResponse{}, client.getMinDateHandleError(resp)
	}
	return client.getMinDateHandleResponse(resp)
}

// getMinDateCreateRequest creates the GetMinDate request.
func (client *DateClient) getMinDateCreateRequest(ctx context.Context, options *DateGetMinDateOptions) (*policy.Request, error) {
	urlPath := "/date/min"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getMinDateHandleResponse handles the GetMinDate response.
func (client *DateClient) getMinDateHandleResponse(resp *http.Response) (DateGetMinDateResponse, error) {
	result := DateGetMinDateResponse{RawResponse: resp}
	var aux *dateType
	if err := runtime.UnmarshalAsJSON(resp, &aux); err != nil {
		return DateGetMinDateResponse{}, err
	}
	result.Value = (*time.Time)(aux)
	return result, nil
}

// getMinDateHandleError handles the GetMinDate error response.
func (client *DateClient) getMinDateHandleError(resp *http.Response) error {
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

// GetNull - Get null date value
// If the operation fails it returns the *Error error type.
func (client *DateClient) GetNull(ctx context.Context, options *DateGetNullOptions) (DateGetNullResponse, error) {
	req, err := client.getNullCreateRequest(ctx, options)
	if err != nil {
		return DateGetNullResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DateGetNullResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DateGetNullResponse{}, client.getNullHandleError(resp)
	}
	return client.getNullHandleResponse(resp)
}

// getNullCreateRequest creates the GetNull request.
func (client *DateClient) getNullCreateRequest(ctx context.Context, options *DateGetNullOptions) (*policy.Request, error) {
	urlPath := "/date/null"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getNullHandleResponse handles the GetNull response.
func (client *DateClient) getNullHandleResponse(resp *http.Response) (DateGetNullResponse, error) {
	result := DateGetNullResponse{RawResponse: resp}
	var aux *dateType
	if err := runtime.UnmarshalAsJSON(resp, &aux); err != nil {
		return DateGetNullResponse{}, err
	}
	result.Value = (*time.Time)(aux)
	return result, nil
}

// getNullHandleError handles the GetNull error response.
func (client *DateClient) getNullHandleError(resp *http.Response) error {
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

// GetOverflowDate - Get overflow date value
// If the operation fails it returns the *Error error type.
func (client *DateClient) GetOverflowDate(ctx context.Context, options *DateGetOverflowDateOptions) (DateGetOverflowDateResponse, error) {
	req, err := client.getOverflowDateCreateRequest(ctx, options)
	if err != nil {
		return DateGetOverflowDateResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DateGetOverflowDateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DateGetOverflowDateResponse{}, client.getOverflowDateHandleError(resp)
	}
	return client.getOverflowDateHandleResponse(resp)
}

// getOverflowDateCreateRequest creates the GetOverflowDate request.
func (client *DateClient) getOverflowDateCreateRequest(ctx context.Context, options *DateGetOverflowDateOptions) (*policy.Request, error) {
	urlPath := "/date/overflowdate"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getOverflowDateHandleResponse handles the GetOverflowDate response.
func (client *DateClient) getOverflowDateHandleResponse(resp *http.Response) (DateGetOverflowDateResponse, error) {
	result := DateGetOverflowDateResponse{RawResponse: resp}
	var aux *dateType
	if err := runtime.UnmarshalAsJSON(resp, &aux); err != nil {
		return DateGetOverflowDateResponse{}, err
	}
	result.Value = (*time.Time)(aux)
	return result, nil
}

// getOverflowDateHandleError handles the GetOverflowDate error response.
func (client *DateClient) getOverflowDateHandleError(resp *http.Response) error {
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

// GetUnderflowDate - Get underflow date value
// If the operation fails it returns the *Error error type.
func (client *DateClient) GetUnderflowDate(ctx context.Context, options *DateGetUnderflowDateOptions) (DateGetUnderflowDateResponse, error) {
	req, err := client.getUnderflowDateCreateRequest(ctx, options)
	if err != nil {
		return DateGetUnderflowDateResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DateGetUnderflowDateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DateGetUnderflowDateResponse{}, client.getUnderflowDateHandleError(resp)
	}
	return client.getUnderflowDateHandleResponse(resp)
}

// getUnderflowDateCreateRequest creates the GetUnderflowDate request.
func (client *DateClient) getUnderflowDateCreateRequest(ctx context.Context, options *DateGetUnderflowDateOptions) (*policy.Request, error) {
	urlPath := "/date/underflowdate"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getUnderflowDateHandleResponse handles the GetUnderflowDate response.
func (client *DateClient) getUnderflowDateHandleResponse(resp *http.Response) (DateGetUnderflowDateResponse, error) {
	result := DateGetUnderflowDateResponse{RawResponse: resp}
	var aux *dateType
	if err := runtime.UnmarshalAsJSON(resp, &aux); err != nil {
		return DateGetUnderflowDateResponse{}, err
	}
	result.Value = (*time.Time)(aux)
	return result, nil
}

// getUnderflowDateHandleError handles the GetUnderflowDate error response.
func (client *DateClient) getUnderflowDateHandleError(resp *http.Response) error {
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

// PutMaxDate - Put max date value 9999-12-31
// If the operation fails it returns the *Error error type.
func (client *DateClient) PutMaxDate(ctx context.Context, dateBody time.Time, options *DatePutMaxDateOptions) (DatePutMaxDateResponse, error) {
	req, err := client.putMaxDateCreateRequest(ctx, dateBody, options)
	if err != nil {
		return DatePutMaxDateResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DatePutMaxDateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DatePutMaxDateResponse{}, client.putMaxDateHandleError(resp)
	}
	return DatePutMaxDateResponse{RawResponse: resp}, nil
}

// putMaxDateCreateRequest creates the PutMaxDate request.
func (client *DateClient) putMaxDateCreateRequest(ctx context.Context, dateBody time.Time, options *DatePutMaxDateOptions) (*policy.Request, error) {
	urlPath := "/date/max"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, dateType(dateBody))
}

// putMaxDateHandleError handles the PutMaxDate error response.
func (client *DateClient) putMaxDateHandleError(resp *http.Response) error {
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

// PutMinDate - Put min date value 0000-01-01
// If the operation fails it returns the *Error error type.
func (client *DateClient) PutMinDate(ctx context.Context, dateBody time.Time, options *DatePutMinDateOptions) (DatePutMinDateResponse, error) {
	req, err := client.putMinDateCreateRequest(ctx, dateBody, options)
	if err != nil {
		return DatePutMinDateResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DatePutMinDateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DatePutMinDateResponse{}, client.putMinDateHandleError(resp)
	}
	return DatePutMinDateResponse{RawResponse: resp}, nil
}

// putMinDateCreateRequest creates the PutMinDate request.
func (client *DateClient) putMinDateCreateRequest(ctx context.Context, dateBody time.Time, options *DatePutMinDateOptions) (*policy.Request, error) {
	urlPath := "/date/min"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, dateType(dateBody))
}

// putMinDateHandleError handles the PutMinDate error response.
func (client *DateClient) putMinDateHandleError(resp *http.Response) error {
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
