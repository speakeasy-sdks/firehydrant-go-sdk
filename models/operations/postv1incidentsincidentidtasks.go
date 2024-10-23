// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type PostV1IncidentsIncidentIDTasksRequest struct {
	IncidentID                     string                                    `pathParam:"style=simple,explode=false,name=incident_id"`
	PostV1IncidentsIncidentIDTasks components.PostV1IncidentsIncidentIDTasks `request:"mediaType=application/json"`
}

func (o *PostV1IncidentsIncidentIDTasksRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *PostV1IncidentsIncidentIDTasksRequest) GetPostV1IncidentsIncidentIDTasks() components.PostV1IncidentsIncidentIDTasks {
	if o == nil {
		return components.PostV1IncidentsIncidentIDTasks{}
	}
	return o.PostV1IncidentsIncidentIDTasks
}

type PostV1IncidentsIncidentIDTasksResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Create a task
	TaskEntity *components.TaskEntity
}

func (o *PostV1IncidentsIncidentIDTasksResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PostV1IncidentsIncidentIDTasksResponse) GetTaskEntity() *components.TaskEntity {
	if o == nil {
		return nil
	}
	return o.TaskEntity
}
