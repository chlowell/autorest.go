//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

type sqlScriptClient struct {
	con *connection
}

// BeginCreateOrUpdateSQLScript - Creates or updates a Sql Script.
// If the operation fails it returns the *CloudError error type.
func (client *sqlScriptClient) BeginCreateOrUpdateSQLScript(ctx context.Context, sqlScriptName string, sqlScript SQLScriptResource, options *SQLScriptBeginCreateOrUpdateSQLScriptOptions) (SQLScriptCreateOrUpdateSQLScriptPollerResponse, error) {
	resp, err := client.createOrUpdateSQLScript(ctx, sqlScriptName, sqlScript, options)
	if err != nil {
		return SQLScriptCreateOrUpdateSQLScriptPollerResponse{}, err
	}
	result := SQLScriptCreateOrUpdateSQLScriptPollerResponse{
		RawResponse: resp,
	}
	pt, err := runtime.NewPoller("sqlScriptClient.CreateOrUpdateSQLScript", resp, client.con.Pipeline(), client.createOrUpdateSQLScriptHandleError)
	if err != nil {
		return SQLScriptCreateOrUpdateSQLScriptPollerResponse{}, err
	}
	result.Poller = &SQLScriptCreateOrUpdateSQLScriptPoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdateSQLScript - Creates or updates a Sql Script.
// If the operation fails it returns the *CloudError error type.
func (client *sqlScriptClient) createOrUpdateSQLScript(ctx context.Context, sqlScriptName string, sqlScript SQLScriptResource, options *SQLScriptBeginCreateOrUpdateSQLScriptOptions) (*http.Response, error) {
	req, err := client.createOrUpdateSQLScriptCreateRequest(ctx, sqlScriptName, sqlScript, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateSQLScriptHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateSQLScriptCreateRequest creates the CreateOrUpdateSQLScript request.
func (client *sqlScriptClient) createOrUpdateSQLScriptCreateRequest(ctx context.Context, sqlScriptName string, sqlScript SQLScriptResource, options *SQLScriptBeginCreateOrUpdateSQLScriptOptions) (*policy.Request, error) {
	urlPath := "/sqlScripts/{sqlScriptName}"
	if sqlScriptName == "" {
		return nil, errors.New("parameter sqlScriptName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlScriptName}", url.PathEscape(sqlScriptName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, sqlScript)
}

// createOrUpdateSQLScriptHandleError handles the CreateOrUpdateSQLScript error response.
func (client *sqlScriptClient) createOrUpdateSQLScriptHandleError(resp *http.Response) error {
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

// BeginDeleteSQLScript - Deletes a Sql Script.
// If the operation fails it returns the *CloudError error type.
func (client *sqlScriptClient) BeginDeleteSQLScript(ctx context.Context, sqlScriptName string, options *SQLScriptBeginDeleteSQLScriptOptions) (SQLScriptDeleteSQLScriptPollerResponse, error) {
	resp, err := client.deleteSQLScript(ctx, sqlScriptName, options)
	if err != nil {
		return SQLScriptDeleteSQLScriptPollerResponse{}, err
	}
	result := SQLScriptDeleteSQLScriptPollerResponse{
		RawResponse: resp,
	}
	pt, err := runtime.NewPoller("sqlScriptClient.DeleteSQLScript", resp, client.con.Pipeline(), client.deleteSQLScriptHandleError)
	if err != nil {
		return SQLScriptDeleteSQLScriptPollerResponse{}, err
	}
	result.Poller = &SQLScriptDeleteSQLScriptPoller{
		pt: pt,
	}
	return result, nil
}

// DeleteSQLScript - Deletes a Sql Script.
// If the operation fails it returns the *CloudError error type.
func (client *sqlScriptClient) deleteSQLScript(ctx context.Context, sqlScriptName string, options *SQLScriptBeginDeleteSQLScriptOptions) (*http.Response, error) {
	req, err := client.deleteSQLScriptCreateRequest(ctx, sqlScriptName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteSQLScriptHandleError(resp)
	}
	return resp, nil
}

// deleteSQLScriptCreateRequest creates the DeleteSQLScript request.
func (client *sqlScriptClient) deleteSQLScriptCreateRequest(ctx context.Context, sqlScriptName string, options *SQLScriptBeginDeleteSQLScriptOptions) (*policy.Request, error) {
	urlPath := "/sqlScripts/{sqlScriptName}"
	if sqlScriptName == "" {
		return nil, errors.New("parameter sqlScriptName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlScriptName}", url.PathEscape(sqlScriptName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteSQLScriptHandleError handles the DeleteSQLScript error response.
func (client *sqlScriptClient) deleteSQLScriptHandleError(resp *http.Response) error {
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

// GetSQLScript - Gets a sql script.
// If the operation fails it returns the *CloudError error type.
func (client *sqlScriptClient) GetSQLScript(ctx context.Context, sqlScriptName string, options *SQLScriptGetSQLScriptOptions) (SQLScriptGetSQLScriptResponse, error) {
	req, err := client.getSQLScriptCreateRequest(ctx, sqlScriptName, options)
	if err != nil {
		return SQLScriptGetSQLScriptResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SQLScriptGetSQLScriptResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNotModified) {
		return SQLScriptGetSQLScriptResponse{}, client.getSQLScriptHandleError(resp)
	}
	return client.getSQLScriptHandleResponse(resp)
}

// getSQLScriptCreateRequest creates the GetSQLScript request.
func (client *sqlScriptClient) getSQLScriptCreateRequest(ctx context.Context, sqlScriptName string, options *SQLScriptGetSQLScriptOptions) (*policy.Request, error) {
	urlPath := "/sqlScripts/{sqlScriptName}"
	if sqlScriptName == "" {
		return nil, errors.New("parameter sqlScriptName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlScriptName}", url.PathEscape(sqlScriptName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getSQLScriptHandleResponse handles the GetSQLScript response.
func (client *sqlScriptClient) getSQLScriptHandleResponse(resp *http.Response) (SQLScriptGetSQLScriptResponse, error) {
	result := SQLScriptGetSQLScriptResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLScriptResource); err != nil {
		return SQLScriptGetSQLScriptResponse{}, err
	}
	return result, nil
}

// getSQLScriptHandleError handles the GetSQLScript error response.
func (client *sqlScriptClient) getSQLScriptHandleError(resp *http.Response) error {
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

// GetSQLScriptsByWorkspace - Lists sql scripts.
// If the operation fails it returns the *CloudError error type.
func (client *sqlScriptClient) GetSQLScriptsByWorkspace(options *SQLScriptGetSQLScriptsByWorkspaceOptions) *SQLScriptGetSQLScriptsByWorkspacePager {
	return &SQLScriptGetSQLScriptsByWorkspacePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.getSQLScriptsByWorkspaceCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp SQLScriptGetSQLScriptsByWorkspaceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SQLScriptsListResponse.NextLink)
		},
	}
}

// getSQLScriptsByWorkspaceCreateRequest creates the GetSQLScriptsByWorkspace request.
func (client *sqlScriptClient) getSQLScriptsByWorkspaceCreateRequest(ctx context.Context, options *SQLScriptGetSQLScriptsByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/sqlScripts"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getSQLScriptsByWorkspaceHandleResponse handles the GetSQLScriptsByWorkspace response.
func (client *sqlScriptClient) getSQLScriptsByWorkspaceHandleResponse(resp *http.Response) (SQLScriptGetSQLScriptsByWorkspaceResponse, error) {
	result := SQLScriptGetSQLScriptsByWorkspaceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLScriptsListResponse); err != nil {
		return SQLScriptGetSQLScriptsByWorkspaceResponse{}, err
	}
	return result, nil
}

// getSQLScriptsByWorkspaceHandleError handles the GetSQLScriptsByWorkspace error response.
func (client *sqlScriptClient) getSQLScriptsByWorkspaceHandleError(resp *http.Response) error {
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

// BeginRenameSQLScript - Renames a sqlScript.
// If the operation fails it returns the *CloudError error type.
func (client *sqlScriptClient) BeginRenameSQLScript(ctx context.Context, sqlScriptName string, request ArtifactRenameRequest, options *SQLScriptBeginRenameSQLScriptOptions) (SQLScriptRenameSQLScriptPollerResponse, error) {
	resp, err := client.renameSQLScript(ctx, sqlScriptName, request, options)
	if err != nil {
		return SQLScriptRenameSQLScriptPollerResponse{}, err
	}
	result := SQLScriptRenameSQLScriptPollerResponse{
		RawResponse: resp,
	}
	pt, err := runtime.NewPoller("sqlScriptClient.RenameSQLScript", resp, client.con.Pipeline(), client.renameSQLScriptHandleError)
	if err != nil {
		return SQLScriptRenameSQLScriptPollerResponse{}, err
	}
	result.Poller = &SQLScriptRenameSQLScriptPoller{
		pt: pt,
	}
	return result, nil
}

// RenameSQLScript - Renames a sqlScript.
// If the operation fails it returns the *CloudError error type.
func (client *sqlScriptClient) renameSQLScript(ctx context.Context, sqlScriptName string, request ArtifactRenameRequest, options *SQLScriptBeginRenameSQLScriptOptions) (*http.Response, error) {
	req, err := client.renameSQLScriptCreateRequest(ctx, sqlScriptName, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.renameSQLScriptHandleError(resp)
	}
	return resp, nil
}

// renameSQLScriptCreateRequest creates the RenameSQLScript request.
func (client *sqlScriptClient) renameSQLScriptCreateRequest(ctx context.Context, sqlScriptName string, request ArtifactRenameRequest, options *SQLScriptBeginRenameSQLScriptOptions) (*policy.Request, error) {
	urlPath := "/sqlScripts/{sqlScriptName}/rename"
	if sqlScriptName == "" {
		return nil, errors.New("parameter sqlScriptName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlScriptName}", url.PathEscape(sqlScriptName))
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

// renameSQLScriptHandleError handles the RenameSQLScript error response.
func (client *sqlScriptClient) renameSQLScriptHandleError(resp *http.Response) error {
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
