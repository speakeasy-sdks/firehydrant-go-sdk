# SDKConnections
(*Integrations.Aws.Connections*)

## Overview

### Available Operations

* [List](#list) - List AWS connections

## List

Lists the available and configured AWS integration connections for the authenticated organization.

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
    res, err := s.Integrations.Aws.Connections.List(ctx, operations.GetV1IntegrationsAwsConnectionsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.IntegrationsAwsConnectionEntityPaginated != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.GetV1IntegrationsAwsConnectionsRequest](../../models/operations/getv1integrationsawsconnectionsrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.GetV1IntegrationsAwsConnectionsResponse](../../models/operations/getv1integrationsawsconnectionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |