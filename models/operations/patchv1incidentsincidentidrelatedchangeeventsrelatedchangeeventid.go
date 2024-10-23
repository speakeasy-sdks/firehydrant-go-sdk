// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDRequest struct {
	RelatedChangeEventID                                              string                                                                       `pathParam:"style=simple,explode=false,name=related_change_event_id"`
	IncidentID                                                        string                                                                       `pathParam:"style=simple,explode=false,name=incident_id"`
	PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventID components.PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventID `request:"mediaType=application/json"`
}

func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDRequest) GetRelatedChangeEventID() string {
	if o == nil {
		return ""
	}
	return o.RelatedChangeEventID
}

func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDRequest) GetPatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventID() components.PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventID {
	if o == nil {
		return components.PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventID{}
	}
	return o.PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventID
}

type PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Update a change attached to an incident
	IncidentsRelatedChangeEventEntity *components.IncidentsRelatedChangeEventEntity
}

func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDResponse) GetIncidentsRelatedChangeEventEntity() *components.IncidentsRelatedChangeEventEntity {
	if o == nil {
		return nil
	}
	return o.IncidentsRelatedChangeEventEntity
}