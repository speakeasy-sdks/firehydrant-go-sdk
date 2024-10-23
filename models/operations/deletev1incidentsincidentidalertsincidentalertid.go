// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDRequest struct {
	IncidentAlertID string `pathParam:"style=simple,explode=false,name=incident_alert_id"`
	IncidentID      string `pathParam:"style=simple,explode=false,name=incident_id"`
}

func (o *DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDRequest) GetIncidentAlertID() string {
	if o == nil {
		return ""
	}
	return o.IncidentAlertID
}

func (o *DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

type DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
