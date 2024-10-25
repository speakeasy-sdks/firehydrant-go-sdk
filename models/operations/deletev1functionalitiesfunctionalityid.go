// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type DeleteV1FunctionalitiesFunctionalityIDRequest struct {
	FunctionalityID string `pathParam:"style=simple,explode=false,name=functionality_id"`
}

func (o *DeleteV1FunctionalitiesFunctionalityIDRequest) GetFunctionalityID() string {
	if o == nil {
		return ""
	}
	return o.FunctionalityID
}

type DeleteV1FunctionalitiesFunctionalityIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Archive a functionality
	FunctionalityEntity *components.FunctionalityEntity
}

func (o *DeleteV1FunctionalitiesFunctionalityIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteV1FunctionalitiesFunctionalityIDResponse) GetFunctionalityEntity() *components.FunctionalityEntity {
	if o == nil {
		return nil
	}
	return o.FunctionalityEntity
}
