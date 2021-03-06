package web

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
// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// GlobalResourceGroupsClient is the use these APIs to manage Azure Websites
// resources through the Azure Resource Manager. All task operations conform
// to the HTTP/1.1 protocol specification and each operation returns an
// x-ms-request-id header that can be used to obtain information about the
// request. You must make sure that requests made to these resources are
// secure. For more information, see <a
// href="https://msdn.microsoft.com/en-us/library/azure/dn790557.aspx">Authenticating
// Azure Resource Manager requests.</a>
type GlobalResourceGroupsClient struct {
	ManagementClient
}

// NewGlobalResourceGroupsClient creates an instance of the
// GlobalResourceGroupsClient client.
func NewGlobalResourceGroupsClient(subscriptionID string) GlobalResourceGroupsClient {
	return NewGlobalResourceGroupsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewGlobalResourceGroupsClientWithBaseURI creates an instance of the
// GlobalResourceGroupsClient client.
func NewGlobalResourceGroupsClientWithBaseURI(baseURI string, subscriptionID string) GlobalResourceGroupsClient {
	return GlobalResourceGroupsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// MoveResources sends the move resources request.
//
func (client GlobalResourceGroupsClient) MoveResources(resourceGroupName string, moveResourceEnvelope CsmMoveResourceEnvelope) (result autorest.Response, err error) {
	req, err := client.MoveResourcesPreparer(resourceGroupName, moveResourceEnvelope)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.GlobalResourceGroupsClient", "MoveResources", nil, "Failure preparing request")
	}

	resp, err := client.MoveResourcesSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "web.GlobalResourceGroupsClient", "MoveResources", resp, "Failure sending request")
	}

	result, err = client.MoveResourcesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.GlobalResourceGroupsClient", "MoveResources", resp, "Failure responding to request")
	}

	return
}

// MoveResourcesPreparer prepares the MoveResources request.
func (client GlobalResourceGroupsClient) MoveResourcesPreparer(resourceGroupName string, moveResourceEnvelope CsmMoveResourceEnvelope) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/moveResources", pathParameters),
		autorest.WithJSON(moveResourceEnvelope),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// MoveResourcesSender sends the MoveResources request. The method will close the
// http.Response Body if it receives an error.
func (client GlobalResourceGroupsClient) MoveResourcesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// MoveResourcesResponder handles the response to the MoveResources request. The method always
// closes the http.Response Body.
func (client GlobalResourceGroupsClient) MoveResourcesResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
