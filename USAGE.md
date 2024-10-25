<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"firehydrant"
	"firehydrant/models/components"
	"log"
)

func main() {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Incidents.Create(ctx, components.CreateIncident{
		Name: "<value>",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.IncidentEntity != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->