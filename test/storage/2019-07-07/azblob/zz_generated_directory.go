// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azblob

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"strconv"
	"time"
)

type directoryClient struct {
	con            *connection
	pathRenameMode *PathRenameMode
}

// Pipeline returns the pipeline associated with this client.
func (client directoryClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// Create - Create a directory. By default, the destination is overwritten and if the destination already exists and has a lease the lease is broken. This
// operation supports conditional HTTP requests. For more
// information, see Specifying Conditional Headers for Blob Service Operations [https://docs.microsoft.com/en-us/rest/api/storageservices/specifying-conditional-headers-for-blob-service-operations].
// To
// fail if the destination already exists, use a conditional request with If-None-Match: "*".
func (client directoryClient) Create(ctx context.Context, directoryCreateOptions *DirectoryCreateOptions, directoryHttpHeaders *DirectoryHttpHeaders, leaseAccessConditions *LeaseAccessConditions, modifiedAccessConditions *ModifiedAccessConditions) (*DirectoryCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, directoryCreateOptions, directoryHttpHeaders, leaseAccessConditions, modifiedAccessConditions)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusCreated) {
		return nil, client.createHandleError(resp)
	}
	result, err := client.createHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// createCreateRequest creates the Create request.
func (client directoryClient) createCreateRequest(ctx context.Context, directoryCreateOptions *DirectoryCreateOptions, directoryHttpHeaders *DirectoryHttpHeaders, leaseAccessConditions *LeaseAccessConditions, modifiedAccessConditions *ModifiedAccessConditions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodPut, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("resource", "directory")
	if directoryCreateOptions != nil && directoryCreateOptions.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*directoryCreateOptions.Timeout), 10))
	}
	req.URL.RawQuery = query.Encode()
	if directoryCreateOptions != nil && directoryCreateOptions.DirectoryProperties != nil {
		req.Header.Set("x-ms-properties", *directoryCreateOptions.DirectoryProperties)
	}
	if directoryCreateOptions != nil && directoryCreateOptions.PosixPermissions != nil {
		req.Header.Set("x-ms-permissions", *directoryCreateOptions.PosixPermissions)
	}
	if directoryCreateOptions != nil && directoryCreateOptions.PosixUmask != nil {
		req.Header.Set("x-ms-umask", *directoryCreateOptions.PosixUmask)
	}
	if directoryHttpHeaders != nil && directoryHttpHeaders.CacheControl != nil {
		req.Header.Set("x-ms-cache-control", *directoryHttpHeaders.CacheControl)
	}
	if directoryHttpHeaders != nil && directoryHttpHeaders.ContentType != nil {
		req.Header.Set("x-ms-content-type", *directoryHttpHeaders.ContentType)
	}
	if directoryHttpHeaders != nil && directoryHttpHeaders.ContentEncoding != nil {
		req.Header.Set("x-ms-content-encoding", *directoryHttpHeaders.ContentEncoding)
	}
	if directoryHttpHeaders != nil && directoryHttpHeaders.ContentLanguage != nil {
		req.Header.Set("x-ms-content-language", *directoryHttpHeaders.ContentLanguage)
	}
	if directoryHttpHeaders != nil && directoryHttpHeaders.ContentDisposition != nil {
		req.Header.Set("x-ms-content-disposition", *directoryHttpHeaders.ContentDisposition)
	}
	if leaseAccessConditions != nil && leaseAccessConditions.LeaseId != nil {
		req.Header.Set("x-ms-lease-id", *leaseAccessConditions.LeaseId)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfModifiedSince != nil {
		req.Header.Set("If-Modified-Since", modifiedAccessConditions.IfModifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfUnmodifiedSince != nil {
		req.Header.Set("If-Unmodified-Since", modifiedAccessConditions.IfUnmodifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfMatch != nil {
		req.Header.Set("If-Match", *modifiedAccessConditions.IfMatch)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *modifiedAccessConditions.IfNoneMatch)
	}
	req.Header.Set("x-ms-version", "2019-07-07")
	if directoryCreateOptions != nil && directoryCreateOptions.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *directoryCreateOptions.RequestId)
	}
	req.Header.Set("Accept", "application/xml")
	return req, nil
}

// createHandleResponse handles the Create response.
func (client directoryClient) createHandleResponse(resp *azcore.Response) (*DirectoryCreateResponse, error) {
	result := DirectoryCreateResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if val := resp.Header.Get("Last-Modified"); val != "" {
		lastModified, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return nil, err
		}
		result.LastModified = &lastModified
	}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Content-Length"); val != "" {
		contentLength, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return nil, err
		}
		result.ContentLength = &contentLength
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return nil, err
		}
		result.Date = &date
	}
	return &result, nil
}

