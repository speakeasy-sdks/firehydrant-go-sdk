// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type PatchV1IncidentsIncidentIDRequest struct {
	IncidentID                 string                                `pathParam:"style=simple,explode=false,name=incident_id"`
	PatchV1IncidentsIncidentID components.PatchV1IncidentsIncidentID `request:"mediaType=application/json"`
}

func (o *PatchV1IncidentsIncidentIDRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *PatchV1IncidentsIncidentIDRequest) GetPatchV1IncidentsIncidentID() components.PatchV1IncidentsIncidentID {
	if o == nil {
		return components.PatchV1IncidentsIncidentID{}
	}
	return o.PatchV1IncidentsIncidentID
}

type PatchV1IncidentsIncidentIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Updates an incident with provided parameters
	IncidentEntity *components.IncidentEntity
}

func (o *PatchV1IncidentsIncidentIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PatchV1IncidentsIncidentIDResponse) GetIncidentEntity() *components.IncidentEntity {
	if o == nil {
		return nil
	}
	return o.IncidentEntity
}
