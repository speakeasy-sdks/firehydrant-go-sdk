// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type GetV1IncidentsIncidentIDEventsEventIDRequest struct {
	IncidentID string `pathParam:"style=simple,explode=false,name=incident_id"`
	EventID    string `pathParam:"style=simple,explode=false,name=event_id"`
}

func (o *GetV1IncidentsIncidentIDEventsEventIDRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *GetV1IncidentsIncidentIDEventsEventIDRequest) GetEventID() string {
	if o == nil {
		return ""
	}
	return o.EventID
}

type GetV1IncidentsIncidentIDEventsEventIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Retrieve a single event for an incident
	IncidentEventEntity *components.IncidentEventEntity
}

func (o *GetV1IncidentsIncidentIDEventsEventIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetV1IncidentsIncidentIDEventsEventIDResponse) GetIncidentEventEntity() *components.IncidentEventEntity {
	if o == nil {
		return nil
	}
	return o.IncidentEventEntity
}
