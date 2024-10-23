// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"openapi/internal/utils"
	"time"
)

type MetricsMilestonesFunnelEntityDataBucketEntity struct {
	// The start datetime for the period
	TimeBucket      *time.Time                                                    `json:"time_bucket,omitempty"`
	FilterParams    *MetricsMilestonesFunnelEntityDataBucketFilterParamsEntity    `json:"filter_params,omitempty"`
	MilestoneCounts []MetricsMilestonesFunnelEntityDataBucketMilestoneCountEntity `json:"milestone_counts,omitempty"`
}

func (m MetricsMilestonesFunnelEntityDataBucketEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MetricsMilestonesFunnelEntityDataBucketEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MetricsMilestonesFunnelEntityDataBucketEntity) GetTimeBucket() *time.Time {
	if o == nil {
		return nil
	}
	return o.TimeBucket
}

func (o *MetricsMilestonesFunnelEntityDataBucketEntity) GetFilterParams() *MetricsMilestonesFunnelEntityDataBucketFilterParamsEntity {
	if o == nil {
		return nil
	}
	return o.FilterParams
}

func (o *MetricsMilestonesFunnelEntityDataBucketEntity) GetMilestoneCounts() []MetricsMilestonesFunnelEntityDataBucketMilestoneCountEntity {
	if o == nil {
		return nil
	}
	return o.MilestoneCounts
}
