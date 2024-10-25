// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package tests

import (
	"context"
	"firehydrant"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIncidentTags_PostV1IncidentTagsValidate(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.IncidentTags.Validate(ctx, []string{
		"<value>",
		"<value>",
	})
	require.NoError(t, err)
	assert.Equal(t, 201, res.HTTPMeta.Response.StatusCode)
}

func TestIncidentTags_GetV1IncidentTags(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.IncidentTags.List(ctx, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}
