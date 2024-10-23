// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type PutV1ScimV2UsersIDRequest struct {
	ID                 string                        `pathParam:"style=simple,explode=false,name=id"`
	PutV1ScimV2UsersID components.PutV1ScimV2UsersID `request:"mediaType=application/scim+json"`
}

func (o *PutV1ScimV2UsersIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PutV1ScimV2UsersIDRequest) GetPutV1ScimV2UsersID() components.PutV1ScimV2UsersID {
	if o == nil {
		return components.PutV1ScimV2UsersID{}
	}
	return o.PutV1ScimV2UsersID
}

type PutV1ScimV2UsersIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *PutV1ScimV2UsersIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}