# PostV1LifecyclesMilestonesRequestBody


## Fields

| Field                                                                                                                            | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `Name`                                                                                                                           | *string*                                                                                                                         | :heavy_check_mark:                                                                                                               | The name of the milestone                                                                                                        |
| `Description`                                                                                                                    | *string*                                                                                                                         | :heavy_check_mark:                                                                                                               | A long-form description of the milestone's purpose                                                                               |
| `Slug`                                                                                                                           | **string*                                                                                                                        | :heavy_minus_sign:                                                                                                               | A unique identifier for the milestone. If not provided, one will be generated from the name.                                     |
| `PhaseID`                                                                                                                        | *string*                                                                                                                         | :heavy_check_mark:                                                                                                               | The ID of the phase to which the milestone should belong                                                                         |
| `Position`                                                                                                                       | **int*                                                                                                                           | :heavy_minus_sign:                                                                                                               | The position of the milestone within the phase. If not provided, the milestone will be added as the last milestone in the phase. |
| `RequiredAtMilestoneID`                                                                                                          | **string*                                                                                                                        | :heavy_minus_sign:                                                                                                               | The ID of a later milestone that cannot be started until this milestone has a timestamp populated                                |