// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type DeleteV1ConversationsConversationIDCommentsCommentIDRequest struct {
	CommentID      string `pathParam:"style=simple,explode=false,name=comment_id"`
	ConversationID string `pathParam:"style=simple,explode=false,name=conversation_id"`
}

func (o *DeleteV1ConversationsConversationIDCommentsCommentIDRequest) GetCommentID() string {
	if o == nil {
		return ""
	}
	return o.CommentID
}

func (o *DeleteV1ConversationsConversationIDCommentsCommentIDRequest) GetConversationID() string {
	if o == nil {
		return ""
	}
	return o.ConversationID
}

type DeleteV1ConversationsConversationIDCommentsCommentIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *DeleteV1ConversationsConversationIDCommentsCommentIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}