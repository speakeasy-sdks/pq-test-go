# pq-c1-test

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/pq-test-go
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
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
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Auth](docs/auth/README.md)

* [Introspect](docs/auth/README.md#introspect) - Invokes the c1.api.auth.v1.Auth.Introspect method.

### [User](docs/user/README.md)

* [Get](docs/user/README.md#get) - Invokes the c1.api.user.v1.UserService.Get method.
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta and therefore, we recommend pinning usage to a specific package version.
This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated and maintained programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
