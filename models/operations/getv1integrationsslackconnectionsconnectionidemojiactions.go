// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsRequest struct {
	// Slack Connection UUID
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	Page         *int   `queryParam:"style=form,explode=true,name=page"`
	PerPage      *int   `queryParam:"style=form,explode=true,name=per_page"`
}

func (o *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *GetV1IntegrationsSlackConnectionsConnectionIDEmojiActionsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