// createHandleError handles the Create error response.
func (client directoryClient) createHandleError(resp *azcore.Response) error {
	var err DataLakeStorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Delete - Deletes the directory
func (client directoryClient) Delete(ctx context.Context, recursiveDirectoryDelete bool, directoryDeleteOptions *DirectoryDeleteOptions, leaseAccessConditions *LeaseAccessConditions, modifiedAccessConditions *ModifiedAccessConditions) (*DirectoryDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, recursiveDirectoryDelete, directoryDeleteOptions, leaseAccessConditions, modifiedAccessConditions)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.deleteHandleError(resp)
	}
	result, err := client.deleteHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// deleteCreateRequest creates the Delete request.
func (client directoryClient) deleteCreateRequest(ctx context.Context, recursiveDirectoryDelete bool, directoryDeleteOptions *DirectoryDeleteOptions, leaseAccessConditions *LeaseAccessConditions, modifiedAccessConditions *ModifiedAccessConditions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodDelete, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	if directoryDeleteOptions != nil && directoryDeleteOptions.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*directoryDeleteOptions.Timeout), 10))
	}
	query.Set("recursive", strconv.FormatBool(recursiveDirectoryDelete))
	if directoryDeleteOptions != nil && directoryDeleteOptions.Marker != nil {
		query.Set("continuation", *directoryDeleteOptions.Marker)
	}
	req.URL.RawQuery = query.Encode()
	if leaseAccessConditions != nil && leaseAccessConditions.LeaseId != nil {
		req.Header.Set("x-ms-lease-id", *leaseAccessConditions.LeaseId)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfModifiedSince != nil {
		req.Header.Set("If-Modified-Since", modifiedAccessConditions.IfModifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfUnmodifiedSince != nil {
		req.Header.Set("If-Unmodified-Since", modifiedAccessConditions.IfUnmodifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfMatch != nil {
		req.Header.Set("If-Match", *modifiedAccessConditions.IfMatch)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *modifiedAccessConditions.IfNoneMatch)
	}
	req.Header.Set("x-ms-version", "2019-07-07")
	if directoryDeleteOptions != nil && directoryDeleteOptions.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *directoryDeleteOptions.RequestId)
	}
	req.Header.Set("Accept", "application/xml")
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client directoryClient) deleteHandleResponse(resp *azcore.Response) (*DirectoryDeleteResponse, error) {
	result := DirectoryDeleteResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-continuation"); val != "" {
		result.Marker = &val
	}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return nil, err
		}
		result.Date = &date
	}
	return &result, nil
}

