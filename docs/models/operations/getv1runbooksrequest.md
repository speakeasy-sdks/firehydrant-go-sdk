# GetV1RunbooksRequest


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `Page`                                                                  | **int*                                                                  | :heavy_minus_sign:                                                      | N/A                                                                     |
| `PerPage`                                                               | **int*                                                                  | :heavy_minus_sign:                                                      | N/A                                                                     |
| `Name`                                                                  | **string*                                                               | :heavy_minus_sign:                                                      | A query to search runbooks by their name                                |
| `Owners`                                                                | **string*                                                               | :heavy_minus_sign:                                                      | A query to search runbooks by their owners                              |
| `Sort`                                                                  | [*operations.QueryParamSort](../../models/operations/queryparamsort.md) | :heavy_minus_sign:                                                      | Sort runbooks by their updated date. Accepts 'asc', 'desc'              |