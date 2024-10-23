# Zendesk
(*Integrations.Zendesk*)

## Overview

### Available Operations

* [Search](#search)

## Search

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
    res, err := s.Integrations.Zendesk.Search(ctx, "<id>", nil)
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
| `ticketID`                                               | *string*                                                 | :heavy_check_mark:                                       | Zendesk ticket ID                                        |
| `include`                                                | **string*                                                | :heavy_minus_sign:                                       | Use to include attached_incidents                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetV1IntegrationsZendeskSearchResponse](../../models/operations/getv1integrationszendesksearchresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |