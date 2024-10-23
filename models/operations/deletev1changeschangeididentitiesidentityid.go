// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type DeleteV1ChangesChangeIDIdentitiesIdentityIDRequest struct {
	IdentityID string `pathParam:"style=simple,explode=false,name=identity_id"`
	ChangeID   string `pathParam:"style=simple,explode=false,name=change_id"`
}

func (o *DeleteV1ChangesChangeIDIdentitiesIdentityIDRequest) GetIdentityID() string {
	if o == nil {
		return ""
	}
	return o.IdentityID
}

func (o *DeleteV1ChangesChangeIDIdentitiesIdentityIDRequest) GetChangeID() string {
	if o == nil {
		return ""
	}
	return o.ChangeID
}

type DeleteV1ChangesChangeIDIdentitiesIdentityIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *DeleteV1ChangesChangeIDIdentitiesIdentityIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}