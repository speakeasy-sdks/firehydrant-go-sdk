// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type DeleteChangeRequest struct {
	ChangeID string `pathParam:"style=simple,explode=false,name=change_id"`
}

func (o *DeleteChangeRequest) GetChangeID() string {
	if o == nil {
		return ""
	}
	return o.ChangeID
}

type DeleteChangeResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *DeleteChangeResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}