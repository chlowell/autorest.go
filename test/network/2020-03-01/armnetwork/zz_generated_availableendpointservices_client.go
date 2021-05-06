// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// AvailableEndpointServicesClient contains the methods for the AvailableEndpointServices group.
// Don't use this type directly, use NewAvailableEndpointServicesClient() instead.
type AvailableEndpointServicesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewAvailableEndpointServicesClient creates a new instance of AvailableEndpointServicesClient with the specified values.
func NewAvailableEndpointServicesClient(con *armcore.Connection, subscriptionID string) *AvailableEndpointServicesClient {
	return &AvailableEndpointServicesClient{con: con, subscriptionID: subscriptionID}
}

// List - List what values of endpoint services are available for use.
// If the operation fails it returns the *CloudError error type.
func (client *AvailableEndpointServicesClient) List(location string, options *AvailableEndpointServicesListOptions) EndpointServicesListResultPager {
	return &endpointServicesListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, location, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp EndpointServicesListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.EndpointServicesListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *AvailableEndpointServicesClient) listCreateRequest(ctx context.Context, location string, options *AvailableEndpointServicesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/virtualNetworkAvailableEndpointServices"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AvailableEndpointServicesClient) listHandleResponse(resp *azcore.Response) (EndpointServicesListResultResponse, error) {
	var val *EndpointServicesListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return EndpointServicesListResultResponse{}, err
	}
	return EndpointServicesListResultResponse{RawResponse: resp.Response, EndpointServicesListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *AvailableEndpointServicesClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
