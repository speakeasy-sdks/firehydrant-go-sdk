// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type GetV1FormConfigurationsSlugRequest struct {
	Slug string `pathParam:"style=simple,explode=false,name=slug"`
}

func (o *GetV1FormConfigurationsSlugRequest) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

type GetV1FormConfigurationsSlugResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *GetV1FormConfigurationsSlugResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
