// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type UpdateIntegrationConnectionRequest struct {
	Slug         string `pathParam:"style=simple,explode=false,name=slug"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
}

func (o *UpdateIntegrationConnectionRequest) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

func (o *UpdateIntegrationConnectionRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type UpdateIntegrationConnectionResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *UpdateIntegrationConnectionResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}