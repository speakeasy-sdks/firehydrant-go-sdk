# FireHydrantSignalsTeams
(*Signals.Teams*)

## Overview

### Available Operations

* [GetEscalationPolicy](#getescalationpolicy) - Get an escalation policy for a team

## GetEscalationPolicy

Get a Signals escalation policy by ID

### Example Usage

```go
package main

import(
	"context"
	"firehydrant"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Signals.Teams.GetEscalationPolicy(ctx, "<id>", "<id>")
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
| `teamID`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetTeamEscalationPolicyResponse](../../models/operations/getteamescalationpolicyresponse.md), error**

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