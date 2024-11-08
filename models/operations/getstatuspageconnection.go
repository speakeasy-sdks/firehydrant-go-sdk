// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type GetStatuspageConnectionRequest struct {
	// Connection UUID
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
}

func (o *GetStatuspageConnectionRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type GetStatuspageConnectionResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Retrieve the information about the Statuspage connection.
	IntegrationsStatuspageConnectionEntity *components.IntegrationsStatuspageConnectionEntity
}

func (o *GetStatuspageConnectionResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetStatuspageConnectionResponse) GetIntegrationsStatuspageConnectionEntity() *components.IntegrationsStatuspageConnectionEntity {
	if o == nil {
		return nil
	}
	return o.IntegrationsStatuspageConnectionEntity
}