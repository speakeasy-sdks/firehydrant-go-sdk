// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type PatchV1IntegrationsConnectionsSlugConnectionIDRefreshRequest struct {
	Slug         string `pathParam:"style=simple,explode=false,name=slug"`
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
}

func (o *PatchV1IntegrationsConnectionsSlugConnectionIDRefreshRequest) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

func (o *PatchV1IntegrationsConnectionsSlugConnectionIDRefreshRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type PatchV1IntegrationsConnectionsSlugConnectionIDRefreshResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *PatchV1IntegrationsConnectionsSlugConnectionIDRefreshResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
