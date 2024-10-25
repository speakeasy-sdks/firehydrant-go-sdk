// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type GetV1TeamsTeamIDOnCallSchedulesRequest struct {
	TeamID string `pathParam:"style=simple,explode=false,name=team_id"`
	// A query string for searching through the list of on-call schedules.
	Query   *string `queryParam:"style=form,explode=true,name=query"`
	Page    *int    `queryParam:"style=form,explode=true,name=page"`
	PerPage *int    `queryParam:"style=form,explode=true,name=per_page"`
}

func (o *GetV1TeamsTeamIDOnCallSchedulesRequest) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *GetV1TeamsTeamIDOnCallSchedulesRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *GetV1TeamsTeamIDOnCallSchedulesRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetV1TeamsTeamIDOnCallSchedulesRequest) GetPerPage() *int {
	if o == nil {
		return nil
	}
	return o.PerPage
}

type GetV1TeamsTeamIDOnCallSchedulesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *GetV1TeamsTeamIDOnCallSchedulesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
