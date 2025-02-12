//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

type dataFlowDebugSessionClient struct {
	con *connection
}

// AddDataFlow - Add a data flow into debug session.
// If the operation fails it returns the *CloudError error type.
func (client *dataFlowDebugSessionClient) AddDataFlow(ctx context.Context, request DataFlowDebugPackage, options *DataFlowDebugSessionAddDataFlowOptions) (DataFlowDebugSessionAddDataFlowResponse, error) {
	req, err := client.addDataFlowCreateRequest(ctx, request, options)
	if err != nil {
		return DataFlowDebugSessionAddDataFlowResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DataFlowDebugSessionAddDataFlowResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataFlowDebugSessionAddDataFlowResponse{}, client.addDataFlowHandleError(resp)
	}
	return client.addDataFlowHandleResponse(resp)
}

// addDataFlowCreateRequest creates the AddDataFlow request.
func (client *dataFlowDebugSessionClient) addDataFlowCreateRequest(ctx context.Context, request DataFlowDebugPackage, options *DataFlowDebugSessionAddDataFlowOptions) (*policy.Request, error) {
	urlPath := "/addDataFlowToDebugSession"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, request)
}

// addDataFlowHandleResponse handles the AddDataFlow response.
func (client *dataFlowDebugSessionClient) addDataFlowHandleResponse(resp *http.Response) (DataFlowDebugSessionAddDataFlowResponse, error) {
	result := DataFlowDebugSessionAddDataFlowResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AddDataFlowToDebugSessionResponse); err != nil {
		return DataFlowDebugSessionAddDataFlowResponse{}, err
	}
	return result, nil
}

// addDataFlowHandleError handles the AddDataFlow error response.
func (client *dataFlowDebugSessionClient) addDataFlowHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginCreateDataFlowDebugSession - Creates a data flow debug session.
// If the operation fails it returns the *CloudError error type.
func (client *dataFlowDebugSessionClient) BeginCreateDataFlowDebugSession(ctx context.Context, request CreateDataFlowDebugSessionRequest, options *DataFlowDebugSessionBeginCreateDataFlowDebugSessionOptions) (DataFlowDebugSessionCreateDataFlowDebugSessionPollerResponse, error) {
	resp, err := client.createDataFlowDebugSession(ctx, request, options)
	if err != nil {
		return DataFlowDebugSessionCreateDataFlowDebugSessionPollerResponse{}, err
	}
	result := DataFlowDebugSessionCreateDataFlowDebugSessionPollerResponse{
		RawResponse: resp,
	}
	pt, err := runtime.NewPoller("dataFlowDebugSessionClient.CreateDataFlowDebugSession", resp, client.con.Pipeline(), client.createDataFlowDebugSessionHandleError)
	if err != nil {
		return DataFlowDebugSessionCreateDataFlowDebugSessionPollerResponse{}, err
	}
	result.Poller = &DataFlowDebugSessionCreateDataFlowDebugSessionPoller{
		pt: pt,
	}
	return result, nil
}

// CreateDataFlowDebugSession - Creates a data flow debug session.
// If the operation fails it returns the *CloudError error type.
func (client *dataFlowDebugSessionClient) createDataFlowDebugSession(ctx context.Context, request CreateDataFlowDebugSessionRequest, options *DataFlowDebugSessionBeginCreateDataFlowDebugSessionOptions) (*http.Response, error) {
	req, err := client.createDataFlowDebugSessionCreateRequest(ctx, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.createDataFlowDebugSessionHandleError(resp)
	}
	return resp, nil
}

// createDataFlowDebugSessionCreateRequest creates the CreateDataFlowDebugSession request.
func (client *dataFlowDebugSessionClient) createDataFlowDebugSessionCreateRequest(ctx context.Context, request CreateDataFlowDebugSessionRequest, options *DataFlowDebugSessionBeginCreateDataFlowDebugSessionOptions) (*policy.Request, error) {
	urlPath := "/createDataFlowDebugSession"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, request)
}

// createDataFlowDebugSessionHandleError handles the CreateDataFlowDebugSession error response.
func (client *dataFlowDebugSessionClient) createDataFlowDebugSessionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// DeleteDataFlowDebugSession - Deletes a data flow debug session.
// If the operation fails it returns the *CloudError error type.
func (client *dataFlowDebugSessionClient) DeleteDataFlowDebugSession(ctx context.Context, request DeleteDataFlowDebugSessionRequest, options *DataFlowDebugSessionDeleteDataFlowDebugSessionOptions) (DataFlowDebugSessionDeleteDataFlowDebugSessionResponse, error) {
	req, err := client.deleteDataFlowDebugSessionCreateRequest(ctx, request, options)
	if err != nil {
		return DataFlowDebugSessionDeleteDataFlowDebugSessionResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DataFlowDebugSessionDeleteDataFlowDebugSessionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataFlowDebugSessionDeleteDataFlowDebugSessionResponse{}, client.deleteDataFlowDebugSessionHandleError(resp)
	}
	return DataFlowDebugSessionDeleteDataFlowDebugSessionResponse{RawResponse: resp}, nil
}

