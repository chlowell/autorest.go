//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package xmlgroup

import "net/http"

// XMLGetACLsResponse contains the response from method XML.GetACLs.
type XMLGetACLsResponse struct {
	XMLGetACLsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLGetACLsResult contains the result from method XML.GetACLs.
type XMLGetACLsResult struct {
	// a collection of signed identifiers
	SignedIdentifiers []*SignedIdentifier `xml:"SignedIdentifier"`
}

// XMLGetBytesResponse contains the response from method XML.GetBytes.
type XMLGetBytesResponse struct {
	XMLGetBytesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLGetBytesResult contains the result from method XML.GetBytes.
type XMLGetBytesResult struct {
	ModelWithByteProperty
}

// XMLGetComplexTypeRefNoMetaResponse contains the response from method XML.GetComplexTypeRefNoMeta.
type XMLGetComplexTypeRefNoMetaResponse struct {
	XMLGetComplexTypeRefNoMetaResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLGetComplexTypeRefNoMetaResult contains the result from method XML.GetComplexTypeRefNoMeta.
type XMLGetComplexTypeRefNoMetaResult struct {
	RootWithRefAndNoMeta
}

// XMLGetComplexTypeRefWithMetaResponse contains the response from method XML.GetComplexTypeRefWithMeta.
type XMLGetComplexTypeRefWithMetaResponse struct {
	XMLGetComplexTypeRefWithMetaResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLGetComplexTypeRefWithMetaResult contains the result from method XML.GetComplexTypeRefWithMeta.
type XMLGetComplexTypeRefWithMetaResult struct {
	RootWithRefAndMeta
}

// XMLGetEmptyChildElementResponse contains the response from method XML.GetEmptyChildElement.
type XMLGetEmptyChildElementResponse struct {
	XMLGetEmptyChildElementResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLGetEmptyChildElementResult contains the result from method XML.GetEmptyChildElement.
type XMLGetEmptyChildElementResult struct {
	Banana
}

// XMLGetEmptyListResponse contains the response from method XML.GetEmptyList.
type XMLGetEmptyListResponse struct {
	XMLGetEmptyListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLGetEmptyListResult contains the result from method XML.GetEmptyList.
type XMLGetEmptyListResult struct {
	Slideshow
}

// XMLGetEmptyRootListResponse contains the response from method XML.GetEmptyRootList.
type XMLGetEmptyRootListResponse struct {
	XMLGetEmptyRootListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLGetEmptyRootListResult contains the result from method XML.GetEmptyRootList.
type XMLGetEmptyRootListResult struct {
	// Array of Banana
	Bananas []*Banana `xml:"banana"`
}

// XMLGetEmptyWrappedListsResponse contains the response from method XML.GetEmptyWrappedLists.
type XMLGetEmptyWrappedListsResponse struct {
	XMLGetEmptyWrappedListsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLGetEmptyWrappedListsResult contains the result from method XML.GetEmptyWrappedLists.
type XMLGetEmptyWrappedListsResult struct {
	AppleBarrel
}

// XMLGetHeadersResponse contains the response from method XML.GetHeaders.
type XMLGetHeadersResponse struct {
	XMLGetHeadersResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLGetHeadersResult contains the result from method XML.GetHeaders.
type XMLGetHeadersResult struct {
	// CustomHeader contains the information returned from the Custom-Header header response.
	CustomHeader *string
}

// XMLGetRootListResponse contains the response from method XML.GetRootList.
type XMLGetRootListResponse struct {
	XMLGetRootListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLGetRootListResult contains the result from method XML.GetRootList.
type XMLGetRootListResult struct {
	// Array of Banana
	Bananas []*Banana `xml:"banana"`
}

// XMLGetRootListSingleItemResponse contains the response from method XML.GetRootListSingleItem.
type XMLGetRootListSingleItemResponse struct {
	XMLGetRootListSingleItemResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLGetRootListSingleItemResult contains the result from method XML.GetRootListSingleItem.
type XMLGetRootListSingleItemResult struct {
	// Array of Banana
	Bananas []*Banana `xml:"banana"`
}

// XMLGetServicePropertiesResponse contains the response from method XML.GetServiceProperties.
type XMLGetServicePropertiesResponse struct {
	XMLGetServicePropertiesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLGetServicePropertiesResult contains the result from method XML.GetServiceProperties.
type XMLGetServicePropertiesResult struct {
	StorageServiceProperties
}

// XMLGetSimpleResponse contains the response from method XML.GetSimple.
type XMLGetSimpleResponse struct {
	XMLGetSimpleResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLGetSimpleResult contains the result from method XML.GetSimple.
type XMLGetSimpleResult struct {
	Slideshow
}

// XMLGetURIResponse contains the response from method XML.GetURI.
type XMLGetURIResponse struct {
	XMLGetURIResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLGetURIResult contains the result from method XML.GetURI.
type XMLGetURIResult struct {
	ModelWithURLProperty
}

// XMLGetWrappedListsResponse contains the response from method XML.GetWrappedLists.
type XMLGetWrappedListsResponse struct {
	XMLGetWrappedListsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLGetWrappedListsResult contains the result from method XML.GetWrappedLists.
type XMLGetWrappedListsResult struct {
	AppleBarrel
}

// XMLGetXMsTextResponse contains the response from method XML.GetXMsText.
type XMLGetXMsTextResponse struct {
	XMLGetXMsTextResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLGetXMsTextResult contains the result from method XML.GetXMsText.
type XMLGetXMsTextResult struct {
	ObjectWithXMsTextProperty
}

// XMLJSONInputResponse contains the response from method XML.JSONInput.
type XMLJSONInputResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLJSONOutputResponse contains the response from method XML.JSONOutput.
type XMLJSONOutputResponse struct {
	XMLJSONOutputResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLJSONOutputResult contains the result from method XML.JSONOutput.
type XMLJSONOutputResult struct {
	JSONOutput
}

// XMLListBlobsResponse contains the response from method XML.ListBlobs.
type XMLListBlobsResponse struct {
	XMLListBlobsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLListBlobsResult contains the result from method XML.ListBlobs.
type XMLListBlobsResult struct {
	ListBlobsResponse
}

// XMLListContainersResponse contains the response from method XML.ListContainers.
type XMLListContainersResponse struct {
	XMLListContainersResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLListContainersResult contains the result from method XML.ListContainers.
type XMLListContainersResult struct {
	ListContainersResponse
}

// XMLPutACLsResponse contains the response from method XML.PutACLs.
type XMLPutACLsResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLPutBinaryResponse contains the response from method XML.PutBinary.
type XMLPutBinaryResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLPutComplexTypeRefNoMetaResponse contains the response from method XML.PutComplexTypeRefNoMeta.
type XMLPutComplexTypeRefNoMetaResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLPutComplexTypeRefWithMetaResponse contains the response from method XML.PutComplexTypeRefWithMeta.
type XMLPutComplexTypeRefWithMetaResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLPutEmptyChildElementResponse contains the response from method XML.PutEmptyChildElement.
type XMLPutEmptyChildElementResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLPutEmptyListResponse contains the response from method XML.PutEmptyList.
type XMLPutEmptyListResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLPutEmptyRootListResponse contains the response from method XML.PutEmptyRootList.
type XMLPutEmptyRootListResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLPutEmptyWrappedListsResponse contains the response from method XML.PutEmptyWrappedLists.
type XMLPutEmptyWrappedListsResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLPutRootListResponse contains the response from method XML.PutRootList.
type XMLPutRootListResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLPutRootListSingleItemResponse contains the response from method XML.PutRootListSingleItem.
type XMLPutRootListSingleItemResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLPutServicePropertiesResponse contains the response from method XML.PutServiceProperties.
type XMLPutServicePropertiesResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLPutSimpleResponse contains the response from method XML.PutSimple.
type XMLPutSimpleResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLPutURIResponse contains the response from method XML.PutURI.
type XMLPutURIResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// XMLPutWrappedListsResponse contains the response from method XML.PutWrappedLists.
type XMLPutWrappedListsResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}
