// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"firehydrant/internal/utils"
	"time"
)

type IncidentsLifecycleMeasurementEntity struct {
	ID                *string    `json:"id,omitempty"`
	Name              *string    `json:"name,omitempty"`
	Description       *string    `json:"description,omitempty"`
	Slug              *string    `json:"slug,omitempty"`
	StartsAtMilestone *string    `json:"starts_at_milestone,omitempty"`
	EndsAtMilestone   *string    `json:"ends_at_milestone,omitempty"`
	Value             *string    `json:"value,omitempty"`
	CalculatedAt      *time.Time `json:"calculated_at,omitempty"`
}

func (i IncidentsLifecycleMeasurementEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *IncidentsLifecycleMeasurementEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *IncidentsLifecycleMeasurementEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *IncidentsLifecycleMeasurementEntity) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *IncidentsLifecycleMeasurementEntity) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *IncidentsLifecycleMeasurementEntity) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *IncidentsLifecycleMeasurementEntity) GetStartsAtMilestone() *string {
	if o == nil {
		return nil
	}
	return o.StartsAtMilestone
}

func (o *IncidentsLifecycleMeasurementEntity) GetEndsAtMilestone() *string {
	if o == nil {
		return nil
	}
	return o.EndsAtMilestone
}

func (o *IncidentsLifecycleMeasurementEntity) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

func (o *IncidentsLifecycleMeasurementEntity) GetCalculatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CalculatedAt
}
