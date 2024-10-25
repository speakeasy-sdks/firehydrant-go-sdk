// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type PostV1SeveritiesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Create a new severity
	SeverityEntity *components.SeverityEntity
}

func (o *PostV1SeveritiesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PostV1SeveritiesResponse) GetSeverityEntity() *components.SeverityEntity {
	if o == nil {
		return nil
	}
	return o.SeverityEntity
}
