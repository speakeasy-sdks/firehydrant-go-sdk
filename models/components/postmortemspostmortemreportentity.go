// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"openapi/internal/utils"
	"time"
)

// PostMortemsPostMortemReportEntity - PostMortems_PostMortemReportEntity model
type PostMortemsPostMortemReportEntity struct {
	ID                *string    `json:"id,omitempty"`
	Name              *string    `json:"name,omitempty"`
	Summary           *string    `json:"summary,omitempty"`
	IncidentID        *string    `json:"incident_id,omitempty"`
	CreatedAt         *time.Time `json:"created_at,omitempty"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty"`
	TagList           []string   `json:"tag_list,omitempty"`
	AdditionalDetails []string   `json:"additional_details,omitempty"`
	// IncidentEntity model
	Incident       *IncidentEntity            `json:"incident,omitempty"`
	Questions      *PostMortemsQuestionEntity `json:"questions,omitempty"`
	CalendarEvents *CalendarsEventEntity      `json:"calendar_events,omitempty"`
}

func (p PostMortemsPostMortemReportEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PostMortemsPostMortemReportEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PostMortemsPostMortemReportEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PostMortemsPostMortemReportEntity) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *PostMortemsPostMortemReportEntity) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *PostMortemsPostMortemReportEntity) GetIncidentID() *string {
	if o == nil {
		return nil
	}
	return o.IncidentID
}

func (o *PostMortemsPostMortemReportEntity) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *PostMortemsPostMortemReportEntity) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *PostMortemsPostMortemReportEntity) GetTagList() []string {
	if o == nil {
		return nil
	}
	return o.TagList
}

func (o *PostMortemsPostMortemReportEntity) GetAdditionalDetails() []string {
	if o == nil {
		return nil
	}
	return o.AdditionalDetails
}

func (o *PostMortemsPostMortemReportEntity) GetIncident() *IncidentEntity {
	if o == nil {
		return nil
	}
	return o.Incident
}

func (o *PostMortemsPostMortemReportEntity) GetQuestions() *PostMortemsQuestionEntity {
	if o == nil {
		return nil
	}
	return o.Questions
}

func (o *PostMortemsPostMortemReportEntity) GetCalendarEvents() *CalendarsEventEntity {
	if o == nil {
		return nil
	}
	return o.CalendarEvents
}
