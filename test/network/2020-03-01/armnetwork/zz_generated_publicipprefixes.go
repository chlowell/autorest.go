// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// PublicIPPrefixesClient contains the methods for the PublicIPPrefixes group.
// Don't use this type directly, use NewPublicIPPrefixesClient() instead.
type PublicIPPrefixesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewPublicIPPrefixesClient creates a new instance of PublicIPPrefixesClient with the specified values.
func NewPublicIPPrefixesClient(con *armcore.Connection, subscriptionID string) PublicIPPrefixesClient {
	return PublicIPPrefixesClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client PublicIPPrefixesClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// BeginCreateOrUpdate - Creates or updates a static or dynamic public IP prefix.
func (client PublicIPPrefixesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, publicIPPrefixName string, parameters PublicIPPrefix, options *PublicIPPrefixesCreateOrUpdateOptions) (*PublicIPPrefixPollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, publicIPPrefixName, parameters, options)
	if err != nil {
		return nil, err
	}
	result := &PublicIPPrefixPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("PublicIPPrefixesClient.CreateOrUpdate", "location", resp, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &publicIPPrefixPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*PublicIPPrefixResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new PublicIPPrefixPoller from the specified resume token.
// token - The value must come from a previous call to PublicIPPrefixPoller.ResumeToken().
func (client PublicIPPrefixesClient) ResumeCreateOrUpdate(token string) (PublicIPPrefixPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("PublicIPPrefixesClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &publicIPPrefixPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates a static or dynamic public IP prefix.
func (client PublicIPPrefixesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, publicIPPrefixName string, parameters PublicIPPrefix, options *PublicIPPrefixesCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, publicIPPrefixName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client PublicIPPrefixesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, publicIPPrefixName string, parameters PublicIPPrefix, options *PublicIPPrefixesCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIpPrefixName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{publicIpPrefixName}", url.PathEscape(publicIPPrefixName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client PublicIPPrefixesClient) createOrUpdateHandleResponse(resp *azcore.Response) (*PublicIPPrefixResponse, error) {
	result := PublicIPPrefixResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PublicIPPrefix)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client PublicIPPrefixesClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Deletes the specified public IP prefix.
func (client PublicIPPrefixesClient) BeginDelete(ctx context.Context, resourceGroupName string, publicIPPrefixName string, options *PublicIPPrefixesDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, publicIPPrefixName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("PublicIPPrefixesClient.Delete", "location", resp, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client PublicIPPrefixesClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("PublicIPPrefixesClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified public IP prefix.
func (client PublicIPPrefixesClient) Delete(ctx context.Context, resourceGroupName string, publicIPPrefixName string, options *PublicIPPrefixesDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, publicIPPrefixName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client PublicIPPrefixesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, publicIPPrefixName string, options *PublicIPPrefixesDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIpPrefixName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{publicIpPrefixName}", url.PathEscape(publicIPPrefixName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client PublicIPPrefixesClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified public IP prefix in a specified resource group.
func (client PublicIPPrefixesClient) Get(ctx context.Context, resourceGroupName string, publicIPPrefixName string, options *PublicIPPrefixesGetOptions) (*PublicIPPrefixResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, publicIPPrefixName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client PublicIPPrefixesClient) getCreateRequest(ctx context.Context, resourceGroupName string, publicIPPrefixName string, options *PublicIPPrefixesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIpPrefixName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{publicIpPrefixName}", url.PathEscape(publicIPPrefixName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	if options != nil && options.Expand != nil {
		query.Set("$expand", *options.Expand)
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client PublicIPPrefixesClient) getHandleResponse(resp *azcore.Response) (*PublicIPPrefixResponse, error) {
	result := PublicIPPrefixResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PublicIPPrefix)
}

// getHandleError handles the Get error response.
func (client PublicIPPrefixesClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all public IP prefixes in a resource group.
func (client PublicIPPrefixesClient) List(resourceGroupName string, options *PublicIPPrefixesListOptions) PublicIPPrefixListResultPager {
	return &publicIPPrefixListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp *PublicIPPrefixListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.PublicIPPrefixListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client PublicIPPrefixesClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *PublicIPPrefixesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client PublicIPPrefixesClient) listHandleResponse(resp *azcore.Response) (*PublicIPPrefixListResultResponse, error) {
	result := PublicIPPrefixListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PublicIPPrefixListResult)
}

// listHandleError handles the List error response.
func (client PublicIPPrefixesClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListAll - Gets all the public IP prefixes in a subscription.
func (client PublicIPPrefixesClient) ListAll(options *PublicIPPrefixesListAllOptions) PublicIPPrefixListResultPager {
	return &publicIPPrefixListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listAllCreateRequest(ctx, options)
		},
		responder: client.listAllHandleResponse,
		errorer:   client.listAllHandleError,
		advancer: func(ctx context.Context, resp *PublicIPPrefixListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.PublicIPPrefixListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listAllCreateRequest creates the ListAll request.
func (client PublicIPPrefixesClient) listAllCreateRequest(ctx context.Context, options *PublicIPPrefixesListAllOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/publicIPPrefixes"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client PublicIPPrefixesClient) listAllHandleResponse(resp *azcore.Response) (*PublicIPPrefixListResultResponse, error) {
	result := PublicIPPrefixListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PublicIPPrefixListResult)
}

// listAllHandleError handles the ListAll error response.
func (client PublicIPPrefixesClient) listAllHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// UpdateTags - Updates public IP prefix tags.
func (client PublicIPPrefixesClient) UpdateTags(ctx context.Context, resourceGroupName string, publicIPPrefixName string, parameters TagsObject, options *PublicIPPrefixesUpdateTagsOptions) (*PublicIPPrefixResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, publicIPPrefixName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.updateTagsHandleError(resp)
	}
	result, err := client.updateTagsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client PublicIPPrefixesClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, publicIPPrefixName string, parameters TagsObject, options *PublicIPPrefixesUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIpPrefixName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{publicIpPrefixName}", url.PathEscape(publicIPPrefixName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client PublicIPPrefixesClient) updateTagsHandleResponse(resp *azcore.Response) (*PublicIPPrefixResponse, error) {
	result := PublicIPPrefixResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.PublicIPPrefix)
}

// updateTagsHandleError handles the UpdateTags error response.
func (client PublicIPPrefixesClient) updateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
