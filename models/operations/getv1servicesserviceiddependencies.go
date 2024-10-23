// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type GetV1ServicesServiceIDDependenciesRequest struct {
	ServiceID string `pathParam:"style=simple,explode=false,name=service_id"`
	// If true, returns all dependencies in one array. If false, splits dependencies into different arrays for child and parent dependencies
	Flatten *bool `queryParam:"style=form,explode=true,name=flatten"`
}

func (o *GetV1ServicesServiceIDDependenciesRequest) GetServiceID() string {
	if o == nil {
		return ""
	}
	return o.ServiceID
}

func (o *GetV1ServicesServiceIDDependenciesRequest) GetFlatten() *bool {
	if o == nil {
		return nil
	}
	return o.Flatten
}

type GetV1ServicesServiceIDDependenciesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Retrieves a service's dependencies
	ServiceWithAllDependenciesEntity *components.ServiceWithAllDependenciesEntity
}

func (o *GetV1ServicesServiceIDDependenciesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetV1ServicesServiceIDDependenciesResponse) GetServiceWithAllDependenciesEntity() *components.ServiceWithAllDependenciesEntity {
	if o == nil {
		return nil
	}
	return o.ServiceWithAllDependenciesEntity
}
