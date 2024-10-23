# Steps
(*Runbooks.Steps*)

## Overview

### Available Operations

* [UpdateVotes](#updatevotes) - Update the votes on an object

## UpdateVotes

Allows for upvoting or downvoting an event

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
    res, err := s.Runbooks.Steps.UpdateVotes(ctx, "<id>", "<id>", components.PatchV1RunbooksExecutionsExecutionIDStepsStepIDVotes{
        Direction: components.PatchV1RunbooksExecutionsExecutionIDStepsStepIDVotesDirectionDown,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VotesEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                          | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                              | :heavy_check_mark:                                                                                                                                 | The context to use for the request.                                                                                                                |
| `executionID`                                                                                                                                      | *string*                                                                                                                                           | :heavy_check_mark:                                                                                                                                 | N/A                                                                                                                                                |
| `stepID`                                                                                                                                           | *string*                                                                                                                                           | :heavy_check_mark:                                                                                                                                 | N/A                                                                                                                                                |
| `patchV1RunbooksExecutionsExecutionIDStepsStepIDVotes`                                                                                             | [components.PatchV1RunbooksExecutionsExecutionIDStepsStepIDVotes](../../models/components/patchv1runbooksexecutionsexecutionidstepsstepidvotes.md) | :heavy_check_mark:                                                                                                                                 | N/A                                                                                                                                                |
| `opts`                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                           | :heavy_minus_sign:                                                                                                                                 | The options for this request.                                                                                                                      |

### Response

**[*operations.PatchV1RunbooksExecutionsExecutionIDStepsStepIDVotesResponse](../../models/operations/patchv1runbooksexecutionsexecutionidstepsstepidvotesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |