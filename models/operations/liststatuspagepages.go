// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type ListStatuspagePagesRequest struct {
	// Connection UUID
	ConnectionID string `pathParam:"style=simple,explode=false,name=connection_id"`
}

func (o *ListStatuspagePagesRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type ListStatuspagePagesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *ListStatuspagePagesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
