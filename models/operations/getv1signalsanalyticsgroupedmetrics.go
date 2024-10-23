// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"openapi/internal/utils"
	"openapi/models/components"
	"time"
)

// GetV1SignalsAnalyticsGroupedMetricsQueryParamGroupBy - String that determines how records are grouped
type GetV1SignalsAnalyticsGroupedMetricsQueryParamGroupBy string

const (
	GetV1SignalsAnalyticsGroupedMetricsQueryParamGroupBySignalRules  GetV1SignalsAnalyticsGroupedMetricsQueryParamGroupBy = "signal_rules"
	GetV1SignalsAnalyticsGroupedMetricsQueryParamGroupByTeams        GetV1SignalsAnalyticsGroupedMetricsQueryParamGroupBy = "teams"
	GetV1SignalsAnalyticsGroupedMetricsQueryParamGroupByServices     GetV1SignalsAnalyticsGroupedMetricsQueryParamGroupBy = "services"
	GetV1SignalsAnalyticsGroupedMetricsQueryParamGroupByEnvironments GetV1SignalsAnalyticsGroupedMetricsQueryParamGroupBy = "environments"
	GetV1SignalsAnalyticsGroupedMetricsQueryParamGroupByTags         GetV1SignalsAnalyticsGroupedMetricsQueryParamGroupBy = "tags"
)

func (e GetV1SignalsAnalyticsGroupedMetricsQueryParamGroupBy) ToPointer() *GetV1SignalsAnalyticsGroupedMetricsQueryParamGroupBy {
	return &e
}
func (e *GetV1SignalsAnalyticsGroupedMetricsQueryParamGroupBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "signal_rules":
		fallthrough
	case "teams":
		fallthrough
	case "services":
		fallthrough
	case "environments":
		fallthrough
	case "tags":
		*e = GetV1SignalsAnalyticsGroupedMetricsQueryParamGroupBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetV1SignalsAnalyticsGroupedMetricsQueryParamGroupBy: %v", v)
	}
}

// GetV1SignalsAnalyticsGroupedMetricsQueryParamSortBy - String that determines how records are sorted
type GetV1SignalsAnalyticsGroupedMetricsQueryParamSortBy string

const (
	GetV1SignalsAnalyticsGroupedMetricsQueryParamSortByTotalOpenedAlerts   GetV1SignalsAnalyticsGroupedMetricsQueryParamSortBy = "total_opened_alerts"
	GetV1SignalsAnalyticsGroupedMetricsQueryParamSortByTotalAckedAlerts    GetV1SignalsAnalyticsGroupedMetricsQueryParamSortBy = "total_acked_alerts"
	GetV1SignalsAnalyticsGroupedMetricsQueryParamSortByTotalIncidents      GetV1SignalsAnalyticsGroupedMetricsQueryParamSortBy = "total_incidents"
	GetV1SignalsAnalyticsGroupedMetricsQueryParamSortByAckedPercentage     GetV1SignalsAnalyticsGroupedMetricsQueryParamSortBy = "acked_percentage"
	GetV1SignalsAnalyticsGroupedMetricsQueryParamSortByIncidentsPercentage GetV1SignalsAnalyticsGroupedMetricsQueryParamSortBy = "incidents_percentage"
)

func (e GetV1SignalsAnalyticsGroupedMetricsQueryParamSortBy) ToPointer() *GetV1SignalsAnalyticsGroupedMetricsQueryParamSortBy {
	return &e
}
func (e *GetV1SignalsAnalyticsGroupedMetricsQueryParamSortBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "total_opened_alerts":
		fallthrough
	case "total_acked_alerts":
		fallthrough
	case "total_incidents":
		fallthrough
	case "acked_percentage":
		fallthrough
	case "incidents_percentage":
		*e = GetV1SignalsAnalyticsGroupedMetricsQueryParamSortBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetV1SignalsAnalyticsGroupedMetricsQueryParamSortBy: %v", v)
	}
}

// GetV1SignalsAnalyticsGroupedMetricsQueryParamSortDirection - String that determines how records are sorted
type GetV1SignalsAnalyticsGroupedMetricsQueryParamSortDirection string

const (
	GetV1SignalsAnalyticsGroupedMetricsQueryParamSortDirectionAsc  GetV1SignalsAnalyticsGroupedMetricsQueryParamSortDirection = "asc"
	GetV1SignalsAnalyticsGroupedMetricsQueryParamSortDirectionDesc GetV1SignalsAnalyticsGroupedMetricsQueryParamSortDirection = "desc"
)

