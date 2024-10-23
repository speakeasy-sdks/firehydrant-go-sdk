// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type DeleteV1IncidentRolesIncidentRoleIDRequest struct {
	IncidentRoleID string `pathParam:"style=simple,explode=false,name=incident_role_id"`
}

func (o *DeleteV1IncidentRolesIncidentRoleIDRequest) GetIncidentRoleID() string {
	if o == nil {
		return ""
	}
	return o.IncidentRoleID
}

type DeleteV1IncidentRolesIncidentRoleIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Archives an incident role which will hide it from lists and metrics
	IncidentRoleEntity *components.IncidentRoleEntity
}

func (o *DeleteV1IncidentRolesIncidentRoleIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteV1IncidentRolesIncidentRoleIDResponse) GetIncidentRoleEntity() *components.IncidentRoleEntity {
	if o == nil {
		return nil
	}
	return o.IncidentRoleEntity
}