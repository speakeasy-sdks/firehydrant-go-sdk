# CustomFields
(*CustomFields*)

## Overview

### Available Operations

* [DeleteDefinition](#deletedefinition) - Delete a custom field definition
* [UpdateDefinition](#updatedefinition) - Update custom field definition
* [CreateDefinition](#createdefinition) - Create custom field definition
* [ListDefinitions](#listdefinitions) - List custom field definitions
* [GetSelectOptions](#getselectoptions) - Get custom field permissible values

## DeleteDefinition

Delete a custom field definition

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
    res, err := s.CustomFields.DeleteDefinition(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.OrganizationsCustomFieldDefinitionEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `fieldID`                                                | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteV1CustomFieldsDefinitionsFieldIDResponse](../../models/operations/deletev1customfieldsdefinitionsfieldidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateDefinition

Update a single custom field definition

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
    res, err := s.CustomFields.UpdateDefinition(ctx, "<id>", components.PatchV1CustomFieldsDefinitionsFieldID{})
    if err != nil {
        log.Fatal(err)
    }
    if res.OrganizationsCustomFieldDefinitionEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `fieldID`                                                                                                            | *string*                                                                                                             | :heavy_check_mark:                                                                                                   | N/A                                                                                                                  |
| `patchV1CustomFieldsDefinitionsFieldID`                                                                              | [components.PatchV1CustomFieldsDefinitionsFieldID](../../models/components/patchv1customfieldsdefinitionsfieldid.md) | :heavy_check_mark:                                                                                                   | N/A                                                                                                                  |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.PatchV1CustomFieldsDefinitionsFieldIDResponse](../../models/operations/patchv1customfieldsdefinitionsfieldidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateDefinition

Create a new custom field definition

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
    res, err := s.CustomFields.CreateDefinition(ctx, components.PostV1CustomFieldsDefinitions{
        DisplayName: "Hank_Blick88",
        FieldType: "<value>",
        Required: true,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OrganizationsCustomFieldDefinitionEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [components.PostV1CustomFieldsDefinitions](../../models/components/postv1customfieldsdefinitions.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.PostV1CustomFieldsDefinitionsResponse](../../models/operations/postv1customfieldsdefinitionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListDefinitions

List all custom field definitions

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
    res, err := s.CustomFields.ListDefinitions(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.OrganizationsCustomFieldDefinitionEntity != nil {
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

**[*operations.GetV1CustomFieldsDefinitionsResponse](../../models/operations/getv1customfieldsdefinitionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetSelectOptions

Get the permissible values for the a currently active custom select or multi-select field.

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
    res, err := s.CustomFields.GetSelectOptions(ctx, "<id>", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.OrganizationsCustomFieldDefinitionEntity != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `fieldID`                                                    | *string*                                                     | :heavy_check_mark:                                           | N/A                                                          |
| `query`                                                      | **string*                                                    | :heavy_minus_sign:                                           | Text string of a query for filtering values.                 |
| `allVersions`                                                | **bool*                                                      | :heavy_minus_sign:                                           | If true, return all versions of the custom field definition. |
| `opts`                                                       | [][operations.Option](../../models/operations/option.md)     | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*operations.GetV1CustomFieldsDefinitionsFieldIDSelectOptionsResponse](../../models/operations/getv1customfieldsdefinitionsfieldidselectoptionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |