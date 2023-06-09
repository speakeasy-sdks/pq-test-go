<!-- Start SDK Example Usage -->
```go
package main

import(
	"context"
	"log"
	"pq-c1-test"
	"pq-c1-test/pkg/models/operations"
)

func main() {
    s := pqc1test.New()

    ctx := context.Background()
    res, err := s.App.C1APIAppV1AppEntitlementUserBindingServiceListAppUsersForIdentityWithGrant(ctx, "corrupti", "provident", "distinctio")
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APIAppV1ListAppUsersForIdentityWithGrantResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->