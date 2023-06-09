# Auth

### Available Operations

* [C1APIAuthV1AuthIntrospect](#c1apiauthv1authintrospect) - Invokes the c1.api.auth.v1.Auth.Introspect method.

## C1APIAuthV1AuthIntrospect

Invokes the c1.api.auth.v1.Auth.Introspect method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"pq-c1-test"
)

func main() {
    s := pqc1test.New()

    ctx := context.Background()
    res, err := s.Auth.C1APIAuthV1AuthIntrospect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APIAuthV1IntrospectResponse != nil {
        // handle response
    }
}
```
