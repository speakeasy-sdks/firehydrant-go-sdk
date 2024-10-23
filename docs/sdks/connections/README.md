# Connections
(*Nunc.Connections*)

## Overview

### Available Operations

* [Create](#create) - Create a FireHydrant hosted status page

## Create

Create a new FireHydrant hosted status page for customer facing statuses.

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
    res, err := s.Nunc.Connections.Create(ctx, operations.PostV1NuncConnectionsRequestBody{
        Domain: "low-packaging.info",
        ConditionsNuncCondition: []string{
            "<value>",
        },
        ConditionsConditionID: []string{
            "<value>",
        },
        ComponentsInfrastructureType: []string{
            "<value>",
            "<value>",
            "<value>",
        },
        ComponentsInfrastructureID: []string{
            "<value>",
            "<value>",
            "<value>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.NuncConnectionEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.PostV1NuncConnectionsRequestBody](../../models/operations/postv1nuncconnectionsrequestbody.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.PostV1NuncConnectionsResponse](../../models/operations/postv1nuncconnectionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |