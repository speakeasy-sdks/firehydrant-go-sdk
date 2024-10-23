// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type PatchV1IncidentsIncidentIDImpactRequest struct {
	IncidentID                       string                                      `pathParam:"style=simple,explode=false,name=incident_id"`
	PatchV1IncidentsIncidentIDImpact components.PatchV1IncidentsIncidentIDImpact `request:"mediaType=application/json"`
}

func (o *PatchV1IncidentsIncidentIDImpactRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *PatchV1IncidentsIncidentIDImpactRequest) GetPatchV1IncidentsIncidentIDImpact() components.PatchV1IncidentsIncidentIDImpact {
	if o == nil {
		return components.PatchV1IncidentsIncidentIDImpact{}
	}
	return o.PatchV1IncidentsIncidentIDImpact
}

type PatchV1IncidentsIncidentIDImpactResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Allows updating an incident's impacted infrastructure, with the option to
	// move the incident into a different milestone and provide a note to update
	// the incident timeline and any attached status pages. If this method is
	// requested with the PUT verb, impacts will be completely replaced with the
	// information in the request body, even if not provided (effectively clearing
	// all impacts). If this method is requested with the PATCH verb, the provided
	// impacts will be added or updated, but no impacts will be removed.
	//
	IncidentEntity *components.IncidentEntity
}

func (o *PatchV1IncidentsIncidentIDImpactResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PatchV1IncidentsIncidentIDImpactResponse) GetIncidentEntity() *components.IncidentEntity {
	if o == nil {
		return nil
	}
	return o.IncidentEntity
}
