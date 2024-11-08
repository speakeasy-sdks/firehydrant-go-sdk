// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type CreateServiceChecklistResponseRequest struct {
	ServiceID                                           string                                                         `pathParam:"style=simple,explode=false,name=service_id"`
	ChecklistID                                         string                                                         `pathParam:"style=simple,explode=false,name=checklist_id"`
	PostV1ServicesServiceIDChecklistResponseChecklistID components.PostV1ServicesServiceIDChecklistResponseChecklistID `request:"mediaType=application/json"`
}

func (o *CreateServiceChecklistResponseRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *CreateServiceChecklistResponseRequest) GetChecklistID() string {
	if o == nil {
		return ""
	}
	return o.ChecklistID
}

func (o *CreateServiceChecklistResponseRequest) GetPostV1ServicesServiceIDChecklistResponseChecklistID() components.PostV1ServicesServiceIDChecklistResponseChecklistID {
	if o == nil {
		return components.PostV1ServicesServiceIDChecklistResponseChecklistID{}
	}
	return o.PostV1ServicesServiceIDChecklistResponseChecklistID
}

type CreateServiceChecklistResponseResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *CreateServiceChecklistResponseResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}