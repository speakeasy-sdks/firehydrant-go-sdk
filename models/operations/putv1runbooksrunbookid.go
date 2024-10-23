// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type PutV1RunbooksRunbookIDRequest struct {
	RunbookID              string                            `pathParam:"style=simple,explode=false,name=runbook_id"`
	PutV1RunbooksRunbookID components.PutV1RunbooksRunbookID `request:"mediaType=application/json"`
}

func (o *PutV1RunbooksRunbookIDRequest) GetRunbookID() string {
	if o == nil {
		return ""
	}
	return o.RunbookID
}

func (o *PutV1RunbooksRunbookIDRequest) GetPutV1RunbooksRunbookID() components.PutV1RunbooksRunbookID {
	if o == nil {
		return components.PutV1RunbooksRunbookID{}
	}
	return o.PutV1RunbooksRunbookID
}

type PutV1RunbooksRunbookIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Update a runbook and any attachment rules associated with it. This endpoint is used to configure nearly everything
	// about a runbook, including but not limited to the steps, environments, attachment rules, and severities.
	//
	RunbookEntity *components.RunbookEntity
}

func (o *PutV1RunbooksRunbookIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PutV1RunbooksRunbookIDResponse) GetRunbookEntity() *components.RunbookEntity {
	if o == nil {
		return nil
	}
	return o.RunbookEntity
}
