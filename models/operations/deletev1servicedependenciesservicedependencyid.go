// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type DeleteV1ServiceDependenciesServiceDependencyIDRequest struct {
	ServiceDependencyID string `pathParam:"style=simple,explode=false,name=service_dependency_id"`
}

func (o *DeleteV1ServiceDependenciesServiceDependencyIDRequest) GetServiceDependencyID() string {
	if o == nil {
		return ""
	}
	return o.ServiceDependencyID
}

type DeleteV1ServiceDependenciesServiceDependencyIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Deletes a single service dependency
	ServiceDependencyEntity *components.ServiceDependencyEntity
}

func (o *DeleteV1ServiceDependenciesServiceDependencyIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteV1ServiceDependenciesServiceDependencyIDResponse) GetServiceDependencyEntity() *components.ServiceDependencyEntity {
	if o == nil {
		return nil
	}
	return o.ServiceDependencyEntity
}
