// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type GetV1IncidentsIncidentIDEventsRequest struct {
	IncidentID string `pathParam:"style=simple,explode=false,name=incident_id"`
	// A comma separated list of types of events to filter by
	Types   *string `queryParam:"style=form,explode=true,name=types"`
	Page    *int    `queryParam:"style=form,explode=true,name=page"`
	PerPage *int    `queryParam:"style=form,explode=true,name=per_page"`
}

func (o *GetV1IncidentsIncidentIDEventsRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *GetV1IncidentsIncidentIDEventsRequest) GetTypes() *string {
	if o == nil {
		return nil
	}
	return o.Types
}

func (o *GetV1IncidentsIncidentIDEventsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetV1IncidentsIncidentIDEventsRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type GetV1IncidentsIncidentIDEventsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// List all events for an incident. An event is a timeline entry. This can be filtered with params to retrieve events of a certain type.
	IncidentEventEntityPaginated *components.IncidentEventEntityPaginated
}

func (o *GetV1IncidentsIncidentIDEventsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetV1IncidentsIncidentIDEventsResponse) GetIncidentEventEntityPaginated() *components.IncidentEventEntityPaginated {
	if o == nil {
		return nil
	}
	return o.IncidentEventEntityPaginated
}
