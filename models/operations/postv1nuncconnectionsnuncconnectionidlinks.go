// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type PostV1NuncConnectionsNuncConnectionIDLinksRequest struct {
	NuncConnectionID string `pathParam:"style=simple,explode=false,name=nunc_connection_id"`
}

func (o *PostV1NuncConnectionsNuncConnectionIDLinksRequest) GetNuncConnectionID() string {
	if o == nil {
		return ""
	}
	return o.NuncConnectionID
}

type PostV1NuncConnectionsNuncConnectionIDLinksResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Add a link to be displayed on a FireHydrant status page
	NuncConnectionEntity *components.NuncConnectionEntity
}

func (o *PostV1NuncConnectionsNuncConnectionIDLinksResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PostV1NuncConnectionsNuncConnectionIDLinksResponse) GetNuncConnectionEntity() *components.NuncConnectionEntity {
	if o == nil {
		return nil
	}
	return o.NuncConnectionEntity
}