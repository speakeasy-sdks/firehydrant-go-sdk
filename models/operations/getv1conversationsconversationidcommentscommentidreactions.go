// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type GetV1ConversationsConversationIDCommentsCommentIDReactionsRequest struct {
	ConversationID string `pathParam:"style=simple,explode=false,name=conversation_id"`
	CommentID      string `pathParam:"style=simple,explode=false,name=comment_id"`
}

func (o *GetV1ConversationsConversationIDCommentsCommentIDReactionsRequest) GetConversationID() string {
	if o == nil {
		return ""
	}
	return o.ConversationID
}

func (o *GetV1ConversationsConversationIDCommentsCommentIDReactionsRequest) GetCommentID() string {
	if o == nil {
		return ""
	}
	return o.CommentID
}

type GetV1ConversationsConversationIDCommentsCommentIDReactionsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *GetV1ConversationsConversationIDCommentsCommentIDReactionsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