// deleteDataFlowDebugSessionCreateRequest creates the DeleteDataFlowDebugSession request.
func (client *dataFlowDebugSessionClient) deleteDataFlowDebugSessionCreateRequest(ctx context.Context, request DeleteDataFlowDebugSessionRequest, options *DataFlowDebugSessionDeleteDataFlowDebugSessionOptions) (*policy.Request, error) {
	urlPath := "/deleteDataFlowDebugSession"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, request)
}

// deleteDataFlowDebugSessionHandleError handles the DeleteDataFlowDebugSession error response.
func (client *dataFlowDebugSessionClient) deleteDataFlowDebugSessionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginExecuteCommand - Execute a data flow debug command.
// If the operation fails it returns the *CloudError error type.
func (client *dataFlowDebugSessionClient) BeginExecuteCommand(ctx context.Context, request DataFlowDebugCommandRequest, options *DataFlowDebugSessionBeginExecuteCommandOptions) (DataFlowDebugSessionExecuteCommandPollerResponse, error) {
	resp, err := client.executeCommand(ctx, request, options)
	if err != nil {
		return DataFlowDebugSessionExecuteCommandPollerResponse{}, err
	}
	result := DataFlowDebugSessionExecuteCommandPollerResponse{
		RawResponse: resp,
	}
	pt, err := runtime.NewPoller("dataFlowDebugSessionClient.ExecuteCommand", resp, client.con.Pipeline(), client.executeCommandHandleError)
	if err != nil {
		return DataFlowDebugSessionExecuteCommandPollerResponse{}, err
	}
	result.Poller = &DataFlowDebugSessionExecuteCommandPoller{
		pt: pt,
	}
	return result, nil
}

// ExecuteCommand - Execute a data flow debug command.
// If the operation fails it returns the *CloudError error type.
func (client *dataFlowDebugSessionClient) executeCommand(ctx context.Context, request DataFlowDebugCommandRequest, options *DataFlowDebugSessionBeginExecuteCommandOptions) (*http.Response, error) {
	req, err := client.executeCommandCreateRequest(ctx, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.executeCommandHandleError(resp)
	}
	return resp, nil
}

// executeCommandCreateRequest creates the ExecuteCommand request.
func (client *dataFlowDebugSessionClient) executeCommandCreateRequest(ctx context.Context, request DataFlowDebugCommandRequest, options *DataFlowDebugSessionBeginExecuteCommandOptions) (*policy.Request, error) {
	urlPath := "/executeDataFlowDebugCommand"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, request)
}

// executeCommandHandleError handles the ExecuteCommand error response.
func (client *dataFlowDebugSessionClient) executeCommandHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// QueryDataFlowDebugSessionsByWorkspace - Query all active data flow debug sessions.
// If the operation fails it returns the *CloudError error type.
func (client *dataFlowDebugSessionClient) QueryDataFlowDebugSessionsByWorkspace(options *DataFlowDebugSessionQueryDataFlowDebugSessionsByWorkspaceOptions) *DataFlowDebugSessionQueryDataFlowDebugSessionsByWorkspacePager {
	return &DataFlowDebugSessionQueryDataFlowDebugSessionsByWorkspacePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.queryDataFlowDebugSessionsByWorkspaceCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp DataFlowDebugSessionQueryDataFlowDebugSessionsByWorkspaceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.QueryDataFlowDebugSessionsResponse.NextLink)
		},
	}
}

// queryDataFlowDebugSessionsByWorkspaceCreateRequest creates the QueryDataFlowDebugSessionsByWorkspace request.
func (client *dataFlowDebugSessionClient) queryDataFlowDebugSessionsByWorkspaceCreateRequest(ctx context.Context, options *DataFlowDebugSessionQueryDataFlowDebugSessionsByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/queryDataFlowDebugSessions"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// queryDataFlowDebugSessionsByWorkspaceHandleResponse handles the QueryDataFlowDebugSessionsByWorkspace response.
func (client *dataFlowDebugSessionClient) queryDataFlowDebugSessionsByWorkspaceHandleResponse(resp *http.Response) (DataFlowDebugSessionQueryDataFlowDebugSessionsByWorkspaceResponse, error) {
	result := DataFlowDebugSessionQueryDataFlowDebugSessionsByWorkspaceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.QueryDataFlowDebugSessionsResponse); err != nil {
		return DataFlowDebugSessionQueryDataFlowDebugSessionsByWorkspaceResponse{}, err
	}
	return result, nil
}

// queryDataFlowDebugSessionsByWorkspaceHandleError handles the QueryDataFlowDebugSessionsByWorkspace error response.
func (client *dataFlowDebugSessionClient) queryDataFlowDebugSessionsByWorkspaceHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
