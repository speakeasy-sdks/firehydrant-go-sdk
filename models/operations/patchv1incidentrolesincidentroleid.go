// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type PatchV1IncidentRolesIncidentRoleIDRequest struct {
	IncidentRoleID                     string                                        `pathParam:"style=simple,explode=false,name=incident_role_id"`
	PatchV1IncidentRolesIncidentRoleID components.PatchV1IncidentRolesIncidentRoleID `request:"mediaType=application/json"`
}

func (o *PatchV1IncidentRolesIncidentRoleIDRequest) GetIncidentRoleID() string {
	if o == nil {
		return ""
	}
	return o.IncidentRoleID
}

func (o *PatchV1IncidentRolesIncidentRoleIDRequest) GetPatchV1IncidentRolesIncidentRoleID() components.PatchV1IncidentRolesIncidentRoleID {
	if o == nil {
		return components.PatchV1IncidentRolesIncidentRoleID{}
	}
	return o.PatchV1IncidentRolesIncidentRoleID
}

type PatchV1IncidentRolesIncidentRoleIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Update a single incident role from its ID
	IncidentRoleEntity *components.IncidentRoleEntity
}

func (o *PatchV1IncidentRolesIncidentRoleIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PatchV1IncidentRolesIncidentRoleIDResponse) GetIncidentRoleEntity() *components.IncidentRoleEntity {
	if o == nil {
		return nil
	}
	return o.IncidentRoleEntity
}