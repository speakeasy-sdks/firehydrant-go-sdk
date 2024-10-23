// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type DeleteV1ScheduledMaintenancesScheduledMaintenanceIDRequest struct {
	ScheduledMaintenanceID string `pathParam:"style=simple,explode=false,name=scheduled_maintenance_id"`
}

func (o *DeleteV1ScheduledMaintenancesScheduledMaintenanceIDRequest) GetScheduledMaintenanceID() string {
	if o == nil {
		return ""
	}
	return o.ScheduledMaintenanceID
}

type DeleteV1ScheduledMaintenancesScheduledMaintenanceIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *DeleteV1ScheduledMaintenancesScheduledMaintenanceIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
