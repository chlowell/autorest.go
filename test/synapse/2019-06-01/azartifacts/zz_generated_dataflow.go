// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

type dataFlowClient struct {
	*client
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *dataFlowClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// CreateOrUpdateDataFlow - Creates or updates a data flow.
func (client *dataFlowClient) CreateOrUpdateDataFlow(ctx context.Context, dataFlowName string, dataFlow DataFlowResource, dataFlowCreateOrUpdateDataFlowOptions *DataFlowCreateOrUpdateDataFlowOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateDataFlowCreateRequest(ctx, dataFlowName, dataFlow, dataFlowCreateOrUpdateDataFlowOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.CreateOrUpdateDataFlowHandleError(resp)
	}
	return resp, nil
}

// CreateOrUpdateDataFlowCreateRequest creates the CreateOrUpdateDataFlow request.
func (client *dataFlowClient) CreateOrUpdateDataFlowCreateRequest(ctx context.Context, dataFlowName string, dataFlow DataFlowResource, dataFlowCreateOrUpdateDataFlowOptions *DataFlowCreateOrUpdateDataFlowOptions) (*azcore.Request, error) {
	urlPath := "/dataflows/{dataFlowName}"
	urlPath = strings.ReplaceAll(urlPath, "{dataFlowName}", url.PathEscape(dataFlowName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	if dataFlowCreateOrUpdateDataFlowOptions != nil && dataFlowCreateOrUpdateDataFlowOptions.IfMatch != nil {
		req.Header.Set("If-Match", *dataFlowCreateOrUpdateDataFlowOptions.IfMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(dataFlow)
}

// CreateOrUpdateDataFlowHandleResponse handles the CreateOrUpdateDataFlow response.
func (client *dataFlowClient) CreateOrUpdateDataFlowHandleResponse(resp *azcore.Response) (*DataFlowResourceResponse, error) {
	result := DataFlowResourceResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DataFlowResource)
}

// CreateOrUpdateDataFlowHandleError handles the CreateOrUpdateDataFlow error response.
func (client *dataFlowClient) CreateOrUpdateDataFlowHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// DeleteDataFlow - Deletes a data flow.
func (client *dataFlowClient) DeleteDataFlow(ctx context.Context, dataFlowName string) (*azcore.Response, error) {
	req, err := client.DeleteDataFlowCreateRequest(ctx, dataFlowName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.DeleteDataFlowHandleError(resp)
	}
	return resp, nil
}

// DeleteDataFlowCreateRequest creates the DeleteDataFlow request.
func (client *dataFlowClient) DeleteDataFlowCreateRequest(ctx context.Context, dataFlowName string) (*azcore.Request, error) {
	urlPath := "/dataflows/{dataFlowName}"
	urlPath = strings.ReplaceAll(urlPath, "{dataFlowName}", url.PathEscape(dataFlowName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DeleteDataFlowHandleError handles the DeleteDataFlow error response.
func (client *dataFlowClient) DeleteDataFlowHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetDataFlow - Gets a data flow.
func (client *dataFlowClient) GetDataFlow(ctx context.Context, dataFlowName string, dataFlowGetDataFlowOptions *DataFlowGetDataFlowOptions) (*DataFlowResourceResponse, error) {
	req, err := client.GetDataFlowCreateRequest(ctx, dataFlowName, dataFlowGetDataFlowOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetDataFlowHandleError(resp)
	}
	result, err := client.GetDataFlowHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetDataFlowCreateRequest creates the GetDataFlow request.
func (client *dataFlowClient) GetDataFlowCreateRequest(ctx context.Context, dataFlowName string, dataFlowGetDataFlowOptions *DataFlowGetDataFlowOptions) (*azcore.Request, error) {
	urlPath := "/dataflows/{dataFlowName}"
	urlPath = strings.ReplaceAll(urlPath, "{dataFlowName}", url.PathEscape(dataFlowName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	if dataFlowGetDataFlowOptions != nil && dataFlowGetDataFlowOptions.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *dataFlowGetDataFlowOptions.IfNoneMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetDataFlowHandleResponse handles the GetDataFlow response.
func (client *dataFlowClient) GetDataFlowHandleResponse(resp *azcore.Response) (*DataFlowResourceResponse, error) {
	result := DataFlowResourceResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DataFlowResource)
}

// GetDataFlowHandleError handles the GetDataFlow error response.
func (client *dataFlowClient) GetDataFlowHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetDataFlowsByWorkspace - Lists data flows.
func (client *dataFlowClient) GetDataFlowsByWorkspace() DataFlowListResponsePager {
	return &dataFlowListResponsePager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.GetDataFlowsByWorkspaceCreateRequest(ctx)
		},
		responder: client.GetDataFlowsByWorkspaceHandleResponse,
		errorer:   client.GetDataFlowsByWorkspaceHandleError,
		advancer: func(ctx context.Context, resp *DataFlowListResponseResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.DataFlowListResponse.NextLink)
		},
	}
}

// GetDataFlowsByWorkspaceCreateRequest creates the GetDataFlowsByWorkspace request.
func (client *dataFlowClient) GetDataFlowsByWorkspaceCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/dataflows"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetDataFlowsByWorkspaceHandleResponse handles the GetDataFlowsByWorkspace response.
func (client *dataFlowClient) GetDataFlowsByWorkspaceHandleResponse(resp *azcore.Response) (*DataFlowListResponseResponse, error) {
	result := DataFlowListResponseResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DataFlowListResponse)
}

// GetDataFlowsByWorkspaceHandleError handles the GetDataFlowsByWorkspace error response.
func (client *dataFlowClient) GetDataFlowsByWorkspaceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}