# IncidentsAlertEntity

Incidents_AlertEntity model


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `ID`                                                                          | **string*                                                                     | :heavy_minus_sign:                                                            | N/A                                                                           |
| `Alert`                                                                       | [*components.AlertsAlertEntity](../../models/components/alertsalertentity.md) | :heavy_minus_sign:                                                            | Alerts_AlertEntity model                                                      |
| `Primary`                                                                     | **bool*                                                                       | :heavy_minus_sign:                                                            | whether or not this is the primary alert for this incident                    |