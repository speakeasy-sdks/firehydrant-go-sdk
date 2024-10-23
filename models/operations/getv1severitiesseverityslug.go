// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type GetV1SeveritiesSeveritySlugRequest struct {
	SeveritySlug string `pathParam:"style=simple,explode=false,name=severity_slug"`
}

func (o *GetV1SeveritiesSeveritySlugRequest) GetSeveritySlug() string {
	if o == nil {
		return ""
	}
	return o.SeveritySlug
}

type GetV1SeveritiesSeveritySlugResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Retrieve a specific severity
	SeverityEntity *components.SeverityEntity
}

func (o *GetV1SeveritiesSeveritySlugResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetV1SeveritiesSeveritySlugResponse) GetSeverityEntity() *components.SeverityEntity {
	if o == nil {
		return nil
	}
	return o.SeverityEntity
}