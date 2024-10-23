// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type GetV1TaskListsRequest struct {
	Page    *int `queryParam:"style=form,explode=true,name=page"`
	PerPage *int `queryParam:"style=form,explode=true,name=per_page"`
}

func (o *GetV1TaskListsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetV1TaskListsRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type GetV1TaskListsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Lists all task lists for your organization
	TaskListEntity *components.TaskListEntity
}

func (o *GetV1TaskListsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetV1TaskListsResponse) GetTaskListEntity() *components.TaskListEntity {
	if o == nil {
		return nil
	}
	return o.TaskListEntity
}
