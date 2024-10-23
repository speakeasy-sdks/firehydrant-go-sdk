// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type GetV1IntegrationsSlackConnectionsConnectionIDWorkspacesRequest struct {
	// Connection UUID
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
}

func (o *GetV1IntegrationsSlackConnectionsConnectionIDWorkspacesRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type GetV1IntegrationsSlackConnectionsConnectionIDWorkspacesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *GetV1IntegrationsSlackConnectionsConnectionIDWorkspacesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}