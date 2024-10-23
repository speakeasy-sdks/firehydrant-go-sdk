# PutV1ScimV2UsersID

PUT SCIM endpoint to update a User. This endpoint is used to replace a resource's attributes.


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `UserName`                                               | **string*                                                | :heavy_minus_sign:                                       | A service provider's unique identifier for the user      |
| `Name`                                                   | [*components.Name](../../models/components/name.md)      | :heavy_minus_sign:                                       | The components of the user's name                        |
| `Emails`                                                 | [][components.Emails](../../models/components/emails.md) | :heavy_minus_sign:                                       | Email addresses for the User                             |
| `Roles`                                                  | []*string*                                               | :heavy_minus_sign:                                       | Roles for the User                                       |
| `Active`                                                 | **bool*                                                  | :heavy_minus_sign:                                       | Boolean that represents whether user is active           |