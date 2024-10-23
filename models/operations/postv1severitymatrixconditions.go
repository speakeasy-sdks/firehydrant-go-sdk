// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type PostV1SeverityMatrixConditionsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Create a new condition
	SeverityMatrixConditionEntity *components.SeverityMatrixConditionEntity
}

func (o *PostV1SeverityMatrixConditionsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PostV1SeverityMatrixConditionsResponse) GetSeverityMatrixConditionEntity() *components.SeverityMatrixConditionEntity {
	if o == nil {
		return nil
	}
	return o.SeverityMatrixConditionEntity
}
