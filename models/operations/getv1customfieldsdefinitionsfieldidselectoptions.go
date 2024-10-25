// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type GetV1CustomFieldsDefinitionsFieldIDSelectOptionsRequest struct {
	FieldID string `pathParam:"style=simple,explode=false,name=field_id"`
	// Text string of a query for filtering values.
	Query *string `queryParam:"style=form,explode=true,name=query"`
	// If true, return all versions of the custom field definition.
	AllVersions *bool `queryParam:"style=form,explode=true,name=all_versions"`
}

func (o *GetV1CustomFieldsDefinitionsFieldIDSelectOptionsRequest) GetFieldID() string {
	if o == nil {
		return ""
	}
	return o.FieldID
}

func (o *GetV1CustomFieldsDefinitionsFieldIDSelectOptionsRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *GetV1CustomFieldsDefinitionsFieldIDSelectOptionsRequest) GetAllVersions() *bool {
	if o == nil {
		return nil
	}
	return o.AllVersions
}

type GetV1CustomFieldsDefinitionsFieldIDSelectOptionsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Get the permissible values for the a currently active custom select or multi-select field.
	OrganizationsCustomFieldDefinitionEntity *components.OrganizationsCustomFieldDefinitionEntity
}

func (o *GetV1CustomFieldsDefinitionsFieldIDSelectOptionsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetV1CustomFieldsDefinitionsFieldIDSelectOptionsResponse) GetOrganizationsCustomFieldDefinitionEntity() *components.OrganizationsCustomFieldDefinitionEntity {
	if o == nil {
		return nil
	}
	return o.OrganizationsCustomFieldDefinitionEntity
}
