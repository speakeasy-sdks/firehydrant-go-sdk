// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/internal/utils"
	"openapi/models/components"
	"openapi/types"
)

type GetV1MetricsRetrospectivesRequest struct {
	// The start date to return metrics from
	StartDate *types.Date `queryParam:"style=form,explode=true,name=start_date"`
	// The end date to return metrics from
	EndDate *types.Date `queryParam:"style=form,explode=true,name=end_date"`
}

func (g GetV1MetricsRetrospectivesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetV1MetricsRetrospectivesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetV1MetricsRetrospectivesRequest) GetStartDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *GetV1MetricsRetrospectivesRequest) GetEndDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.EndDate
}

type GetV1MetricsRetrospectivesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns a report with retrospective analytics data
	MetricsRetrospectiveEntity *components.MetricsRetrospectiveEntity
}

func (o *GetV1MetricsRetrospectivesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetV1MetricsRetrospectivesResponse) GetMetricsRetrospectiveEntity() *components.MetricsRetrospectiveEntity {
	if o == nil {
		return nil
	}
	return o.MetricsRetrospectiveEntity
}