// deleteHandleError handles the Delete error response.
func (client directoryClient) deleteHandleError(resp *azcore.Response) error {
	var err DataLakeStorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetAccessControl - Get the owner, group, permissions, or access control list for a directory.
func (client directoryClient) GetAccessControl(ctx context.Context, directoryGetAccessControlOptions *DirectoryGetAccessControlOptions, leaseAccessConditions *LeaseAccessConditions, modifiedAccessConditions *ModifiedAccessConditions) (*DirectoryGetAccessControlResponse, error) {
	req, err := client.getAccessControlCreateRequest(ctx, directoryGetAccessControlOptions, leaseAccessConditions, modifiedAccessConditions)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getAccessControlHandleError(resp)
	}
	result, err := client.getAccessControlHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getAccessControlCreateRequest creates the GetAccessControl request.
func (client directoryClient) getAccessControlCreateRequest(ctx context.Context, directoryGetAccessControlOptions *DirectoryGetAccessControlOptions, leaseAccessConditions *LeaseAccessConditions, modifiedAccessConditions *ModifiedAccessConditions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodHead, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("action", "getAccessControl")
	if directoryGetAccessControlOptions != nil && directoryGetAccessControlOptions.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*directoryGetAccessControlOptions.Timeout), 10))
	}
	if directoryGetAccessControlOptions != nil && directoryGetAccessControlOptions.Upn != nil {
		query.Set("upn", strconv.FormatBool(*directoryGetAccessControlOptions.Upn))
	}
	req.URL.RawQuery = query.Encode()
	if leaseAccessConditions != nil && leaseAccessConditions.LeaseId != nil {
		req.Header.Set("x-ms-lease-id", *leaseAccessConditions.LeaseId)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfMatch != nil {
		req.Header.Set("If-Match", *modifiedAccessConditions.IfMatch)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *modifiedAccessConditions.IfNoneMatch)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfModifiedSince != nil {
		req.Header.Set("If-Modified-Since", modifiedAccessConditions.IfModifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfUnmodifiedSince != nil {
		req.Header.Set("If-Unmodified-Since", modifiedAccessConditions.IfUnmodifiedSince.Format(time.RFC1123))
	}
	if directoryGetAccessControlOptions != nil && directoryGetAccessControlOptions.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *directoryGetAccessControlOptions.RequestId)
	}
	req.Header.Set("x-ms-version", "2019-07-07")
	req.Header.Set("Accept", "application/xml")
	return req, nil
}

// getAccessControlHandleResponse handles the GetAccessControl response.
func (client directoryClient) getAccessControlHandleResponse(resp *azcore.Response) (*DirectoryGetAccessControlResponse, error) {
	result := DirectoryGetAccessControlResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return nil, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if val := resp.Header.Get("Last-Modified"); val != "" {
		lastModified, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return nil, err
		}
		result.LastModified = &lastModified
	}
	if val := resp.Header.Get("x-ms-owner"); val != "" {
		result.Owner = &val
	}
	if val := resp.Header.Get("x-ms-group"); val != "" {
		result.Group = &val
	}
	if val := resp.Header.Get("x-ms-permissions"); val != "" {
		result.Permissions = &val
	}
	if val := resp.Header.Get("x-ms-acl"); val != "" {
		result.ACL = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	return &result, nil
}

