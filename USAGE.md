<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"firehydrant"
	"log"
)

func main() {
	s := firehydrant.New(
		firehydrant.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.AccountSettings.GetAiPreferences(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.AIEntitiesPreferencesEntity != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->