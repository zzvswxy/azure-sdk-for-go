//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpolicy

import "net/http"

// DataPolicyManifestsGetByPolicyModeResponse contains the response from method DataPolicyManifests.GetByPolicyMode.
type DataPolicyManifestsGetByPolicyModeResponse struct {
	DataPolicyManifestsGetByPolicyModeResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DataPolicyManifestsGetByPolicyModeResult contains the result from method DataPolicyManifests.GetByPolicyMode.
type DataPolicyManifestsGetByPolicyModeResult struct {
	DataPolicyManifest
}

// DataPolicyManifestsListResponse contains the response from method DataPolicyManifests.List.
type DataPolicyManifestsListResponse struct {
	DataPolicyManifestsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DataPolicyManifestsListResult contains the result from method DataPolicyManifests.List.
type DataPolicyManifestsListResult struct {
	DataPolicyManifestListResult
}

// PolicyAssignmentsCreateByIDResponse contains the response from method PolicyAssignments.CreateByID.
type PolicyAssignmentsCreateByIDResponse struct {
	PolicyAssignmentsCreateByIDResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicyAssignmentsCreateByIDResult contains the result from method PolicyAssignments.CreateByID.
type PolicyAssignmentsCreateByIDResult struct {
	PolicyAssignment
}

// PolicyAssignmentsCreateResponse contains the response from method PolicyAssignments.Create.
type PolicyAssignmentsCreateResponse struct {
	PolicyAssignmentsCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicyAssignmentsCreateResult contains the result from method PolicyAssignments.Create.
type PolicyAssignmentsCreateResult struct {
	PolicyAssignment
}

// PolicyAssignmentsDeleteByIDResponse contains the response from method PolicyAssignments.DeleteByID.
type PolicyAssignmentsDeleteByIDResponse struct {
	PolicyAssignmentsDeleteByIDResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicyAssignmentsDeleteByIDResult contains the result from method PolicyAssignments.DeleteByID.
type PolicyAssignmentsDeleteByIDResult struct {
	PolicyAssignment
}

// PolicyAssignmentsDeleteResponse contains the response from method PolicyAssignments.Delete.
type PolicyAssignmentsDeleteResponse struct {
	PolicyAssignmentsDeleteResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicyAssignmentsDeleteResult contains the result from method PolicyAssignments.Delete.
type PolicyAssignmentsDeleteResult struct {
	PolicyAssignment
}

// PolicyAssignmentsGetByIDResponse contains the response from method PolicyAssignments.GetByID.
type PolicyAssignmentsGetByIDResponse struct {
	PolicyAssignmentsGetByIDResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicyAssignmentsGetByIDResult contains the result from method PolicyAssignments.GetByID.
type PolicyAssignmentsGetByIDResult struct {
	PolicyAssignment
}

// PolicyAssignmentsGetResponse contains the response from method PolicyAssignments.Get.
type PolicyAssignmentsGetResponse struct {
	PolicyAssignmentsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicyAssignmentsGetResult contains the result from method PolicyAssignments.Get.
type PolicyAssignmentsGetResult struct {
	PolicyAssignment
}

// PolicyAssignmentsListForManagementGroupResponse contains the response from method PolicyAssignments.ListForManagementGroup.
type PolicyAssignmentsListForManagementGroupResponse struct {
	PolicyAssignmentsListForManagementGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicyAssignmentsListForManagementGroupResult contains the result from method PolicyAssignments.ListForManagementGroup.
type PolicyAssignmentsListForManagementGroupResult struct {
	PolicyAssignmentListResult
}

// PolicyAssignmentsListForResourceGroupResponse contains the response from method PolicyAssignments.ListForResourceGroup.
type PolicyAssignmentsListForResourceGroupResponse struct {
	PolicyAssignmentsListForResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicyAssignmentsListForResourceGroupResult contains the result from method PolicyAssignments.ListForResourceGroup.
type PolicyAssignmentsListForResourceGroupResult struct {
	PolicyAssignmentListResult
}

// PolicyAssignmentsListForResourceResponse contains the response from method PolicyAssignments.ListForResource.
type PolicyAssignmentsListForResourceResponse struct {
	PolicyAssignmentsListForResourceResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicyAssignmentsListForResourceResult contains the result from method PolicyAssignments.ListForResource.
type PolicyAssignmentsListForResourceResult struct {
	PolicyAssignmentListResult
}

// PolicyAssignmentsListResponse contains the response from method PolicyAssignments.List.
type PolicyAssignmentsListResponse struct {
	PolicyAssignmentsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicyAssignmentsListResult contains the result from method PolicyAssignments.List.
type PolicyAssignmentsListResult struct {
	PolicyAssignmentListResult
}

// PolicyAssignmentsUpdateByIDResponse contains the response from method PolicyAssignments.UpdateByID.
type PolicyAssignmentsUpdateByIDResponse struct {
	PolicyAssignmentsUpdateByIDResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicyAssignmentsUpdateByIDResult contains the result from method PolicyAssignments.UpdateByID.
type PolicyAssignmentsUpdateByIDResult struct {
	PolicyAssignment
}

// PolicyAssignmentsUpdateResponse contains the response from method PolicyAssignments.Update.
type PolicyAssignmentsUpdateResponse struct {
	PolicyAssignmentsUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicyAssignmentsUpdateResult contains the result from method PolicyAssignments.Update.
type PolicyAssignmentsUpdateResult struct {
	PolicyAssignment
}

// PolicyDefinitionsCreateOrUpdateAtManagementGroupResponse contains the response from method PolicyDefinitions.CreateOrUpdateAtManagementGroup.
type PolicyDefinitionsCreateOrUpdateAtManagementGroupResponse struct {
	PolicyDefinitionsCreateOrUpdateAtManagementGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicyDefinitionsCreateOrUpdateAtManagementGroupResult contains the result from method PolicyDefinitions.CreateOrUpdateAtManagementGroup.
type PolicyDefinitionsCreateOrUpdateAtManagementGroupResult struct {
	PolicyDefinition
}

// PolicyDefinitionsCreateOrUpdateResponse contains the response from method PolicyDefinitions.CreateOrUpdate.
type PolicyDefinitionsCreateOrUpdateResponse struct {
	PolicyDefinitionsCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicyDefinitionsCreateOrUpdateResult contains the result from method PolicyDefinitions.CreateOrUpdate.
type PolicyDefinitionsCreateOrUpdateResult struct {
	PolicyDefinition
}

// PolicyDefinitionsDeleteAtManagementGroupResponse contains the response from method PolicyDefinitions.DeleteAtManagementGroup.
type PolicyDefinitionsDeleteAtManagementGroupResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicyDefinitionsDeleteResponse contains the response from method PolicyDefinitions.Delete.
type PolicyDefinitionsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicyDefinitionsGetAtManagementGroupResponse contains the response from method PolicyDefinitions.GetAtManagementGroup.
type PolicyDefinitionsGetAtManagementGroupResponse struct {
	PolicyDefinitionsGetAtManagementGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicyDefinitionsGetAtManagementGroupResult contains the result from method PolicyDefinitions.GetAtManagementGroup.
type PolicyDefinitionsGetAtManagementGroupResult struct {
	PolicyDefinition
}

// PolicyDefinitionsGetBuiltInResponse contains the response from method PolicyDefinitions.GetBuiltIn.
type PolicyDefinitionsGetBuiltInResponse struct {
	PolicyDefinitionsGetBuiltInResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicyDefinitionsGetBuiltInResult contains the result from method PolicyDefinitions.GetBuiltIn.
type PolicyDefinitionsGetBuiltInResult struct {
	PolicyDefinition
}

// PolicyDefinitionsGetResponse contains the response from method PolicyDefinitions.Get.
type PolicyDefinitionsGetResponse struct {
	PolicyDefinitionsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicyDefinitionsGetResult contains the result from method PolicyDefinitions.Get.
type PolicyDefinitionsGetResult struct {
	PolicyDefinition
}

// PolicyDefinitionsListBuiltInResponse contains the response from method PolicyDefinitions.ListBuiltIn.
type PolicyDefinitionsListBuiltInResponse struct {
	PolicyDefinitionsListBuiltInResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicyDefinitionsListBuiltInResult contains the result from method PolicyDefinitions.ListBuiltIn.
type PolicyDefinitionsListBuiltInResult struct {
	PolicyDefinitionListResult
}

// PolicyDefinitionsListByManagementGroupResponse contains the response from method PolicyDefinitions.ListByManagementGroup.
type PolicyDefinitionsListByManagementGroupResponse struct {
	PolicyDefinitionsListByManagementGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicyDefinitionsListByManagementGroupResult contains the result from method PolicyDefinitions.ListByManagementGroup.
type PolicyDefinitionsListByManagementGroupResult struct {
	PolicyDefinitionListResult
}

// PolicyDefinitionsListResponse contains the response from method PolicyDefinitions.List.
type PolicyDefinitionsListResponse struct {
	PolicyDefinitionsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicyDefinitionsListResult contains the result from method PolicyDefinitions.List.
type PolicyDefinitionsListResult struct {
	PolicyDefinitionListResult
}

// PolicyExemptionsCreateOrUpdateResponse contains the response from method PolicyExemptions.CreateOrUpdate.
type PolicyExemptionsCreateOrUpdateResponse struct {
	PolicyExemptionsCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicyExemptionsCreateOrUpdateResult contains the result from method PolicyExemptions.CreateOrUpdate.
type PolicyExemptionsCreateOrUpdateResult struct {
	PolicyExemption
}

// PolicyExemptionsDeleteResponse contains the response from method PolicyExemptions.Delete.
type PolicyExemptionsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicyExemptionsGetResponse contains the response from method PolicyExemptions.Get.
type PolicyExemptionsGetResponse struct {
	PolicyExemptionsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicyExemptionsGetResult contains the result from method PolicyExemptions.Get.
type PolicyExemptionsGetResult struct {
	PolicyExemption
}

// PolicyExemptionsListForManagementGroupResponse contains the response from method PolicyExemptions.ListForManagementGroup.
type PolicyExemptionsListForManagementGroupResponse struct {
	PolicyExemptionsListForManagementGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicyExemptionsListForManagementGroupResult contains the result from method PolicyExemptions.ListForManagementGroup.
type PolicyExemptionsListForManagementGroupResult struct {
	PolicyExemptionListResult
}

// PolicyExemptionsListForResourceGroupResponse contains the response from method PolicyExemptions.ListForResourceGroup.
type PolicyExemptionsListForResourceGroupResponse struct {
	PolicyExemptionsListForResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicyExemptionsListForResourceGroupResult contains the result from method PolicyExemptions.ListForResourceGroup.
type PolicyExemptionsListForResourceGroupResult struct {
	PolicyExemptionListResult
}

// PolicyExemptionsListForResourceResponse contains the response from method PolicyExemptions.ListForResource.
type PolicyExemptionsListForResourceResponse struct {
	PolicyExemptionsListForResourceResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicyExemptionsListForResourceResult contains the result from method PolicyExemptions.ListForResource.
type PolicyExemptionsListForResourceResult struct {
	PolicyExemptionListResult
}

// PolicyExemptionsListResponse contains the response from method PolicyExemptions.List.
type PolicyExemptionsListResponse struct {
	PolicyExemptionsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicyExemptionsListResult contains the result from method PolicyExemptions.List.
type PolicyExemptionsListResult struct {
	PolicyExemptionListResult
}

// PolicySetDefinitionsCreateOrUpdateAtManagementGroupResponse contains the response from method PolicySetDefinitions.CreateOrUpdateAtManagementGroup.
type PolicySetDefinitionsCreateOrUpdateAtManagementGroupResponse struct {
	PolicySetDefinitionsCreateOrUpdateAtManagementGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicySetDefinitionsCreateOrUpdateAtManagementGroupResult contains the result from method PolicySetDefinitions.CreateOrUpdateAtManagementGroup.
type PolicySetDefinitionsCreateOrUpdateAtManagementGroupResult struct {
	PolicySetDefinition
}

// PolicySetDefinitionsCreateOrUpdateResponse contains the response from method PolicySetDefinitions.CreateOrUpdate.
type PolicySetDefinitionsCreateOrUpdateResponse struct {
	PolicySetDefinitionsCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicySetDefinitionsCreateOrUpdateResult contains the result from method PolicySetDefinitions.CreateOrUpdate.
type PolicySetDefinitionsCreateOrUpdateResult struct {
	PolicySetDefinition
}

// PolicySetDefinitionsDeleteAtManagementGroupResponse contains the response from method PolicySetDefinitions.DeleteAtManagementGroup.
type PolicySetDefinitionsDeleteAtManagementGroupResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicySetDefinitionsDeleteResponse contains the response from method PolicySetDefinitions.Delete.
type PolicySetDefinitionsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicySetDefinitionsGetAtManagementGroupResponse contains the response from method PolicySetDefinitions.GetAtManagementGroup.
type PolicySetDefinitionsGetAtManagementGroupResponse struct {
	PolicySetDefinitionsGetAtManagementGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicySetDefinitionsGetAtManagementGroupResult contains the result from method PolicySetDefinitions.GetAtManagementGroup.
type PolicySetDefinitionsGetAtManagementGroupResult struct {
	PolicySetDefinition
}

// PolicySetDefinitionsGetBuiltInResponse contains the response from method PolicySetDefinitions.GetBuiltIn.
type PolicySetDefinitionsGetBuiltInResponse struct {
	PolicySetDefinitionsGetBuiltInResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicySetDefinitionsGetBuiltInResult contains the result from method PolicySetDefinitions.GetBuiltIn.
type PolicySetDefinitionsGetBuiltInResult struct {
	PolicySetDefinition
}

// PolicySetDefinitionsGetResponse contains the response from method PolicySetDefinitions.Get.
type PolicySetDefinitionsGetResponse struct {
	PolicySetDefinitionsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicySetDefinitionsGetResult contains the result from method PolicySetDefinitions.Get.
type PolicySetDefinitionsGetResult struct {
	PolicySetDefinition
}

// PolicySetDefinitionsListBuiltInResponse contains the response from method PolicySetDefinitions.ListBuiltIn.
type PolicySetDefinitionsListBuiltInResponse struct {
	PolicySetDefinitionsListBuiltInResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicySetDefinitionsListBuiltInResult contains the result from method PolicySetDefinitions.ListBuiltIn.
type PolicySetDefinitionsListBuiltInResult struct {
	PolicySetDefinitionListResult
}

// PolicySetDefinitionsListByManagementGroupResponse contains the response from method PolicySetDefinitions.ListByManagementGroup.
type PolicySetDefinitionsListByManagementGroupResponse struct {
	PolicySetDefinitionsListByManagementGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicySetDefinitionsListByManagementGroupResult contains the result from method PolicySetDefinitions.ListByManagementGroup.
type PolicySetDefinitionsListByManagementGroupResult struct {
	PolicySetDefinitionListResult
}

// PolicySetDefinitionsListResponse contains the response from method PolicySetDefinitions.List.
type PolicySetDefinitionsListResponse struct {
	PolicySetDefinitionsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PolicySetDefinitionsListResult contains the result from method PolicySetDefinitions.List.
type PolicySetDefinitionsListResult struct {
	PolicySetDefinitionListResult
}