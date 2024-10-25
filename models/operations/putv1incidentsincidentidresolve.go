// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type PutV1IncidentsIncidentIDResolveRequestBody struct {
	// The slug of any milestone in the post-incident or closed phase to set on the incident (and its children, if `resolve_children` os set). Must be one of the configured milestones available on this incident.
	Milestone *string `form:"name=milestone"`
}

func (o *PutV1IncidentsIncidentIDResolveRequestBody) GetMilestone() *string {
	if o == nil {
		return nil
	}
	return o.Milestone
}

type PutV1IncidentsIncidentIDResolveRequest struct {
	IncidentID  string                                      `pathParam:"style=simple,explode=false,name=incident_id"`
	RequestBody *PutV1IncidentsIncidentIDResolveRequestBody `request:"mediaType=application/x-www-form-urlencoded"`
}

func (o *PutV1IncidentsIncidentIDResolveRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *PutV1IncidentsIncidentIDResolveRequest) GetRequestBody() *PutV1IncidentsIncidentIDResolveRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type PutV1IncidentsIncidentIDResolveResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Resolves a currently active incident
	IncidentEntity *components.IncidentEntity
}

func (o *PutV1IncidentsIncidentIDResolveResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PutV1IncidentsIncidentIDResolveResponse) GetIncidentEntity() *components.IncidentEntity {
	if o == nil {
		return nil
	}
	return o.IncidentEntity
}
