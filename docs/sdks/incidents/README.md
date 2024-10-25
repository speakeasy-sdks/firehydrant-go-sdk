# Incidents
(*Incidents*)

## Overview

Operations about incidents

### Available Operations

* [Create](#create) - Create an incident
* [List](#list) - List all incidents
* [GetChannel](#getchannel) - Retrieve chat channel information for an incident
* [Close](#close) - Close an incident
* [Resolve](#resolve) - Resolve an active incident
* [Archive](#archive) - Archive an incident
* [Update](#update) - Update an incident
* [Get](#get) - Retrieve an incident
* [Unarchive](#unarchive) - Unarchives an incident
* [AttachAlert](#attachalert) - Attach an alert to an incident
* [ListAlerts](#listalerts) - List alerts on an incident
* [SetPrimaryAlert](#setprimaryalert) - Set an alert as primary
* [RemoveAlert](#removealert) - Remove an alert
* [BulkUpdateMilestones](#bulkupdatemilestones) - Update milestone times
* [ListMilestones](#listmilestones) - List milestones on an incident
* [AddRelatedChange](#addrelatedchange) - Add a related change to an incident
* [GetRelatedChanges](#getrelatedchanges) - List related changes on an incident
* [UpdateRelatedChangeEvent](#updaterelatedchangeevent) - Update a change attached to an incident
* [AddStatusPage](#addstatuspage) - Add a status page to an incident
* [ListStatusPages](#liststatuspages) - List status pages on an incident
* [DeleteStatusPage](#deletestatuspage) - Remove a status page incident attached to an incident
* [AddTasksFromList](#addtasksfromlist) - Adds tasks to incident from task list
* [CreateTask](#createtask) - Creates a new task
* [GetTasks](#gettasks) - Retrieve tasks
* [DeleteTask](#deletetask) - Deletes a task
* [UpdateTask](#updatetask) - Update a task
* [GetTask](#gettask) - Retrieve a single task for an incident
* [ConvertTask](#converttask) - Creates a follow-up from a task, removing the task in the process
* [AddLink](#addlink) - Add a link to the incident
* [DeleteLink](#deletelink) - Deletes the external incident link
* [UpdateLink](#updatelink) - Updates the external incident link
* [GetTranscript](#gettranscript) - Lists all of the messages in the incident's transcript
* [DeleteTranscript](#deletetranscript)
* [GetSimilarIncidents](#getsimilarincidents)
* [GetAttachments](#getattachments)
* [AddAttachment](#addattachment) - Add an attachment to the incident timeline
* [ListEvents](#listevents) - List events for an incident
* [UpdateEvent](#updateevent) - Update a single event for an incident
* [GetEvent](#getevent) - Retrieve a single event for an incident
* [UpdateVotes](#updatevotes) - Update the votes on an object
* [UpdateImpact](#updateimpact) - Create a status update for an incident
* [ReplaceImpact](#replaceimpact) - Create a status update for an incident
* [AddImpact](#addimpact) - Add impacted infrastructure
* [ListImpact](#listimpact) - List impacted infrastructure
* [RemoveImpact](#removeimpact) - Remove impacted infrastructure on an incident
* [AddNote](#addnote) - Add a note to an incident
* [UpdateNote](#updatenote) - Update a note
* [AddGenericChatMessage](#addgenericchatmessage) - Add a generic chat message to the incident timeline.
* [DeleteChatMessage](#deletechatmessage) - Delete an existing generic chat message on an incident.
* [UpdateChatMessage](#updatechatmessage) - Update an existing generic chat message on an incident.
* [ListRoleAssignments](#listroleassignments) - List all assignees
* [UnassignRole](#unassignrole) - Unassign a role
* [AssignTeam](#assignteam) - Assign a team
* [DeleteTeamAssignment](#deleteteamassignment) - Unassign a team
* [GetUserRole](#getuserrole) - Get current role for user
* [GetRelationships](#getrelationships) - List any parent/child relationships for an incident

## Create

Create a new incident

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"firehydrant/models/components"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.Create(ctx, components.CreateIncident{
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IncidentEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [components.CreateIncident](../../models/components/createincident.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.PostV1IncidentsResponse](../../models/operations/postv1incidentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## List

List all of the incidents in the organization

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"firehydrant/models/operations"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.List(ctx, operations.GetV1IncidentsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.IncidentEntityPaginated != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetV1IncidentsRequest](../../models/operations/getv1incidentsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.GetV1IncidentsResponse](../../models/operations/getv1incidentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetChannel

Gives chat channel information for the specified incident

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.GetChannel(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.IncidentsChannelEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `incidentID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetV1IncidentsIncidentIDChannelResponse](../../models/operations/getv1incidentsincidentidchannelresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Close

Closes an incident and optionally close all children

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.Close(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.IncidentEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `incidentID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.PutV1IncidentsIncidentIDCloseResponse](../../models/operations/putv1incidentsincidentidcloseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Resolve

Resolves a currently active incident

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.Resolve(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.IncidentEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                       | Type                                                                                                                            | Required                                                                                                                        | Description                                                                                                                     |
| ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                                           | :heavy_check_mark:                                                                                                              | The context to use for the request.                                                                                             |
| `incidentID`                                                                                                                    | *string*                                                                                                                        | :heavy_check_mark:                                                                                                              | N/A                                                                                                                             |
| `requestBody`                                                                                                                   | [*operations.PutV1IncidentsIncidentIDResolveRequestBody](../../models/operations/putv1incidentsincidentidresolverequestbody.md) | :heavy_minus_sign:                                                                                                              | N/A                                                                                                                             |
| `opts`                                                                                                                          | [][operations.Option](../../models/operations/option.md)                                                                        | :heavy_minus_sign:                                                                                                              | The options for this request.                                                                                                   |

### Response

**[*operations.PutV1IncidentsIncidentIDResolveResponse](../../models/operations/putv1incidentsincidentidresolveresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Archive

Archives an incident which will hide it from lists and metrics

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.Archive(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.IncidentEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `incidentID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteV1IncidentsIncidentIDResponse](../../models/operations/deletev1incidentsincidentidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Update

Updates an incident with provided parameters

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"firehydrant/models/components"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.Update(ctx, "<id>", components.PatchV1IncidentsIncidentID{})
    if err != nil {
        log.Fatal(err)
    }
    if res.IncidentEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `incidentID`                                                                                   | *string*                                                                                       | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `patchV1IncidentsIncidentID`                                                                   | [components.PatchV1IncidentsIncidentID](../../models/components/patchv1incidentsincidentid.md) | :heavy_check_mark:                                                                             | N/A                                                                                            |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.PatchV1IncidentsIncidentIDResponse](../../models/operations/patchv1incidentsincidentidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Get

Retrieve a single incident from its ID

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.Get(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.IncidentEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `incidentID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetV1IncidentsIncidentIDResponse](../../models/operations/getv1incidentsincidentidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Unarchive

Unarchives an incident

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.Unarchive(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.IncidentEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `incidentID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.PostV1IncidentsIncidentIDUnarchiveResponse](../../models/operations/postv1incidentsincidentidunarchiveresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## AttachAlert

Add an alert to an incident. FireHydrant needs to have ingested the alert from a third party system in order to attach it to the incident.

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.AttachAlert(ctx, "<id>", []string{
        "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `incidentID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `requestBody`                                            | []*string*                                               | :heavy_check_mark:                                       | Array of alert IDs to be assigned to the incident        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.PostV1IncidentsIncidentIDAlertsResponse](../../models/operations/postv1incidentsincidentidalertsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAlerts

List alerts that have been attached to an incident

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.ListAlerts(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.IncidentsAlertEntityPaginated != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `incidentID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetV1IncidentsIncidentIDAlertsResponse](../../models/operations/getv1incidentsincidentidalertsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SetPrimaryAlert

Setting an alert as primary will overwrite milestone times in the FireHydrant incident with times included in the primary alert. Services attached to the primary alert will also be automatically added to the incident.

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"firehydrant/models/components"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.SetPrimaryAlert(ctx, "<id>", "<id>", components.PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimary{
        Primary: true,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IncidentsAlertEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                              | Type                                                                                                                                                   | Required                                                                                                                                               | Description                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                  | :heavy_check_mark:                                                                                                                                     | The context to use for the request.                                                                                                                    |
| `incidentAlertID`                                                                                                                                      | *string*                                                                                                                                               | :heavy_check_mark:                                                                                                                                     | N/A                                                                                                                                                    |
| `incidentID`                                                                                                                                           | *string*                                                                                                                                               | :heavy_check_mark:                                                                                                                                     | N/A                                                                                                                                                    |
| `patchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimary`                                                                                               | [components.PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimary](../../models/components/patchv1incidentsincidentidalertsincidentalertidprimary.md) | :heavy_check_mark:                                                                                                                                     | N/A                                                                                                                                                    |
| `opts`                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                               | :heavy_minus_sign:                                                                                                                                     | The options for this request.                                                                                                                          |

### Response

**[*operations.PatchV1IncidentsIncidentIDAlertsIncidentAlertIDPrimaryResponse](../../models/operations/patchv1incidentsincidentidalertsincidentalertidprimaryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveAlert

Remove an alert from an incident

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.RemoveAlert(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `incidentAlertID`                                        | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `incidentID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteV1IncidentsIncidentIDAlertsIncidentAlertIDResponse](../../models/operations/deletev1incidentsincidentidalertsincidentalertidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## BulkUpdateMilestones

Update milestone times in bulk for a given incident. All milestone
times for an incident must occur in chronological order
corresponding to the configured order of milestones. If the result
of this request would cause any milestone(s) to appear out of place,
a 422 response will instead be returned. This includes milestones
not explicitly submitted or updated in this request.


### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"firehydrant/types"
	"firehydrant/models/components"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.BulkUpdateMilestones(ctx, "<id>", components.PutV1IncidentsIncidentIDMilestonesBulkUpdate{
        Milestones: []components.PutV1IncidentsIncidentIDMilestonesBulkUpdateMilestones{
            components.PutV1IncidentsIncidentIDMilestonesBulkUpdateMilestones{
                Type: "<value>",
                OccurredAt: types.MustTimeFromString("2024-11-07T10:44:41.426Z"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IncidentsMilestoneEntityPaginated != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `incidentID`                                                                                                                       | *string*                                                                                                                           | :heavy_check_mark:                                                                                                                 | N/A                                                                                                                                |
| `putV1IncidentsIncidentIDMilestonesBulkUpdate`                                                                                     | [components.PutV1IncidentsIncidentIDMilestonesBulkUpdate](../../models/components/putv1incidentsincidentidmilestonesbulkupdate.md) | :heavy_check_mark:                                                                                                                 | N/A                                                                                                                                |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.PutV1IncidentsIncidentIDMilestonesBulkUpdateResponse](../../models/operations/putv1incidentsincidentidmilestonesbulkupdateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListMilestones

List times and durations for each milestone on an incident

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.ListMilestones(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.IncidentsMilestoneEntityPaginated != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `incidentID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetV1IncidentsIncidentIDMilestonesResponse](../../models/operations/getv1incidentsincidentidmilestonesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## AddRelatedChange

Add a related change to an incident. Changes added to an incident can be causes, fixes, or suspects. To remove a change from an incident, the type field should be set to dismissed.

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"firehydrant/models/components"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.AddRelatedChange(ctx, "<id>", components.PostV1IncidentsIncidentIDRelatedChangeEvents{
        ChangeEventID: "<id>",
        Type: components.PostV1IncidentsIncidentIDRelatedChangeEventsTypeCaused,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IncidentsRelatedChangeEventEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `incidentID`                                                                                                                       | *string*                                                                                                                           | :heavy_check_mark:                                                                                                                 | N/A                                                                                                                                |
| `postV1IncidentsIncidentIDRelatedChangeEvents`                                                                                     | [components.PostV1IncidentsIncidentIDRelatedChangeEvents](../../models/components/postv1incidentsincidentidrelatedchangeevents.md) | :heavy_check_mark:                                                                                                                 | N/A                                                                                                                                |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.PostV1IncidentsIncidentIDRelatedChangeEventsResponse](../../models/operations/postv1incidentsincidentidrelatedchangeeventsresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorEntity | 400, 409              | application/json      |
| sdkerrors.SDKError    | 4XX, 5XX              | \*/\*                 |

## GetRelatedChanges

List related changes that have been attached to an incident

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.GetRelatedChanges(ctx, "<id>", nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.IncidentsRelatedChangeEventEntityPaginated != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                               | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `ctx`                                                                   | [context.Context](https://pkg.go.dev/context#Context)                   | :heavy_check_mark:                                                      | The context to use for the request.                                     |
| `incidentID`                                                            | *string*                                                                | :heavy_check_mark:                                                      | N/A                                                                     |
| `page`                                                                  | **int*                                                                  | :heavy_minus_sign:                                                      | N/A                                                                     |
| `perPage`                                                               | **int*                                                                  | :heavy_minus_sign:                                                      | N/A                                                                     |
| `type_`                                                                 | [*operations.QueryParamType](../../models/operations/queryparamtype.md) | :heavy_minus_sign:                                                      | The type of the relation to the incident                                |
| `opts`                                                                  | [][operations.Option](../../models/operations/option.md)                | :heavy_minus_sign:                                                      | The options for this request.                                           |

### Response

**[*operations.GetV1IncidentsIncidentIDRelatedChangeEventsResponse](../../models/operations/getv1incidentsincidentidrelatedchangeeventsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateRelatedChangeEvent

Update a change attached to an incident

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"firehydrant/models/components"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.UpdateRelatedChangeEvent(ctx, "<id>", "<id>", components.PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventID{})
    if err != nil {
        log.Fatal(err)
    }
    if res.IncidentsRelatedChangeEventEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                    | Type                                                                                                                                                                         | Required                                                                                                                                                                     | Description                                                                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                        | :heavy_check_mark:                                                                                                                                                           | The context to use for the request.                                                                                                                                          |
| `relatedChangeEventID`                                                                                                                                                       | *string*                                                                                                                                                                     | :heavy_check_mark:                                                                                                                                                           | N/A                                                                                                                                                                          |
| `incidentID`                                                                                                                                                                 | *string*                                                                                                                                                                     | :heavy_check_mark:                                                                                                                                                           | N/A                                                                                                                                                                          |
| `patchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventID`                                                                                                          | [components.PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventID](../../models/components/patchv1incidentsincidentidrelatedchangeeventsrelatedchangeeventid.md) | :heavy_check_mark:                                                                                                                                                           | N/A                                                                                                                                                                          |
| `opts`                                                                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                                                                     | :heavy_minus_sign:                                                                                                                                                           | The options for this request.                                                                                                                                                |

### Response

**[*operations.PatchV1IncidentsIncidentIDRelatedChangeEventsRelatedChangeEventIDResponse](../../models/operations/patchv1incidentsincidentidrelatedchangeeventsrelatedchangeeventidresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| sdkerrors.ErrorEntity1 | 400, 409               | application/json       |
| sdkerrors.SDKError     | 4XX, 5XX               | \*/\*                  |

## AddStatusPage

Add a status page to an incident.

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"firehydrant/models/components"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.AddStatusPage(ctx, "<id>", components.PostV1IncidentsIncidentIDStatusPages{
        IntegrationSlug: "<value>",
        IntegrationID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IncidentsStatusPageEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `incidentID`                                                                                                       | *string*                                                                                                           | :heavy_check_mark:                                                                                                 | N/A                                                                                                                |
| `postV1IncidentsIncidentIDStatusPages`                                                                             | [components.PostV1IncidentsIncidentIDStatusPages](../../models/components/postv1incidentsincidentidstatuspages.md) | :heavy_check_mark:                                                                                                 | N/A                                                                                                                |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.PostV1IncidentsIncidentIDStatusPagesResponse](../../models/operations/postv1incidentsincidentidstatuspagesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListStatusPages

List status pages that are attached to an incident

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.ListStatusPages(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.IncidentsStatusPageEntityPaginated != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `incidentID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetV1IncidentsIncidentIDStatusPagesResponse](../../models/operations/getv1incidentsincidentidstatuspagesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteStatusPage

Remove a status page incident attached to an incident

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.DeleteStatusPage(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `statusPageID`                                           | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `incidentID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteV1IncidentsIncidentIDStatusPagesStatusPageIDResponse](../../models/operations/deletev1incidentsincidentidstatuspagesstatuspageidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## AddTasksFromList

Add all tasks from list to incident

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"firehydrant/models/components"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.AddTasksFromList(ctx, "<id>", components.PostV1IncidentsIncidentIDTaskLists{
        TaskListID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `incidentID`                                                                                                   | *string*                                                                                                       | :heavy_check_mark:                                                                                             | N/A                                                                                                            |
| `postV1IncidentsIncidentIDTaskLists`                                                                           | [components.PostV1IncidentsIncidentIDTaskLists](../../models/components/postv1incidentsincidentidtasklists.md) | :heavy_check_mark:                                                                                             | N/A                                                                                                            |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.PostV1IncidentsIncidentIDTaskListsResponse](../../models/operations/postv1incidentsincidentidtasklistsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateTask

Create a task

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"firehydrant/models/components"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.CreateTask(ctx, "<id>", components.PostV1IncidentsIncidentIDTasks{
        Title: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `incidentID`                                                                                           | *string*                                                                                               | :heavy_check_mark:                                                                                     | N/A                                                                                                    |
| `postV1IncidentsIncidentIDTasks`                                                                       | [components.PostV1IncidentsIncidentIDTasks](../../models/components/postv1incidentsincidentidtasks.md) | :heavy_check_mark:                                                                                     | N/A                                                                                                    |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.PostV1IncidentsIncidentIDTasksResponse](../../models/operations/postv1incidentsincidentidtasksresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetTasks

Retrieve a list of all tasks for a specific incident

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.GetTasks(ctx, "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskEntityPaginated != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `incidentID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `page`                                                   | **int*                                                   | :heavy_minus_sign:                                       | N/A                                                      |
| `perPage`                                                | **int*                                                   | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetV1IncidentsIncidentIDTasksResponse](../../models/operations/getv1incidentsincidentidtasksresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteTask

Delete a task

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.DeleteTask(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `taskID`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `incidentID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteV1IncidentsIncidentIDTasksTaskIDResponse](../../models/operations/deletev1incidentsincidentidtaskstaskidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateTask

Update a task's attributes

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"firehydrant/models/components"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.UpdateTask(ctx, "<id>", "<id>", components.PatchV1IncidentsIncidentIDTasksTaskID{})
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `taskID`                                                                                                             | *string*                                                                                                             | :heavy_check_mark:                                                                                                   | N/A                                                                                                                  |
| `incidentID`                                                                                                         | *string*                                                                                                             | :heavy_check_mark:                                                                                                   | N/A                                                                                                                  |
| `patchV1IncidentsIncidentIDTasksTaskID`                                                                              | [components.PatchV1IncidentsIncidentIDTasksTaskID](../../models/components/patchv1incidentsincidentidtaskstaskid.md) | :heavy_check_mark:                                                                                                   | N/A                                                                                                                  |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.PatchV1IncidentsIncidentIDTasksTaskIDResponse](../../models/operations/patchv1incidentsincidentidtaskstaskidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetTask

Retrieve a single task for an incident

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.GetTask(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `taskID`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `incidentID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetV1IncidentsIncidentIDTasksTaskIDResponse](../../models/operations/getv1incidentsincidentidtaskstaskidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ConvertTask

Convert a task to a follow-up

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"firehydrant/models/components"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.ConvertTask(ctx, "<id>", "<id>", components.PostV1IncidentsIncidentIDTasksTaskIDConvert{})
    if err != nil {
        log.Fatal(err)
    }
    if res.TaskEntityPaginated != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `taskID`                                                                                                                         | *string*                                                                                                                         | :heavy_check_mark:                                                                                                               | N/A                                                                                                                              |
| `incidentID`                                                                                                                     | *string*                                                                                                                         | :heavy_check_mark:                                                                                                               | N/A                                                                                                                              |
| `postV1IncidentsIncidentIDTasksTaskIDConvert`                                                                                    | [components.PostV1IncidentsIncidentIDTasksTaskIDConvert](../../models/components/postv1incidentsincidentidtaskstaskidconvert.md) | :heavy_check_mark:                                                                                                               | N/A                                                                                                                              |
| `opts`                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.PostV1IncidentsIncidentIDTasksTaskIDConvertResponse](../../models/operations/postv1incidentsincidentidtaskstaskidconvertresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## AddLink

Allows adding adhoc links to an incident as an attachment

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"firehydrant/models/components"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.AddLink(ctx, "<id>", components.PostV1IncidentsIncidentIDLinks{
        Href: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AttachmentsLinkEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `incidentID`                                                                                           | *string*                                                                                               | :heavy_check_mark:                                                                                     | N/A                                                                                                    |
| `postV1IncidentsIncidentIDLinks`                                                                       | [components.PostV1IncidentsIncidentIDLinks](../../models/components/postv1incidentsincidentidlinks.md) | :heavy_check_mark:                                                                                     | N/A                                                                                                    |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.PostV1IncidentsIncidentIDLinksResponse](../../models/operations/postv1incidentsincidentidlinksresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteLink

Deletes the external incident link

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.DeleteLink(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `linkID`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `incidentID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteV1IncidentsIncidentIDLinksLinkIDResponse](../../models/operations/deletev1incidentsincidentidlinkslinkidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateLink

Updates the external incident link attributes

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"firehydrant/models/components"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.UpdateLink(ctx, "<id>", "<id>", components.PutV1IncidentsIncidentIDLinksLinkID{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `linkID`                                                                                                         | *string*                                                                                                         | :heavy_check_mark:                                                                                               | N/A                                                                                                              |
| `incidentID`                                                                                                     | *string*                                                                                                         | :heavy_check_mark:                                                                                               | N/A                                                                                                              |
| `putV1IncidentsIncidentIDLinksLinkID`                                                                            | [components.PutV1IncidentsIncidentIDLinksLinkID](../../models/components/putv1incidentsincidentidlinkslinkid.md) | :heavy_check_mark:                                                                                               | N/A                                                                                                              |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.PutV1IncidentsIncidentIDLinksLinkIDResponse](../../models/operations/putv1incidentsincidentidlinkslinkidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetTranscript

Retrieve the transcript for a specific incident

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.GetTranscript(ctx, "<id>", nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.PublicApiv1IncidentsTranscriptEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `incidentID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `after`                                                  | **string*                                                | :heavy_minus_sign:                                       | The ID of the transcript entry to start after.           |
| `before`                                                 | **string*                                                | :heavy_minus_sign:                                       | The ID of the transcript entry to start before.          |
| `sort`                                                   | [*operations.Sort](../../models/operations/sort.md)      | :heavy_minus_sign:                                       | The order to sort the transcript entries.                |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetV1IncidentsIncidentIDTranscriptResponse](../../models/operations/getv1incidentsincidentidtranscriptresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteTranscript

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.DeleteTranscript(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `transcriptID`                                           | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `incidentID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteV1IncidentsIncidentIDTranscriptTranscriptIDResponse](../../models/operations/deletev1incidentsincidentidtranscripttranscriptidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetSimilarIncidents

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.GetSimilarIncidents(ctx, "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `incidentID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `threshold`                                              | **float32*                                               | :heavy_minus_sign:                                       | N/A                                                      |
| `limit`                                                  | **int*                                                   | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetV1IncidentsIncidentIDSimilarResponse](../../models/operations/getv1incidentsincidentidsimilarresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetAttachments

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.GetAttachments(ctx, "<id>", nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `incidentID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `attachableType`                                         | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `page`                                                   | **int*                                                   | :heavy_minus_sign:                                       | N/A                                                      |
| `perPage`                                                | **int*                                                   | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetV1IncidentsIncidentIDAttachmentsResponse](../../models/operations/getv1incidentsincidentidattachmentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## AddAttachment

Allows adding image attachments to an incident

### Example Usage

```go
package main

import(
	"firehydrant"
	"os"
	"context"
	"firehydrant/models/operations"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    content, fileErr := os.Open("example.file")
    if fileErr != nil {
        panic(fileErr)
    }

    ctx := context.Background()
    res, err := s.Incidents.AddAttachment(ctx, "<id>", operations.PostV1IncidentsIncidentIDAttachmentsRequestBody{
        File: operations.File{
            FileName: "example.file",
            Content: content,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IncidentAttachmentEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `incidentID`                                                                                                                             | *string*                                                                                                                                 | :heavy_check_mark:                                                                                                                       | N/A                                                                                                                                      |
| `requestBody`                                                                                                                            | [operations.PostV1IncidentsIncidentIDAttachmentsRequestBody](../../models/operations/postv1incidentsincidentidattachmentsrequestbody.md) | :heavy_check_mark:                                                                                                                       | N/A                                                                                                                                      |
| `opts`                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                 | :heavy_minus_sign:                                                                                                                       | The options for this request.                                                                                                            |

### Response

**[*operations.PostV1IncidentsIncidentIDAttachmentsResponse](../../models/operations/postv1incidentsincidentidattachmentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListEvents

List all events for an incident. An event is a timeline entry. This can be filtered with params to retrieve events of a certain type.

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.ListEvents(ctx, "<id>", nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.IncidentEventEntityPaginated != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `incidentID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `types`                                                  | **string*                                                | :heavy_minus_sign:                                       | A comma separated list of types of events to filter by   |
| `page`                                                   | **int*                                                   | :heavy_minus_sign:                                       | N/A                                                      |
| `perPage`                                                | **int*                                                   | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetV1IncidentsIncidentIDEventsResponse](../../models/operations/getv1incidentsincidentideventsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateEvent

Update a single event for an incident

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.UpdateEvent(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.IncidentEventEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `incidentID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `eventID`                                                | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.PatchV1IncidentsIncidentIDEventsEventIDResponse](../../models/operations/patchv1incidentsincidentideventseventidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetEvent

Retrieve a single event for an incident

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.GetEvent(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.IncidentEventEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `incidentID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `eventID`                                                | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetV1IncidentsIncidentIDEventsEventIDResponse](../../models/operations/getv1incidentsincidentideventseventidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateVotes

Allows for upvoting or downvoting an event

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"firehydrant/models/components"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.UpdateVotes(ctx, "<id>", "<id>", components.PatchV1IncidentsIncidentIDEventsEventIDVotes{
        Direction: components.DirectionDown,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VotesEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `incidentID`                                                                                                                       | *string*                                                                                                                           | :heavy_check_mark:                                                                                                                 | N/A                                                                                                                                |
| `eventID`                                                                                                                          | *string*                                                                                                                           | :heavy_check_mark:                                                                                                                 | N/A                                                                                                                                |
| `patchV1IncidentsIncidentIDEventsEventIDVotes`                                                                                     | [components.PatchV1IncidentsIncidentIDEventsEventIDVotes](../../models/components/patchv1incidentsincidentideventseventidvotes.md) | :heavy_check_mark:                                                                                                                 | N/A                                                                                                                                |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.PatchV1IncidentsIncidentIDEventsEventIDVotesResponse](../../models/operations/patchv1incidentsincidentideventseventidvotesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateImpact

Allows updating an incident's impacted infrastructure, with the option to
move the incident into a different milestone and provide a note to update
the incident timeline and any attached status pages. If this method is
requested with the PUT verb, impacts will be completely replaced with the
information in the request body, even if not provided (effectively clearing
all impacts). If this method is requested with the PATCH verb, the provided
impacts will be added or updated, but no impacts will be removed.


### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"firehydrant/models/components"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.UpdateImpact(ctx, "<id>", components.PatchV1IncidentsIncidentIDImpact{})
    if err != nil {
        log.Fatal(err)
    }
    if res.IncidentEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `incidentID`                                                                                               | *string*                                                                                                   | :heavy_check_mark:                                                                                         | N/A                                                                                                        |
| `patchV1IncidentsIncidentIDImpact`                                                                         | [components.PatchV1IncidentsIncidentIDImpact](../../models/components/patchv1incidentsincidentidimpact.md) | :heavy_check_mark:                                                                                         | N/A                                                                                                        |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.PatchV1IncidentsIncidentIDImpactResponse](../../models/operations/patchv1incidentsincidentidimpactresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ReplaceImpact

Allows updating an incident's impacted infrastructure, with the option to
move the incident into a different milestone and provide a note to update
the incident timeline and any attached status pages. If this method is
requested with the PUT verb, impacts will be completely replaced with the
information in the request body, even if not provided (effectively clearing
all impacts). If this method is requested with the PATCH verb, the provided
impacts will be added or updated, but no impacts will be removed.


### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"firehydrant/models/components"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.ReplaceImpact(ctx, "<id>", components.PutV1IncidentsIncidentIDImpact{})
    if err != nil {
        log.Fatal(err)
    }
    if res.IncidentEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `incidentID`                                                                                           | *string*                                                                                               | :heavy_check_mark:                                                                                     | N/A                                                                                                    |
| `putV1IncidentsIncidentIDImpact`                                                                       | [components.PutV1IncidentsIncidentIDImpact](../../models/components/putv1incidentsincidentidimpact.md) | :heavy_check_mark:                                                                                     | N/A                                                                                                    |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.PutV1IncidentsIncidentIDImpactResponse](../../models/operations/putv1incidentsincidentidimpactresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## AddImpact

Add impacted infrastructure to an incident

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"firehydrant/models/operations"
	"firehydrant/models/components"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.AddImpact(ctx, "<id>", operations.PathParamTypeCustomers, components.PostV1IncidentsIncidentIDImpactType{
        ID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IncidentImpactEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `incidentID`                                                                                                     | *string*                                                                                                         | :heavy_check_mark:                                                                                               | N/A                                                                                                              |
| `type_`                                                                                                          | [operations.PathParamType](../../models/operations/pathparamtype.md)                                             | :heavy_check_mark:                                                                                               | N/A                                                                                                              |
| `postV1IncidentsIncidentIDImpactType`                                                                            | [components.PostV1IncidentsIncidentIDImpactType](../../models/components/postv1incidentsincidentidimpacttype.md) | :heavy_check_mark:                                                                                               | N/A                                                                                                              |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.PostV1IncidentsIncidentIDImpactTypeResponse](../../models/operations/postv1incidentsincidentidimpacttyperesponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| sdkerrors.ErrorEntity2 | 400                    | application/json       |
| sdkerrors.SDKError     | 4XX, 5XX               | \*/\*                  |

## ListImpact

List impacted infrastructure on an incident by specifying type

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"firehydrant/models/operations"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.ListImpact(ctx, "<id>", operations.GetV1IncidentsIncidentIDImpactTypePathParamTypeEnvironments)
    if err != nil {
        log.Fatal(err)
    }
    if res.IncidentImpactEntityPaginated != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                    | :heavy_check_mark:                                                                                                                       | The context to use for the request.                                                                                                      |
| `incidentID`                                                                                                                             | *string*                                                                                                                                 | :heavy_check_mark:                                                                                                                       | N/A                                                                                                                                      |
| `type_`                                                                                                                                  | [operations.GetV1IncidentsIncidentIDImpactTypePathParamType](../../models/operations/getv1incidentsincidentidimpacttypepathparamtype.md) | :heavy_check_mark:                                                                                                                       | N/A                                                                                                                                      |
| `opts`                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                 | :heavy_minus_sign:                                                                                                                       | The options for this request.                                                                                                            |

### Response

**[*operations.GetV1IncidentsIncidentIDImpactTypeResponse](../../models/operations/getv1incidentsincidentidimpacttyperesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## RemoveImpact

Remove impacted infrastructure on an incident

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"firehydrant/models/operations"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.RemoveImpact(ctx, "<id>", operations.DeleteV1IncidentsIncidentIDImpactTypeIDPathParamTypeFunctionalities, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `incidentID`                                                                                                                                       | *string*                                                                                                                                           | :heavy_check_mark:                                                                                                                                 | N/A                                                                                                                                                |
| `type_`                                                                                                                                            | [operations.DeleteV1IncidentsIncidentIDImpactTypeIDPathParamType](../../models/operations/deletev1incidentsincidentidimpacttypeidpathparamtype.md) | :heavy_check_mark:                                                                                                                                 | N/A                                                                                                                                                |
| `id`                                                                                                                                               | *string*                                                                                                                                           | :heavy_check_mark:                                                                                                                                 | N/A                                                                                                                                                |
| `opts`                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                           | :heavy_minus_sign:                                                                                                                                 | The options for this request.                                                                                                                      |

### Response

**[*operations.DeleteV1IncidentsIncidentIDImpactTypeIDResponse](../../models/operations/deletev1incidentsincidentidimpacttypeidresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| sdkerrors.ErrorEntity3 | 400                    | application/json       |
| sdkerrors.SDKError     | 4XX, 5XX               | \*/\*                  |

## AddNote

Create a new note on for an incident. The visibility field on a note determines where it gets posted.

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"firehydrant/models/components"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.AddNote(ctx, "<id>", components.PostV1IncidentsIncidentIDNotes{
        Body: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EventNoteEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `incidentID`                                                                                           | *string*                                                                                               | :heavy_check_mark:                                                                                     | N/A                                                                                                    |
| `postV1IncidentsIncidentIDNotes`                                                                       | [components.PostV1IncidentsIncidentIDNotes](../../models/components/postv1incidentsincidentidnotes.md) | :heavy_check_mark:                                                                                     | N/A                                                                                                    |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.PostV1IncidentsIncidentIDNotesResponse](../../models/operations/postv1incidentsincidentidnotesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateNote

Updates the body of a note

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"firehydrant/models/components"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.UpdateNote(ctx, "<id>", "<id>", components.PatchV1IncidentsIncidentIDNotesNoteID{
        Body: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EventNoteEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `noteID`                                                                                                             | *string*                                                                                                             | :heavy_check_mark:                                                                                                   | N/A                                                                                                                  |
| `incidentID`                                                                                                         | *string*                                                                                                             | :heavy_check_mark:                                                                                                   | N/A                                                                                                                  |
| `patchV1IncidentsIncidentIDNotesNoteID`                                                                              | [components.PatchV1IncidentsIncidentIDNotesNoteID](../../models/components/patchv1incidentsincidentidnotesnoteid.md) | :heavy_check_mark:                                                                                                   | N/A                                                                                                                  |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.PatchV1IncidentsIncidentIDNotesNoteIDResponse](../../models/operations/patchv1incidentsincidentidnotesnoteidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## AddGenericChatMessage

Create a new generic chat message on an incident timeline. These are independent of any specific chat provider.

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"firehydrant/models/components"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.AddGenericChatMessage(ctx, "<id>", components.PostV1IncidentsIncidentIDGenericChatMessages{
        Body: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EventGenericChatMessageEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                          | Type                                                                                                                               | Required                                                                                                                           | Description                                                                                                                        |
| ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                              | :heavy_check_mark:                                                                                                                 | The context to use for the request.                                                                                                |
| `incidentID`                                                                                                                       | *string*                                                                                                                           | :heavy_check_mark:                                                                                                                 | N/A                                                                                                                                |
| `postV1IncidentsIncidentIDGenericChatMessages`                                                                                     | [components.PostV1IncidentsIncidentIDGenericChatMessages](../../models/components/postv1incidentsincidentidgenericchatmessages.md) | :heavy_check_mark:                                                                                                                 | N/A                                                                                                                                |
| `opts`                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                           | :heavy_minus_sign:                                                                                                                 | The options for this request.                                                                                                      |

### Response

**[*operations.PostV1IncidentsIncidentIDGenericChatMessagesResponse](../../models/operations/postv1incidentsincidentidgenericchatmessagesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteChatMessage

Delete an existing generic chat message on an incident.

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.DeleteChatMessage(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.EventGenericChatMessageEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `messageID`                                              | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `incidentID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteV1IncidentsIncidentIDGenericChatMessagesMessageIDResponse](../../models/operations/deletev1incidentsincidentidgenericchatmessagesmessageidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateChatMessage

Update an existing generic chat message on an incident.

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"firehydrant/models/components"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.UpdateChatMessage(ctx, "<id>", "<id>", components.PatchV1IncidentsIncidentIDGenericChatMessagesMessageID{
        Body: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EventGenericChatMessageEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                              | Type                                                                                                                                                   | Required                                                                                                                                               | Description                                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                  | :heavy_check_mark:                                                                                                                                     | The context to use for the request.                                                                                                                    |
| `messageID`                                                                                                                                            | *string*                                                                                                                                               | :heavy_check_mark:                                                                                                                                     | N/A                                                                                                                                                    |
| `incidentID`                                                                                                                                           | *string*                                                                                                                                               | :heavy_check_mark:                                                                                                                                     | N/A                                                                                                                                                    |
| `patchV1IncidentsIncidentIDGenericChatMessagesMessageID`                                                                                               | [components.PatchV1IncidentsIncidentIDGenericChatMessagesMessageID](../../models/components/patchv1incidentsincidentidgenericchatmessagesmessageid.md) | :heavy_check_mark:                                                                                                                                     | N/A                                                                                                                                                    |
| `opts`                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                               | :heavy_minus_sign:                                                                                                                                     | The options for this request.                                                                                                                          |

### Response

**[*operations.PatchV1IncidentsIncidentIDGenericChatMessagesMessageIDResponse](../../models/operations/patchv1incidentsincidentidgenericchatmessagesmessageidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListRoleAssignments

Retrieve a list of all of the current role assignments for the incident

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.ListRoleAssignments(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.IncidentsRoleAssignmentEntityPaginated != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `incidentID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `status`                                                 | [*operations.Status](../../models/operations/status.md)  | :heavy_minus_sign:                                       | Filter on status of the role assignment                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetV1IncidentsIncidentIDRoleAssignmentsResponse](../../models/operations/getv1incidentsincidentidroleassignmentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UnassignRole

Unassign a role from a user

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.UnassignRole(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.IncidentsRoleAssignmentEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `incidentID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `roleAssignmentID`                                       | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteV1IncidentsIncidentIDRoleAssignmentsRoleAssignmentIDResponse](../../models/operations/deletev1incidentsincidentidroleassignmentsroleassignmentidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## AssignTeam

Assign a team for this incident

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"firehydrant/models/components"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.AssignTeam(ctx, "<id>", components.PostV1IncidentsIncidentIDTeamAssignments{
        TeamID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `incidentID`                                                                                                               | *string*                                                                                                                   | :heavy_check_mark:                                                                                                         | N/A                                                                                                                        |
| `postV1IncidentsIncidentIDTeamAssignments`                                                                                 | [components.PostV1IncidentsIncidentIDTeamAssignments](../../models/components/postv1incidentsincidentidteamassignments.md) | :heavy_check_mark:                                                                                                         | N/A                                                                                                                        |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.PostV1IncidentsIncidentIDTeamAssignmentsResponse](../../models/operations/postv1incidentsincidentidteamassignmentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteTeamAssignment

Unassign a team from an incident

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.DeleteTeamAssignment(ctx, "<id>", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                             | Type                                                                                                                                                                                  | Required                                                                                                                                                                              | Description                                                                                                                                                                           |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                 | :heavy_check_mark:                                                                                                                                                                    | The context to use for the request.                                                                                                                                                   |
| `incidentID`                                                                                                                                                                          | *string*                                                                                                                                                                              | :heavy_check_mark:                                                                                                                                                                    | N/A                                                                                                                                                                                   |
| `teamAssignmentID`                                                                                                                                                                    | *string*                                                                                                                                                                              | :heavy_check_mark:                                                                                                                                                                    | N/A                                                                                                                                                                                   |
| `requestBody`                                                                                                                                                                         | [*operations.DeleteV1IncidentsIncidentIDTeamAssignmentsTeamAssignmentIDRequestBody](../../models/operations/deletev1incidentsincidentidteamassignmentsteamassignmentidrequestbody.md) | :heavy_minus_sign:                                                                                                                                                                    | N/A                                                                                                                                                                                   |
| `opts`                                                                                                                                                                                | [][operations.Option](../../models/operations/option.md)                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                    | The options for this request.                                                                                                                                                         |

### Response

**[*operations.DeleteV1IncidentsIncidentIDTeamAssignmentsTeamAssignmentIDResponse](../../models/operations/deletev1incidentsincidentidteamassignmentsteamassignmentidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetUserRole

Retrieve a user with current roles for an incident

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.GetUserRole(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.IncidentsRoleAssignmentEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `incidentID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `userID`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetV1IncidentsIncidentIDUsersUserIDResponse](../../models/operations/getv1incidentsincidentidusersuseridresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetRelationships

List any parent/child relationships for an incident

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Incidents.GetRelationships(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.IncidentsRelationshipsEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `incidentID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetV1IncidentsIncidentIDRelationshipsResponse](../../models/operations/getv1incidentsincidentidrelationshipsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |