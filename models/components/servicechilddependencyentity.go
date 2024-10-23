// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"openapi/internal/utils"
	"time"
)

type ServiceChildDependencyEntity struct {
	ID        *string    `json:"id,omitempty"`
	Notes     *string    `json:"notes,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// ServiceEntity model
	Service *ServiceEntity `json:"service,omitempty"`
	Type    *string        `json:"type,omitempty"`
}

func (s ServiceChildDependencyEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *ServiceChildDependencyEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ServiceChildDependencyEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ServiceChildDependencyEntity) GetNotes() *string {
	if o == nil {
		return nil
	}
	return o.Notes
}

func (o *ServiceChildDependencyEntity) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *ServiceChildDependencyEntity) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *ServiceChildDependencyEntity) GetService() *ServiceEntity {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *ServiceChildDependencyEntity) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}
