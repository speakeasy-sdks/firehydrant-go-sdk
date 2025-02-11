// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"firehydrant/models/components"
)

type CreateScheduledMaintenanceResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Create a new scheduled maintenance event
	ScheduledMaintenanceEntity *components.ScheduledMaintenanceEntity
}

func (o *CreateScheduledMaintenanceResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateScheduledMaintenanceResponse) GetScheduledMaintenanceEntity() *components.ScheduledMaintenanceEntity {
	if o == nil {
		return nil
	}
	return o.ScheduledMaintenanceEntity
}
