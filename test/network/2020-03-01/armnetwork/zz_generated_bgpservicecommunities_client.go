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

// BgpServiceCommunitiesClient contains the methods for the BgpServiceCommunities group.
// Don't use this type directly, use NewBgpServiceCommunitiesClient() instead.
type BgpServiceCommunitiesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewBgpServiceCommunitiesClient creates a new instance of BgpServiceCommunitiesClient with the specified values.
func NewBgpServiceCommunitiesClient(con *armcore.Connection, subscriptionID string) *BgpServiceCommunitiesClient {
	return &BgpServiceCommunitiesClient{con: con, subscriptionID: subscriptionID}
}

// List - Gets all the available bgp service communities.
// If the operation fails it returns the *CloudError error type.
func (client *BgpServiceCommunitiesClient) List(options *BgpServiceCommunitiesListOptions) BgpServiceCommunityListResultPager {
	return &bgpServiceCommunityListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp BgpServiceCommunityListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.BgpServiceCommunityListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *BgpServiceCommunitiesClient) listCreateRequest(ctx context.Context, options *BgpServiceCommunitiesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/bgpServiceCommunities"
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
func (client *BgpServiceCommunitiesClient) listHandleResponse(resp *azcore.Response) (BgpServiceCommunityListResultResponse, error) {
	var val *BgpServiceCommunityListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return BgpServiceCommunityListResultResponse{}, err
	}
	return BgpServiceCommunityListResultResponse{RawResponse: resp.Response, BgpServiceCommunityListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *BgpServiceCommunitiesClient) listHandleError(resp *azcore.Response) error {
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
