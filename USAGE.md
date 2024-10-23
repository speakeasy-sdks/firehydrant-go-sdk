<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"log"
	"openapi"
)

func main() {
	s := openapi.New(
		openapi.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Ping.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if res.PongEntity != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->