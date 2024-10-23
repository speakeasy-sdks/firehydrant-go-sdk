// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type GroupBy string

const (
	GroupByStartedDay   GroupBy = "started_day"
	GroupByStartedWeek  GroupBy = "started_week"
	GroupByStartedMonth GroupBy = "started_month"
	GroupByAllTime      GroupBy = "all_time"
)

func (e GroupBy) ToPointer() *GroupBy {
	return &e
}
func (e *GroupBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "started_day":
		fallthrough
	case "started_week":
		fallthrough
	case "started_month":
		fallthrough
	case "all_time":
		*e = GroupBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GroupBy: %v", v)
	}
}

type GetV1MetricsTicketFunnel struct {
	GroupBy []GroupBy `form:"name=group_by"`
}

func (o *GetV1MetricsTicketFunnel) GetGroupBy() []GroupBy {
	if o == nil {
		return nil
	}
	return o.GroupBy
}