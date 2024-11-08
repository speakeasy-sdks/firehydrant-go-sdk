// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type UpdateIncidentTypeRequest struct {
	ID                     string                            `pathParam:"style=simple,explode=false,name=id"`
	PatchV1IncidentTypesID components.PatchV1IncidentTypesID `request:"mediaType=application/json"`
}

func (o *UpdateIncidentTypeRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateIncidentTypeRequest) GetPatchV1IncidentTypesID() components.PatchV1IncidentTypesID {
	if o == nil {
		return components.PatchV1IncidentTypesID{}
	}
	return o.PatchV1IncidentTypesID
}

type UpdateIncidentTypeResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Update a single incident type from its ID
	IncidentTypeEntity *components.IncidentTypeEntity
}

func (o *UpdateIncidentTypeResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateIncidentTypeResponse) GetIncidentTypeEntity() *components.IncidentTypeEntity {
	if o == nil {
		return nil
	}
	return o.IncidentTypeEntity
}