// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type GetV1NoauthPingResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Simple endpoint to verify your API connection is working
	PongEntity *components.PongEntity
}

func (o *GetV1NoauthPingResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetV1NoauthPingResponse) GetPongEntity() *components.PongEntity {
	if o == nil {
		return nil
	}
	return o.PongEntity
}
