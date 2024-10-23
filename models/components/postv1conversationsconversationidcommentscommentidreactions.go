// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// PostV1ConversationsConversationIDCommentsCommentIDReactions - ALPHA - Create a reaction on a comment
type PostV1ConversationsConversationIDCommentsCommentIDReactions struct {
	// CLDR short name of Unicode emojis
	Reaction string `json:"reaction"`
}

func (o *PostV1ConversationsConversationIDCommentsCommentIDReactions) GetReaction() string {
	if o == nil {
		return ""
	}
	return o.Reaction
}
