// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
)

// BadRequest - A collection of codes that generally means the end user got something wrong in making the request
type BadRequest struct {
	Message              *string        `json:"message,omitempty"`
	AdditionalProperties map[string]any `additionalProperties:"true" json:"-"`
}

var _ error = &BadRequest{}

func (e *BadRequest) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
