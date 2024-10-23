# IncidentTypeEntityTemplateEntity


## Fields

| Field                                                                                                                    | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `IncidentName`                                                                                                           | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | N/A                                                                                                                      |
| `Summary`                                                                                                                | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | N/A                                                                                                                      |
| `Description`                                                                                                            | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | N/A                                                                                                                      |
| `CustomerImpactSummary`                                                                                                  | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | N/A                                                                                                                      |
| `Labels`                                                                                                                 | [*components.IncidentTypeEntityTemplateEntityLabels](../../models/components/incidenttypeentitytemplateentitylabels.md)  | :heavy_minus_sign:                                                                                                       | Arbitrary key:value pairs of labels for your incidents.                                                                  |
| `Severity`                                                                                                               | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | N/A                                                                                                                      |
| `Priority`                                                                                                               | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | N/A                                                                                                                      |
| `TagList`                                                                                                                | []*string*                                                                                                               | :heavy_minus_sign:                                                                                                       | N/A                                                                                                                      |
| `RunbookIds`                                                                                                             | []*string*                                                                                                               | :heavy_minus_sign:                                                                                                       | N/A                                                                                                                      |
| `TeamIds`                                                                                                                | []*string*                                                                                                               | :heavy_minus_sign:                                                                                                       | N/A                                                                                                                      |
| `PrivateIncident`                                                                                                        | **bool*                                                                                                                  | :heavy_minus_sign:                                                                                                       | N/A                                                                                                                      |
| `CustomFields`                                                                                                           | **string*                                                                                                                | :heavy_minus_sign:                                                                                                       | N/A                                                                                                                      |
| `Impacts`                                                                                                                | [][components.IncidentTypeEntityTemplateImpactEntity](../../models/components/incidenttypeentitytemplateimpactentity.md) | :heavy_minus_sign:                                                                                                       | N/A                                                                                                                      |