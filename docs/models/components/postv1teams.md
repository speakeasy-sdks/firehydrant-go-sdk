# PostV1Teams

Create a new team


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `Name`                                                                  | *string*                                                                | :heavy_check_mark:                                                      | N/A                                                                     |
| `Description`                                                           | **string*                                                               | :heavy_minus_sign:                                                      | N/A                                                                     |
| `Slug`                                                                  | **string*                                                               | :heavy_minus_sign:                                                      | N/A                                                                     |
| `SlackChannelID`                                                        | **string*                                                               | :heavy_minus_sign:                                                      | The Slack channel ID that this team is associated with                  |
| `MsTeamsChannel`                                                        | [*components.MsTeamsChannel](../../models/components/msteamschannel.md) | :heavy_minus_sign:                                                      | MS Teams channel identity for channel associated with this team         |
| `Memberships`                                                           | [][components.Memberships](../../models/components/memberships.md)      | :heavy_minus_sign:                                                      | N/A                                                                     |