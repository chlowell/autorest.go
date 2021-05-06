// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package formdatagroup

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// FormdataClient contains the methods for the Formdata group.
// Don't use this type directly, use NewFormdataClient() instead.
type FormdataClient struct {
	con *Connection
}

// NewFormdataClient creates a new instance of FormdataClient with the specified values.
func NewFormdataClient(con *Connection) *FormdataClient {
	return &FormdataClient{con: con}
}

// UploadFile - Upload file
// If the operation fails it returns the *Error error type.
func (client *FormdataClient) UploadFile(ctx context.Context, fileContent azcore.ReadSeekCloser, fileName string, options *FormdataUploadFileOptions) (*http.Response, error) {
	req, err := client.uploadFileCreateRequest(ctx, fileContent, fileName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.uploadFileHandleError(resp)
	}
	return resp.Response, nil
}

// uploadFileCreateRequest creates the UploadFile request.
func (client *FormdataClient) uploadFileCreateRequest(ctx context.Context, fileContent azcore.ReadSeekCloser, fileName string, options *FormdataUploadFileOptions) (*azcore.Request, error) {
	urlPath := "/formdata/stream/uploadfile"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.SkipBodyDownload()
	req.Header.Set("Accept", "application/octet-stream, application/json")
	if err := req.SetMultipartFormData(map[string]interface{}{
		"fileContent": fileContent,
		"fileName":    fileName,
	}); err != nil {
		return nil, err
	}
	return req, nil
}

// uploadFileHandleError handles the UploadFile error response.
func (client *FormdataClient) uploadFileHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := Error{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// UploadFileViaBody - Upload file
// If the operation fails it returns the *Error error type.
func (client *FormdataClient) UploadFileViaBody(ctx context.Context, fileContent azcore.ReadSeekCloser, options *FormdataUploadFileViaBodyOptions) (*http.Response, error) {
	req, err := client.uploadFileViaBodyCreateRequest(ctx, fileContent, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.uploadFileViaBodyHandleError(resp)
	}
	return resp.Response, nil
}

// uploadFileViaBodyCreateRequest creates the UploadFileViaBody request.
func (client *FormdataClient) uploadFileViaBodyCreateRequest(ctx context.Context, fileContent azcore.ReadSeekCloser, options *FormdataUploadFileViaBodyOptions) (*azcore.Request, error) {
	urlPath := "/formdata/stream/uploadfile"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.SkipBodyDownload()
	req.Header.Set("Accept", "application/octet-stream, application/json")
	return req, req.SetBody(fileContent, "application/octet-stream")
}

// uploadFileViaBodyHandleError handles the UploadFileViaBody error response.
func (client *FormdataClient) uploadFileViaBodyHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := Error{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// UploadFiles - Upload multiple files
// If the operation fails it returns the *Error error type.
func (client *FormdataClient) UploadFiles(ctx context.Context, fileContent []azcore.ReadSeekCloser, options *FormdataUploadFilesOptions) (*http.Response, error) {
	req, err := client.uploadFilesCreateRequest(ctx, fileContent, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.uploadFilesHandleError(resp)
	}
	return resp.Response, nil
}

// uploadFilesCreateRequest creates the UploadFiles request.
func (client *FormdataClient) uploadFilesCreateRequest(ctx context.Context, fileContent []azcore.ReadSeekCloser, options *FormdataUploadFilesOptions) (*azcore.Request, error) {
	urlPath := "/formdata/stream/uploadfiles"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.SkipBodyDownload()
	req.Header.Set("Accept", "application/octet-stream, application/json")
	if err := req.SetMultipartFormData(map[string]interface{}{
		"fileContent": fileContent,
	}); err != nil {
		return nil, err
	}
	return req, nil
}

// uploadFilesHandleError handles the UploadFiles error response.
func (client *FormdataClient) uploadFilesHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := Error{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
