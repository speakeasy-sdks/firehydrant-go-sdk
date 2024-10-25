// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type PatchV1SeverityMatrixConditionsConditionIDRequest struct {
	ConditionID                                string                                                `pathParam:"style=simple,explode=false,name=condition_id"`
	PatchV1SeverityMatrixConditionsConditionID components.PatchV1SeverityMatrixConditionsConditionID `request:"mediaType=application/json"`
}

func (o *PatchV1SeverityMatrixConditionsConditionIDRequest) GetConditionID() string {
	if o == nil {
		return ""
	}
	return o.ConditionID
}

func (o *PatchV1SeverityMatrixConditionsConditionIDRequest) GetPatchV1SeverityMatrixConditionsConditionID() components.PatchV1SeverityMatrixConditionsConditionID {
	if o == nil {
		return components.PatchV1SeverityMatrixConditionsConditionID{}
	}
	return o.PatchV1SeverityMatrixConditionsConditionID
}

type PatchV1SeverityMatrixConditionsConditionIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Update a specific condition
	SeverityMatrixConditionEntity *components.SeverityMatrixConditionEntity
}

func (o *PatchV1SeverityMatrixConditionsConditionIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PatchV1SeverityMatrixConditionsConditionIDResponse) GetSeverityMatrixConditionEntity() *components.SeverityMatrixConditionEntity {
	if o == nil {
		return nil
	}
	return o.SeverityMatrixConditionEntity
}
