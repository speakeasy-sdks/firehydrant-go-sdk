# NuncConnections
(*NuncConnections*)

## Overview

### Available Operations

* [List](#list) - Lists the FireHydrant hosted status pages
* [Subscribe](#subscribe) - Subscribes a comma-separated string of emails to status page updates
* [ListSubscribers](#listsubscribers) - Retrieves the list of subscribers for a status page
* [Delete](#delete) - Delete a FireHydrant hosted status page
* [Get](#get) - Retrieve information about a FireHydrant hosted status page
* [DeleteComponentGroup](#deletecomponentgroup) - Delete a component group displayed on a FireHydrant status page
* [UpdateComponentGroup](#updatecomponentgroup) - Update a component group to be displayed on a FireHydrant status page
* [AddComponentGroup](#addcomponentgroup) - Add a component group to be displayed on a FireHydrant status page
* [UpdateLink](#updatelink) - Update a link to be displayed on a FireHydrant status page
* [AddLink](#addlink) - Update a link to be displayed on a FireHydrant status page
* [UpdateImage](#updateimage) - Add or replace an image attached to a FireHydrant status page
* [UnsubscribeSubscribers](#unsubscribesubscribers) - Unsubscribes one or more status page subscribers.
* [Update](#update) - Update a FireHydrant hosted status page
* [DeleteLink](#deletelink) - Delete a link displayed on a FireHydrant status page
* [DeleteImage](#deleteimage) - Delete an image attached to a FireHydrant status page

## List

Lists the information displayed as part of your FireHydrant hosted status pages.

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
    res, err := s.NuncConnections.List(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.NuncConnectionEntityPaginated != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetV1NuncConnectionsResponse](../../models/operations/getv1nuncconnectionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Subscribe

Subscribes a comma-separated string of emails to status page updates

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"firehydrant/models/operations"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.NuncConnections.Subscribe(ctx, "<id>", operations.PostV1NuncConnectionsNuncConnectionIDSubscribersRequestBody{
        Emails: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.NuncEmailSubscribersEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                        | Type                                                                                                                                                             | Required                                                                                                                                                         | Description                                                                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                                            | :heavy_check_mark:                                                                                                                                               | The context to use for the request.                                                                                                                              |
| `nuncConnectionID`                                                                                                                                               | *string*                                                                                                                                                         | :heavy_check_mark:                                                                                                                                               | N/A                                                                                                                                                              |
| `requestBody`                                                                                                                                                    | [operations.PostV1NuncConnectionsNuncConnectionIDSubscribersRequestBody](../../models/operations/postv1nuncconnectionsnuncconnectionidsubscribersrequestbody.md) | :heavy_check_mark:                                                                                                                                               | N/A                                                                                                                                                              |
| `opts`                                                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                                                         | :heavy_minus_sign:                                                                                                                                               | The options for this request.                                                                                                                                    |

### Response

**[*operations.PostV1NuncConnectionsNuncConnectionIDSubscribersResponse](../../models/operations/postv1nuncconnectionsnuncconnectionidsubscribersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListSubscribers

Retrieves the list of subscribers for a status page.

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
    res, err := s.NuncConnections.ListSubscribers(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.NuncEmailSubscribersEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `nuncConnectionID`                                       | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetV1NuncConnectionsNuncConnectionIDSubscribersResponse](../../models/operations/getv1nuncconnectionsnuncconnectionidsubscribersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Delete

Delete a FireHydrant hosted status page, stopping updates of your incidents to it.

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
    res, err := s.NuncConnections.Delete(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.NuncConnectionEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `nuncConnectionID`                                       | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteV1NuncConnectionsNuncConnectionIDResponse](../../models/operations/deletev1nuncconnectionsnuncconnectionidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Get

Retrieve the information displayed as part of your FireHydrant hosted status page.

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
    res, err := s.NuncConnections.Get(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.NuncConnectionEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `nuncConnectionID`                                       | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetV1NuncConnectionsNuncConnectionIDResponse](../../models/operations/getv1nuncconnectionsnuncconnectionidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteComponentGroup

Delete a component group displayed on a FireHydrant status page

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
    res, err := s.NuncConnections.DeleteComponentGroup(ctx, "<id>", "<id>")
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
| `nuncConnectionID`                                       | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `groupID`                                                | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDResponse](../../models/operations/deletev1nuncconnectionsnuncconnectionidcomponentgroupsgroupidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateComponentGroup

Update a component group to be displayed on a FireHydrant status page

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
    res, err := s.NuncConnections.UpdateComponentGroup(ctx, "<id>", "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                 | Type                                                                                                                                                                                      | Required                                                                                                                                                                                  | Description                                                                                                                                                                               |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                        | The context to use for the request.                                                                                                                                                       |
| `nuncConnectionID`                                                                                                                                                                        | *string*                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                        | N/A                                                                                                                                                                                       |
| `groupID`                                                                                                                                                                                 | *string*                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                        | N/A                                                                                                                                                                                       |
| `requestBody`                                                                                                                                                                             | [*operations.PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDRequestBody](../../models/operations/patchv1nuncconnectionsnuncconnectionidcomponentgroupsgroupidrequestbody.md) | :heavy_minus_sign:                                                                                                                                                                        | N/A                                                                                                                                                                                       |
| `opts`                                                                                                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                        | The options for this request.                                                                                                                                                             |

### Response

**[*operations.PatchV1NuncConnectionsNuncConnectionIDComponentGroupsGroupIDResponse](../../models/operations/patchv1nuncconnectionsnuncconnectionidcomponentgroupsgroupidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## AddComponentGroup

Add a component group to be displayed on a FireHydrant status page

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"firehydrant/models/operations"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.NuncConnections.AddComponentGroup(ctx, "<id>", operations.PostV1NuncConnectionsNuncConnectionIDComponentGroupsRequestBody{
        Name: "<value>",
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

| Parameter                                                                                                                                                                | Type                                                                                                                                                                     | Required                                                                                                                                                                 | Description                                                                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                    | :heavy_check_mark:                                                                                                                                                       | The context to use for the request.                                                                                                                                      |
| `nuncConnectionID`                                                                                                                                                       | *string*                                                                                                                                                                 | :heavy_check_mark:                                                                                                                                                       | N/A                                                                                                                                                                      |
| `requestBody`                                                                                                                                                            | [operations.PostV1NuncConnectionsNuncConnectionIDComponentGroupsRequestBody](../../models/operations/postv1nuncconnectionsnuncconnectionidcomponentgroupsrequestbody.md) | :heavy_check_mark:                                                                                                                                                       | N/A                                                                                                                                                                      |
| `opts`                                                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                                                 | :heavy_minus_sign:                                                                                                                                                       | The options for this request.                                                                                                                                            |

### Response

**[*operations.PostV1NuncConnectionsNuncConnectionIDComponentGroupsResponse](../../models/operations/postv1nuncconnectionsnuncconnectionidcomponentgroupsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateLink

Update a link to be displayed on a FireHydrant status page

### Example Usage

```go
package main

import(
	"firehydrant"
	"context"
	"firehydrant/models/components"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.NuncConnections.UpdateLink(ctx, "<id>", "<id>", components.PatchV1NuncConnectionsNuncConnectionIDLinksLinkID{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `nuncConnectionID`                                                                                                                           | *string*                                                                                                                                     | :heavy_check_mark:                                                                                                                           | N/A                                                                                                                                          |
| `linkID`                                                                                                                                     | *string*                                                                                                                                     | :heavy_check_mark:                                                                                                                           | N/A                                                                                                                                          |
| `patchV1NuncConnectionsNuncConnectionIDLinksLinkID`                                                                                          | [components.PatchV1NuncConnectionsNuncConnectionIDLinksLinkID](../../models/components/patchv1nuncconnectionsnuncconnectionidlinkslinkid.md) | :heavy_check_mark:                                                                                                                           | N/A                                                                                                                                          |
| `opts`                                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                                     | :heavy_minus_sign:                                                                                                                           | The options for this request.                                                                                                                |

### Response

**[*operations.PatchV1NuncConnectionsNuncConnectionIDLinksLinkIDResponse](../../models/operations/patchv1nuncconnectionsnuncconnectionidlinkslinkidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## AddLink

Add a link to be displayed on a FireHydrant status page

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
    res, err := s.NuncConnections.AddLink(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.NuncConnectionEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `nuncConnectionID`                                       | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.PostV1NuncConnectionsNuncConnectionIDLinksResponse](../../models/operations/postv1nuncconnectionsnuncconnectionidlinksresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateImage

Add or replace an image attached to a FireHydrant status page

### Example Usage

```go
package main

import(
	"firehydrant"
	"os"
	"context"
	"log"
)

func main() {
    s := firehydrant.New(
        firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    content, fileErr := os.Open("example.file")
    if fileErr != nil {
        panic(fileErr)
    }

    ctx := context.Background()
    res, err := s.NuncConnections.UpdateImage(ctx, "<id>", "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.NuncConnectionEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                     | Type                                                                                                                                                          | Required                                                                                                                                                      | Description                                                                                                                                                   |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                                                                                         | :heavy_check_mark:                                                                                                                                            | The context to use for the request.                                                                                                                           |
| `nuncConnectionID`                                                                                                                                            | *string*                                                                                                                                                      | :heavy_check_mark:                                                                                                                                            | N/A                                                                                                                                                           |
| `type_`                                                                                                                                                       | *string*                                                                                                                                                      | :heavy_check_mark:                                                                                                                                            | N/A                                                                                                                                                           |
| `requestBody`                                                                                                                                                 | [*operations.PutV1NuncConnectionsNuncConnectionIDImagesTypeRequestBody](../../models/operations/putv1nuncconnectionsnuncconnectionidimagestyperequestbody.md) | :heavy_minus_sign:                                                                                                                                            | N/A                                                                                                                                                           |
| `opts`                                                                                                                                                        | [][operations.Option](../../models/operations/option.md)                                                                                                      | :heavy_minus_sign:                                                                                                                                            | The options for this request.                                                                                                                                 |

### Response

**[*operations.PutV1NuncConnectionsNuncConnectionIDImagesTypeResponse](../../models/operations/putv1nuncconnectionsnuncconnectionidimagestyperesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UnsubscribeSubscribers

Unsubscribes one or more status page subscribers.

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
    res, err := s.NuncConnections.UnsubscribeSubscribers(ctx, "<id>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.NuncEmailSubscribersEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `nuncConnectionID`                                       | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `subscriberIds`                                          | *string*                                                 | :heavy_check_mark:                                       | A list of subscriber IDs to unsubscribe.                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteV1NuncConnectionsNuncConnectionIDSubscribersResponse](../../models/operations/deletev1nuncconnectionsnuncconnectionidsubscribersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## Update

Update your company's information and other components in the specified FireHydrant hosted status page.

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
    res, err := s.NuncConnections.Update(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.NuncConnectionEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                 | Type                                                                                                                                      | Required                                                                                                                                  | Description                                                                                                                               |
| ----------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                                                     | :heavy_check_mark:                                                                                                                        | The context to use for the request.                                                                                                       |
| `nuncConnectionID`                                                                                                                        | *string*                                                                                                                                  | :heavy_check_mark:                                                                                                                        | N/A                                                                                                                                       |
| `requestBody`                                                                                                                             | [*operations.PutV1NuncConnectionsNuncConnectionIDRequestBody](../../models/operations/putv1nuncconnectionsnuncconnectionidrequestbody.md) | :heavy_minus_sign:                                                                                                                        | N/A                                                                                                                                       |
| `opts`                                                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                                                  | :heavy_minus_sign:                                                                                                                        | The options for this request.                                                                                                             |

### Response

**[*operations.PutV1NuncConnectionsNuncConnectionIDResponse](../../models/operations/putv1nuncconnectionsnuncconnectionidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteLink

Delete a link displayed on a FireHydrant status page

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
    res, err := s.NuncConnections.DeleteLink(ctx, "<id>", "<id>")
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
| `nuncConnectionID`                                       | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `linkID`                                                 | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteV1NuncConnectionsNuncConnectionIDLinksLinkIDResponse](../../models/operations/deletev1nuncconnectionsnuncconnectionidlinkslinkidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteImage

Delete an image attached to a FireHydrant status page

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
    res, err := s.NuncConnections.DeleteImage(ctx, "<id>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.NuncConnectionEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `nuncConnectionID`                                       | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `type_`                                                  | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteV1NuncConnectionsNuncConnectionIDImagesTypeResponse](../../models/operations/deletev1nuncconnectionsnuncconnectionidimagestyperesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |