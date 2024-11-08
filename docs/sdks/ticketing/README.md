# Ticketing
(*Ticketing*)

## Overview

### Available Operations

* [GetFieldMap](#getfieldmap) - Get a field map for a ticketing project
* [UpdateProjectConfig](#updateprojectconfig) - Update a ticketing project configuration

## GetFieldMap

Retrieve field map for a ticketing project

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
    res, err := s.Ticketing.GetFieldMap(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.TicketingProjectFieldMapEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `mapID`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `ticketingProjectID`                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetTicketingProjectFieldMapResponse](../../models/operations/getticketingprojectfieldmapresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateProjectConfig

Update configuration for a ticketing project

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
    res, err := s.Ticketing.UpdateProjectConfig(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.TicketingProjectConfigEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `ticketingProjectID`                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `configID`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.UpdateTicketingProjectConfigResponse](../../models/operations/updateticketingprojectconfigresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |