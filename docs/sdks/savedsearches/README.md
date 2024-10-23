# SavedSearches
(*SavedSearches*)

## Overview

Operations about saved_searches

### Available Operations

* [Delete](#delete) - Delete a specific saved search
* [Update](#update) - Update a specific saved search
* [List](#list) - Lists save searches
* [Get](#get) - Retrieve a specific save search
* [Create](#create) - Create saved search

## Delete

Delete a specific saved search

### Example Usage

```go
package main

import(
	"openapi"
	"context"
	"openapi/models/operations"
	"log"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.SavedSearches.Delete(ctx, operations.ResourceTypeTicketFollowUps, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `resourceType`                                                     | [operations.ResourceType](../../models/operations/resourcetype.md) | :heavy_check_mark:                                                 | N/A                                                                |
| `savedSearchID`                                                    | *string*                                                           | :heavy_check_mark:                                                 | N/A                                                                |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.DeleteV1SavedSearchesResourceTypeSavedSearchIDResponse](../../models/operations/deletev1savedsearchesresourcetypesavedsearchidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Update

Update a specific saved search

### Example Usage

```go
package main

import(
	"openapi"
	"context"
	"openapi/models/operations"
	"openapi/models/components"
	"log"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.SavedSearches.Update(ctx, operations.PathParamResourceTypeAlerts, "<id>", components.PatchV1SavedSearchesResourceTypeSavedSearchID{})
    if err != nil {
        log.Fatal(err)
    }
    if res.SavedSearchEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                            | Type                                                                                                                                 | Required                                                                                                                             | Description                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                | :heavy_check_mark:                                                                                                                   | The context to use for the request.                                                                                                  |
| `resourceType`                                                                                                                       | [operations.PathParamResourceType](../../models/operations/pathparamresourcetype.md)                                                 | :heavy_check_mark:                                                                                                                   | N/A                                                                                                                                  |
| `savedSearchID`                                                                                                                      | *string*                                                                                                                             | :heavy_check_mark:                                                                                                                   | N/A                                                                                                                                  |
| `patchV1SavedSearchesResourceTypeSavedSearchID`                                                                                      | [components.PatchV1SavedSearchesResourceTypeSavedSearchID](../../models/components/patchv1savedsearchesresourcetypesavedsearchid.md) | :heavy_check_mark:                                                                                                                   | N/A                                                                                                                                  |
| `opts`                                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                                             | :heavy_minus_sign:                                                                                                                   | The options for this request.                                                                                                        |

### Response

**[*operations.PatchV1SavedSearchesResourceTypeSavedSearchIDResponse](../../models/operations/patchv1savedsearchesresourcetypesavedsearchidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## List

Lists save searches

### Example Usage

```go
package main

import(
	"openapi"
	"context"
	"openapi/models/operations"
	"log"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.SavedSearches.List(ctx, operations.GetV1SavedSearchesResourceTypeRequest{
        ResourceType: operations.GetV1SavedSearchesResourceTypePathParamResourceTypeImpactAnalytics,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SavedSearchEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.GetV1SavedSearchesResourceTypeRequest](../../models/operations/getv1savedsearchesresourcetyperequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.GetV1SavedSearchesResourceTypeResponse](../../models/operations/getv1savedsearchesresourcetyperesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Get

Retrieve a specific save search

### Example Usage

```go
package main

import(
	"openapi"
	"context"
	"openapi/models/operations"
	"log"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.SavedSearches.Get(ctx, operations.GetV1SavedSearchesResourceTypeSavedSearchIDPathParamResourceTypeImpactAnalytics, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.SavedSearchEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                  | Type                                                                                                                                                                       | Required                                                                                                                                                                   | Description                                                                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                      | :heavy_check_mark:                                                                                                                                                         | The context to use for the request.                                                                                                                                        |
| `resourceType`                                                                                                                                                             | [operations.GetV1SavedSearchesResourceTypeSavedSearchIDPathParamResourceType](../../models/operations/getv1savedsearchesresourcetypesavedsearchidpathparamresourcetype.md) | :heavy_check_mark:                                                                                                                                                         | N/A                                                                                                                                                                        |
| `savedSearchID`                                                                                                                                                            | *string*                                                                                                                                                                   | :heavy_check_mark:                                                                                                                                                         | N/A                                                                                                                                                                        |
| `opts`                                                                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                                                                   | :heavy_minus_sign:                                                                                                                                                         | The options for this request.                                                                                                                                              |

### Response

**[*operations.GetV1SavedSearchesResourceTypeSavedSearchIDResponse](../../models/operations/getv1savedsearchesresourcetypesavedsearchidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Create

Create a new saved search for a particular resource type

### Example Usage

```go
package main

import(
	"openapi"
	"context"
	"openapi/models/operations"
	"openapi/models/components"
	"log"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.SavedSearches.Create(ctx, operations.PostV1SavedSearchesResourceTypePathParamResourceTypeAnalytics, components.PostV1SavedSearchesResourceType{
        Name: "<value>",
        FilterValues: map[string]any{
            "key": "<value>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SavedSearchEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `resourceType`                                                                                                                                     | [operations.PostV1SavedSearchesResourceTypePathParamResourceType](../../models/operations/postv1savedsearchesresourcetypepathparamresourcetype.md) | :heavy_check_mark:                                                                                                                                 | N/A                                                                                                                                                |
| `postV1SavedSearchesResourceType`                                                                                                                  | [components.PostV1SavedSearchesResourceType](../../models/components/postv1savedsearchesresourcetype.md)                                           | :heavy_check_mark:                                                                                                                                 | N/A                                                                                                                                                |
| `opts`                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                           | :heavy_minus_sign:                                                                                                                                 | The options for this request.                                                                                                                      |

### Response

**[*operations.PostV1SavedSearchesResourceTypeResponse](../../models/operations/postv1savedsearchesresourcetyperesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |