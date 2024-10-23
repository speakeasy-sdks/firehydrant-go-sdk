// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type GetV1PostMortemsQuestionsRequest struct {
	Page    *int `queryParam:"style=form,explode=true,name=page"`
	PerPage *int `queryParam:"style=form,explode=true,name=per_page"`
}

func (o *GetV1PostMortemsQuestionsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetV1PostMortemsQuestionsRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type GetV1PostMortemsQuestionsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// List the questions configured to be provided and filled out on each retrospective report.
	PostMortemsQuestionTypeEntityPaginated *components.PostMortemsQuestionTypeEntityPaginated
}

func (o *GetV1PostMortemsQuestionsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetV1PostMortemsQuestionsResponse) GetPostMortemsQuestionTypeEntityPaginated() *components.PostMortemsQuestionTypeEntityPaginated {
	if o == nil {
		return nil
	}
	return o.PostMortemsQuestionTypeEntityPaginated
}
