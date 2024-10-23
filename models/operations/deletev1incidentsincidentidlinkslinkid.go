// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type DeleteV1IncidentsIncidentIDLinksLinkIDRequest struct {
	LinkID     string `pathParam:"style=simple,explode=false,name=link_id"`
	IncidentID string `pathParam:"style=simple,explode=false,name=incident_id"`
}

func (o *DeleteV1IncidentsIncidentIDLinksLinkIDRequest) GetLinkID() string {
	if o == nil {
		return ""
	}
	return o.LinkID
}

func (o *DeleteV1IncidentsIncidentIDLinksLinkIDRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

type DeleteV1IncidentsIncidentIDLinksLinkIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *DeleteV1IncidentsIncidentIDLinksLinkIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}