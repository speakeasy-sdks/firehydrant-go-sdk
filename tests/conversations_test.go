// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package tests

import (
	"context"
	"firehydrant"
	"firehydrant/models/components"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConversations_ListConversationComments(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Conversations.ListComments(ctx, "<id>", nil, nil, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestConversations_CreateConversationComment(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Conversations.CreateComment(ctx, "<id>", components.PostV1ConversationsConversationIDComments{
		Body: "<value>",
	})
	require.NoError(t, err)
	assert.Equal(t, 201, res.HTTPMeta.Response.StatusCode)
}

func TestConversations_GetConversationComment(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Conversations.GetComment(ctx, "<id>", "<id>")
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestConversations_DeleteConversationComment(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Conversations.DeleteComment(ctx, "<id>", "<id>")
	require.NoError(t, err)
	assert.Equal(t, 204, res.HTTPMeta.Response.StatusCode)
}

func TestConversations_UpdateConversationComment(t *testing.T) {
	assert.Fail(t, "incomplete test found please make sure to address the following errors: [`missing request body values for required request body`]")
}

func TestConversations_ListConversationCommentReactions(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Conversations.ListCommentReactions(ctx, "<id>", "<id>")
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestConversations_CreateConversationCommentReaction(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Conversations.CreateCommentReaction(ctx, "<id>", "<id>", components.PostV1ConversationsConversationIDCommentsCommentIDReactions{
		Reaction: "<value>",
	})
	require.NoError(t, err)
	assert.Equal(t, 201, res.HTTPMeta.Response.StatusCode)
}

func TestConversations_DeleteConversationCommentReaction(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Conversations.DeleteCommentReaction(ctx, "<id>", "<id>", "<id>")
	require.NoError(t, err)
	assert.Equal(t, 204, res.HTTPMeta.Response.StatusCode)
}
