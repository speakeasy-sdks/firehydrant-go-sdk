// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package tests

import (
	"context"
	"firehydrant"
	"firehydrant/models/components"
	"firehydrant/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMaintenances_ListScheduledMaintenances(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Maintenances.List(ctx, nil, nil, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestMaintenances_CreateScheduledMaintenance(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Maintenances.Create(ctx, components.PostV1ScheduledMaintenances{
		Name:     "<value>",
		StartsAt: types.MustTimeFromString("2023-04-18T00:41:46.120Z"),
		EndsAt:   types.MustTimeFromString("2024-01-18T06:46:27.056Z"),
	})
	require.NoError(t, err)
	assert.Equal(t, 201, res.HTTPMeta.Response.StatusCode)
}

func TestMaintenances_DeleteScheduledMaintenance(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Maintenances.Delete(ctx, "<id>")
	require.NoError(t, err)
	assert.Equal(t, 204, res.HTTPMeta.Response.StatusCode)
}
