# Infrastructures
(*Infrastructures*)

## Overview

### Available Operations

* [List](#list) - List catalog entries

## List

Lists functionality, service and environment objects

### Example Usage

```go
package main

import(
	"context"
	"firehydrant"
	"firehydrant/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Infrastructures.List(ctx, operations.ListInfrastructuresRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.InfrastructureSearchEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListInfrastructuresRequest](../../models/operations/listinfrastructuresrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.ListInfrastructuresResponse](../../models/operations/listinfrastructuresresponse.md), error**

### Errors

| Error Type                    | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.BadRequest          | 400, 413, 414, 415, 422, 431  | application/json              |
| sdkerrors.Unauthorized        | 401, 403, 407                 | application/json              |
| sdkerrors.NotFound            | 404                           | application/json              |
| sdkerrors.Timeout             | 408                           | application/json              |
| sdkerrors.RateLimited         | 429                           | application/json              |
| sdkerrors.InternalServerError | 500, 502, 503, 506, 507, 508  | application/json              |
| sdkerrors.NotFound            | 501, 505                      | application/json              |
| sdkerrors.Timeout             | 504                           | application/json              |
| sdkerrors.BadRequest          | 510                           | application/json              |
| sdkerrors.Unauthorized        | 511                           | application/json              |
| sdkerrors.SDKError            | 4XX, 5XX                      | \*/\*                         |