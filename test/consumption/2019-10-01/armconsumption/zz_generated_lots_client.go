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
	"strings"
)

// LotsClient contains the methods for the Lots group.
// Don't use this type directly, use NewLotsClient() instead.
type LotsClient struct {
	ep string
	pl runtime.Pipeline
}

// NewLotsClient creates a new instance of LotsClient with the specified values.
func NewLotsClient(con *arm.Connection) *LotsClient {
	return &LotsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version)}
}

// List - Lists the lots by billingAccountId and billingProfileId.
// If the operation fails it returns the *ErrorResponse error type.
func (client *LotsClient) List(scope string, options *LotsListOptions) *LotsListPager {
	return &LotsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, scope, options)
		},
		advancer: func(ctx context.Context, resp LotsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.Lots.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *LotsClient) listCreateRequest(ctx context.Context, scope string, options *LotsListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Consumption/lots"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LotsClient) listHandleResponse(resp *http.Response) (LotsListResponse, error) {
	result := LotsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Lots); err != nil {
		return LotsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *LotsClient) listHandleError(resp *http.Response) error {
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
