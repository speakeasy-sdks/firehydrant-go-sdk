# IncidentTypes
(*IncidentTypes*)

## Overview

### Available Operations

* [Create](#create) - Create an incident type
* [List](#list) - List all incident types
* [Archive](#archive) - Archive an incident type
* [Update](#update) - Update an incident type
* [Get](#get) - Retrieve an incident type

## Create

Create a new incident type

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
    res, err := s.IncidentTypes.Create(ctx, components.PostV1IncidentTypes{
        Name: "<value>",
        Template: components.Template{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IncidentTypeEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [components.PostV1IncidentTypes](../../models/components/postv1incidenttypes.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.PostV1IncidentTypesResponse](../../models/operations/postv1incidenttypesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## List

List all of the incident types in the organization

### Example Usage

```go
package main

import(
	"openapi"
	"context"
	"log"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.IncidentTypes.List(ctx, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.IncidentTypeEntityPaginated != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `query`                                                  | **string*                                                | :heavy_minus_sign:                                       | A query to search incident types by their name           |
| `page`                                                   | **int*                                                   | :heavy_minus_sign:                                       | N/A                                                      |
| `perPage`                                                | **int*                                                   | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetV1IncidentTypesResponse](../../models/operations/getv1incidenttypesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Archive

Archives an incident type which will hide it from lists and metrics

### Example Usage

```go
package main

import(
	"openapi"
	"context"
	"log"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.IncidentTypes.Archive(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.IncidentTypeEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteV1IncidentTypesIDResponse](../../models/operations/deletev1incidenttypesidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Update

Update a single incident type from its ID

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
    res, err := s.IncidentTypes.Update(ctx, "<id>", components.PatchV1IncidentTypesID{
        Name: "<value>",
        Template: components.PatchV1IncidentTypesIDTemplate{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.IncidentTypeEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `id`                                                                                   | *string*                                                                               | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `patchV1IncidentTypesID`                                                               | [components.PatchV1IncidentTypesID](../../models/components/patchv1incidenttypesid.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.PatchV1IncidentTypesIDResponse](../../models/operations/patchv1incidenttypesidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Get

Retrieve a single incident type from its ID

### Example Usage

```go
package main

import(
	"openapi"
	"context"
	"log"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.IncidentTypes.Get(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.IncidentTypeEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetV1IncidentTypesIDResponse](../../models/operations/getv1incidenttypesidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |