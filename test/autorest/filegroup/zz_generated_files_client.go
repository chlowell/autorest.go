//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package filegroup

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// FilesClient contains the methods for the Files group.
// Don't use this type directly, use NewFilesClient() instead.
type FilesClient struct {
	con *Connection
}

// NewFilesClient creates a new instance of FilesClient with the specified values.
func NewFilesClient(con *Connection) *FilesClient {
	return &FilesClient{con: con}
}

// GetEmptyFile - Get empty file
// If the operation fails it returns the *Error error type.
func (client *FilesClient) GetEmptyFile(ctx context.Context, options *FilesGetEmptyFileOptions) (FilesGetEmptyFileResponse, error) {
	req, err := client.getEmptyFileCreateRequest(ctx, options)
	if err != nil {
		return FilesGetEmptyFileResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return FilesGetEmptyFileResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FilesGetEmptyFileResponse{}, client.getEmptyFileHandleError(resp)
	}
	return FilesGetEmptyFileResponse{RawResponse: resp}, nil
}

// getEmptyFileCreateRequest creates the GetEmptyFile request.
func (client *FilesClient) getEmptyFileCreateRequest(ctx context.Context, options *FilesGetEmptyFileOptions) (*policy.Request, error) {
	urlPath := "/files/stream/empty"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.SkipBodyDownload()
	req.Raw().Header.Set("Accept", "image/png, application/json")
	return req, nil
}

// getEmptyFileHandleError handles the GetEmptyFile error response.
func (client *FilesClient) getEmptyFileHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetFile - Get file
// If the operation fails it returns the *Error error type.
func (client *FilesClient) GetFile(ctx context.Context, options *FilesGetFileOptions) (FilesGetFileResponse, error) {
	req, err := client.getFileCreateRequest(ctx, options)
	if err != nil {
		return FilesGetFileResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return FilesGetFileResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FilesGetFileResponse{}, client.getFileHandleError(resp)
	}
	return FilesGetFileResponse{RawResponse: resp}, nil
}

// getFileCreateRequest creates the GetFile request.
func (client *FilesClient) getFileCreateRequest(ctx context.Context, options *FilesGetFileOptions) (*policy.Request, error) {
	urlPath := "/files/stream/nonempty"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.SkipBodyDownload()
	req.Raw().Header.Set("Accept", "image/png, application/json")
	return req, nil
}

// getFileHandleError handles the GetFile error response.
func (client *FilesClient) getFileHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetFileLarge - Get a large file
// If the operation fails it returns the *Error error type.
func (client *FilesClient) GetFileLarge(ctx context.Context, options *FilesGetFileLargeOptions) (FilesGetFileLargeResponse, error) {
	req, err := client.getFileLargeCreateRequest(ctx, options)
	if err != nil {
		return FilesGetFileLargeResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return FilesGetFileLargeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FilesGetFileLargeResponse{}, client.getFileLargeHandleError(resp)
	}
	return FilesGetFileLargeResponse{RawResponse: resp}, nil
}

// getFileLargeCreateRequest creates the GetFileLarge request.
func (client *FilesClient) getFileLargeCreateRequest(ctx context.Context, options *FilesGetFileLargeOptions) (*policy.Request, error) {
	urlPath := "/files/stream/verylarge"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.SkipBodyDownload()
	req.Raw().Header.Set("Accept", "image/png, application/json")
	return req, nil
}

// getFileLargeHandleError handles the GetFileLarge error response.
func (client *FilesClient) getFileLargeHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
