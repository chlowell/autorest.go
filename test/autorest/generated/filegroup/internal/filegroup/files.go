// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package filegroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"path"
)

type FilesOperations struct{}

// GetEmptyFileCreateRequest creates the GetEmptyFile request.
func (FilesOperations) GetEmptyFileCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/files/stream/empty"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	req.SkipBodyDownload()
	return req, nil
}

// GetEmptyFileHandleResponse handles the GetEmptyFile response.
func (FilesOperations) GetEmptyFileHandleResponse(resp *azcore.Response) (*FilesGetEmptyFileResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &FilesGetEmptyFileResponse{RawResponse: resp.Response}, nil
}

// GetFileCreateRequest creates the GetFile request.
func (FilesOperations) GetFileCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/files/stream/nonempty"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	req.SkipBodyDownload()
	return req, nil
}

// GetFileHandleResponse handles the GetFile response.
func (FilesOperations) GetFileHandleResponse(resp *azcore.Response) (*FilesGetFileResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &FilesGetFileResponse{RawResponse: resp.Response}, nil
}

// GetFileLargeCreateRequest creates the GetFileLarge request.
func (FilesOperations) GetFileLargeCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/files/stream/verylarge"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	req.SkipBodyDownload()
	return req, nil
}

// GetFileLargeHandleResponse handles the GetFileLarge response.
func (FilesOperations) GetFileLargeHandleResponse(resp *azcore.Response) (*FilesGetFileLargeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &FilesGetFileLargeResponse{RawResponse: resp.Response}, nil
}