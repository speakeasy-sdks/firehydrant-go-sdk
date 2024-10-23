# IncidentsRoleAssignmentEntity

Incidents_RoleAssignmentEntity model


## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `ID`                                                                            | **string*                                                                       | :heavy_minus_sign:                                                              | N/A                                                                             |
| `Status`                                                                        | [*components.Status](../../models/components/status.md)                         | :heavy_minus_sign:                                                              | N/A                                                                             |
| `CreatedAt`                                                                     | [*time.Time](https://pkg.go.dev/time#Time)                                      | :heavy_minus_sign:                                                              | N/A                                                                             |
| `UpdatedAt`                                                                     | [*time.Time](https://pkg.go.dev/time#Time)                                      | :heavy_minus_sign:                                                              | N/A                                                                             |
| `IncidentRole`                                                                  | [*components.IncidentRoleEntity](../../models/components/incidentroleentity.md) | :heavy_minus_sign:                                                              | IncidentRoleEntity model                                                        |
| `User`                                                                          | [*components.UserEntity](../../models/components/userentity.md)                 | :heavy_minus_sign:                                                              | N/A                                                                             |