func (e GetV1SignalsAnalyticsGroupedMetricsQueryParamSortDirection) ToPointer() *GetV1SignalsAnalyticsGroupedMetricsQueryParamSortDirection {
	return &e
}
func (e *GetV1SignalsAnalyticsGroupedMetricsQueryParamSortDirection) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = GetV1SignalsAnalyticsGroupedMetricsQueryParamSortDirection(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetV1SignalsAnalyticsGroupedMetricsQueryParamSortDirection: %v", v)
	}
}

type GetV1SignalsAnalyticsGroupedMetricsRequest struct {
	// A comma separated list of signal rule IDs
	SignalRules *string `queryParam:"style=form,explode=true,name=signal_rules"`
	// A comma separated list of team IDs
	Teams *string `queryParam:"style=form,explode=true,name=teams"`
	// A comma separated list of environment IDs
	Environments *string `queryParam:"style=form,explode=true,name=environments"`
	// A comma separated list of service IDs
	Services *string `queryParam:"style=form,explode=true,name=services"`
	// A comma separated list of tags
	Tags *string `queryParam:"style=form,explode=true,name=tags"`
	// A comma separated list of user IDs
	Users *string `queryParam:"style=form,explode=true,name=users"`
	// String that determines how records are grouped
	GroupBy *GetV1SignalsAnalyticsGroupedMetricsQueryParamGroupBy `queryParam:"style=form,explode=true,name=group_by"`
	// String that determines how records are sorted
	SortBy *GetV1SignalsAnalyticsGroupedMetricsQueryParamSortBy `queryParam:"style=form,explode=true,name=sort_by"`
	// String that determines how records are sorted
	SortDirection *GetV1SignalsAnalyticsGroupedMetricsQueryParamSortDirection `queryParam:"style=form,explode=true,name=sort_direction"`
	// The start date to return metrics from
	StartDate *time.Time `queryParam:"style=form,explode=true,name=start_date"`
	// The end date to return metrics from
	EndDate *time.Time `queryParam:"style=form,explode=true,name=end_date"`
}

func (g GetV1SignalsAnalyticsGroupedMetricsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetV1SignalsAnalyticsGroupedMetricsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetV1SignalsAnalyticsGroupedMetricsRequest) GetSignalRules() *string {
	if o == nil {
		return nil
	}
	return o.SignalRules
}

func (o *GetV1SignalsAnalyticsGroupedMetricsRequest) GetTeams() *string {
	if o == nil {
		return nil
	}
	return o.Teams
}

func (o *GetV1SignalsAnalyticsGroupedMetricsRequest) GetEnvironments() *string {
	if o == nil {
		return nil
	}
	return o.Environments
}

func (o *GetV1SignalsAnalyticsGroupedMetricsRequest) GetServices() *string {
	if o == nil {
		return nil
	}
	return o.Services
}

func (o *GetV1SignalsAnalyticsGroupedMetricsRequest) GetTags() *string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *GetV1SignalsAnalyticsGroupedMetricsRequest) GetUsers() *string {
	if o == nil {
		return nil
	}
	return o.Users
}

func (o *GetV1SignalsAnalyticsGroupedMetricsRequest) GetGroupBy() *GetV1SignalsAnalyticsGroupedMetricsQueryParamGroupBy {
	if o == nil {
		return nil
	}
	return o.GroupBy
}

func (o *GetV1SignalsAnalyticsGroupedMetricsRequest) GetSortBy() *GetV1SignalsAnalyticsGroupedMetricsQueryParamSortBy {
	if o == nil {
		return nil
	}
	return o.SortBy
}

func (o *GetV1SignalsAnalyticsGroupedMetricsRequest) GetSortDirection() *GetV1SignalsAnalyticsGroupedMetricsQueryParamSortDirection {
	if o == nil {
		return nil
	}
	return o.SortDirection
}

func (o *GetV1SignalsAnalyticsGroupedMetricsRequest) GetStartDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *GetV1SignalsAnalyticsGroupedMetricsRequest) GetEndDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndDate
}

type GetV1SignalsAnalyticsGroupedMetricsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *GetV1SignalsAnalyticsGroupedMetricsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
