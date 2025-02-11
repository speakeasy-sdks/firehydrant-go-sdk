// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type ListRunbookActionsRequest struct {
	Page    *int `queryParam:"style=form,explode=true,name=page"`
	PerPage *int `queryParam:"style=form,explode=true,name=per_page"`
	// List actions supporting this specific Runbook type
	Type *string `queryParam:"style=form,explode=true,name=type"`
	// Boolean to determine whether to return a slimified version of the action object's integration
	Lite *bool `queryParam:"style=form,explode=true,name=lite"`
}

func (o *ListRunbookActionsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *ListRunbookActionsRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *ListRunbookActionsRequest) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ListRunbookActionsRequest) GetLite() *bool {
	if o == nil {
		return nil
	}
	return o.Lite
}

type ListRunbookActionsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// List all Runbook actions available through your connected integrations
	RunbooksActionsEntityPaginated *components.RunbooksActionsEntityPaginated
}

func (o *ListRunbookActionsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListRunbookActionsResponse) GetRunbooksActionsEntityPaginated() *components.RunbooksActionsEntityPaginated {
	if o == nil {
		return nil
	}
	return o.RunbooksActionsEntityPaginated
}
