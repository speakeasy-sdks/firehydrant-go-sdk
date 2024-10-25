// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"firehydrant/internal/utils"
	"time"
)

// EventGenericChatMessageEntity - Event_GenericChatMessageEntity model
type EventGenericChatMessageEntity struct {
	ID        *string    `json:"id,omitempty"`
	Body      *string    `json:"body,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

func (e EventGenericChatMessageEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *EventGenericChatMessageEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *EventGenericChatMessageEntity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *EventGenericChatMessageEntity) GetBody() *string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *EventGenericChatMessageEntity) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}
