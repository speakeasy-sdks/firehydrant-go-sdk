# Entitlements
(*Entitlements*)

## Overview

Operations about entitlements

### Available Operations

* [List](#list) - Retrieve all entitlements

## List

Retrieve all entitlements

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
    res, err := s.Entitlements.List(ctx, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.EntitlementEntityPaginated != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `name`                                                   | **string*                                                | :heavy_minus_sign:                                       | Name of Entitlement                                      |
| `type_`                                                  | [*operations.Type](../../models/operations/type.md)      | :heavy_minus_sign:                                       | Type of Entitlement                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetV1EntitlementsResponse](../../models/operations/getv1entitlementsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |