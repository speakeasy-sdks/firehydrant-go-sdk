// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type PatchV1ChangesChangeIDIdentitiesIdentityIDRequest struct {
	IdentityID                                 string                                                `pathParam:"style=simple,explode=false,name=identity_id"`
	ChangeID                                   string                                                `pathParam:"style=simple,explode=false,name=change_id"`
	PatchV1ChangesChangeIDIdentitiesIdentityID components.PatchV1ChangesChangeIDIdentitiesIdentityID `request:"mediaType=application/json"`
}

func (o *PatchV1ChangesChangeIDIdentitiesIdentityIDRequest) GetIdentityID() string {
	if o == nil {
		return ""
	}
	return o.IdentityID
}

func (o *PatchV1ChangesChangeIDIdentitiesIdentityIDRequest) GetChangeID() string {
	if o == nil {
		return ""
	}
	return o.ChangeID
}

func (o *PatchV1ChangesChangeIDIdentitiesIdentityIDRequest) GetPatchV1ChangesChangeIDIdentitiesIdentityID() components.PatchV1ChangesChangeIDIdentitiesIdentityID {
	if o == nil {
		return components.PatchV1ChangesChangeIDIdentitiesIdentityID{}
	}
	return o.PatchV1ChangesChangeIDIdentitiesIdentityID
}

type PatchV1ChangesChangeIDIdentitiesIdentityIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Update an identity
	ChangeIdentityEntity *components.ChangeIdentityEntity
}

func (o *PatchV1ChangesChangeIDIdentitiesIdentityIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PatchV1ChangesChangeIDIdentitiesIdentityIDResponse) GetChangeIdentityEntity() *components.ChangeIdentityEntity {
	if o == nil {
		return nil
	}
	return o.ChangeIdentityEntity
}
