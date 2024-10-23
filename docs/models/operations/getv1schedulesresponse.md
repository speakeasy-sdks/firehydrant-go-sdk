# GetV1SchedulesResponse


## Fields

| Field                                                                                     | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `HTTPMeta`                                                                                | [components.HTTPMetadata](../../models/components/httpmetadata.md)                        | :heavy_check_mark:                                                                        | N/A                                                                                       |
| `ScheduleEntityPaginated`                                                                 | [*components.ScheduleEntityPaginated](../../models/components/scheduleentitypaginated.md) | :heavy_minus_sign:                                                                        | List all known schedules in FireHydrant as pulled from external sources                   |