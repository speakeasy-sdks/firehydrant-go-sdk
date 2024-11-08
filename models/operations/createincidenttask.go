// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type CreateIncidentTaskRequest struct {
	IncidentID                     string                                    `pathParam:"style=simple,explode=false,name=incident_id"`
	PostV1IncidentsIncidentIDTasks components.PostV1IncidentsIncidentIDTasks `request:"mediaType=application/json"`
}

func (o *CreateIncidentTaskRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *CreateIncidentTaskRequest) GetPostV1IncidentsIncidentIDTasks() components.PostV1IncidentsIncidentIDTasks {
	if o == nil {
		return components.PostV1IncidentsIncidentIDTasks{}
	}
	return o.PostV1IncidentsIncidentIDTasks
}

type CreateIncidentTaskResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Create a task
	TaskEntity *components.TaskEntity
}

func (o *CreateIncidentTaskResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateIncidentTaskResponse) GetTaskEntity() *components.TaskEntity {
	if o == nil {
		return nil
	}
	return o.TaskEntity
}