// getAccessControlHandleError handles the GetAccessControl error response.
func (client directoryClient) getAccessControlHandleError(resp *azcore.Response) error {
	var err DataLakeStorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Rename - Rename a directory. By default, the destination is overwritten and if the destination already exists and has a lease the lease is broken. This
// operation supports conditional HTTP requests. For more
// information, see Specifying Conditional Headers for Blob Service Operations [https://docs.microsoft.com/en-us/rest/api/storageservices/specifying-conditional-headers-for-blob-service-operations].
// To
// fail if the destination already exists, use a conditional request with If-None-Match: "*".
func (client directoryClient) Rename(ctx context.Context, renameSource string, directoryRenameOptions *DirectoryRenameOptions, directoryHttpHeaders *DirectoryHttpHeaders, leaseAccessConditions *LeaseAccessConditions, modifiedAccessConditions *ModifiedAccessConditions, sourceModifiedAccessConditions *SourceModifiedAccessConditions) (*DirectoryRenameResponse, error) {
	req, err := client.renameCreateRequest(ctx, renameSource, directoryRenameOptions, directoryHttpHeaders, leaseAccessConditions, modifiedAccessConditions, sourceModifiedAccessConditions)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusCreated) {
		return nil, client.renameHandleError(resp)
	}
	result, err := client.renameHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// renameCreateRequest creates the Rename request.
func (client directoryClient) renameCreateRequest(ctx context.Context, renameSource string, directoryRenameOptions *DirectoryRenameOptions, directoryHttpHeaders *DirectoryHttpHeaders, leaseAccessConditions *LeaseAccessConditions, modifiedAccessConditions *ModifiedAccessConditions, sourceModifiedAccessConditions *SourceModifiedAccessConditions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodPut, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	if directoryRenameOptions != nil && directoryRenameOptions.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*directoryRenameOptions.Timeout), 10))
	}
	if directoryRenameOptions != nil && directoryRenameOptions.Marker != nil {
		query.Set("continuation", *directoryRenameOptions.Marker)
	}
	if client.pathRenameMode != nil {
		query.Set("mode", string(*client.pathRenameMode))
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("x-ms-rename-source", renameSource)
	if directoryRenameOptions != nil && directoryRenameOptions.DirectoryProperties != nil {
		req.Header.Set("x-ms-properties", *directoryRenameOptions.DirectoryProperties)
	}
	if directoryRenameOptions != nil && directoryRenameOptions.PosixPermissions != nil {
		req.Header.Set("x-ms-permissions", *directoryRenameOptions.PosixPermissions)
	}
	if directoryRenameOptions != nil && directoryRenameOptions.PosixUmask != nil {
		req.Header.Set("x-ms-umask", *directoryRenameOptions.PosixUmask)
	}
	if directoryHttpHeaders != nil && directoryHttpHeaders.CacheControl != nil {
		req.Header.Set("x-ms-cache-control", *directoryHttpHeaders.CacheControl)
	}
	if directoryHttpHeaders != nil && directoryHttpHeaders.ContentType != nil {
		req.Header.Set("x-ms-content-type", *directoryHttpHeaders.ContentType)
	}
	if directoryHttpHeaders != nil && directoryHttpHeaders.ContentEncoding != nil {
		req.Header.Set("x-ms-content-encoding", *directoryHttpHeaders.ContentEncoding)
	}
	if directoryHttpHeaders != nil && directoryHttpHeaders.ContentLanguage != nil {
		req.Header.Set("x-ms-content-language", *directoryHttpHeaders.ContentLanguage)
	}
	if directoryHttpHeaders != nil && directoryHttpHeaders.ContentDisposition != nil {
		req.Header.Set("x-ms-content-disposition", *directoryHttpHeaders.ContentDisposition)
	}
	if leaseAccessConditions != nil && leaseAccessConditions.LeaseId != nil {
		req.Header.Set("x-ms-lease-id", *leaseAccessConditions.LeaseId)
	}
	if directoryRenameOptions != nil && directoryRenameOptions.SourceLeaseId != nil {
		req.Header.Set("x-ms-source-lease-id", *directoryRenameOptions.SourceLeaseId)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfModifiedSince != nil {
		req.Header.Set("If-Modified-Since", modifiedAccessConditions.IfModifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfUnmodifiedSince != nil {
		req.Header.Set("If-Unmodified-Since", modifiedAccessConditions.IfUnmodifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfMatch != nil {
		req.Header.Set("If-Match", *modifiedAccessConditions.IfMatch)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *modifiedAccessConditions.IfNoneMatch)
	}
	if sourceModifiedAccessConditions != nil && sourceModifiedAccessConditions.SourceIfModifiedSince != nil {
		req.Header.Set("x-ms-source-if-modified-since", sourceModifiedAccessConditions.SourceIfModifiedSince.Format(time.RFC1123))
	}
	if sourceModifiedAccessConditions != nil && sourceModifiedAccessConditions.SourceIfUnmodifiedSince != nil {
		req.Header.Set("x-ms-source-if-unmodified-since", sourceModifiedAccessConditions.SourceIfUnmodifiedSince.Format(time.RFC1123))
	}
	if sourceModifiedAccessConditions != nil && sourceModifiedAccessConditions.SourceIfMatch != nil {
		req.Header.Set("x-ms-source-if-match", *sourceModifiedAccessConditions.SourceIfMatch)
	}
	if sourceModifiedAccessConditions != nil && sourceModifiedAccessConditions.SourceIfNoneMatch != nil {
		req.Header.Set("x-ms-source-if-none-match", *sourceModifiedAccessConditions.SourceIfNoneMatch)
	}
	req.Header.Set("x-ms-version", "2019-07-07")
	if directoryRenameOptions != nil && directoryRenameOptions.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *directoryRenameOptions.RequestId)
	}
	req.Header.Set("Accept", "application/xml")
	return req, nil
}

// renameHandleResponse handles the Rename response.
func (client directoryClient) renameHandleResponse(resp *azcore.Response) (*DirectoryRenameResponse, error) {
	result := DirectoryRenameResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-continuation"); val != "" {
		result.Marker = &val
	}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if val := resp.Header.Get("Last-Modified"); val != "" {
		lastModified, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return nil, err
		}
		result.LastModified = &lastModified
	}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Content-Length"); val != "" {
		contentLength, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return nil, err
		}
		result.ContentLength = &contentLength
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return nil, err
		}
		result.Date = &date
	}
	return &result, nil
}

