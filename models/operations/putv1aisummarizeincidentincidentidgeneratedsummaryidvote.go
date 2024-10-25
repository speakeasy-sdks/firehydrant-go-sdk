// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"firehydrant/models/components"
	"fmt"
)

type Direction string

const (
	DirectionUp   Direction = "up"
	DirectionDown Direction = "down"
)

func (e Direction) ToPointer() *Direction {
	return &e
}
func (e *Direction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "up":
		fallthrough
	case "down":
		*e = Direction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Direction: %v", v)
	}
}

type PutV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVoteRequestBody struct {
	Direction Direction `form:"name=direction"`
}

func (o *PutV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVoteRequestBody) GetDirection() Direction {
	if o == nil {
		return Direction("")
	}
	return o.Direction
}

type PutV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVoteRequest struct {
	IncidentID         string                                                              `pathParam:"style=simple,explode=false,name=incident_id"`
	GeneratedSummaryID string                                                              `pathParam:"style=simple,explode=false,name=generated_summary_id"`
	RequestBody        PutV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVoteRequestBody `request:"mediaType=application/x-www-form-urlencoded"`
}

func (o *PutV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVoteRequest) GetIncidentID() string {
	if o == nil {
		return ""
	}
	return o.IncidentID
}

func (o *PutV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVoteRequest) GetGeneratedSummaryID() string {
	if o == nil {
		return ""
	}
	return o.GeneratedSummaryID
}

func (o *PutV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVoteRequest) GetRequestBody() PutV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVoteRequestBody {
	if o == nil {
		return PutV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVoteRequestBody{}
	}
	return o.RequestBody
}

type PutV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVoteResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *PutV1AiSummarizeIncidentIncidentIDGeneratedSummaryIDVoteResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
