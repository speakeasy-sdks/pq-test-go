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
    res, err := s.C1APIAppV1AppResourceServiceGet(ctx, operations.C1APIAppV1AppResourceServiceGetRequest{
        AppID: "corrupti",
        AppResourceTypeID: "provident",
        ID: "bd9d8d69-a674-4e0f-867c-c8796ed151a0",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APIAppV1AppResourceServiceGetResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->