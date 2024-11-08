# PartialUpdateIncidentImpactsResponse


## Fields

| Field                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | Type                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | Required                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `HTTPMeta`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | [components.HTTPMetadata](../../models/components/httpmetadata.md)                                                                                                                                                                                                                                                                                                                                                                                                                                                              | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | N/A                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| `IncidentEntity`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                | [*components.IncidentEntity](../../models/components/incidententity.md)                                                                                                                                                                                                                                                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | Allows updating an incident's impacted infrastructure, with the option to<br/>move the incident into a different milestone and provide a note to update<br/>the incident timeline and any attached status pages. If this method is<br/>requested with the PUT verb, impacts will be completely replaced with the<br/>information in the request body, even if not provided (effectively clearing<br/>all impacts). If this method is requested with the PATCH verb, the provided<br/>impacts will be added or updated, but no impacts will be removed.<br/> |