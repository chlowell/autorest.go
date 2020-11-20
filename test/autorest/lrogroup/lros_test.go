// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package lrogroup

import (
	"context"
	"errors"
	"generatortests/helpers"
	"net/http"
	"net/http/cookiejar"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/to"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

func newLROSClient() LrOSClient {
	options := DefaultConnectionOptions()
	options.Retry.RetryDelay = 10 * time.Millisecond
	options.HTTPClient = httpClientWithCookieJar()
	return NewLrOSClient(NewDefaultConnection(&options))
}

func httpClientWithCookieJar() azcore.Transport {
	j, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	http.DefaultClient.Jar = j
	return azcore.TransportFunc(func(req *http.Request) (*http.Response, error) {
		return http.DefaultClient.Do(req)
	})
}

func TestLROResumeWrongPoller(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDelete202NoRetry204(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	diffPoller, err := op.ResumePost200WithPayload(rt)
	if err == nil {
		t.Fatal("expected an error but did not find receive one")
	}
	if diffPoller != nil {
		t.Fatal("expected a nil poller from the resume operation")
	}
}

func TestLROBeginDelete202NoRetry204(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDelete202NoRetry204(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDelete202NoRetry204(rt)
	if err != nil {
		t.Fatal(err)
	}
	prodResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, prodResp.RawResponse, 204)
}

func TestLROBeginDelete202Retry200(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDelete202Retry200(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDelete202Retry200(rt)
	if err != nil {
		t.Fatal(err)
	}
	prodResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, prodResp.RawResponse, 200)
}

func TestLROBeginDelete204Succeeded(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDelete204Succeeded(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	_, err = poller.ResumeToken()
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, res, 204)
}

func TestLROBeginDeleteAsyncNoHeaderInRetry(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDeleteAsyncNoHeaderInRetry(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDeleteAsyncNoHeaderInRetry(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, res, 200)
}

func TestLROBeginDeleteAsyncNoRetrySucceeded(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDeleteAsyncNoRetrySucceeded(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDeleteAsyncNoRetrySucceeded(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, res, 200)
}

func TestLROBeginDeleteAsyncRetryFailed(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDeleteAsyncRetryFailed(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDeleteAsyncRetryFailed(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if res != nil {
		t.Fatal("expected a nil response from the polling operation")
	}
	var cloudErr *CloudError
	if !errors.As(err, &cloudErr) {
		t.Fatal("expected a CloudError but did not receive one")
	}
	var httpResp azcore.HTTPResponse
	if !errors.As(err, &httpResp) {
		t.Fatal("expected azcore.HTTPResponse error")
	} else if sc := httpResp.RawResponse().StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
}

func TestLROBeginDeleteAsyncRetrySucceeded(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDeleteAsyncRetrySucceeded(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDeleteAsyncRetrySucceeded(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, res, 200)
}

func TestLROBeginDeleteAsyncRetrycanceled(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDeleteAsyncRetrycanceled(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDeleteAsyncRetrycanceled(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if res != nil {
		t.Fatal("expected a nil response from the polling operation")
	}
	var cloudErr *CloudError
	if !errors.As(err, &cloudErr) {
		t.Fatal("expected a CloudError but did not receive one")
	}
	var httpResp azcore.HTTPResponse
	if !errors.As(err, &httpResp) {
		t.Fatal("expected azcore.HTTPResponse error")
	} else if sc := httpResp.RawResponse().StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
}

func TestLROBeginDeleteNoHeaderInRetry(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDeleteNoHeaderInRetry(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDeleteNoHeaderInRetry(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, res, 204)
}

func TestLROBeginDeleteProvisioning202Accepted200Succeeded(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDeleteProvisioning202Accepted200Succeeded(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDeleteProvisioning202Accepted200Succeeded(rt)
	if err != nil {
		t.Fatal(err)
	}
	prodResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, prodResp.RawResponse, 200)
}

func TestLROBeginDeleteProvisioning202DeletingFailed200(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDeleteProvisioning202DeletingFailed200(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDeleteProvisioning202DeletingFailed200(rt)
	if err != nil {
		t.Fatal(err)
	}
	_, err = resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROBeginDeleteProvisioning202Deletingcanceled200(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginDeleteProvisioning202Deletingcanceled200(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumeDeleteProvisioning202Deletingcanceled200(rt)
	if err != nil {
		t.Fatal(err)
	}
	_, err = resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
}

func TestLROBeginPost200WithPayload(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPost200WithPayload(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePost200WithPayload(rt)
	if err != nil {
		t.Fatal(err)
	}
	skuResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, skuResp.RawResponse, 200)
	helpers.DeepEqualOrFatal(t, skuResp.SKU, &SKU{
		ID:   to.StringPtr("1"),
		Name: to.StringPtr("product"),
	})
}

func TestLROBeginPost202List(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPost202List(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePost202List(rt)
	if err != nil {
		t.Fatal(err)
	}
	prodArrayResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, prodArrayResp.RawResponse, 200)
	helpers.DeepEqualOrFatal(t, prodArrayResp.ProductArray, &[]Product{
		{
			Resource: Resource{
				ID:   to.StringPtr("100"),
				Name: to.StringPtr("foo"),
			},
		},
	})
}

func TestLROBeginPost202NoRetry204(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPost202NoRetry204(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePost202NoRetry204(rt)
	if err != nil {
		t.Fatal(err)
	}
	prodResp, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, prodResp.RawResponse, 204)
}

func TestLROBeginPost202Retry200(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPost202Retry200(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePost202Retry200(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, res, 200)
}

func TestLROBeginPostAsyncNoRetrySucceeded(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPostAsyncNoRetrySucceeded(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePostAsyncNoRetrySucceeded(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, res.RawResponse, 200)
	helpers.DeepEqualOrFatal(t, res.Product, &Product{
		Resource: Resource{
			ID:   to.StringPtr("100"),
			Name: to.StringPtr("foo"),
		},
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	})
}

func TestLROBeginPostAsyncRetryFailed(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPostAsyncRetryFailed(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePostAsyncRetryFailed(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if res != nil {
		t.Fatal("expected a nil response from the polling operation")
	}
	var cloudErr *CloudError
	if !errors.As(err, &cloudErr) {
		t.Fatal("expected a CloudError but did not receive one")
	}
	var httpResp azcore.HTTPResponse
	if !errors.As(err, &httpResp) {
		t.Fatal("expected azcore.HTTPResponse error")
	} else if sc := httpResp.RawResponse().StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
}

func TestLROBeginPostAsyncRetrySucceeded(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPostAsyncRetrySucceeded(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePostAsyncRetrySucceeded(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, res.RawResponse, 200)
	helpers.DeepEqualOrFatal(t, res.Product, &Product{
		Resource: Resource{
			ID:   to.StringPtr("100"),
			Name: to.StringPtr("foo"),
		},
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	})
}

func TestLROBeginPostAsyncRetrycanceled(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPostAsyncRetrycanceled(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePostAsyncRetrycanceled(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if res != nil {
		t.Fatal("expected a nil response from the polling operation")
	}
	var cloudErr *CloudError
	if !errors.As(err, &cloudErr) {
		t.Fatal("expected a CloudError but did not receive one")
	}
	var httpResp azcore.HTTPResponse
	if !errors.As(err, &httpResp) {
		t.Fatal("expected azcore.HTTPResponse error")
	} else if sc := httpResp.RawResponse().StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
}

func TestLROBeginPostDoubleHeadersFinalAzureHeaderGet(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPostDoubleHeadersFinalAzureHeaderGet(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePostDoubleHeadersFinalAzureHeaderGet(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, res.RawResponse, 200)
	helpers.DeepEqualOrFatal(t, res.Product, &Product{
		Resource: Resource{
			ID: to.StringPtr("100"),
		},
	})
}

func TestLROBeginPostDoubleHeadersFinalAzureHeaderGetDefault(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPostDoubleHeadersFinalAzureHeaderGetDefault(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePostDoubleHeadersFinalAzureHeaderGetDefault(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, res.RawResponse, 200)
	helpers.DeepEqualOrFatal(t, res.Product, &Product{
		Resource: Resource{
			ID:   to.StringPtr("100"),
			Name: to.StringPtr("foo"),
		},
	})
}

func TestLROBeginPostDoubleHeadersFinalLocationGet(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPostDoubleHeadersFinalLocationGet(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePostDoubleHeadersFinalLocationGet(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, res.RawResponse, 200)
	helpers.DeepEqualOrFatal(t, res.Product, &Product{})
}

func TestLROBeginPut200Acceptedcanceled200(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPut200Acceptedcanceled200(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePut200Acceptedcanceled200(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if res != nil {
		t.Fatal("expected a nil response from the polling operation")
	}
	var cloudErr *CloudError
	if !errors.As(err, &cloudErr) {
		t.Fatal("expected a CloudError but did not receive one")
	}
	var httpResp azcore.HTTPResponse
	if !errors.As(err, &httpResp) {
		t.Fatal("expected azcore.HTTPResponse error")
	} else if sc := httpResp.RawResponse().StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
}

func TestLROBeginPut200Succeeded(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPut200Succeeded(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	_, err = poller.ResumeToken()
	if err == nil {
		t.Fatal("Expected an error but did not receive one")
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, res.RawResponse, 200)
	helpers.DeepEqualOrFatal(t, res.Product, &Product{
		Resource: Resource{
			ID:   to.StringPtr("100"),
			Name: to.StringPtr("foo"),
		},
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	})
}

func TestLROBeginPut200SucceededNoState(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPut200SucceededNoState(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	_, err = poller.ResumeToken()
	if err == nil {
		t.Fatal("Expected an error but did not receive one")
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, res.RawResponse, 200)
	helpers.DeepEqualOrFatal(t, res.Product, &Product{
		Resource: Resource{
			ID:   to.StringPtr("100"),
			Name: to.StringPtr("foo"),
		},
	})
}

// TODO check if this test should actually be returning a 200 or a 204
func TestLROBeginPut200UpdatingSucceeded204(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPut200UpdatingSucceeded204(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePut200UpdatingSucceeded204(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, res.RawResponse, 200)
	helpers.DeepEqualOrFatal(t, res.Product, &Product{
		Resource: Resource{
			ID:   to.StringPtr("100"),
			Name: to.StringPtr("foo"),
		},
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	})
}

func TestLROBeginPut201CreatingFailed200(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPut201CreatingFailed200(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePut201CreatingFailed200(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if res != nil {
		t.Fatal("expected a nil response from the polling operation")
	}
	var cloudErr *CloudError
	if !errors.As(err, &cloudErr) {
		t.Fatal("expected a CloudError but did not receive one")
	}
	var httpResp azcore.HTTPResponse
	if !errors.As(err, &httpResp) {
		t.Fatal("expected azcore.HTTPResponse error")
	} else if sc := httpResp.RawResponse().StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
}

func TestLROBeginPut201CreatingSucceeded200(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPut201CreatingSucceeded200(context.Background(), &LrOSPut201CreatingSucceeded200Options{Product: &Product{}})
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePut201CreatingSucceeded200(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, res.RawResponse, 200)
	helpers.DeepEqualOrFatal(t, res.Product, &Product{
		Resource: Resource{
			ID:   to.StringPtr("100"),
			Name: to.StringPtr("foo"),
		},
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	})
}

func TestLROBeginPut202Retry200(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPut202Retry200(context.Background(), &LrOSPut202Retry200Options{Product: &Product{}})
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePut202Retry200(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, res.RawResponse, 200)
	helpers.DeepEqualOrFatal(t, res.Product, &Product{
		Resource: Resource{
			ID:   to.StringPtr("100"),
			Name: to.StringPtr("foo"),
		},
	})
}

func TestLROBeginPutAsyncNoHeaderInRetry(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPutAsyncNoHeaderInRetry(context.Background(), &LrOSPutAsyncNoHeaderInRetryOptions{Product: &Product{}})
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePutAsyncNoHeaderInRetry(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, res.RawResponse, 200)
	helpers.DeepEqualOrFatal(t, res.Product, &Product{
		Resource: Resource{
			ID:   to.StringPtr("100"),
			Name: to.StringPtr("foo"),
		},
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	})
}

func TestLROBeginPutAsyncNoRetrySucceeded(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPutAsyncNoRetrySucceeded(context.Background(), &LrOSPutAsyncNoRetrySucceededOptions{Product: &Product{}})
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePutAsyncNoRetrySucceeded(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, res.RawResponse, 200)
	helpers.DeepEqualOrFatal(t, res.Product, &Product{
		Resource: Resource{
			ID:   to.StringPtr("100"),
			Name: to.StringPtr("foo"),
		},
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	})
}

func TestLROBeginPutAsyncNoRetrycanceled(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPutAsyncNoRetrycanceled(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePutAsyncNoRetrycanceled(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if res != nil {
		t.Fatal("expected a nil response from the polling operation")
	}
	var cloudErr *CloudError
	if !errors.As(err, &cloudErr) {
		t.Fatal("expected a CloudError but did not receive one")
	}
	var httpResp azcore.HTTPResponse
	if !errors.As(err, &httpResp) {
		t.Fatal("expected azcore.HTTPResponse error")
	} else if sc := httpResp.RawResponse().StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
}

func TestLROBeginPutAsyncNonResource(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPutAsyncNonResource(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePutAsyncNonResource(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, res.RawResponse, 200)
	helpers.DeepEqualOrFatal(t, res.SKU, &SKU{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("sku"),
	})
}

func TestLROBeginPutAsyncRetryFailed(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPutAsyncRetryFailed(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePutAsyncRetryFailed(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err == nil {
		t.Fatal("expected an error but did not receive one")
	}
	if res != nil {
		t.Fatal("expected a nil response from the polling operation")
	}
	var cloudErr *CloudError
	if !errors.As(err, &cloudErr) {
		t.Fatal("expected a CloudError but did not receive one")
	}
	var httpResp azcore.HTTPResponse
	if !errors.As(err, &httpResp) {
		t.Fatal("expected azcore.HTTPResponse error")
	} else if sc := httpResp.RawResponse().StatusCode; sc != http.StatusOK {
		t.Fatalf("unexpected status code %d", sc)
	}
}

func TestLROBeginPutAsyncRetrySucceeded(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPutAsyncRetrySucceeded(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePutAsyncRetrySucceeded(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, res.RawResponse, 200)
	helpers.DeepEqualOrFatal(t, res.Product, &Product{
		Resource: Resource{
			ID:   to.StringPtr("100"),
			Name: to.StringPtr("foo"),
		},
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	})
}

func TestLROBeginPutAsyncSubResource(t *testing.T) {
	op := newLROSClient()
	resp, err := op.BeginPutAsyncSubResource(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePutAsyncSubResource(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, res.RawResponse, 200)
	helpers.DeepEqualOrFatal(t, res.SubProduct, &SubProduct{
		SubResource: SubResource{
			ID: to.StringPtr("100"),
		},
		Properties: &SubProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	})
}

func TestLROBeginPutNoHeaderInRetry(t *testing.T) {
	t.Skip("problem with put flow")
	op := newLROSClient()
	resp, err := op.BeginPutNoHeaderInRetry(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePutNoHeaderInRetry(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, res.RawResponse, 200)
	helpers.DeepEqualOrFatal(t, res.Product, &Product{
		Resource: Resource{
			ID: to.StringPtr("100"),
		},
		Properties: &ProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	})
}

func TestLROBeginPutNonResource(t *testing.T) {
	t.Skip("problem with put flow")
	op := newLROSClient()
	resp, err := op.BeginPutNonResource(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePutNonResource(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, res.RawResponse, 200)
	helpers.DeepEqualOrFatal(t, res.SKU, &SKU{
		ID:   to.StringPtr("100"),
		Name: to.StringPtr("sku"),
	})
}

func TestLROBeginPutSubResource(t *testing.T) {
	t.Skip("problem with put flow")
	op := newLROSClient()
	resp, err := op.BeginPutSubResource(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	poller := resp.Poller
	rt, err := poller.ResumeToken()
	if err != nil {
		t.Fatal(err)
	}
	poller, err = op.ResumePutSubResource(rt)
	if err != nil {
		t.Fatal(err)
	}
	res, err := resp.PollUntilDone(context.Background(), 1*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, res.RawResponse, 200)
	helpers.DeepEqualOrFatal(t, res.SubProduct, &SubProduct{
		SubResource: SubResource{
			ID: to.StringPtr("100"),
		},
		Properties: &SubProductProperties{
			ProvisioningState: to.StringPtr("Succeeded"),
		},
	})
}
