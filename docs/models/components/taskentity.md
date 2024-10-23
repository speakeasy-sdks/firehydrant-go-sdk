# TaskEntity

TaskEntity model


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `ID`                                                                | **string*                                                           | :heavy_minus_sign:                                                  | N/A                                                                 |
| `Title`                                                             | **string*                                                           | :heavy_minus_sign:                                                  | N/A                                                                 |
| `Description`                                                       | **string*                                                           | :heavy_minus_sign:                                                  | N/A                                                                 |
| `State`                                                             | **string*                                                           | :heavy_minus_sign:                                                  | N/A                                                                 |
| `Assignee`                                                          | [*components.AuthorEntity](../../models/components/authorentity.md) | :heavy_minus_sign:                                                  | N/A                                                                 |
| `CreatedBy`                                                         | [*components.AuthorEntity](../../models/components/authorentity.md) | :heavy_minus_sign:                                                  | N/A                                                                 |
| `CreatedAt`                                                         | [*time.Time](https://pkg.go.dev/time#Time)                          | :heavy_minus_sign:                                                  | N/A                                                                 |
| `UpdatedAt`                                                         | [*time.Time](https://pkg.go.dev/time#Time)                          | :heavy_minus_sign:                                                  | N/A                                                                 |
| `DueAt`                                                             | [*time.Time](https://pkg.go.dev/time#Time)                          | :heavy_minus_sign:                                                  | N/A                                                                 |