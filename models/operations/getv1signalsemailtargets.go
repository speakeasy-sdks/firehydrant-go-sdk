// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type GetV1SignalsEmailTargetsRequest struct {
	// A query string to search the list of targets by.
	Query *string `queryParam:"style=form,explode=true,name=query"`
}

func (o *GetV1SignalsEmailTargetsRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

type GetV1SignalsEmailTargetsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *GetV1SignalsEmailTargetsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}