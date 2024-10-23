# RoleAssignments
(*Incidents.RoleAssignments*)

## Overview

### Available Operations

* [Assign](#assign) - Assign a role

## Assign

Assign a role to a user for this incident

### Example Usage

```go
package main

import(
	"openapi"
	"context"
	"openapi/models/components"
	"log"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.RoleAssignments.Assign(ctx, "<id>", components.PostV1IncidentsIncidentIDRoleAssignments{
        UserID: "<id>",
        IncidentRoleID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IncidentsRoleAssignmentEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `incidentID`                                                                                                               | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | N/A                                                                                                                        |
| `postV1IncidentsIncidentIDRoleAssignments`                                                                                 | [components.PostV1IncidentsIncidentIDRoleAssignments](../../models/components/postv1incidentsincidentidroleassignments.md) | :heavy_check_mark:                                                                                                         | N/A                                                                                                                        |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.PostV1IncidentsIncidentIDRoleAssignmentsResponse](../../models/operations/postv1incidentsincidentidroleassignmentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |