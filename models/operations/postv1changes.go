// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type PostV1ChangesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Create a new change entry
	ChangeEntity *components.ChangeEntity
}

func (o *PostV1ChangesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PostV1ChangesResponse) GetChangeEntity() *components.ChangeEntity {
	if o == nil {
		return nil
	}
	return o.ChangeEntity
}
