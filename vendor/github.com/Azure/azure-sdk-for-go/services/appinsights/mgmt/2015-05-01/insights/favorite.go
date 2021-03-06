package insights

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// FavoriteClient is the composite Swagger for Application Insights Management Client
type FavoriteClient struct {
	BaseClient
}

// NewFavoriteClient creates an instance of the FavoriteClient client.
func NewFavoriteClient(subscriptionID string) FavoriteClient {
	return NewFavoriteClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewFavoriteClientWithBaseURI creates an instance of the FavoriteClient client.
func NewFavoriteClientWithBaseURI(baseURI string, subscriptionID string) FavoriteClient {
	return FavoriteClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Add adds a new favorites to an Application Insights component.
//
// resourceGroupName is the name of the resource group. resourceName is the name of the Application Insights
// component resource. favoriteID is the Id of a specific favorite defined in the Application Insights component
// favoriteProperties is properties that need to be specified to create a new favorite and add it to an Application
// Insights component.
func (client FavoriteClient) Add(ctx context.Context, resourceGroupName string, resourceName string, favoriteID string, favoriteProperties ApplicationInsightsComponentFavorite) (result ApplicationInsightsComponentFavorite, err error) {
	req, err := client.AddPreparer(ctx, resourceGroupName, resourceName, favoriteID, favoriteProperties)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.FavoriteClient", "Add", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.FavoriteClient", "Add", resp, "Failure sending request")
		return
	}

	result, err = client.AddResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.FavoriteClient", "Add", resp, "Failure responding to request")
	}

	return
}

// AddPreparer prepares the Add request.
func (client FavoriteClient) AddPreparer(ctx context.Context, resourceGroupName string, resourceName string, favoriteID string, favoriteProperties ApplicationInsightsComponentFavorite) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"favoriteId":        autorest.Encode("path", favoriteID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/favorites/{favoriteId}", pathParameters),
		autorest.WithJSON(favoriteProperties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AddSender sends the Add request. The method will close the
// http.Response Body if it receives an error.
func (client FavoriteClient) AddSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// AddResponder handles the response to the Add request. The method always
// closes the http.Response Body.
func (client FavoriteClient) AddResponder(resp *http.Response) (result ApplicationInsightsComponentFavorite, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete remove a favorite that is associated to an Application Insights component.
//
// resourceGroupName is the name of the resource group. resourceName is the name of the Application Insights
// component resource. favoriteID is the Id of a specific favorite defined in the Application Insights component
func (client FavoriteClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, favoriteID string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(ctx, resourceGroupName, resourceName, favoriteID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.FavoriteClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "insights.FavoriteClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.FavoriteClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client FavoriteClient) DeletePreparer(ctx context.Context, resourceGroupName string, resourceName string, favoriteID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"favoriteId":        autorest.Encode("path", favoriteID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/favorites/{favoriteId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client FavoriteClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client FavoriteClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get a single favorite by its FavoriteId, defined within an Application Insights component.
//
// resourceGroupName is the name of the resource group. resourceName is the name of the Application Insights
// component resource. favoriteID is the Id of a specific favorite defined in the Application Insights component
func (client FavoriteClient) Get(ctx context.Context, resourceGroupName string, resourceName string, favoriteID string) (result ApplicationInsightsComponentFavorite, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, resourceName, favoriteID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.FavoriteClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.FavoriteClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.FavoriteClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client FavoriteClient) GetPreparer(ctx context.Context, resourceGroupName string, resourceName string, favoriteID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"favoriteId":        autorest.Encode("path", favoriteID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/favorites/{favoriteId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client FavoriteClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client FavoriteClient) GetResponder(resp *http.Response) (result ApplicationInsightsComponentFavorite, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates a favorite that has already been added to an Application Insights component.
//
// resourceGroupName is the name of the resource group. resourceName is the name of the Application Insights
// component resource. favoriteID is the Id of a specific favorite defined in the Application Insights component
// favoriteProperties is properties that need to be specified to update the existing favorite.
func (client FavoriteClient) Update(ctx context.Context, resourceGroupName string, resourceName string, favoriteID string, favoriteProperties ApplicationInsightsComponentFavorite) (result ApplicationInsightsComponentFavorite, err error) {
	req, err := client.UpdatePreparer(ctx, resourceGroupName, resourceName, favoriteID, favoriteProperties)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.FavoriteClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.FavoriteClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.FavoriteClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client FavoriteClient) UpdatePreparer(ctx context.Context, resourceGroupName string, resourceName string, favoriteID string, favoriteProperties ApplicationInsightsComponentFavorite) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"favoriteId":        autorest.Encode("path", favoriteID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/favorites/{favoriteId}", pathParameters),
		autorest.WithJSON(favoriteProperties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client FavoriteClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client FavoriteClient) UpdateResponder(resp *http.Response) (result ApplicationInsightsComponentFavorite, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
