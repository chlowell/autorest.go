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

// AzureFirewallFqdnTagsClient contains the methods for the AzureFirewallFqdnTags group.
// Don't use this type directly, use NewAzureFirewallFqdnTagsClient() instead.
type AzureFirewallFqdnTagsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewAzureFirewallFqdnTagsClient creates a new instance of AzureFirewallFqdnTagsClient with the specified values.
func NewAzureFirewallFqdnTagsClient(con *armcore.Connection, subscriptionID string) *AzureFirewallFqdnTagsClient {
	return &AzureFirewallFqdnTagsClient{con: con, subscriptionID: subscriptionID}
}

// ListAll - Gets all the Azure Firewall FQDN Tags in a subscription.
// If the operation fails it returns the *CloudError error type.
func (client *AzureFirewallFqdnTagsClient) ListAll(options *AzureFirewallFqdnTagsListAllOptions) AzureFirewallFqdnTagListResultPager {
	return &azureFirewallFqdnTagListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listAllCreateRequest(ctx, options)
		},
		responder: client.listAllHandleResponse,
		errorer:   client.listAllHandleError,
		advancer: func(ctx context.Context, resp AzureFirewallFqdnTagListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.AzureFirewallFqdnTagListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listAllCreateRequest creates the ListAll request.
func (client *AzureFirewallFqdnTagsClient) listAllCreateRequest(ctx context.Context, options *AzureFirewallFqdnTagsListAllOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/azureFirewallFqdnTags"
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

// listAllHandleResponse handles the ListAll response.
func (client *AzureFirewallFqdnTagsClient) listAllHandleResponse(resp *azcore.Response) (AzureFirewallFqdnTagListResultResponse, error) {
	var val *AzureFirewallFqdnTagListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return AzureFirewallFqdnTagListResultResponse{}, err
	}
	return AzureFirewallFqdnTagListResultResponse{RawResponse: resp.Response, AzureFirewallFqdnTagListResult: val}, nil
}

// listAllHandleError handles the ListAll error response.
func (client *AzureFirewallFqdnTagsClient) listAllHandleError(resp *azcore.Response) error {
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
