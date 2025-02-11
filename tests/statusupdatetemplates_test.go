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

func TestStatusUpdateTemplates_CreateStatusUpdateTemplate(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.StatusUpdateTemplates.Create(ctx, components.PostV1StatusUpdateTemplates{
		Name: "<value>",
		Body: "<value>",
	})
	require.NoError(t, err)
	assert.Equal(t, 201, res.HTTPMeta.Response.StatusCode)
}
