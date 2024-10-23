// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type GetV1ChangesChangeIDIdentitiesRequest struct {
	Page     *int   `queryParam:"style=form,explode=true,name=page"`
	PerPage  *int   `queryParam:"style=form,explode=true,name=per_page"`
	ChangeID string `pathParam:"style=simple,explode=false,name=change_id"`
}

func (o *GetV1ChangesChangeIDIdentitiesRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetV1ChangesChangeIDIdentitiesRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *GetV1ChangesChangeIDIdentitiesRequest) GetChangeID() string {
	if o == nil {
		return ""
	}
	return o.ChangeID
}

type GetV1ChangesChangeIDIdentitiesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Retrieve all identities for the change
	ChangeIdentityEntityPaginated *components.ChangeIdentityEntityPaginated
}

func (o *GetV1ChangesChangeIDIdentitiesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetV1ChangesChangeIDIdentitiesResponse) GetChangeIdentityEntityPaginated() *components.ChangeIdentityEntityPaginated {
	if o == nil {
		return nil
	}
	return o.ChangeIdentityEntityPaginated
}
