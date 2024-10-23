// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type GetV1RunbooksSelectOptionsIntegrationSlugActionSlugFieldRequest struct {
	IntegrationSlug string `pathParam:"style=simple,explode=false,name=integration_slug"`
	ActionSlug      string `pathParam:"style=simple,explode=false,name=action_slug"`
	Field           string `pathParam:"style=simple,explode=false,name=field"`
	// Text string of a query for filtering values.
	Query *string `queryParam:"style=form,explode=true,name=query"`
	// Generic params used to add specificity (eg an id of some kind) to the select options request
	Scope *string `queryParam:"style=form,explode=true,name=scope"`
	// Maximum number of items to return.
	PerPage *int `queryParam:"style=form,explode=true,name=per_page"`
}

func (o *GetV1RunbooksSelectOptionsIntegrationSlugActionSlugFieldRequest) GetIntegrationSlug() string {
	if o == nil {
		return ""
	}
	return o.IntegrationSlug
}

func (o *GetV1RunbooksSelectOptionsIntegrationSlugActionSlugFieldRequest) GetActionSlug() string {
	if o == nil {
		return ""
	}
	return o.ActionSlug
}

func (o *GetV1RunbooksSelectOptionsIntegrationSlugActionSlugFieldRequest) GetField() string {
	if o == nil {
		return ""
	}
	return o.Field
}

func (o *GetV1RunbooksSelectOptionsIntegrationSlugActionSlugFieldRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *GetV1RunbooksSelectOptionsIntegrationSlugActionSlugFieldRequest) GetScope() *string {
	if o == nil {
		return nil
	}
	return o.Scope
}

func (o *GetV1RunbooksSelectOptionsIntegrationSlugActionSlugFieldRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type GetV1RunbooksSelectOptionsIntegrationSlugActionSlugFieldResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *GetV1RunbooksSelectOptionsIntegrationSlugActionSlugFieldResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}