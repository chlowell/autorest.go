//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconsumption

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strconv"
	"strings"
)

// UsageDetailsClient contains the methods for the UsageDetails group.
// Don't use this type directly, use NewUsageDetailsClient() instead.
type UsageDetailsClient struct {
	ep string
	pl runtime.Pipeline
}

// NewUsageDetailsClient creates a new instance of UsageDetailsClient with the specified values.
func NewUsageDetailsClient(con *arm.Connection) *UsageDetailsClient {
	return &UsageDetailsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version)}
}

// List - Lists the usage details for the defined scope. Usage details are available via this API only for May 1, 2014 or later. For more information on
// using this API, including how to specify a date range,
// please see: https://docs.microsoft.com/en-us/azure/cost-management-billing/costs/manage-automation
// If the operation fails it returns the *ErrorResponse error type.
func (client *UsageDetailsClient) List(scope string, options *UsageDetailsListOptions) *UsageDetailsListPager {
	return &UsageDetailsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, scope, options)
		},
		advancer: func(ctx context.Context, resp UsageDetailsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.UsageDetailsListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *UsageDetailsClient) listCreateRequest(ctx context.Context, scope string, options *UsageDetailsListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Consumption/usageDetails"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Skiptoken != nil {
		reqQP.Set("$skiptoken", *options.Skiptoken)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2019-10-01")
	if options != nil && options.Metric != nil {
		reqQP.Set("metric", string(*options.Metric))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *UsageDetailsClient) listHandleResponse(resp *http.Response) (UsageDetailsListResponse, error) {
	result := UsageDetailsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.UsageDetailsListResult); err != nil {
		return UsageDetailsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *UsageDetailsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
