// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type PatchV1TeamsTeamIDRequest struct {
	TeamID             string                        `pathParam:"style=simple,explode=false,name=team_id"`
	PatchV1TeamsTeamID components.PatchV1TeamsTeamID `request:"mediaType=application/json"`
}

func (o *PatchV1TeamsTeamIDRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *PatchV1TeamsTeamIDRequest) GetPatchV1TeamsTeamID() components.PatchV1TeamsTeamID {
	if o == nil {
		return components.PatchV1TeamsTeamID{}
	}
	return o.PatchV1TeamsTeamID
}

type PatchV1TeamsTeamIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Update a single team from its ID
	TeamEntity *components.TeamEntity
}

func (o *PatchV1TeamsTeamIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PatchV1TeamsTeamIDResponse) GetTeamEntity() *components.TeamEntity {
	if o == nil {
		return nil
	}
	return o.TeamEntity
}
