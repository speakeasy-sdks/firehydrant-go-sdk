// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
)

// NotFound - Status codes relating to the resource/entity they are requesting not being found or endpoints/routes not existing
type NotFound struct {
	Message              *string        `json:"message,omitempty"`
	AdditionalProperties map[string]any `additionalProperties:"true" json:"-"`
}

var _ error = &NotFound{}

func (e *NotFound) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}