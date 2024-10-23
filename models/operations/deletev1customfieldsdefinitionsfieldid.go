// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type DeleteV1CustomFieldsDefinitionsFieldIDRequest struct {
	FieldID string `pathParam:"style=simple,explode=false,name=field_id"`
}

func (o *DeleteV1CustomFieldsDefinitionsFieldIDRequest) GetFieldID() string {
	if o == nil {
		return ""
	}
	return o.FieldID
}

type DeleteV1CustomFieldsDefinitionsFieldIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Delete a custom field definition
	OrganizationsCustomFieldDefinitionEntity *components.OrganizationsCustomFieldDefinitionEntity
}

func (o *DeleteV1CustomFieldsDefinitionsFieldIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteV1CustomFieldsDefinitionsFieldIDResponse) GetOrganizationsCustomFieldDefinitionEntity() *components.OrganizationsCustomFieldDefinitionEntity {
	if o == nil {
		return nil
	}
	return o.OrganizationsCustomFieldDefinitionEntity
}