# RunbookAudits
(*RunbookAudits*)

## Overview

Operations about runbook_audits

### Available Operations

* [List](#list) - List runbook audits

## List

Please contact support to enable audit logging for your account.

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
    res, err := s.RunbookAudits.List(ctx, nil, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                         | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `ctx`                                                                             | [context.Context](https://pkg.go.dev/context#Context)                             | :heavy_check_mark:                                                                | The context to use for the request.                                               |
| `page`                                                                            | **int*                                                                            | :heavy_minus_sign:                                                                | N/A                                                                               |
| `perPage`                                                                         | **int*                                                                            | :heavy_minus_sign:                                                                | N/A                                                                               |
| `auditableType`                                                                   | [*operations.AuditableType](../../models/operations/auditabletype.md)             | :heavy_minus_sign:                                                                | A query to filter audits by type                                                  |
| `sort`                                                                            | **string*                                                                         | :heavy_minus_sign:                                                                | A query to sort audits by their created_at timestamp. Options are 'asc' or 'desc' |
| `opts`                                                                            | [][operations.Option](../../models/operations/option.md)                          | :heavy_minus_sign:                                                                | The options for this request.                                                     |

### Response

**[*operations.GetV1RunbookAuditsResponse](../../models/operations/getv1runbookauditsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |