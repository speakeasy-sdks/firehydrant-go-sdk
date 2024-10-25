# Lifecycles
(*Lifecycles*)

## Overview

Operations about lifecycles

### Available Operations

* [CreateMeasurementDefinition](#createmeasurementdefinition) - Create a measurement definition
* [ListMeasurementDefinitions](#listmeasurementdefinitions) - List all measurement definitions
* [ArchiveMeasurementDefinition](#archivemeasurementdefinition) - Archive a measurement definition
* [UpdateMeasurementDefinition](#updatemeasurementdefinition) - Update a measurement definition
* [GetMeasurementDefinition](#getmeasurementdefinition) - Retrieve a measurement definition
* [ListPhases](#listphases) - List all phases and milestones
* [CreateMilestone](#createmilestone) - Create a milestone
* [DeleteMilestone](#deletemilestone) - Delete a milestone
* [UpdateMilestone](#updatemilestone) - Update a milestone

## CreateMeasurementDefinition

Create a new measurement definition

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
    res, err := s.Lifecycles.CreateMeasurementDefinition(ctx, operations.PostV1LifecyclesMeasurementDefinitionsRequestBody{
        Name: "<value>",
        StartsAtMilestoneID: "<id>",
        EndsAtMilestoneID: "<id>",
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

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `request`                                                                                                                                    | [operations.PostV1LifecyclesMeasurementDefinitionsRequestBody](../../models/operations/postv1lifecyclesmeasurementdefinitionsrequestbody.md) | :heavy_check_mark:                                                                                                                           | The request object to use for the request.                                                                                                   |
| `opts`                                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                                     | :heavy_minus_sign:                                                                                                                           | The options for this request.                                                                                                                |

### Response

**[*operations.PostV1LifecyclesMeasurementDefinitionsResponse](../../models/operations/postv1lifecyclesmeasurementdefinitionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListMeasurementDefinitions

List all of the measurement definitions in the organization

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
    res, err := s.Lifecycles.ListMeasurementDefinitions(ctx, nil, nil)
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
| `page`                                                   | **int*                                                   | :heavy_minus_sign:                                       | N/A                                                      |
| `perPage`                                                | **int*                                                   | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetV1LifecyclesMeasurementDefinitionsResponse](../../models/operations/getv1lifecyclesmeasurementdefinitionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ArchiveMeasurementDefinition

Archives a measurement definition which will hide it from lists and metrics

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
    res, err := s.Lifecycles.ArchiveMeasurementDefinition(ctx, "<id>")
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
| `measurementDefinitionID`                                | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDResponse](../../models/operations/deletev1lifecyclesmeasurementdefinitionsmeasurementdefinitionidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateMeasurementDefinition

Update a single measurement definition from its ID

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
    res, err := s.Lifecycles.UpdateMeasurementDefinition(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                     | Type                                                                                                                                                                                          | Required                                                                                                                                                                                      | Description                                                                                                                                                                                   |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                         | :heavy_check_mark:                                                                                                                                                                            | The context to use for the request.                                                                                                                                                           |
| `measurementDefinitionID`                                                                                                                                                                     | *string*                                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                            | N/A                                                                                                                                                                                           |
| `requestBody`                                                                                                                                                                                 | [*operations.PatchV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDRequestBody](../../models/operations/patchv1lifecyclesmeasurementdefinitionsmeasurementdefinitionidrequestbody.md) | :heavy_minus_sign:                                                                                                                                                                            | N/A                                                                                                                                                                                           |
| `opts`                                                                                                                                                                                        | [][operations.Option](../../models/operations/option.md)                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                            | The options for this request.                                                                                                                                                                 |

### Response

**[*operations.PatchV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDResponse](../../models/operations/patchv1lifecyclesmeasurementdefinitionsmeasurementdefinitionidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetMeasurementDefinition

Retrieve a single measurement definition from its ID

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
    res, err := s.Lifecycles.GetMeasurementDefinition(ctx, "<id>")
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
| `measurementDefinitionID`                                | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetV1LifecyclesMeasurementDefinitionsMeasurementDefinitionIDResponse](../../models/operations/getv1lifecyclesmeasurementdefinitionsmeasurementdefinitionidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListPhases

List all of the lifecycle phases and milestones in the organization

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
    res, err := s.Lifecycles.ListPhases(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.LifecyclesPhaseEntityList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetV1LifecyclesPhasesResponse](../../models/operations/getv1lifecyclesphasesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateMilestone

Create a new milestone

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
    res, err := s.Lifecycles.CreateMilestone(ctx, operations.PostV1LifecyclesMilestonesRequestBody{
        Name: "<value>",
        Description: "consequently scoff caring inhibit entice wherever",
        PhaseID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.LifecyclesPhaseEntityList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.PostV1LifecyclesMilestonesRequestBody](../../models/operations/postv1lifecyclesmilestonesrequestbody.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.PostV1LifecyclesMilestonesResponse](../../models/operations/postv1lifecyclesmilestonesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteMilestone

Delete a milestone

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
    res, err := s.Lifecycles.DeleteMilestone(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.LifecyclesPhaseEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `milestoneID`                                            | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteV1LifecyclesMilestonesMilestoneIDResponse](../../models/operations/deletev1lifecyclesmilestonesmilestoneidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateMilestone

Update a milestone

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
    res, err := s.Lifecycles.UpdateMilestone(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.LifecyclesPhaseEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                     | Type                                                                                                                                          | Required                                                                                                                                      | Description                                                                                                                                   |
| --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                                                                         | :heavy_check_mark:                                                                                                                            | The context to use for the request.                                                                                                           |
| `milestoneID`                                                                                                                                 | *string*                                                                                                                                      | :heavy_check_mark:                                                                                                                            | N/A                                                                                                                                           |
| `requestBody`                                                                                                                                 | [*operations.PatchV1LifecyclesMilestonesMilestoneIDRequestBody](../../models/operations/patchv1lifecyclesmilestonesmilestoneidrequestbody.md) | :heavy_minus_sign:                                                                                                                            | N/A                                                                                                                                           |
| `opts`                                                                                                                                        | [][operations.Option](../../models/operations/option.md)                                                                                      | :heavy_minus_sign:                                                                                                                            | The options for this request.                                                                                                                 |

### Response

**[*operations.PatchV1LifecyclesMilestonesMilestoneIDResponse](../../models/operations/patchv1lifecyclesmilestonesmilestoneidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |