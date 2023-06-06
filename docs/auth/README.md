# Auth

### Available Operations

* [Introspect](#introspect) - Invokes the c1.api.auth.v1.Auth.Introspect method.

## Introspect

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
    res, err := s.Auth.Introspect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.IntrospectResponse != nil {
        // handle response
    }
}
```
