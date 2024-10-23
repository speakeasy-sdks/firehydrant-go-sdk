// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type PostV1IncidentsIncidentIDTasksTaskIDConvertRequest struct {
	TaskID                                      string                                                 `pathParam:"style=simple,explode=false,name=task_id"`
	IncidentID                                  string                                                 `pathParam:"style=simple,explode=false,name=incident_id"`
	PostV1IncidentsIncidentIDTasksTaskIDConvert components.PostV1IncidentsIncidentIDTasksTaskIDConvert `request:"mediaType=application/json"`
}

func (o *PostV1IncidentsIncidentIDTasksTaskIDConvertRequest) GetTaskID() string {
	if o == nil {
		return ""
	}
	return o.TaskID
}

func (o *PostV1IncidentsIncidentIDTasksTaskIDConvertRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *PostV1IncidentsIncidentIDTasksTaskIDConvertRequest) GetPostV1IncidentsIncidentIDTasksTaskIDConvert() components.PostV1IncidentsIncidentIDTasksTaskIDConvert {
	if o == nil {
		return components.PostV1IncidentsIncidentIDTasksTaskIDConvert{}
	}
	return o.PostV1IncidentsIncidentIDTasksTaskIDConvert
}

type PostV1IncidentsIncidentIDTasksTaskIDConvertResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Convert a task to a follow-up
	TaskEntityPaginated *components.TaskEntityPaginated
}

func (o *PostV1IncidentsIncidentIDTasksTaskIDConvertResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PostV1IncidentsIncidentIDTasksTaskIDConvertResponse) GetTaskEntityPaginated() *components.TaskEntityPaginated {
	if o == nil {
		return nil
	}
	return o.TaskEntityPaginated
}