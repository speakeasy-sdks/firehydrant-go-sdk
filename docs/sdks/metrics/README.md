# Metrics
(*Metrics*)

## Overview

Operations about metrics

### Available Operations

* [GetTicketFunnel](#getticketfunnel) - List ticket task and follow up creation and completion metrics
* [ListRetrospectives](#listretrospectives) - List retrospective metrics
* [ListMilestoneFunnel](#listmilestonefunnel) - List funnel metrics
* [GetUserInvolvements](#getuserinvolvements) - List user metrics
* [GetIncidentMetrics](#getincidentmetrics) - List incident metrics
* [GetMttx](#getmttx) - Fetch infrastructure metrics based on custom query
* [List](#list) - List metrics for a component type
* [Get](#get) - Show metrics for a component

## GetTicketFunnel

Returns a report with task and follow up creation and completion data

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
    res, err := s.Metrics.GetTicketFunnel(ctx, operations.GetV1MetricsTicketFunnelRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.MetricsTicketFunnelMetricsEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetV1MetricsTicketFunnelRequest](../../models/operations/getv1metricsticketfunnelrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.GetV1MetricsTicketFunnelResponse](../../models/operations/getv1metricsticketfunnelresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListRetrospectives

Returns a report with retrospective analytics data

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
    res, err := s.Metrics.ListRetrospectives(ctx, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.MetricsRetrospectiveEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `startDate`                                              | [*types.Date](../../types/date.md)                       | :heavy_minus_sign:                                       | The start date to return metrics from                    |
| `endDate`                                                | [*types.Date](../../types/date.md)                       | :heavy_minus_sign:                                       | The end date to return metrics from                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetV1MetricsRetrospectivesResponse](../../models/operations/getv1metricsretrospectivesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListMilestoneFunnel

Returns a report with time bucketed milestone data

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
    res, err := s.Metrics.ListMilestoneFunnel(ctx, operations.GetV1MetricsMilestoneFunnelRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.MetricsMilestonesFunnelEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetV1MetricsMilestoneFunnelRequest](../../models/operations/getv1metricsmilestonefunnelrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.GetV1MetricsMilestoneFunnelResponse](../../models/operations/getv1metricsmilestonefunnelresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetUserInvolvements

Returns a report with time bucketed analytics data

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
    res, err := s.Metrics.GetUserInvolvements(ctx, operations.GetV1MetricsUserInvolvementsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.MetricsMetricsEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetV1MetricsUserInvolvementsRequest](../../models/operations/getv1metricsuserinvolvementsrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.GetV1MetricsUserInvolvementsResponse](../../models/operations/getv1metricsuserinvolvementsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetIncidentMetrics

Returns a report with time bucketed analytics data

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
    res, err := s.Metrics.GetIncidentMetrics(ctx, operations.GetV1MetricsIncidentsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.MetricsMetricsEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetV1MetricsIncidentsRequest](../../models/operations/getv1metricsincidentsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.GetV1MetricsIncidentsResponse](../../models/operations/getv1metricsincidentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetMttx

Fetch infrastructure metrics based on custom query

### Example Usage

```go
package main

import(
	"openapi"
	"context"
	"openapi/types"
	"openapi/models/operations"
	"log"
)

func main() {
    s := openapi.New(
        openapi.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Metrics.GetMttx(ctx, operations.GetV1MetricsMttxRequest{
        StartDate: types.MustDateFromString("2024-12-27"),
        EndDate: types.MustDateFromString("2024-02-03"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MetricsMttxDataEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetV1MetricsMttxRequest](../../models/operations/getv1metricsmttxrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetV1MetricsMttxResponse](../../models/operations/getv1metricsmttxresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## List

Returns metrics for all components of a given type

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
    res, err := s.Metrics.List(ctx, operations.InfraTypeCustomers, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.MetricsInfrastructureListEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `infraType`                                                    | [operations.InfraType](../../models/operations/infratype.md)   | :heavy_check_mark:                                             | N/A                                                            |
| `startDate`                                                    | [*types.Date](../../types/date.md)                             | :heavy_minus_sign:                                             | The start date to return metrics from; defaults to 30 days ago |
| `endDate`                                                      | [*types.Date](../../types/date.md)                             | :heavy_minus_sign:                                             | The end date to return metrics from, defaults to today         |
| `opts`                                                         | [][operations.Option](../../models/operations/option.md)       | :heavy_minus_sign:                                             | The options for this request.                                  |

### Response

**[*operations.GetV1MetricsInfraTypeResponse](../../models/operations/getv1metricsinfratyperesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Get

Return metrics for a specific component

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
    res, err := s.Metrics.Get(ctx, operations.PathParamInfraTypeServices, "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.MetricsInfrastructureMetricsEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `infraType`                                                                    | [operations.PathParamInfraType](../../models/operations/pathparaminfratype.md) | :heavy_check_mark:                                                             | N/A                                                                            |
| `infraID`                                                                      | *string*                                                                       | :heavy_check_mark:                                                             | Component UUID                                                                 |
| `startDate`                                                                    | [*types.Date](../../types/date.md)                                             | :heavy_minus_sign:                                                             | The start date to return metrics from; defaults to 30 days ago                 |
| `endDate`                                                                      | [*types.Date](../../types/date.md)                                             | :heavy_minus_sign:                                                             | The end date to return metrics from, defaults to today                         |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.GetV1MetricsInfraTypeInfraIDResponse](../../models/operations/getv1metricsinfratypeinfraidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |