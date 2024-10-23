// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type GetV1IncidentsIncidentIDStatusPagesRequest struct {
	IncidentID string `pathParam:"style=simple,explode=false,name=incident_id"`
}

func (o *GetV1IncidentsIncidentIDStatusPagesRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

type GetV1IncidentsIncidentIDStatusPagesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// List status pages that are attached to an incident
	IncidentsStatusPageEntityPaginated *components.IncidentsStatusPageEntityPaginated
}

func (o *GetV1IncidentsIncidentIDStatusPagesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetV1IncidentsIncidentIDStatusPagesResponse) GetIncidentsStatusPageEntityPaginated() *components.IncidentsStatusPageEntityPaginated {
	if o == nil {
		return nil
	}
	return o.IncidentsStatusPageEntityPaginated
}
