# User

### Available Operations

* [Get](#get) - Invokes the c1.api.user.v1.UserService.Get method.

## Get

Invokes the c1.api.user.v1.UserService.Get method.

### Example Usage

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
    res, err := s.User.Get(ctx, operations.C1APIUserV1UserServiceGetRequest{
        ID: "89bd9d8d-69a6-474e-8f46-7cc8796ed151",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UserServiceGetResponse != nil {
        // handle response
    }
}
```
