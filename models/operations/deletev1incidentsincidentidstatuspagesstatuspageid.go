// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDRequest struct {
	StatusPageID string `pathParam:"style=simple,explode=false,name=status_page_id"`
	IncidentID   string `pathParam:"style=simple,explode=false,name=incident_id"`
}

func (o *DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDRequest) GetStatusPageID() string {
	if o == nil {
		return ""
	}
	return o.StatusPageID
}

func (o *DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

type DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
