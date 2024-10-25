// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type GetV1ChangeTypesRequest struct {
	Page    *int `queryParam:"style=form,explode=true,name=page"`
	PerPage *int `queryParam:"style=form,explode=true,name=per_page"`
}

func (o *GetV1ChangeTypesRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetV1ChangeTypesRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type GetV1ChangeTypesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Lists all change types
	ChangeTypeEntityPaginated *components.ChangeTypeEntityPaginated
}

func (o *GetV1ChangeTypesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetV1ChangeTypesResponse) GetChangeTypeEntityPaginated() *components.ChangeTypeEntityPaginated {
	if o == nil {
		return nil
	}
	return o.ChangeTypeEntityPaginated
}
