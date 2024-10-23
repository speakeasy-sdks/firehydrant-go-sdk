# Scim
(*Scim*)

## Overview

Operations about scims

### Available Operations

* [UpdateGroup](#updategroup) - Updates a Team (via the Group protocol) and assigns members to that team.
* [DeleteGroup](#deletegroup) - Deletes a Team (via the Group protocol)
* [GetGroup](#getgroup) - Lists a Team (via the Group protocol)
* [CreateGroup](#creategroup) - Creates a new Team (via the Group protocol) and assigns members to that team.
* [ListGroups](#listgroups) - List all Teams (via the Group protocol)
* [DeleteUser](#deleteuser) - Deletes a User using SCIM protocol
* [UpdateUser](#updateuser) - Updates a User using SCIM protocol
* [GetUser](#getuser) - Lists a User (via the User protocol)
* [CreateUser](#createuser) - Creates a new User using SCIM protocol
* [ListUsers](#listusers) - Gets a list of Users using SCIM protocol

## UpdateGroup

SCIM endpoint to update a Team (Colloquial for Group in the SCIM protocol). Any members defined in the payload will be assigned to the team with no defined role, any missing members will be removed from the team.

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
    res, err := s.Scim.UpdateGroup(ctx, "<id>", components.PutV1ScimV2GroupsID{
        DisplayName: "Ibrahim.Kuhic40",
        Members: []components.PutV1ScimV2GroupsIDMembers{
            components.PutV1ScimV2GroupsIDMembers{
                Value: "<value>",
            },
            components.PutV1ScimV2GroupsIDMembers{
                Value: "<value>",
            },
            components.PutV1ScimV2GroupsIDMembers{
                Value: "<value>",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `id`                                                                             | *string*                                                                         | :heavy_check_mark:                                                               | N/A                                                                              |
| `putV1ScimV2GroupsID`                                                            | [components.PutV1ScimV2GroupsID](../../models/components/putv1scimv2groupsid.md) | :heavy_check_mark:                                                               | N/A                                                                              |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.PutV1ScimV2GroupsIDResponse](../../models/operations/putv1scimv2groupsidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteGroup

SCIM endpoint to delete a Team (Colloquial for Group in the SCIM protocol).

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
    res, err := s.Scim.DeleteGroup(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteV1ScimV2GroupsIDResponse](../../models/operations/deletev1scimv2groupsidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetGroup

SCIM endpoint that lists a Team (Colloquial for Group in the SCIM protocol)

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
    res, err := s.Scim.GetGroup(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetV1ScimV2GroupsIDResponse](../../models/operations/getv1scimv2groupsidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateGroup

SCIM endpoint to create a new Team (Colloquial for Group in the SCIM protocol). Any members defined in the payload will be assigned to the team with no defined role.

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
    res, err := s.Scim.CreateGroup(ctx, components.PostV1ScimV2Groups{
        DisplayName: "Germaine.Sipes31",
        Members: []components.PostV1ScimV2GroupsMembers{
            components.PostV1ScimV2GroupsMembers{
                Value: "<value>",
            },
            components.PostV1ScimV2GroupsMembers{
                Value: "<value>",
            },
            components.PostV1ScimV2GroupsMembers{
                Value: "<value>",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.PostV1ScimV2Groups](../../models/components/postv1scimv2groups.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.PostV1ScimV2GroupsResponse](../../models/operations/postv1scimv2groupsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListGroups

SCIM endpoint that lists all Teams (Colloquial for Group in the SCIM protocol)

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
    res, err := s.Scim.ListGroups(ctx, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                           | Type                                                                                                                                                                                                | Required                                                                                                                                                                                            | Description                                                                                                                                                                                         |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                  | The context to use for the request.                                                                                                                                                                 |
| `startIndex`                                                                                                                                                                                        | **int*                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                  | N/A                                                                                                                                                                                                 |
| `count`                                                                                                                                                                                             | **int*                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                  | N/A                                                                                                                                                                                                 |
| `filter`                                                                                                                                                                                            | **string*                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                  | This is a string used to query groups by displayName.<br/><br/>        Proper example syntax for this would be `?filter=displayName eq "My Team Name"`.<br/>        Currently we only support the `eq` operator |
| `opts`                                                                                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                  | The options for this request.                                                                                                                                                                       |

### Response

**[*operations.GetV1ScimV2GroupsResponse](../../models/operations/getv1scimv2groupsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteUser

SCIM endpoint to delete a User. This endpoint will deactivate a confirmed User record in our system.

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
    res, err := s.Scim.DeleteUser(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteV1ScimV2UsersIDResponse](../../models/operations/deletev1scimv2usersidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UpdateUser

PUT SCIM endpoint to update a User. This endpoint is used to replace a resource's attributes.

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
    res, err := s.Scim.UpdateUser(ctx, "<id>", components.PutV1ScimV2UsersID{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `id`                                                                           | *string*                                                                       | :heavy_check_mark:                                                             | N/A                                                                            |
| `putV1ScimV2UsersID`                                                           | [components.PutV1ScimV2UsersID](../../models/components/putv1scimv2usersid.md) | :heavy_check_mark:                                                             | N/A                                                                            |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.PutV1ScimV2UsersIDResponse](../../models/operations/putv1scimv2usersidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetUser

SCIM endpoint that lists a User

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
    res, err := s.Scim.GetUser(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetV1ScimV2UsersIDResponse](../../models/operations/getv1scimv2usersidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## CreateUser

SCIM endpoint to create and provision a new User. This endpoint will provision the User, which allows them to accept their account throught their IDP or via the Forgot Password flow.

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
    res, err := s.Scim.CreateUser(ctx, components.PostV1ScimV2Users{
        UserName: "Jimmy12",
        Name: components.PostV1ScimV2UsersName{
            FamilyName: "<value>",
            GivenName: "<value>",
        },
        Emails: []components.PostV1ScimV2UsersEmails{
            components.PostV1ScimV2UsersEmails{
                Value: "<value>",
                Primary: false,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [components.PostV1ScimV2Users](../../models/components/postv1scimv2users.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.PostV1ScimV2UsersResponse](../../models/operations/postv1scimv2usersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## ListUsers

SCIM endpoint that lists users. This endpoint will display a list of Users currently in the system.

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
    res, err := s.Scim.ListUsers(ctx, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                          | Type                                                                                                                                                                                                                                               | Required                                                                                                                                                                                                                                           | Description                                                                                                                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                              | :heavy_check_mark:                                                                                                                                                                                                                                 | The context to use for the request.                                                                                                                                                                                                                |
| `filter`                                                                                                                                                                                                                                           | **string*                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                 | This is a string used to query users by either userName or email.<br/><br/>        Proper example syntax for this would be `?filter=userName eq john` or `?filter=userName eq "john@firehydrant.com"`.<br/>        Currently we only support the `eq` operator |
| `startIndex`                                                                                                                                                                                                                                       | **int*                                                                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                                                 | This is an integer which represents a pagination offset                                                                                                                                                                                            |
| `count`                                                                                                                                                                                                                                            | **int*                                                                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                                                 | This is an integer which represents the number of items per page in the response                                                                                                                                                                   |
| `opts`                                                                                                                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                                                 | The options for this request.                                                                                                                                                                                                                      |

### Response

**[*operations.GetV1ScimV2UsersResponse](../../models/operations/getv1scimv2usersresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |