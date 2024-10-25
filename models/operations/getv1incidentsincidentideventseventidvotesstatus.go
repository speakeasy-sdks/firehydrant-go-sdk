// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type GetV1IncidentsIncidentIDEventsEventIDVotesStatusRequest struct {
	IncidentID string `pathParam:"style=simple,explode=false,name=incident_id"`
	EventID    string `pathParam:"style=simple,explode=false,name=event_id"`
}

func (o *GetV1IncidentsIncidentIDEventsEventIDVotesStatusRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *GetV1IncidentsIncidentIDEventsEventIDVotesStatusRequest) GetEventID() string {
	if o == nil {
		return ""
	}
	return o.EventID
}

type GetV1IncidentsIncidentIDEventsEventIDVotesStatusResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns the current vote counts for an object
	VotesEntity *components.VotesEntity
}

func (o *GetV1IncidentsIncidentIDEventsEventIDVotesStatusResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetV1IncidentsIncidentIDEventsEventIDVotesStatusResponse) GetVotesEntity() *components.VotesEntity {
	if o == nil {
		return nil
	}
	return o.VotesEntity
}
