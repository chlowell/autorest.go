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

// TagsClient contains the methods for the Tags group.
// Don't use this type directly, use NewTagsClient() instead.
type TagsClient struct {
	ep string
	pl runtime.Pipeline
}

// NewTagsClient creates a new instance of TagsClient with the specified values.
func NewTagsClient(con *arm.Connection) *TagsClient {
	return &TagsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version)}
}

// Get - Get all available tag keys for the defined scope
// If the operation fails it returns the *ErrorResponse error type.
func (client *TagsClient) Get(ctx context.Context, scope string, options *TagsGetOptions) (TagsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, scope, options)
	if err != nil {
		return TagsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TagsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return TagsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *TagsClient) getCreateRequest(ctx context.Context, scope string, options *TagsGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Consumption/tags"
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

// getHandleResponse handles the Get response.
func (client *TagsClient) getHandleResponse(resp *http.Response) (TagsGetResponse, error) {
	result := TagsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.TagsResult); err != nil {
		return TagsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *TagsClient) getHandleError(resp *http.Response) error {
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
