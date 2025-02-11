// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type CreateIncidentTypeResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Create a new incident type
	IncidentTypeEntity *components.IncidentTypeEntity
}

func (o *CreateIncidentTypeResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateIncidentTypeResponse) GetIncidentTypeEntity() *components.IncidentTypeEntity {
	if o == nil {
		return nil
	}
	return o.IncidentTypeEntity
}
