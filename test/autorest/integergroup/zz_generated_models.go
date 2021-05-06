// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package integergroup

import (
	"net/http"
	"time"
)

// Implements the error and azcore.HTTPResponse interfaces.
type Error struct {
	raw     string
	Message *string `json:"message,omitempty"`
	Status  *int32  `json:"status,omitempty"`
}

// Error implements the error interface for type Error.
// The contents of the error text are not contractual and subject to change.
func (e Error) Error() string {
	return e.raw
}

// Int32Response is the response envelope for operations that return a int32 type.
type Int32Response struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
	Value       *int32
}

// Int64Response is the response envelope for operations that return a int64 type.
type Int64Response struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
	Value       *int64
}

// IntGetInvalidOptions contains the optional parameters for the Int.GetInvalid method.
type IntGetInvalidOptions struct {
	// placeholder for future optional parameters
}

// IntGetInvalidUnixTimeOptions contains the optional parameters for the Int.GetInvalidUnixTime method.
type IntGetInvalidUnixTimeOptions struct {
	// placeholder for future optional parameters
}

// IntGetNullOptions contains the optional parameters for the Int.GetNull method.
type IntGetNullOptions struct {
	// placeholder for future optional parameters
}

// IntGetNullUnixTimeOptions contains the optional parameters for the Int.GetNullUnixTime method.
type IntGetNullUnixTimeOptions struct {
	// placeholder for future optional parameters
}

// IntGetOverflowInt32Options contains the optional parameters for the Int.GetOverflowInt32 method.
type IntGetOverflowInt32Options struct {
	// placeholder for future optional parameters
}

// IntGetOverflowInt64Options contains the optional parameters for the Int.GetOverflowInt64 method.
type IntGetOverflowInt64Options struct {
	// placeholder for future optional parameters
}

// IntGetUnderflowInt32Options contains the optional parameters for the Int.GetUnderflowInt32 method.
type IntGetUnderflowInt32Options struct {
	// placeholder for future optional parameters
}

// IntGetUnderflowInt64Options contains the optional parameters for the Int.GetUnderflowInt64 method.
type IntGetUnderflowInt64Options struct {
	// placeholder for future optional parameters
}

// IntGetUnixTimeOptions contains the optional parameters for the Int.GetUnixTime method.
type IntGetUnixTimeOptions struct {
	// placeholder for future optional parameters
}

// IntPutMax32Options contains the optional parameters for the Int.PutMax32 method.
type IntPutMax32Options struct {
	// placeholder for future optional parameters
}

// IntPutMax64Options contains the optional parameters for the Int.PutMax64 method.
type IntPutMax64Options struct {
	// placeholder for future optional parameters
}

// IntPutMin32Options contains the optional parameters for the Int.PutMin32 method.
type IntPutMin32Options struct {
	// placeholder for future optional parameters
}

// IntPutMin64Options contains the optional parameters for the Int.PutMin64 method.
type IntPutMin64Options struct {
	// placeholder for future optional parameters
}

// IntPutUnixTimeDateOptions contains the optional parameters for the Int.PutUnixTimeDate method.
type IntPutUnixTimeDateOptions struct {
	// placeholder for future optional parameters
}

// TimeResponse is the response envelope for operations that return a time.Time type.
type TimeResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// date in seconds since 1970-01-01T00:00:00Z.
	Value *time.Time
}
