// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type PostV1ScimV2UsersResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *PostV1ScimV2UsersResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
