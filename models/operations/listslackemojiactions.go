// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type ListSlackEmojiActionsRequest struct {
	// Slack Connection UUID
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
	Page         *int   `queryParam:"style=form,explode=true,name=page"`
	PerPage      *int   `queryParam:"style=form,explode=true,name=per_page"`
}

func (o *ListSlackEmojiActionsRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *ListSlackEmojiActionsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListSlackEmojiActionsRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type ListSlackEmojiActionsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *ListSlackEmojiActionsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}