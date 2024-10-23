# GetV1MetricsInfraTypeRequest


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `InfraType`                                                    | [operations.InfraType](../../models/operations/infratype.md)   | :heavy_check_mark:                                             | N/A                                                            |
| `StartDate`                                                    | [*types.Date](../../types/date.md)                             | :heavy_minus_sign:                                             | The start date to return metrics from; defaults to 30 days ago |
| `EndDate`                                                      | [*types.Date](../../types/date.md)                             | :heavy_minus_sign:                                             | The end date to return metrics from, defaults to today         |