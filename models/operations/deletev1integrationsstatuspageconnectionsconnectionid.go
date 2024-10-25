// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type DeleteV1IntegrationsStatuspageConnectionsConnectionIDRequest struct {
	// Connection UUID
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
}

func (o *DeleteV1IntegrationsStatuspageConnectionsConnectionIDRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type DeleteV1IntegrationsStatuspageConnectionsConnectionIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Deletes the given Statuspage integration connection.
	IntegrationsStatuspageConnectionEntity *components.IntegrationsStatuspageConnectionEntity
}

func (o *DeleteV1IntegrationsStatuspageConnectionsConnectionIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteV1IntegrationsStatuspageConnectionsConnectionIDResponse) GetIntegrationsStatuspageConnectionEntity() *components.IntegrationsStatuspageConnectionEntity {
	if o == nil {
		return nil
	}
	return o.IntegrationsStatuspageConnectionEntity
}
