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

func TestCustomFields_DeleteV1CustomFieldsDefinitionsFieldID(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.CustomFields.DeleteDefinition(ctx, "<id>")
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestCustomFields_PatchV1CustomFieldsDefinitionsFieldID(t *testing.T) {
	assert.Fail(t, "incomplete test found please make sure to address the following errors: [`missing request body values for required request body`]")
}

func TestCustomFields_PostV1CustomFieldsDefinitions(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.CustomFields.CreateDefinition(ctx, components.PostV1CustomFieldsDefinitions{
		DisplayName: "Hank_Blick88",
		FieldType:   "<value>",
		Required:    true,
	})
	require.NoError(t, err)
	assert.Equal(t, 201, res.HTTPMeta.Response.StatusCode)
}

func TestCustomFields_GetV1CustomFieldsDefinitions(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.CustomFields.ListDefinitions(ctx)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestCustomFields_GetV1CustomFieldsDefinitionsFieldIDSelectOptions(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.CustomFields.GetSelectOptions(ctx, "<id>", nil, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}
