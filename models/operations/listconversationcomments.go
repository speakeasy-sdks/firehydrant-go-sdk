// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"firehydrant/internal/utils"
	"firehydrant/models/components"
	"fmt"
	"time"
)

// Sort - Allows sorting comments by the time they were posted, ascending or descending.
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

type ListConversationCommentsRequest struct {
	// An ISO8601 timestamp that allows filtering for comments posted before the provided time.
	Before *time.Time `queryParam:"style=form,explode=true,name=before"`
	// An ISO8601 timestamp that allows filtering for comments posted after the provided time.
	After *time.Time `queryParam:"style=form,explode=true,name=after"`
	// Allows sorting comments by the time they were posted, ascending or descending.
	Sort           *Sort  `default:"asc" queryParam:"style=form,explode=true,name=sort"`
	ConversationID string `pathParam:"style=simple,explode=false,name=conversation_id"`
}

func (l ListConversationCommentsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListConversationCommentsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListConversationCommentsRequest) GetBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.Before
}

func (o *ListConversationCommentsRequest) GetAfter() *time.Time {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *ListConversationCommentsRequest) GetSort() *Sort {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *ListConversationCommentsRequest) GetConversationID() string {
	if o == nil {
		return ""
	}
	return o.ConversationID
}

type ListConversationCommentsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *ListConversationCommentsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}