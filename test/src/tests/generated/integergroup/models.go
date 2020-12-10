package integergroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
)

// The package's fully qualified name.
const fqdn = "tests/generated/integergroup"

// Error ...
type Error struct {
	Status  *int32  `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
}

// Int32 ...
type Int32 struct {
	autorest.Response `json:"-"`
	Value             *int32 `json:"value,omitempty"`
}

// Int64 ...
type Int64 struct {
	autorest.Response `json:"-"`
	Value             *int64 `json:"value,omitempty"`
}

// UnixTime ...
type UnixTime struct {
	autorest.Response `json:"-"`
	Value             *date.UnixTime `json:"value,omitempty"`
}