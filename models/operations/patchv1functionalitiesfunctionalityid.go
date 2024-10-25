// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type PatchV1FunctionalitiesFunctionalityIDRequest struct {
	FunctionalityID                       string                                           `pathParam:"style=simple,explode=false,name=functionality_id"`
	PatchV1FunctionalitiesFunctionalityID components.PatchV1FunctionalitiesFunctionalityID `request:"mediaType=application/json"`
}

func (o *PatchV1FunctionalitiesFunctionalityIDRequest) GetFunctionalityID() string {
	if o == nil {
		return ""
	}
	return o.FunctionalityID
}

func (o *PatchV1FunctionalitiesFunctionalityIDRequest) GetPatchV1FunctionalitiesFunctionalityID() components.PatchV1FunctionalitiesFunctionalityID {
	if o == nil {
		return components.PatchV1FunctionalitiesFunctionalityID{}
	}
	return o.PatchV1FunctionalitiesFunctionalityID
}

type PatchV1FunctionalitiesFunctionalityIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Update a functionalities attributes
	FunctionalityEntity *components.FunctionalityEntity
}

func (o *PatchV1FunctionalitiesFunctionalityIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PatchV1FunctionalitiesFunctionalityIDResponse) GetFunctionalityEntity() *components.FunctionalityEntity {
	if o == nil {
		return nil
	}
	return o.FunctionalityEntity
}
