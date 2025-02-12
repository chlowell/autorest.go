//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconsumption

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// BudgetsClient contains the methods for the Budgets group.
// Don't use this type directly, use NewBudgetsClient() instead.
type BudgetsClient struct {
	ep string
	pl runtime.Pipeline
}

// NewBudgetsClient creates a new instance of BudgetsClient with the specified values.
func NewBudgetsClient(con *arm.Connection) *BudgetsClient {
	return &BudgetsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version)}
}

// CreateOrUpdate - The operation to create or update a budget. You can optionally provide an eTag if desired as a form of concurrency control. To obtain
// the latest eTag for a given budget, perform a get operation prior
// to your put operation.
// If the operation fails it returns the *ErrorResponse error type.
func (client *BudgetsClient) CreateOrUpdate(ctx context.Context, scope string, budgetName string, parameters Budget, options *BudgetsCreateOrUpdateOptions) (BudgetsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, scope, budgetName, parameters, options)
	if err != nil {
		return BudgetsCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BudgetsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return BudgetsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *BudgetsClient) createOrUpdateCreateRequest(ctx context.Context, scope string, budgetName string, parameters Budget, options *BudgetsCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Consumption/budgets/{budgetName}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if budgetName == "" {
		return nil, errors.New("parameter budgetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{budgetName}", url.PathEscape(budgetName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *BudgetsClient) createOrUpdateHandleResponse(resp *http.Response) (BudgetsCreateOrUpdateResponse, error) {
	result := BudgetsCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Budget); err != nil {
		return BudgetsCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *BudgetsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - The operation to delete a budget.
// If the operation fails it returns the *ErrorResponse error type.
func (client *BudgetsClient) Delete(ctx context.Context, scope string, budgetName string, options *BudgetsDeleteOptions) (BudgetsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, scope, budgetName, options)
	if err != nil {
		return BudgetsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BudgetsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BudgetsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return BudgetsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *BudgetsClient) deleteCreateRequest(ctx context.Context, scope string, budgetName string, options *BudgetsDeleteOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Consumption/budgets/{budgetName}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if budgetName == "" {
		return nil, errors.New("parameter budgetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{budgetName}", url.PathEscape(budgetName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *BudgetsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets the budget for the scope by budget name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *BudgetsClient) Get(ctx context.Context, scope string, budgetName string, options *BudgetsGetOptions) (BudgetsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, scope, budgetName, options)
	if err != nil {
		return BudgetsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return BudgetsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BudgetsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *BudgetsClient) getCreateRequest(ctx context.Context, scope string, budgetName string, options *BudgetsGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Consumption/budgets/{budgetName}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if budgetName == "" {
		return nil, errors.New("parameter budgetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{budgetName}", url.PathEscape(budgetName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BudgetsClient) getHandleResponse(resp *http.Response) (BudgetsGetResponse, error) {
	result := BudgetsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Budget); err != nil {
		return BudgetsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *BudgetsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Lists all budgets for the defined scope.
// If the operation fails it returns the *ErrorResponse error type.
func (client *BudgetsClient) List(scope string, options *BudgetsListOptions) *BudgetsListPager {
	return &BudgetsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, scope, options)
		},
		advancer: func(ctx context.Context, resp BudgetsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.BudgetsListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *BudgetsClient) listCreateRequest(ctx context.Context, scope string, options *BudgetsListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Consumption/budgets"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BudgetsClient) listHandleResponse(resp *http.Response) (BudgetsListResponse, error) {
	result := BudgetsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BudgetsListResult); err != nil {
		return BudgetsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *BudgetsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
