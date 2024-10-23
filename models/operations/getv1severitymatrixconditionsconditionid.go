// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type GetV1SeverityMatrixConditionsConditionIDRequest struct {
	ConditionID string `pathParam:"style=simple,explode=false,name=condition_id"`
}

func (o *GetV1SeverityMatrixConditionsConditionIDRequest) GetConditionID() string {
	if o == nil {
		return ""
	}
	return o.ConditionID
}

type GetV1SeverityMatrixConditionsConditionIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *GetV1SeverityMatrixConditionsConditionIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
