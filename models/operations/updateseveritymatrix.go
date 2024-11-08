// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type UpdateSeverityMatrixResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Update available severities and impacts in your organization's severity matrix.
	SeverityMatrixSeverityMatrixEntity *components.SeverityMatrixSeverityMatrixEntity
}

func (o *UpdateSeverityMatrixResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateSeverityMatrixResponse) GetSeverityMatrixSeverityMatrixEntity() *components.SeverityMatrixSeverityMatrixEntity {
	if o == nil {
		return nil
	}
	return o.SeverityMatrixSeverityMatrixEntity
}