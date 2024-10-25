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

func TestUsers_GetV1Users(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Users.List(ctx, nil, nil, nil, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestUsers_GetV1UsersIDServices(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Users.GetServices(ctx, "<id>", nil, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.Equal(t, []components.TeamEntityPaginated{}, res.TeamEntityPaginateds)
}
