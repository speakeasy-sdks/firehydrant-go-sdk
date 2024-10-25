// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type GetV1ScimV2GroupsRequest struct {
	StartIndex *int `queryParam:"style=form,explode=true,name=startIndex"`
	Count      *int `queryParam:"style=form,explode=true,name=count"`
	// This is a string used to query groups by displayName.
	//
	//         Proper example syntax for this would be `?filter=displayName eq "My Team Name"`.
	//         Currently we only support the `eq` operator
	Filter *string `queryParam:"style=form,explode=true,name=filter"`
}

func (o *GetV1ScimV2GroupsRequest) GetStartIndex() *int {
	if o == nil {
		return nil
	}
	return o.StartIndex
}

func (o *GetV1ScimV2GroupsRequest) GetCount() *int {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *GetV1ScimV2GroupsRequest) GetFilter() *string {
	if o == nil {
		return nil
	}
	return o.Filter
}

type GetV1ScimV2GroupsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *GetV1ScimV2GroupsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
