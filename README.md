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

<!-- Start SDK Available Operations -->
## Available Resources and Operations

### [PqC1Test SDK](docs/pqc1test/README.md)

* [C1APIAppV1AppResourceServiceGet](docs/pqc1test/README.md#c1apiappv1appresourceserviceget) - Invokes the c1.api.app.v1.AppResourceService.Get method.
* [C1APIAppV1AppResourceTypeServiceGet](docs/pqc1test/README.md#c1apiappv1appresourcetypeserviceget) - Invokes the c1.api.app.v1.AppResourceTypeService.Get method.
* [C1APIAppV1AppsGet](docs/pqc1test/README.md#c1apiappv1appsget) - Invokes the c1.api.app.v1.Apps.Get method.
* [C1APIAuthV1AuthIntrospect](docs/pqc1test/README.md#c1apiauthv1authintrospect) - Invokes the c1.api.auth.v1.Auth.Introspect method.
* [C1APIRequestcatalogV1RequestCatalogSearchServiceSearchEntitlements](docs/pqc1test/README.md#c1apirequestcatalogv1requestcatalogsearchservicesearchentitlements) - Invokes the c1.api.requestcatalog.v1.RequestCatalogSearchService.SearchEntitlements method.
* [C1APITaskV1TaskServiceCreateGrantTask](docs/pqc1test/README.md#c1apitaskv1taskservicecreategranttask) - Invokes the c1.api.task.v1.TaskService.CreateGrantTask method.
* [C1APITaskV1TaskServiceCreateRevokeTask](docs/pqc1test/README.md#c1apitaskv1taskservicecreaterevoketask) - Invokes the c1.api.task.v1.TaskService.CreateRevokeTask method.
* [C1APITaskV1TaskServiceGet](docs/pqc1test/README.md#c1apitaskv1taskserviceget) - Invokes the c1.api.task.v1.TaskService.Get method.
* [C1APIUserV1UserServiceGet](docs/pqc1test/README.md#c1apiuserv1userserviceget) - Invokes the c1.api.user.v1.UserService.Get method.
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta and therefore, we recommend pinning usage to a specific package version.
This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated and maintained programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
