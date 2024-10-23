# Services
(*Services*)

## Overview

Operations about services

### Available Operations

* [Create](#create) - Create a service
* [List](#list) - List all services
* [BulkCreateServiceLinks](#bulkcreateservicelinks) - Bulk create services and service links
* [Delete](#delete) - Delete a service
* [Update](#update) - Update a service
* [Get](#get) - Retrieve a single service
* [ListDependencies](#listdependencies) - Retrieve a service's dependencies
* [ListAvailableUpstreamDependencies](#listavailableupstreamdependencies) - Retrieve all available upstream dependencies
* [ListAvailableDownstreamDependencies](#listavailabledownstreamdependencies) - Retrieve all available downstream dependencies
* [DeleteServiceLink](#deleteservicelink) - Delete a service link
* [CreateChecklistResponse](#createchecklistresponse) - Record a response for a checklist item

## Create

Creates a service for the organization, you may also create or attach functionalities to the service on create.

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
    res, err := s.Services.Create(ctx, components.PostV1Services{
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ServiceEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [components.PostV1Services](../../models/components/postv1services.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.PostV1ServicesResponse](../../models/operations/postv1servicesresponse.md), error**

### Errors

| Error Type            | Status Code           | Content Type          |
| --------------------- | --------------------- | --------------------- |
| sdkerrors.ErrorEntity | 400                   | application/json      |
| sdkerrors.SDKError    | 4XX, 5XX              | \*/\*                 |

## List

List all of the services that have been added to the organization.

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
    res, err := s.Services.List(ctx, operations.GetV1ServicesRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ServiceEntityPaginated != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetV1ServicesRequest](../../models/operations/getv1servicesrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.GetV1ServicesResponse](../../models/operations/getv1servicesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## BulkCreateServiceLinks

Creates a service with the appropriate integration for each external service ID passed

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
    res, err := s.Services.BulkCreateServiceLinks(ctx, components.PostV1ServicesServiceLinks{
        ExternalServiceIds: "<value>",
        ConnectionID: "<id>",
        Integration: components.IntegrationVictorops,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ServiceLinkEntities != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [components.PostV1ServicesServiceLinks](../../models/components/postv1servicesservicelinks.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.PostV1ServicesServiceLinksResponse](../../models/operations/postv1servicesservicelinksresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Delete

Deletes the service from FireHydrant.

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
    res, err := s.Services.Delete(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ServiceEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `serviceID`                                              | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteV1ServicesServiceIDResponse](../../models/operations/deletev1servicesserviceidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Update

Update a services attributes, you may also add or remove functionalities from the service as well.
Note: You may not remove or add individual label key/value pairs. You must include the entire object to override label values.


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
    res, err := s.Services.Update(ctx, "<id>", components.PatchV1ServicesServiceID{})
    if err != nil {
        log.Fatal(err)
    }
    if res.ServiceEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `serviceID`                                                                                | *string*                                                                                   | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `patchV1ServicesServiceID`                                                                 | [components.PatchV1ServicesServiceID](../../models/components/patchv1servicesserviceid.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.PatchV1ServicesServiceIDResponse](../../models/operations/patchv1servicesserviceidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Get

Retrieves a single service by ID

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
    res, err := s.Services.Get(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ServiceEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `serviceID`                                              | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetV1ServicesServiceIDResponse](../../models/operations/getv1servicesserviceidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListDependencies

Retrieves a service's dependencies

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
    res, err := s.Services.ListDependencies(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ServiceWithAllDependenciesEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                             | Type                                                                                                                                  | Required                                                                                                                              | Description                                                                                                                           |
| ------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                                                 | :heavy_check_mark:                                                                                                                    | The context to use for the request.                                                                                                   |
| `serviceID`                                                                                                                           | *string*                                                                                                                              | :heavy_check_mark:                                                                                                                    | N/A                                                                                                                                   |
| `flatten`                                                                                                                             | **bool*                                                                                                                               | :heavy_minus_sign:                                                                                                                    | If true, returns all dependencies in one array. If false, splits dependencies into different arrays for child and parent dependencies |
| `opts`                                                                                                                                | [][operations.Option](../../models/operations/option.md)                                                                              | :heavy_minus_sign:                                                                                                                    | The options for this request.                                                                                                         |

### Response

**[*operations.GetV1ServicesServiceIDDependenciesResponse](../../models/operations/getv1servicesserviceiddependenciesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAvailableUpstreamDependencies

Retrieves all services that are available to be upstream dependencies

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
    res, err := s.Services.ListAvailableUpstreamDependencies(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ServiceEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `serviceID`                                              | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetV1ServicesServiceIDAvailableUpstreamDependenciesResponse](../../models/operations/getv1servicesserviceidavailableupstreamdependenciesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListAvailableDownstreamDependencies

Retrieves all services that are available to be downstream dependencies

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
    res, err := s.Services.ListAvailableDownstreamDependencies(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ServiceEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `serviceID`                                              | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetV1ServicesServiceIDAvailableDownstreamDependenciesResponse](../../models/operations/getv1servicesserviceidavailabledownstreamdependenciesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteServiceLink

Deletes a service link from FireHydrant.

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
    res, err := s.Services.DeleteServiceLink(ctx, "<id>", "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                             | Type                                                                                                                                                  | Required                                                                                                                                              | Description                                                                                                                                           |
| ----------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                                                                 | :heavy_check_mark:                                                                                                                                    | The context to use for the request.                                                                                                                   |
| `serviceID`                                                                                                                                           | *string*                                                                                                                                              | :heavy_check_mark:                                                                                                                                    | N/A                                                                                                                                                   |
| `remoteID`                                                                                                                                            | *string*                                                                                                                                              | :heavy_check_mark:                                                                                                                                    | The external service ID which can be found in the JSON<br/>from GET services/:service_id endpoint under<br/>functionalities > external_resources > remote_id<br/> |
| `opts`                                                                                                                                                | [][operations.Option](../../models/operations/option.md)                                                                                              | :heavy_minus_sign:                                                                                                                                    | The options for this request.                                                                                                                         |

### Response

**[*operations.DeleteV1ServicesServiceIDServiceLinksRemoteIDResponse](../../models/operations/deletev1servicesserviceidservicelinksremoteidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateChecklistResponse

Creates a response for a checklist item

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
    res, err := s.Services.CreateChecklistResponse(ctx, "<id>", "<id>", components.PostV1ServicesServiceIDChecklistResponseChecklistID{
        CheckID: "<id>",
        Status: false,
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

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `serviceID`                                                                                                                                      | *string*                                                                                                                                         | :heavy_check_mark:                                                                                                                               | N/A                                                                                                                                              |
| `checklistID`                                                                                                                                    | *string*                                                                                                                                         | :heavy_check_mark:                                                                                                                               | N/A                                                                                                                                              |
| `postV1ServicesServiceIDChecklistResponseChecklistID`                                                                                            | [components.PostV1ServicesServiceIDChecklistResponseChecklistID](../../models/components/postv1servicesserviceidchecklistresponsechecklistid.md) | :heavy_check_mark:                                                                                                                               | N/A                                                                                                                                              |
| `opts`                                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                                         | :heavy_minus_sign:                                                                                                                               | The options for this request.                                                                                                                    |

### Response

**[*operations.PostV1ServicesServiceIDChecklistResponseChecklistIDResponse](../../models/operations/postv1servicesserviceidchecklistresponsechecklistidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |