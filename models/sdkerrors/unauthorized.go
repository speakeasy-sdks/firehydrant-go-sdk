// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
)

// Unauthorized - A collection of codes that generally means the client was not authenticated correctly for the request they want to make
type Unauthorized struct {
	Message              *string        `json:"message,omitempty"`
	AdditionalProperties map[string]any `additionalProperties:"true" json:"-"`
}

var _ error = &Unauthorized{}

func (e *Unauthorized) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
