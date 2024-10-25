// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"firehydrant/internal/utils"
	"fmt"
	"time"
)

type ChangeIdentities struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func (o *ChangeIdentities) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *ChangeIdentities) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type PostV1ChangesEventsType string

const (
	PostV1ChangesEventsTypeLink PostV1ChangesEventsType = "link"
)

func (e PostV1ChangesEventsType) ToPointer() *PostV1ChangesEventsType {
	return &e
}
func (e *PostV1ChangesEventsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "link":
		*e = PostV1ChangesEventsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostV1ChangesEventsType: %v", v)
	}
}

type PostV1ChangesEventsAttachments struct {
	Type PostV1ChangesEventsType `json:"type"`
}

func (o *PostV1ChangesEventsAttachments) GetType() PostV1ChangesEventsType {
	if o == nil {
		return PostV1ChangesEventsType("")
	}
	return o.Type
}

type Authors struct {
	Source   string `json:"source"`
	SourceID string `json:"source_id"`
	Name     string `json:"name"`
}

func (o *Authors) GetSource() string {
	if o == nil {
		return ""
	}
	return o.Source
}

func (o *Authors) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

func (o *Authors) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

// PostV1ChangesEvents - Create a change event
type PostV1ChangesEvents struct {
	Summary     string            `json:"summary"`
	Description *string           `json:"description,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	StartsAt    *time.Time        `json:"starts_at,omitempty"`
	EndsAt      *time.Time        `json:"ends_at,omitempty"`
	// An array of environment IDs
	Environments []string `json:"environments,omitempty"`
	// An array of service IDs
	Services []string `json:"services,omitempty"`
	// An array of change IDs
	Changes []string `json:"changes,omitempty"`
	// The ID of a change event as assigned by an external provider
	ExternalID *string `json:"external_id,omitempty"`
	// If provided and valid, the event will be linked to all changes that have the same identities. Identity *values* must be unique.
	ChangeIdentities []ChangeIdentities `json:"change_identities,omitempty"`
	// JSON objects representing attachments, see attachments documentation for the schema
	Attachments []PostV1ChangesEventsAttachments `json:"attachments,omitempty"`
	// Array of additional authors to add to the change event, the creating actor will automatically be added as an author
	Authors []Authors `json:"authors,omitempty"`
}

func (p PostV1ChangesEvents) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PostV1ChangesEvents) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PostV1ChangesEvents) GetSummary() string {
	if o == nil {
		return ""
	}
	return o.Summary
}

func (o *PostV1ChangesEvents) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *PostV1ChangesEvents) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *PostV1ChangesEvents) GetStartsAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartsAt
}

func (o *PostV1ChangesEvents) GetEndsAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndsAt
}

func (o *PostV1ChangesEvents) GetEnvironments() []string {
	if o == nil {
		return nil
	}
	return o.Environments
}

func (o *PostV1ChangesEvents) GetServices() []string {
	if o == nil {
		return nil
	}
	return o.Services
}

func (o *PostV1ChangesEvents) GetChanges() []string {
	if o == nil {
		return nil
	}
	return o.Changes
}

func (o *PostV1ChangesEvents) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

func (o *PostV1ChangesEvents) GetChangeIdentities() []ChangeIdentities {
	if o == nil {
		return nil
	}
	return o.ChangeIdentities
}

func (o *PostV1ChangesEvents) GetAttachments() []PostV1ChangesEventsAttachments {
	if o == nil {
		return nil
	}
	return o.Attachments
}

func (o *PostV1ChangesEvents) GetAuthors() []Authors {
	if o == nil {
		return nil
	}
	return o.Authors
}
