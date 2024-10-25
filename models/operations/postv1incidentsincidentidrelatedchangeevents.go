// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type PostV1IncidentsIncidentIDRelatedChangeEventsRequest struct {
	IncidentID                                   string                                                  `pathParam:"style=simple,explode=false,name=incident_id"`
	PostV1IncidentsIncidentIDRelatedChangeEvents components.PostV1IncidentsIncidentIDRelatedChangeEvents `request:"mediaType=application/json"`
}

func (o *PostV1IncidentsIncidentIDRelatedChangeEventsRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *PostV1IncidentsIncidentIDRelatedChangeEventsRequest) GetPostV1IncidentsIncidentIDRelatedChangeEvents() components.PostV1IncidentsIncidentIDRelatedChangeEvents {
	if o == nil {
		return components.PostV1IncidentsIncidentIDRelatedChangeEvents{}
	}
	return o.PostV1IncidentsIncidentIDRelatedChangeEvents
}

type PostV1IncidentsIncidentIDRelatedChangeEventsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Add a related change to an incident. Changes added to an incident can be causes, fixes, or suspects. To remove a change from an incident, the type field should be set to dismissed.
	IncidentsRelatedChangeEventEntity *components.IncidentsRelatedChangeEventEntity
}

func (o *PostV1IncidentsIncidentIDRelatedChangeEventsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PostV1IncidentsIncidentIDRelatedChangeEventsResponse) GetIncidentsRelatedChangeEventEntity() *components.IncidentsRelatedChangeEventEntity {
	if o == nil {
		return nil
	}
	return o.IncidentsRelatedChangeEventEntity
}
