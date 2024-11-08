// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type DeleteServiceDependencyRequest struct {
	ServiceDependencyID string `pathParam:"style=simple,explode=false,name=service_dependency_id"`
}

func (o *DeleteServiceDependencyRequest) GetServiceDependencyID() string {
	if o == nil {
		return ""
	}
	return o.ServiceDependencyID
}

type DeleteServiceDependencyResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Deletes a single service dependency
	ServiceDependencyEntity *components.ServiceDependencyEntity
}

func (o *DeleteServiceDependencyResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteServiceDependencyResponse) GetServiceDependencyEntity() *components.ServiceDependencyEntity {
	if o == nil {
		return nil
	}
	return o.ServiceDependencyEntity
}