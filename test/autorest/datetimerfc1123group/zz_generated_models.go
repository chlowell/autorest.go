//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package datetimerfc1123group

// Datetimerfc1123GetInvalidOptions contains the optional parameters for the Datetimerfc1123.GetInvalid method.
type Datetimerfc1123GetInvalidOptions struct {
	// placeholder for future optional parameters
}

// Datetimerfc1123GetNullOptions contains the optional parameters for the Datetimerfc1123.GetNull method.
type Datetimerfc1123GetNullOptions struct {
	// placeholder for future optional parameters
}

// Datetimerfc1123GetOverflowOptions contains the optional parameters for the Datetimerfc1123.GetOverflow method.
type Datetimerfc1123GetOverflowOptions struct {
	// placeholder for future optional parameters
}

// Datetimerfc1123GetUTCLowercaseMaxDateTimeOptions contains the optional parameters for the Datetimerfc1123.GetUTCLowercaseMaxDateTime method.
type Datetimerfc1123GetUTCLowercaseMaxDateTimeOptions struct {
	// placeholder for future optional parameters
}

// Datetimerfc1123GetUTCMinDateTimeOptions contains the optional parameters for the Datetimerfc1123.GetUTCMinDateTime method.
type Datetimerfc1123GetUTCMinDateTimeOptions struct {
	// placeholder for future optional parameters
}

// Datetimerfc1123GetUTCUppercaseMaxDateTimeOptions contains the optional parameters for the Datetimerfc1123.GetUTCUppercaseMaxDateTime method.
type Datetimerfc1123GetUTCUppercaseMaxDateTimeOptions struct {
	// placeholder for future optional parameters
}

// Datetimerfc1123GetUnderflowOptions contains the optional parameters for the Datetimerfc1123.GetUnderflow method.
type Datetimerfc1123GetUnderflowOptions struct {
	// placeholder for future optional parameters
}

// Datetimerfc1123PutUTCMaxDateTimeOptions contains the optional parameters for the Datetimerfc1123.PutUTCMaxDateTime method.
type Datetimerfc1123PutUTCMaxDateTimeOptions struct {
	// placeholder for future optional parameters
}

// Datetimerfc1123PutUTCMinDateTimeOptions contains the optional parameters for the Datetimerfc1123.PutUTCMinDateTime method.
type Datetimerfc1123PutUTCMinDateTimeOptions struct {
	// placeholder for future optional parameters
}

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
