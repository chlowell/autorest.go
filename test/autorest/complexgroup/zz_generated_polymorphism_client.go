// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// PolymorphismClient contains the methods for the Polymorphism group.
// Don't use this type directly, use NewPolymorphismClient() instead.
type PolymorphismClient struct {
	con *Connection
}

// NewPolymorphismClient creates a new instance of PolymorphismClient with the specified values.
func NewPolymorphismClient(con *Connection) *PolymorphismClient {
	return &PolymorphismClient{con: con}
}

// GetComplicated - Get complex types that are polymorphic, but not at the root of the hierarchy; also have additional properties
// If the operation fails it returns the *Error error type.
func (client *PolymorphismClient) GetComplicated(ctx context.Context, options *PolymorphismGetComplicatedOptions) (SalmonResponse, error) {
	req, err := client.getComplicatedCreateRequest(ctx, options)
	if err != nil {
		return SalmonResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SalmonResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SalmonResponse{}, client.getComplicatedHandleError(resp)
	}
	return client.getComplicatedHandleResponse(resp)
}

// getComplicatedCreateRequest creates the GetComplicated request.
func (client *PolymorphismClient) getComplicatedCreateRequest(ctx context.Context, options *PolymorphismGetComplicatedOptions) (*azcore.Request, error) {
	urlPath := "/complex/polymorphism/complicated"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getComplicatedHandleResponse handles the GetComplicated response.
func (client *PolymorphismClient) getComplicatedHandleResponse(resp *azcore.Response) (SalmonResponse, error) {
	result := SalmonResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result); err != nil {
		return SalmonResponse{}, err
	}
	return result, nil
}

