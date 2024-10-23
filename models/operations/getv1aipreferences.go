// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type GetV1AiPreferencesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Retrieves the current AI preferences
	AIEntitiesPreferencesEntity *components.AIEntitiesPreferencesEntity
}

func (o *GetV1AiPreferencesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetV1AiPreferencesResponse) GetAIEntitiesPreferencesEntity() *components.AIEntitiesPreferencesEntity {
	if o == nil {
		return nil
	}
	return o.AIEntitiesPreferencesEntity
}