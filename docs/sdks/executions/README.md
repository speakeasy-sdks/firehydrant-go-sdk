# Executions
(*Runbooks.Executions*)

## Overview

### Available Operations

* [Delete](#delete) - Terminate a runbook execution

## Delete

Terminates a runbook execution, preventing any further steps from being executed

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
    res, err := s.Runbooks.Executions.Delete(ctx, "<id>", "<value>")
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
| `executionID`                                            | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `reason`                                                 | *string*                                                 | :heavy_check_mark:                                       | The reason for terminating the runbook execution         |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteRunbookExecutionResponse](../../models/operations/deleterunbookexecutionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |