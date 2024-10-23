# PublicAPIV1IncidentsTranscriptEntity

PublicAPI_V1_Incidents_TranscriptEntity model


## Fields

| Field                                          | Type                                           | Required                                       | Description                                    |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| `ID`                                           | **string*                                      | :heavy_minus_sign:                             | The unique identifier for the transcript entry |
| `Speaker`                                      | **string*                                      | :heavy_minus_sign:                             | The speaker for the transcript entry           |
| `Start`                                        | **int*                                         | :heavy_minus_sign:                             | The start time for the transcript entry        |
| `Until`                                        | **int*                                         | :heavy_minus_sign:                             | The end time for the transcript entry          |
| `Words`                                        | **string*                                      | :heavy_minus_sign:                             | The words spoken for the transcript entry      |
| `CreatedAt`                                    | [*time.Time](https://pkg.go.dev/time#Time)     | :heavy_minus_sign:                             | The time the transcript entry was created      |
| `UpdatedAt`                                    | [*time.Time](https://pkg.go.dev/time#Time)     | :heavy_minus_sign:                             | The time the transcript entry was last updated |