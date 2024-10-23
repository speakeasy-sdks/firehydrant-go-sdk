// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"openapi/internal/utils"
	"openapi/models/components"
)

// Sort - The order to sort the transcript entries.
type Sort string

const (
	SortAsc  Sort = "asc"
	SortDesc Sort = "desc"
)

func (e Sort) ToPointer() *Sort {
	return &e
}
func (e *Sort) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "asc":
		fallthrough
	case "desc":
		*e = Sort(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Sort: %v", v)
	}
}

type GetV1IncidentsIncidentIDTranscriptRequest struct {
	// The ID of the transcript entry to start after.
	After *string `queryParam:"style=form,explode=true,name=after"`
	// The ID of the transcript entry to start before.
	Before *string `queryParam:"style=form,explode=true,name=before"`
	// The order to sort the transcript entries.
	Sort       *Sort  `default:"asc" queryParam:"style=form,explode=true,name=sort"`
	IncidentID string `pathParam:"style=simple,explode=false,name=incident_id"`
}

func (g GetV1IncidentsIncidentIDTranscriptRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetV1IncidentsIncidentIDTranscriptRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetV1IncidentsIncidentIDTranscriptRequest) GetAfter() *string {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *GetV1IncidentsIncidentIDTranscriptRequest) GetBefore() *string {
	if o == nil {
		return nil
	}
	return o.Before
}

func (o *GetV1IncidentsIncidentIDTranscriptRequest) GetSort() *Sort {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *GetV1IncidentsIncidentIDTranscriptRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

type GetV1IncidentsIncidentIDTranscriptResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Retrieve the transcript for a specific incident
	PublicAPIV1IncidentsTranscriptEntity *components.PublicAPIV1IncidentsTranscriptEntity
}

func (o *GetV1IncidentsIncidentIDTranscriptResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetV1IncidentsIncidentIDTranscriptResponse) GetPublicAPIV1IncidentsTranscriptEntity() *components.PublicAPIV1IncidentsTranscriptEntity {
	if o == nil {
		return nil
	}
	return o.PublicAPIV1IncidentsTranscriptEntity
}