// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type PostV1TeamsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Create a new team
	TeamEntity *components.TeamEntity
}

func (o *PostV1TeamsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PostV1TeamsResponse) GetTeamEntity() *components.TeamEntity {
	if o == nil {
		return nil
	}
	return o.TeamEntity
}
