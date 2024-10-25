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

func TestPostMortems_PostV1PostMortemsReports(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.PostMortems.CreateReport(ctx, components.PostV1PostMortemsReports{
		IncidentID: "<id>",
	})
	require.NoError(t, err)
	assert.Equal(t, 201, res.HTTPMeta.Response.StatusCode)
}

func TestPostMortems_GetV1PostMortemsReportsReportID(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.PostMortems.GetReport(ctx, "<id>")
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestPostMortems_PostV1PostMortemsReportsReportIDReasons(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.PostMortems.CreateReason(ctx, "<id>", components.PostV1PostMortemsReportsReportIDReasons{
		Summary: "<value>",
	})
	require.NoError(t, err)
	assert.Equal(t, 201, res.HTTPMeta.Response.StatusCode)
}

func TestPostMortems_DeleteV1PostMortemsReportsReportIDReasonsReasonID(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.PostMortems.DeleteReason(ctx, "<id>", "<id>")
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestPostMortems_PutV1PostMortemsReportsReportIDReasonsOrder(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.PostMortems.ReorderContributingFactor(ctx, "<id>", components.PutV1PostMortemsReportsReportIDReasonsOrder{
		OldPosition: 825361,
		NewPosition: 334032,
	})
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestPostMortems_PatchV1PostMortemsReportsReportIDFieldsFieldID(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.PostMortems.UpdateField(ctx, "<id>", "<id>", components.PatchV1PostMortemsReportsReportIDFieldsFieldID{
		Value: "<value>",
	})
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestPostMortems_GetV1PostMortemsReports(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.PostMortems.ListReports(ctx, nil, nil, nil, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestPostMortems_PatchV1PostMortemsReportsReportID(t *testing.T) {
	assert.Fail(t, "incomplete test found please make sure to address the following errors: [`missing request body values for required request body`]")
}

func TestPostMortems_GetV1PostMortemsReportsReportIDReasons(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.PostMortems.ListContributingFactors(ctx, "<id>", nil, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestPostMortems_PatchV1PostMortemsReportsReportIDReasonsReasonID(t *testing.T) {
	assert.Fail(t, "incomplete test found please make sure to address the following errors: [`missing request body values for required request body`]")
}

func TestPostMortems_PostV1PostMortemsReportsReportIDPublish(t *testing.T) {
	assert.Fail(t, "incomplete test found please make sure to address the following errors: [`missing request body values for required request body`]")
}

func TestPostMortems_PutV1PostMortemsQuestions(t *testing.T) {
	assert.Fail(t, "incomplete test found please make sure to address the following errors: [`missing request body values for required request body`]")
}

func TestPostMortems_GetV1PostMortemsQuestions(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.PostMortems.ListQuestions(ctx, nil, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}

func TestPostMortems_GetV1PostMortemsQuestionsQuestionID(t *testing.T) {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.PostMortems.GetQuestion(ctx, "<id>")
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
}
