# IncidentAttachmentEntity

IncidentAttachmentEntity model


## Fields

| Field                                                                                                   | Type                                                                                                    | Required                                                                                                | Description                                                                                             |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `FileName`                                                                                              | **string*                                                                                               | :heavy_minus_sign:                                                                                      | N/A                                                                                                     |
| `FileContentType`                                                                                       | **string*                                                                                               | :heavy_minus_sign:                                                                                      | N/A                                                                                                     |
| `SignedURL`                                                                                             | **string*                                                                                               | :heavy_minus_sign:                                                                                      | N/A                                                                                                     |
| `MediaType`                                                                                             | **string*                                                                                               | :heavy_minus_sign:                                                                                      | N/A                                                                                                     |
| `Description`                                                                                           | **string*                                                                                               | :heavy_minus_sign:                                                                                      | N/A                                                                                                     |
| `ExternalID`                                                                                            | **string*                                                                                               | :heavy_minus_sign:                                                                                      | N/A                                                                                                     |
| `FileSize`                                                                                              | **int*                                                                                                  | :heavy_minus_sign:                                                                                      | N/A                                                                                                     |
| `Status`                                                                                                | [*components.IncidentAttachmentEntityStatus](../../models/components/incidentattachmententitystatus.md) | :heavy_minus_sign:                                                                                      | N/A                                                                                                     |
| `Versions`                                                                                              | [*components.Versions](../../models/components/versions.md)                                             | :heavy_minus_sign:                                                                                      | An object with keys that designate a specific version or size of the attachment                         |