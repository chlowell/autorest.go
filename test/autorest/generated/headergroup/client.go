// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package headergroup

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/url"
	"strings"
)

const telemetryInfo = "azsdk-go-headergroup/<version>"

// ClientOptions contains configuration settings for the default client's pipeline.
type ClientOptions struct {
	// HTTPClient sets the transport for making HTTP requests.
	HTTPClient azcore.Transport
	// LogOptions configures the built-in request logging policy behavior.
	LogOptions azcore.RequestLogOptions
	// Retry configures the built-in retry policy behavior.
	Retry azcore.RetryOptions
	// Telemetry configures the built-in telemetry policy behavior.
	Telemetry azcore.TelemetryOptions
	// ApplicationID is an application-specific identification string used in telemetry.
	// It has a maximum length of 24 characters and must not contain any spaces.
	ApplicationID string
}

// DefaultClientOptions creates a ClientOptions type initialized with default values.
func DefaultClientOptions() ClientOptions {
	return ClientOptions{
		HTTPClient: azcore.DefaultHTTPClientTransport(),
		Retry:      azcore.DefaultRetryOptions(),
	}
}

func (c *ClientOptions) telemetryOptions() azcore.TelemetryOptions {
	t := telemetryInfo
	if c.ApplicationID != "" {
		a := strings.ReplaceAll(c.ApplicationID, " ", "/")
		if len(a) > 24 {
			a = a[:24]
		}
		t = fmt.Sprintf("%s %s", a, telemetryInfo)
	}
	if c.Telemetry.Value == "" {
		return azcore.TelemetryOptions{Value: t}
	}
	return azcore.TelemetryOptions{Value: fmt.Sprintf("%s %s", c.Telemetry.Value, t)}
}

// Client - Test Infrastructure for AutoRest
type Client struct {
	u url.URL
	p azcore.Pipeline
}

// DefaultEndpoint is the default service endpoint.
const DefaultEndpoint = "http://localhost:3000"

// NewDefaultClient creates an instance of the Client type using the DefaultEndpoint.
func NewDefaultClient(options *ClientOptions) (*Client, error) {
	return NewClient(DefaultEndpoint, options)
}

// NewClient creates an instance of the Client type with the specified endpoint.
func NewClient(endpoint string, options *ClientOptions) (*Client, error) {
	if options == nil {
		o := DefaultClientOptions()
		options = &o
	}
	p := azcore.NewPipeline(options.HTTPClient,
		azcore.NewTelemetryPolicy(options.telemetryOptions()),
		azcore.NewUniqueRequestIDPolicy(),
		azcore.NewRetryPolicy(&options.Retry),
		azcore.NewRequestLogPolicy(options.LogOptions))
	return NewClientWithPipeline(endpoint, p)
}

// NewClientWithPipeline creates an instance of the Client type with the specified endpoint and pipeline.
func NewClientWithPipeline(endpoint string, p azcore.Pipeline) (*Client, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	if u.Scheme == "" {
		return nil, fmt.Errorf("no scheme detected in endpoint %s", endpoint)
	}
	return &Client{u: *u, p: p}, nil
}

// HeaderOperations returns the HeaderOperations associated with this client.
func (client *Client) HeaderOperations() HeaderOperations {
	return &headerOperations{Client: client}
}
