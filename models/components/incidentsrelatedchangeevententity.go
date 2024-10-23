// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"openapi/internal/utils"
	"time"
)

type IncidentsRelatedChangeEventEntityType string

const (
	IncidentsRelatedChangeEventEntityTypeCaused    IncidentsRelatedChangeEventEntityType = "caused"
	IncidentsRelatedChangeEventEntityTypeFixed     IncidentsRelatedChangeEventEntityType = "fixed"
	IncidentsRelatedChangeEventEntityTypeSuspect   IncidentsRelatedChangeEventEntityType = "suspect"
	IncidentsRelatedChangeEventEntityTypeDismissed IncidentsRelatedChangeEventEntityType = "dismissed"
)

func (e IncidentsRelatedChangeEventEntityType) ToPointer() *IncidentsRelatedChangeEventEntityType {
	return &e
}
func (e *IncidentsRelatedChangeEventEntityType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "caused":
		fallthrough
	case "fixed":
		fallthrough
	case "suspect":
		fallthrough
	case "dismissed":
		*e = IncidentsRelatedChangeEventEntityType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for IncidentsRelatedChangeEventEntityType: %v", v)
	}
}

// IncidentsRelatedChangeEventEntity - Incidents_RelatedChangeEventEntity model
type IncidentsRelatedChangeEventEntity struct {
	ID        *string    `json:"id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The reason why this change event is related to this incident
	Why  *string                                `json:"why,omitempty"`
	Type *IncidentsRelatedChangeEventEntityType `json:"type,omitempty"`
	// ChangeEventEntity model
	ChangeEvent *ChangeEventEntity `json:"change_event,omitempty"`
	IncidentID  *string            `json:"incident_id,omitempty"`
	CreatedBy   *AuthorEntity      `json:"created_by,omitempty"`
}

func (i IncidentsRelatedChangeEventEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *IncidentsRelatedChangeEventEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *IncidentsRelatedChangeEventEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *IncidentsRelatedChangeEventEntity) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *IncidentsRelatedChangeEventEntity) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *IncidentsRelatedChangeEventEntity) GetWhy() *string {
	if o == nil {
		return nil
	}
	return o.Why
}

func (o *IncidentsRelatedChangeEventEntity) GetType() *IncidentsRelatedChangeEventEntityType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *IncidentsRelatedChangeEventEntity) GetChangeEvent() *ChangeEventEntity {
	if o == nil {
		return nil
	}
	return o.ChangeEvent
}

func (o *IncidentsRelatedChangeEventEntity) GetIncidentID() *string {
	if o == nil {
		return nil
	}
	return o.IncidentID
}

func (o *IncidentsRelatedChangeEventEntity) GetCreatedBy() *AuthorEntity {
	if o == nil {
		return nil
	}
	return o.CreatedBy
}
