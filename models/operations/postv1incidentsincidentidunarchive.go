// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type PostV1IncidentsIncidentIDUnarchiveRequest struct {
	IncidentID string `pathParam:"style=simple,explode=false,name=incident_id"`
}

func (o *PostV1IncidentsIncidentIDUnarchiveRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

type PostV1IncidentsIncidentIDUnarchiveResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Unarchives an incident
	IncidentEntity *components.IncidentEntity
}

func (o *PostV1IncidentsIncidentIDUnarchiveResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PostV1IncidentsIncidentIDUnarchiveResponse) GetIncidentEntity() *components.IncidentEntity {
	if o == nil {
		return nil
	}
	return o.IncidentEntity
}