// renameHandleError handles the Rename error response.
func (client directoryClient) renameHandleError(resp *azcore.Response) error {
	var err DataLakeStorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// SetAccessControl - Set the owner, group, permissions, or access control list for a directory.
func (client directoryClient) SetAccessControl(ctx context.Context, directorySetAccessControlOptions *DirectorySetAccessControlOptions, leaseAccessConditions *LeaseAccessConditions, modifiedAccessConditions *ModifiedAccessConditions) (*DirectorySetAccessControlResponse, error) {
	req, err := client.setAccessControlCreateRequest(ctx, directorySetAccessControlOptions, leaseAccessConditions, modifiedAccessConditions)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.setAccessControlHandleError(resp)
	}
	result, err := client.setAccessControlHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// setAccessControlCreateRequest creates the SetAccessControl request.
func (client directoryClient) setAccessControlCreateRequest(ctx context.Context, directorySetAccessControlOptions *DirectorySetAccessControlOptions, leaseAccessConditions *LeaseAccessConditions, modifiedAccessConditions *ModifiedAccessConditions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodPatch, client.con.Endpoint())
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("action", "setAccessControl")
	if directorySetAccessControlOptions != nil && directorySetAccessControlOptions.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*directorySetAccessControlOptions.Timeout), 10))
	}
	req.URL.RawQuery = query.Encode()
	if leaseAccessConditions != nil && leaseAccessConditions.LeaseId != nil {
		req.Header.Set("x-ms-lease-id", *leaseAccessConditions.LeaseId)
	}
	if directorySetAccessControlOptions != nil && directorySetAccessControlOptions.Owner != nil {
		req.Header.Set("x-ms-owner", *directorySetAccessControlOptions.Owner)
	}
	if directorySetAccessControlOptions != nil && directorySetAccessControlOptions.Group != nil {
		req.Header.Set("x-ms-group", *directorySetAccessControlOptions.Group)
	}
	if directorySetAccessControlOptions != nil && directorySetAccessControlOptions.PosixPermissions != nil {
		req.Header.Set("x-ms-permissions", *directorySetAccessControlOptions.PosixPermissions)
	}
	if directorySetAccessControlOptions != nil && directorySetAccessControlOptions.PosixAcl != nil {
		req.Header.Set("x-ms-acl", *directorySetAccessControlOptions.PosixAcl)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfMatch != nil {
		req.Header.Set("If-Match", *modifiedAccessConditions.IfMatch)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *modifiedAccessConditions.IfNoneMatch)
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfModifiedSince != nil {
		req.Header.Set("If-Modified-Since", modifiedAccessConditions.IfModifiedSince.Format(time.RFC1123))
	}
	if modifiedAccessConditions != nil && modifiedAccessConditions.IfUnmodifiedSince != nil {
		req.Header.Set("If-Unmodified-Since", modifiedAccessConditions.IfUnmodifiedSince.Format(time.RFC1123))
	}
	if directorySetAccessControlOptions != nil && directorySetAccessControlOptions.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *directorySetAccessControlOptions.RequestId)
	}
	req.Header.Set("x-ms-version", "2019-07-07")
	req.Header.Set("Accept", "application/xml")
	return req, nil
}

// setAccessControlHandleResponse handles the SetAccessControl response.
func (client directoryClient) setAccessControlHandleResponse(resp *azcore.Response) (*DirectorySetAccessControlResponse, error) {
	result := DirectorySetAccessControlResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return nil, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if val := resp.Header.Get("Last-Modified"); val != "" {
		lastModified, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return nil, err
		}
		result.LastModified = &lastModified
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	return &result, nil
}

// setAccessControlHandleError handles the SetAccessControl error response.
func (client directoryClient) setAccessControlHandleError(resp *azcore.Response) error {
	var err DataLakeStorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
