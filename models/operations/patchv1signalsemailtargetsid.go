// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type PatchV1SignalsEmailTargetsIDRequest struct {
	ID                           string                                  `pathParam:"style=simple,explode=false,name=id"`
	PatchV1SignalsEmailTargetsID components.PatchV1SignalsEmailTargetsID `request:"mediaType=application/json"`
}

func (o *PatchV1SignalsEmailTargetsIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PatchV1SignalsEmailTargetsIDRequest) GetPatchV1SignalsEmailTargetsID() components.PatchV1SignalsEmailTargetsID {
	if o == nil {
		return components.PatchV1SignalsEmailTargetsID{}
	}
	return o.PatchV1SignalsEmailTargetsID
}

type PatchV1SignalsEmailTargetsIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *PatchV1SignalsEmailTargetsIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
