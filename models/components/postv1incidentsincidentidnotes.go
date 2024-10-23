// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
	"openapi/internal/utils"
	"time"
)

type PostV1IncidentsIncidentIDNotesVisibility string

const (
	PostV1IncidentsIncidentIDNotesVisibilityPrivateToOrg       PostV1IncidentsIncidentIDNotesVisibility = "private_to_org"
	PostV1IncidentsIncidentIDNotesVisibilityOpenToPublic       PostV1IncidentsIncidentIDNotesVisibility = "open_to_public"
	PostV1IncidentsIncidentIDNotesVisibilityInternalStatusPage PostV1IncidentsIncidentIDNotesVisibility = "internal_status_page"
)

func (e PostV1IncidentsIncidentIDNotesVisibility) ToPointer() *PostV1IncidentsIncidentIDNotesVisibility {
	return &e
}
func (e *PostV1IncidentsIncidentIDNotesVisibility) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "private_to_org":
		fallthrough
	case "open_to_public":
		fallthrough
	case "internal_status_page":
		*e = PostV1IncidentsIncidentIDNotesVisibility(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PostV1IncidentsIncidentIDNotesVisibility: %v", v)
	}
}

type PostV1IncidentsIncidentIDNotesStatusPages struct {
	ID              string `json:"id"`
	IntegrationSlug string `json:"integration_slug"`
}

func (o *PostV1IncidentsIncidentIDNotesStatusPages) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PostV1IncidentsIncidentIDNotesStatusPages) GetIntegrationSlug() string {
	if o == nil {
		return ""
	}
	return o.IntegrationSlug
}

// PostV1IncidentsIncidentIDNotes - Create a new note on for an incident. The visibility field on a note determines where it gets posted.
type PostV1IncidentsIncidentIDNotes struct {
	Body string `json:"body"`
	// ISO8601 timestamp for when the note occurred
	OccurredAt  *time.Time                                  `json:"occurred_at,omitempty"`
	Visibility  *PostV1IncidentsIncidentIDNotesVisibility   `default:"private_to_org" json:"visibility"`
	StatusPages []PostV1IncidentsIncidentIDNotesStatusPages `json:"status_pages,omitempty"`
}

func (p PostV1IncidentsIncidentIDNotes) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PostV1IncidentsIncidentIDNotes) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PostV1IncidentsIncidentIDNotes) GetBody() string {
	if o == nil {
		return ""
	}
	return o.Body
}

func (o *PostV1IncidentsIncidentIDNotes) GetOccurredAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.OccurredAt
}

func (o *PostV1IncidentsIncidentIDNotes) GetVisibility() *PostV1IncidentsIncidentIDNotesVisibility {
	if o == nil {
		return nil
	}
	return o.Visibility
}

func (o *PostV1IncidentsIncidentIDNotes) GetStatusPages() []PostV1IncidentsIncidentIDNotesStatusPages {
	if o == nil {
		return nil
	}
	return o.StatusPages
}