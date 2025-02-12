//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package mediatypesgroup

import "io"

// MediaTypesClientAnalyzeBodyNoAcceptHeaderOptions contains the optional parameters for the MediaTypesClient.AnalyzeBodyNoAcceptHeader method.
type MediaTypesClientAnalyzeBodyNoAcceptHeaderOptions struct {
	// Input parameter.
	Input io.ReadSeekCloser
}

// MediaTypesClientAnalyzeBodyNoAcceptHeaderWithSourcePathOptions contains the optional parameters for the MediaTypesClient.AnalyzeBodyNoAcceptHeaderWithSourcePath
// method.
type MediaTypesClientAnalyzeBodyNoAcceptHeaderWithSourcePathOptions struct {
	// Input parameter.
	Input *SourcePath
}

// MediaTypesClientAnalyzeBodyOptions contains the optional parameters for the MediaTypesClient.AnalyzeBody method.
type MediaTypesClientAnalyzeBodyOptions struct {
	// Input parameter.
	Input io.ReadSeekCloser
}

// MediaTypesClientAnalyzeBodyWithSourcePathOptions contains the optional parameters for the MediaTypesClient.AnalyzeBodyWithSourcePath method.
type MediaTypesClientAnalyzeBodyWithSourcePathOptions struct {
	// Input parameter.
	Input *SourcePath
}

// MediaTypesClientContentTypeWithEncodingOptions contains the optional parameters for the MediaTypesClient.ContentTypeWithEncoding method.
type MediaTypesClientContentTypeWithEncodingOptions struct {
	// Input parameter.
	Input *string
}

// SourcePath - Uri or local path to source data.
type SourcePath struct {
	// File source path.
	Source *string `json:"source,omitempty"`
}
