// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type GetV1IncidentTypesIDRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetV1IncidentTypesIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetV1IncidentTypesIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Retrieve a single incident type from its ID
	IncidentTypeEntity *components.IncidentTypeEntity
}

func (o *GetV1IncidentTypesIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetV1IncidentTypesIDResponse) GetIncidentTypeEntity() *components.IncidentTypeEntity {
	if o == nil {
		return nil
	}
	return o.IncidentTypeEntity
}