// getComplicatedHandleError handles the GetComplicated error response.
func (client *PolymorphismClient) getComplicatedHandleError(resp *azcore.Response) error {
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

// GetComposedWithDiscriminator - Get complex object composing a polymorphic scalar property and array property with polymorphic element type, with discriminator
// specified. Deserialization must NOT fail and use the discriminator type
// specified on the wire.
// If the operation fails it returns the *Error error type.
func (client *PolymorphismClient) GetComposedWithDiscriminator(ctx context.Context, options *PolymorphismGetComposedWithDiscriminatorOptions) (DotFishMarketResponse, error) {
	req, err := client.getComposedWithDiscriminatorCreateRequest(ctx, options)
	if err != nil {
		return DotFishMarketResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DotFishMarketResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DotFishMarketResponse{}, client.getComposedWithDiscriminatorHandleError(resp)
	}
	return client.getComposedWithDiscriminatorHandleResponse(resp)
}

// getComposedWithDiscriminatorCreateRequest creates the GetComposedWithDiscriminator request.
func (client *PolymorphismClient) getComposedWithDiscriminatorCreateRequest(ctx context.Context, options *PolymorphismGetComposedWithDiscriminatorOptions) (*azcore.Request, error) {
	urlPath := "/complex/polymorphism/composedWithDiscriminator"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getComposedWithDiscriminatorHandleResponse handles the GetComposedWithDiscriminator response.
func (client *PolymorphismClient) getComposedWithDiscriminatorHandleResponse(resp *azcore.Response) (DotFishMarketResponse, error) {
	var val *DotFishMarket
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DotFishMarketResponse{}, err
	}
	return DotFishMarketResponse{RawResponse: resp.Response, DotFishMarket: val}, nil
}

// getComposedWithDiscriminatorHandleError handles the GetComposedWithDiscriminator error response.
func (client *PolymorphismClient) getComposedWithDiscriminatorHandleError(resp *azcore.Response) error {
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

// GetComposedWithoutDiscriminator - Get complex object composing a polymorphic scalar property and array property with polymorphic element type, without
// discriminator specified on wire. Deserialization must NOT fail and use the explicit
// type of the property.
// If the operation fails it returns the *Error error type.
func (client *PolymorphismClient) GetComposedWithoutDiscriminator(ctx context.Context, options *PolymorphismGetComposedWithoutDiscriminatorOptions) (DotFishMarketResponse, error) {
	req, err := client.getComposedWithoutDiscriminatorCreateRequest(ctx, options)
	if err != nil {
		return DotFishMarketResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DotFishMarketResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DotFishMarketResponse{}, client.getComposedWithoutDiscriminatorHandleError(resp)
	}
	return client.getComposedWithoutDiscriminatorHandleResponse(resp)
}

// getComposedWithoutDiscriminatorCreateRequest creates the GetComposedWithoutDiscriminator request.
func (client *PolymorphismClient) getComposedWithoutDiscriminatorCreateRequest(ctx context.Context, options *PolymorphismGetComposedWithoutDiscriminatorOptions) (*azcore.Request, error) {
	urlPath := "/complex/polymorphism/composedWithoutDiscriminator"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getComposedWithoutDiscriminatorHandleResponse handles the GetComposedWithoutDiscriminator response.
func (client *PolymorphismClient) getComposedWithoutDiscriminatorHandleResponse(resp *azcore.Response) (DotFishMarketResponse, error) {
	var val *DotFishMarket
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DotFishMarketResponse{}, err
	}
	return DotFishMarketResponse{RawResponse: resp.Response, DotFishMarket: val}, nil
}

// getComposedWithoutDiscriminatorHandleError handles the GetComposedWithoutDiscriminator error response.
func (client *PolymorphismClient) getComposedWithoutDiscriminatorHandleError(resp *azcore.Response) error {
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

// GetDotSyntax - Get complex types that are polymorphic, JSON key contains a dot
// If the operation fails it returns the *Error error type.
func (client *PolymorphismClient) GetDotSyntax(ctx context.Context, options *PolymorphismGetDotSyntaxOptions) (DotFishResponse, error) {
	req, err := client.getDotSyntaxCreateRequest(ctx, options)
	if err != nil {
		return DotFishResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DotFishResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DotFishResponse{}, client.getDotSyntaxHandleError(resp)
	}
	return client.getDotSyntaxHandleResponse(resp)
}

// getDotSyntaxCreateRequest creates the GetDotSyntax request.
func (client *PolymorphismClient) getDotSyntaxCreateRequest(ctx context.Context, options *PolymorphismGetDotSyntaxOptions) (*azcore.Request, error) {
	urlPath := "/complex/polymorphism/dotsyntax"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getDotSyntaxHandleResponse handles the GetDotSyntax response.
func (client *PolymorphismClient) getDotSyntaxHandleResponse(resp *azcore.Response) (DotFishResponse, error) {
	result := DotFishResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result); err != nil {
		return DotFishResponse{}, err
	}
	return result, nil
}

// getDotSyntaxHandleError handles the GetDotSyntax error response.
func (client *PolymorphismClient) getDotSyntaxHandleError(resp *azcore.Response) error {
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

// GetValid - Get complex types that are polymorphic
// If the operation fails it returns the *Error error type.
func (client *PolymorphismClient) GetValid(ctx context.Context, options *PolymorphismGetValidOptions) (FishResponse, error) {
	req, err := client.getValidCreateRequest(ctx, options)
	if err != nil {
		return FishResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return FishResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return FishResponse{}, client.getValidHandleError(resp)
	}
	return client.getValidHandleResponse(resp)
}

// getValidCreateRequest creates the GetValid request.
func (client *PolymorphismClient) getValidCreateRequest(ctx context.Context, options *PolymorphismGetValidOptions) (*azcore.Request, error) {
	urlPath := "/complex/polymorphism/valid"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getValidHandleResponse handles the GetValid response.
func (client *PolymorphismClient) getValidHandleResponse(resp *azcore.Response) (FishResponse, error) {
	result := FishResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result); err != nil {
		return FishResponse{}, err
	}
	return result, nil
}

// getValidHandleError handles the GetValid error response.
func (client *PolymorphismClient) getValidHandleError(resp *azcore.Response) error {
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

// PutComplicated - Put complex types that are polymorphic, but not at the root of the hierarchy; also have additional properties
// If the operation fails it returns the *Error error type.
func (client *PolymorphismClient) PutComplicated(ctx context.Context, complexBody SalmonClassification, options *PolymorphismPutComplicatedOptions) (*http.Response, error) {
	req, err := client.putComplicatedCreateRequest(ctx, complexBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putComplicatedHandleError(resp)
	}
	return resp.Response, nil
}

// putComplicatedCreateRequest creates the PutComplicated request.
func (client *PolymorphismClient) putComplicatedCreateRequest(ctx context.Context, complexBody SalmonClassification, options *PolymorphismPutComplicatedOptions) (*azcore.Request, error) {
	urlPath := "/complex/polymorphism/complicated"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// putComplicatedHandleError handles the PutComplicated error response.
func (client *PolymorphismClient) putComplicatedHandleError(resp *azcore.Response) error {
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

// PutMissingDiscriminator - Put complex types that are polymorphic, omitting the discriminator
// If the operation fails it returns the *Error error type.
func (client *PolymorphismClient) PutMissingDiscriminator(ctx context.Context, complexBody SalmonClassification, options *PolymorphismPutMissingDiscriminatorOptions) (SalmonResponse, error) {
	req, err := client.putMissingDiscriminatorCreateRequest(ctx, complexBody, options)
	if err != nil {
		return SalmonResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SalmonResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SalmonResponse{}, client.putMissingDiscriminatorHandleError(resp)
	}
	return client.putMissingDiscriminatorHandleResponse(resp)
}

// putMissingDiscriminatorCreateRequest creates the PutMissingDiscriminator request.
func (client *PolymorphismClient) putMissingDiscriminatorCreateRequest(ctx context.Context, complexBody SalmonClassification, options *PolymorphismPutMissingDiscriminatorOptions) (*azcore.Request, error) {
	urlPath := "/complex/polymorphism/missingdiscriminator"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// putMissingDiscriminatorHandleResponse handles the PutMissingDiscriminator response.
func (client *PolymorphismClient) putMissingDiscriminatorHandleResponse(resp *azcore.Response) (SalmonResponse, error) {
	result := SalmonResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result); err != nil {
		return SalmonResponse{}, err
	}
	return result, nil
}

// putMissingDiscriminatorHandleError handles the PutMissingDiscriminator error response.
func (client *PolymorphismClient) putMissingDiscriminatorHandleError(resp *azcore.Response) error {
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

// PutValid - Put complex types that are polymorphic
// If the operation fails it returns the *Error error type.
func (client *PolymorphismClient) PutValid(ctx context.Context, complexBody FishClassification, options *PolymorphismPutValidOptions) (*http.Response, error) {
	req, err := client.putValidCreateRequest(ctx, complexBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putValidHandleError(resp)
	}
	return resp.Response, nil
}

// putValidCreateRequest creates the PutValid request.
func (client *PolymorphismClient) putValidCreateRequest(ctx context.Context, complexBody FishClassification, options *PolymorphismPutValidOptions) (*azcore.Request, error) {
	urlPath := "/complex/polymorphism/valid"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// putValidHandleError handles the PutValid error response.
func (client *PolymorphismClient) putValidHandleError(resp *azcore.Response) error {
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

// PutValidMissingRequired - Put complex types that are polymorphic, attempting to omit required 'birthday' field - the request should not be allowed from
// the client
// If the operation fails it returns the *Error error type.
func (client *PolymorphismClient) PutValidMissingRequired(ctx context.Context, complexBody FishClassification, options *PolymorphismPutValidMissingRequiredOptions) (*http.Response, error) {
	req, err := client.putValidMissingRequiredCreateRequest(ctx, complexBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putValidMissingRequiredHandleError(resp)
	}
	return resp.Response, nil
}

// putValidMissingRequiredCreateRequest creates the PutValidMissingRequired request.
func (client *PolymorphismClient) putValidMissingRequiredCreateRequest(ctx context.Context, complexBody FishClassification, options *PolymorphismPutValidMissingRequiredOptions) (*azcore.Request, error) {
	urlPath := "/complex/polymorphism/missingrequired/invalid"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// putValidMissingRequiredHandleError handles the PutValidMissingRequired error response.
func (client *PolymorphismClient) putValidMissingRequiredHandleError(resp *azcore.Response) error {
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
