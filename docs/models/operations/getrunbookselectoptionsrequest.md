# GetRunbookSelectOptionsRequest


## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `IntegrationSlug`                                                                            | *string*                                                                                     | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `ActionSlug`                                                                                 | *string*                                                                                     | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `Field`                                                                                      | *string*                                                                                     | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `Query`                                                                                      | **string*                                                                                    | :heavy_minus_sign:                                                                           | Text string of a query for filtering values.                                                 |
| `Scope`                                                                                      | **string*                                                                                    | :heavy_minus_sign:                                                                           | Generic params used to add specificity (eg an id of some kind) to the select options request |
| `PerPage`                                                                                    | **int*                                                                                       | :heavy_minus_sign:                                                                           | Maximum number of items to return.                                                           |