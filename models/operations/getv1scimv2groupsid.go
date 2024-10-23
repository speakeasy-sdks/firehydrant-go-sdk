// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type GetV1ScimV2GroupsIDRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetV1ScimV2GroupsIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetV1ScimV2GroupsIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *GetV1ScimV2GroupsIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
