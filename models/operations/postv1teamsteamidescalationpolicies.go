// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type PostV1TeamsTeamIDEscalationPoliciesRequest struct {
	TeamID                              string                                         `pathParam:"style=simple,explode=false,name=team_id"`
	PostV1TeamsTeamIDEscalationPolicies components.PostV1TeamsTeamIDEscalationPolicies `request:"mediaType=application/json"`
}

func (o *PostV1TeamsTeamIDEscalationPoliciesRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *PostV1TeamsTeamIDEscalationPoliciesRequest) GetPostV1TeamsTeamIDEscalationPolicies() components.PostV1TeamsTeamIDEscalationPolicies {
	if o == nil {
		return components.PostV1TeamsTeamIDEscalationPolicies{}
	}
	return o.PostV1TeamsTeamIDEscalationPolicies
}

type PostV1TeamsTeamIDEscalationPoliciesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *PostV1TeamsTeamIDEscalationPoliciesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
