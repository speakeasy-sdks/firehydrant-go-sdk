// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type DeleteChecklistTemplateRequest struct {
	// Checklist Template UUID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteChecklistTemplateRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteChecklistTemplateResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Archive a checklist template
	ChecklistTemplateEntity *components.ChecklistTemplateEntity
}

func (o *DeleteChecklistTemplateResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteChecklistTemplateResponse) GetChecklistTemplateEntity() *components.ChecklistTemplateEntity {
	if o == nil {
		return nil
	}
	return o.ChecklistTemplateEntity
}
