// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type GetV1ChecklistTemplatesIDRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetV1ChecklistTemplatesIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetV1ChecklistTemplatesIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Retrieves a single checklist template by ID
	ChecklistTemplateEntity *components.ChecklistTemplateEntity
}

func (o *GetV1ChecklistTemplatesIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetV1ChecklistTemplatesIDResponse) GetChecklistTemplateEntity() *components.ChecklistTemplateEntity {
	if o == nil {
		return nil
	}
	return o.ChecklistTemplateEntity
}