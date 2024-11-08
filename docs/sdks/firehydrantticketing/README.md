# FireHydrantTicketing
(*Integrations.Ticketing*)

## Overview

### Available Operations

* [GetProject](#getproject) - Get a ticketing project
* [ListTags](#listtags) - List ticket tags

## GetProject

Retrieve a single ticketing project by ID

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
    res, err := s.Integrations.Ticketing.GetProject(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.TicketingProjectsProjectListItemEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `ticketingProjectID`                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetTicketingProjectResponse](../../models/operations/getticketingprojectresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListTags

List all of the ticket tags in the organization

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
    res, err := s.Integrations.Ticketing.ListTags(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.TagEntityPaginated != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `prefix`                                                 | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListTicketTagsResponse](../../models/operations/listtickettagsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |