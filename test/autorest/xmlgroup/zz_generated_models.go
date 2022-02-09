//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package xmlgroup

import "time"

// AccessPolicy - An Access policy
type AccessPolicy struct {
	// REQUIRED; the date-time the policy expires
	Expiry *time.Time `xml:"Expiry"`

	// REQUIRED; the permissions for the acl policy
	Permission *string `xml:"Permission"`

	// REQUIRED; the date-time the policy is active
	Start *time.Time `xml:"Start"`
}

// AppleBarrel - A barrel of apples.
type AppleBarrel struct {
	BadApples  []*string `xml:"BadApples>Apple"`
	GoodApples []*string `xml:"GoodApples>Apple"`
}

// Banana - A banana.
type Banana struct {
	// The time at which you should reconsider eating this banana
	Expiration *time.Time `xml:"expiration"`
	Flavor     *string    `xml:"flavor"`
	Name       *string    `xml:"name"`
}

// Blob - An Azure Storage blob
type Blob struct {
	// REQUIRED
	Deleted *bool `xml:"Deleted"`

	// REQUIRED
	Name *string `xml:"Name"`

	// REQUIRED; Properties of a blob
	Properties *BlobProperties `xml:"Properties"`

	// REQUIRED
	Snapshot *string `xml:"Snapshot"`

	// Dictionary of
	Metadata map[string]*string `xml:"Metadata"`
}

type BlobPrefix struct {
	// REQUIRED
	Name *string `xml:"Name"`
}

// BlobProperties - Properties of a blob
type BlobProperties struct {
	// REQUIRED
	Etag *string `xml:"Etag"`

	// REQUIRED
	LastModified       *time.Time     `xml:"Last-Modified"`
	AccessTier         *AccessTier    `xml:"AccessTier"`
	AccessTierInferred *bool          `xml:"AccessTierInferred"`
	ArchiveStatus      *ArchiveStatus `xml:"ArchiveStatus"`
	BlobSequenceNumber *int32         `xml:"x-ms-blob-sequence-number"`
	BlobType           *BlobType      `xml:"BlobType"`
	CacheControl       *string        `xml:"Cache-Control"`
	ContentDisposition *string        `xml:"Content-Disposition"`
	ContentEncoding    *string        `xml:"Content-Encoding"`
	ContentLanguage    *string        `xml:"Content-Language"`

	// Size in bytes
	ContentLength          *int64             `xml:"Content-Length"`
	ContentMD5             *string            `xml:"Content-MD5"`
	ContentType            *string            `xml:"Content-Type"`
	CopyCompletionTime     *time.Time         `xml:"CopyCompletionTime"`
	CopyID                 *string            `xml:"CopyId"`
	CopyProgress           *string            `xml:"CopyProgress"`
	CopySource             *string            `xml:"CopySource"`
	CopyStatus             *CopyStatusType    `xml:"CopyStatus"`
	CopyStatusDescription  *string            `xml:"CopyStatusDescription"`
	DeletedTime            *time.Time         `xml:"DeletedTime"`
	DestinationSnapshot    *string            `xml:"DestinationSnapshot"`
	IncrementalCopy        *bool              `xml:"IncrementalCopy"`
	LeaseDuration          *LeaseDurationType `xml:"LeaseDuration"`
	LeaseState             *LeaseStateType    `xml:"LeaseState"`
	LeaseStatus            *LeaseStatusType   `xml:"LeaseStatus"`
	RemainingRetentionDays *int32             `xml:"RemainingRetentionDays"`
	ServerEncrypted        *bool              `xml:"ServerEncrypted"`
}

type Blobs struct {
	Blob       []*Blob       `xml:"Blob"`
	BlobPrefix []*BlobPrefix `xml:"BlobPrefix"`
}

// ComplexTypeNoMeta - I am a complex type with no XML node
type ComplexTypeNoMeta struct {
	// The id of the res
	ID *string `xml:"ID"`
}

// ComplexTypeWithMeta - I am a complex type with XML node
type ComplexTypeWithMeta struct {
	// The id of the res
	ID *string `xml:"ID"`
}

// Container - An Azure Storage container
type Container struct {
	// REQUIRED
	Name *string `xml:"Name"`

	// REQUIRED; Properties of a container
	Properties *ContainerProperties `xml:"Properties"`

	// Dictionary of
	Metadata map[string]*string `xml:"Metadata"`
}

// ContainerProperties - Properties of a container
type ContainerProperties struct {
	// REQUIRED
	Etag *string `xml:"Etag"`

	// REQUIRED
	LastModified  *time.Time         `xml:"Last-Modified"`
	LeaseDuration *LeaseDurationType `xml:"LeaseDuration"`
	LeaseState    *LeaseStateType    `xml:"LeaseState"`
	LeaseStatus   *LeaseStatusType   `xml:"LeaseStatus"`
	PublicAccess  *PublicAccessType  `xml:"PublicAccess"`
}

// CorsRule - CORS is an HTTP feature that enables a web application running under one domain to access resources in another
// domain. Web browsers implement a security restriction known as same-origin policy that
// prevents a web page from calling APIs in a different domain; CORS provides a secure way to allow one domain (the origin
// domain) to call APIs in another domain
type CorsRule struct {
	// REQUIRED; the request headers that the origin domain may specify on the CORS request.
	AllowedHeaders *string `xml:"AllowedHeaders"`

	// REQUIRED; The methods (HTTP request verbs) that the origin domain may use for a CORS request. (comma separated)
	AllowedMethods *string `xml:"AllowedMethods"`

	// REQUIRED; The origin domains that are permitted to make a request against the storage service via CORS. The origin domain
	// is the domain from which the request originates. Note that the origin must be an exact
	// case-sensitive match with the origin that the user age sends to the service. You can also use the wildcard character '*'
	// to allow all origin domains to make requests via CORS.
	AllowedOrigins *string `xml:"AllowedOrigins"`

	// REQUIRED; The response headers that may be sent in the response to the CORS request and exposed by the browser to the request
	// issuer
	ExposedHeaders *string `xml:"ExposedHeaders"`

	// REQUIRED; The maximum amount time that a browser should cache the preflight OPTIONS request.
	MaxAgeInSeconds *int32 `xml:"MaxAgeInSeconds"`
}

type Error struct {
	Message *string `json:"message,omitempty"`
	Status  *int32  `json:"status,omitempty"`
}

type JSONInput struct {
	ID *int32 `json:"id,omitempty"`
}

type JSONOutput struct {
	ID *int32 `json:"id,omitempty"`
}

// ListBlobsResponse - An enumeration of blobs
type ListBlobsResponse struct {
	// REQUIRED
	Blobs *Blobs `xml:"Blobs"`

	// REQUIRED
	ContainerName *string `xml:"ContainerName,attr"`

	// REQUIRED
	Delimiter *string `xml:"Delimiter"`

	// REQUIRED
	Marker *string `xml:"Marker"`

	// REQUIRED
	MaxResults *int32 `xml:"MaxResults"`

	// REQUIRED
	NextMarker *string `xml:"NextMarker"`

	// REQUIRED
	Prefix          *string `xml:"Prefix"`
	ServiceEndpoint *string `xml:"ServiceEndpoint,attr"`
}

// ListContainersResponse - An enumeration of containers
type ListContainersResponse struct {
	// REQUIRED
	MaxResults *int32 `xml:"MaxResults"`

	// REQUIRED
	NextMarker *string `xml:"NextMarker"`

	// REQUIRED
	Prefix *string `xml:"Prefix"`

	// REQUIRED
	ServiceEndpoint *string      `xml:"ServiceEndpoint,attr"`
	Containers      []*Container `xml:"Containers>Container"`
	Marker          *string      `xml:"Marker"`
}

// Logging - Azure Analytics Logging settings.
type Logging struct {
	// REQUIRED; Indicates whether all delete requests should be logged.
	Delete *bool `xml:"Delete"`

	// REQUIRED; Indicates whether all read requests should be logged.
	Read *bool `xml:"Read"`

	// REQUIRED; the retention policy
	RetentionPolicy *RetentionPolicy `xml:"RetentionPolicy"`

	// REQUIRED; The version of Storage Analytics to configure.
	Version *string `xml:"Version"`

	// REQUIRED; Indicates whether all write requests should be logged.
	Write *bool `xml:"Write"`
}

type Metrics struct {
	// REQUIRED; Indicates whether metrics are enabled for the Blob service.
	Enabled *bool `xml:"Enabled"`

	// Indicates whether metrics should generate summary statistics for called API operations.
	IncludeAPIs *bool `xml:"IncludeAPIs"`

	// the retention policy
	RetentionPolicy *RetentionPolicy `xml:"RetentionPolicy"`

	// The version of Storage Analytics to configure.
	Version *string `xml:"Version"`
}

type ModelWithByteProperty struct {
	Bytes []byte `xml:"Bytes"`
}

type ModelWithURLProperty struct {
	URL *string `xml:"Url"`
}

// ObjectWithXMsTextProperty - Contans property
type ObjectWithXMsTextProperty struct {
	// Returned value should be 'I am text'
	Content *string `xml:"content"`

	// Returned value should be 'english'
	Language *string `xml:"language,attr"`
}

// RetentionPolicy - the retention policy
type RetentionPolicy struct {
	// REQUIRED; Indicates whether a retention policy is enabled for the storage service
	Enabled *bool `xml:"Enabled"`

	// Indicates the number of days that metrics or logging or soft-deleted data should be retained. All data older than this
	// value will be deleted
	Days *int32 `xml:"Days"`
}

// RootWithRefAndMeta - I am root, and I ref a model WITH meta
type RootWithRefAndMeta struct {
	// XML will use XMLComplexTypeWithMeta
	RefToModel *ComplexTypeWithMeta `xml:"XMLComplexTypeWithMeta"`

	// Something else (just to avoid flattening)
	Something *string `xml:"Something"`
}

// RootWithRefAndNoMeta - I am root, and I ref a model with no meta
type RootWithRefAndNoMeta struct {
	// XML will use RefToModel
	RefToModel *ComplexTypeNoMeta `xml:"RefToModel"`

	// Something else (just to avoid flattening)
	Something *string `xml:"Something"`
}

// SignedIdentifier - signed identifier
type SignedIdentifier struct {
	// REQUIRED; The access policy
	AccessPolicy *AccessPolicy `xml:"AccessPolicy"`

	// REQUIRED; a unique id
	ID *string `xml:"Id"`
}

// Slide - A slide in a slideshow
type Slide struct {
	Items []*string `xml:"item"`
	Title *string   `xml:"title"`
	Type  *string   `xml:"type,attr"`
}

// Slideshow - Data about a slideshow
type Slideshow struct {
	Author *string  `xml:"author,attr"`
	Date   *string  `xml:"date,attr"`
	Slides []*Slide `xml:"slide"`
	Title  *string  `xml:"title,attr"`
}

// StorageServiceProperties - Storage Service Properties.
type StorageServiceProperties struct {
	// The set of CORS rules.
	Cors []*CorsRule `xml:"Cors>CorsRule"`

	// The default version to use for requests to the Blob service if an incoming request's version is not specified. Possible
	// values include version 2008-10-27 and all more recent versions
	DefaultServiceVersion *string `xml:"DefaultServiceVersion"`

	// The Delete Retention Policy for the service
	DeleteRetentionPolicy *RetentionPolicy `xml:"DeleteRetentionPolicy"`

	// A summary of request statistics grouped by API in hourly aggregates for blobs
	HourMetrics *Metrics `xml:"HourMetrics"`

	// Azure Analytics Logging settings
	Logging *Logging `xml:"Logging"`

	// a summary of request statistics grouped by API in minute aggregates for blobs
	MinuteMetrics *Metrics `xml:"MinuteMetrics"`
}

// XMLClientGetACLsOptions contains the optional parameters for the XMLClient.GetACLs method.
type XMLClientGetACLsOptions struct {
	// placeholder for future optional parameters
}

// XMLClientGetBytesOptions contains the optional parameters for the XMLClient.GetBytes method.
type XMLClientGetBytesOptions struct {
	// placeholder for future optional parameters
}

// XMLClientGetComplexTypeRefNoMetaOptions contains the optional parameters for the XMLClient.GetComplexTypeRefNoMeta method.
type XMLClientGetComplexTypeRefNoMetaOptions struct {
	// placeholder for future optional parameters
}

// XMLClientGetComplexTypeRefWithMetaOptions contains the optional parameters for the XMLClient.GetComplexTypeRefWithMeta
// method.
type XMLClientGetComplexTypeRefWithMetaOptions struct {
	// placeholder for future optional parameters
}

// XMLClientGetEmptyChildElementOptions contains the optional parameters for the XMLClient.GetEmptyChildElement method.
type XMLClientGetEmptyChildElementOptions struct {
	// placeholder for future optional parameters
}

// XMLClientGetEmptyListOptions contains the optional parameters for the XMLClient.GetEmptyList method.
type XMLClientGetEmptyListOptions struct {
	// placeholder for future optional parameters
}

// XMLClientGetEmptyRootListOptions contains the optional parameters for the XMLClient.GetEmptyRootList method.
type XMLClientGetEmptyRootListOptions struct {
	// placeholder for future optional parameters
}

// XMLClientGetEmptyWrappedListsOptions contains the optional parameters for the XMLClient.GetEmptyWrappedLists method.
type XMLClientGetEmptyWrappedListsOptions struct {
	// placeholder for future optional parameters
}

// XMLClientGetHeadersOptions contains the optional parameters for the XMLClient.GetHeaders method.
type XMLClientGetHeadersOptions struct {
	// placeholder for future optional parameters
}

// XMLClientGetRootListOptions contains the optional parameters for the XMLClient.GetRootList method.
type XMLClientGetRootListOptions struct {
	// placeholder for future optional parameters
}

// XMLClientGetRootListSingleItemOptions contains the optional parameters for the XMLClient.GetRootListSingleItem method.
type XMLClientGetRootListSingleItemOptions struct {
	// placeholder for future optional parameters
}

// XMLClientGetServicePropertiesOptions contains the optional parameters for the XMLClient.GetServiceProperties method.
type XMLClientGetServicePropertiesOptions struct {
	// placeholder for future optional parameters
}

// XMLClientGetSimpleOptions contains the optional parameters for the XMLClient.GetSimple method.
type XMLClientGetSimpleOptions struct {
	// placeholder for future optional parameters
}

// XMLClientGetURIOptions contains the optional parameters for the XMLClient.GetURI method.
type XMLClientGetURIOptions struct {
	// placeholder for future optional parameters
}

// XMLClientGetWrappedListsOptions contains the optional parameters for the XMLClient.GetWrappedLists method.
type XMLClientGetWrappedListsOptions struct {
	// placeholder for future optional parameters
}

// XMLClientGetXMsTextOptions contains the optional parameters for the XMLClient.GetXMsText method.
type XMLClientGetXMsTextOptions struct {
	// placeholder for future optional parameters
}

// XMLClientJSONInputOptions contains the optional parameters for the XMLClient.JSONInput method.
type XMLClientJSONInputOptions struct {
	// placeholder for future optional parameters
}

// XMLClientJSONOutputOptions contains the optional parameters for the XMLClient.JSONOutput method.
type XMLClientJSONOutputOptions struct {
	// placeholder for future optional parameters
}

// XMLClientListBlobsOptions contains the optional parameters for the XMLClient.ListBlobs method.
type XMLClientListBlobsOptions struct {
	// placeholder for future optional parameters
}

// XMLClientListContainersOptions contains the optional parameters for the XMLClient.ListContainers method.
type XMLClientListContainersOptions struct {
	// placeholder for future optional parameters
}

// XMLClientPutACLsOptions contains the optional parameters for the XMLClient.PutACLs method.
type XMLClientPutACLsOptions struct {
	// placeholder for future optional parameters
}

// XMLClientPutBinaryOptions contains the optional parameters for the XMLClient.PutBinary method.
type XMLClientPutBinaryOptions struct {
	// placeholder for future optional parameters
}

// XMLClientPutComplexTypeRefNoMetaOptions contains the optional parameters for the XMLClient.PutComplexTypeRefNoMeta method.
type XMLClientPutComplexTypeRefNoMetaOptions struct {
	// placeholder for future optional parameters
}

// XMLClientPutComplexTypeRefWithMetaOptions contains the optional parameters for the XMLClient.PutComplexTypeRefWithMeta
// method.
type XMLClientPutComplexTypeRefWithMetaOptions struct {
	// placeholder for future optional parameters
}

// XMLClientPutEmptyChildElementOptions contains the optional parameters for the XMLClient.PutEmptyChildElement method.
type XMLClientPutEmptyChildElementOptions struct {
	// placeholder for future optional parameters
}

// XMLClientPutEmptyListOptions contains the optional parameters for the XMLClient.PutEmptyList method.
type XMLClientPutEmptyListOptions struct {
	// placeholder for future optional parameters
}

// XMLClientPutEmptyRootListOptions contains the optional parameters for the XMLClient.PutEmptyRootList method.
type XMLClientPutEmptyRootListOptions struct {
	// placeholder for future optional parameters
}

// XMLClientPutEmptyWrappedListsOptions contains the optional parameters for the XMLClient.PutEmptyWrappedLists method.
type XMLClientPutEmptyWrappedListsOptions struct {
	// placeholder for future optional parameters
}

// XMLClientPutRootListOptions contains the optional parameters for the XMLClient.PutRootList method.
type XMLClientPutRootListOptions struct {
	// placeholder for future optional parameters
}

// XMLClientPutRootListSingleItemOptions contains the optional parameters for the XMLClient.PutRootListSingleItem method.
type XMLClientPutRootListSingleItemOptions struct {
	// placeholder for future optional parameters
}

// XMLClientPutServicePropertiesOptions contains the optional parameters for the XMLClient.PutServiceProperties method.
type XMLClientPutServicePropertiesOptions struct {
	// placeholder for future optional parameters
}

// XMLClientPutSimpleOptions contains the optional parameters for the XMLClient.PutSimple method.
type XMLClientPutSimpleOptions struct {
	// placeholder for future optional parameters
}

// XMLClientPutURIOptions contains the optional parameters for the XMLClient.PutURI method.
type XMLClientPutURIOptions struct {
	// placeholder for future optional parameters
}

// XMLClientPutWrappedListsOptions contains the optional parameters for the XMLClient.PutWrappedLists method.
type XMLClientPutWrappedListsOptions struct {
	// placeholder for future optional parameters
}
