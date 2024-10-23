// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type DeleteV1IncidentsIncidentIDRequest struct {
	IncidentID string `pathParam:"style=simple,explode=false,name=incident_id"`
}

func (o *DeleteV1IncidentsIncidentIDRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

type DeleteV1IncidentsIncidentIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Archives an incident which will hide it from lists and metrics
	IncidentEntity *components.IncidentEntity
}

func (o *DeleteV1IncidentsIncidentIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteV1IncidentsIncidentIDResponse) GetIncidentEntity() *components.IncidentEntity {
	if o == nil {
		return nil
	}
	return o.IncidentEntity
}
