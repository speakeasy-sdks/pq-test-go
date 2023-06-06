# PqC1Test SDK

## Overview

ConductorOne API: The ConductorOne API is a HTTP API for managing ConductorOne resources.

### Available Operations

* [C1APIAppV1AppResourceServiceGet](#c1apiappv1appresourceserviceget) - Invokes the c1.api.app.v1.AppResourceService.Get method.
* [C1APIAppV1AppResourceTypeServiceGet](#c1apiappv1appresourcetypeserviceget) - Invokes the c1.api.app.v1.AppResourceTypeService.Get method.
* [C1APIAppV1AppsGet](#c1apiappv1appsget) - Invokes the c1.api.app.v1.Apps.Get method.
* [C1APIAuthV1AuthIntrospect](#c1apiauthv1authintrospect) - Invokes the c1.api.auth.v1.Auth.Introspect method.
* [C1APIRequestcatalogV1RequestCatalogSearchServiceSearchEntitlements](#c1apirequestcatalogv1requestcatalogsearchservicesearchentitlements) - Invokes the c1.api.requestcatalog.v1.RequestCatalogSearchService.SearchEntitlements method.
* [C1APITaskV1TaskServiceCreateGrantTask](#c1apitaskv1taskservicecreategranttask) - Invokes the c1.api.task.v1.TaskService.CreateGrantTask method.
* [C1APITaskV1TaskServiceCreateRevokeTask](#c1apitaskv1taskservicecreaterevoketask) - Invokes the c1.api.task.v1.TaskService.CreateRevokeTask method.
* [C1APITaskV1TaskServiceGet](#c1apitaskv1taskserviceget) - Invokes the c1.api.task.v1.TaskService.Get method.
* [C1APIUserV1UserServiceGet](#c1apiuserv1userserviceget) - Invokes the c1.api.user.v1.UserService.Get method.

## C1APIAppV1AppResourceServiceGet

Invokes the c1.api.app.v1.AppResourceService.Get method.

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
    res, err := s.PqC1Test.C1APIAppV1AppResourceServiceGet(ctx, operations.C1APIAppV1AppResourceServiceGetRequest{
        AppID: "ipsam",
        AppResourceTypeID: "repellendus",
        ID: "fc2ddf7c-c78c-4a1b-a928-fc816742cb73",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APIAppV1AppResourceServiceGetResponse != nil {
        // handle response
    }
}
```

## C1APIAppV1AppResourceTypeServiceGet

Invokes the c1.api.app.v1.AppResourceTypeService.Get method.

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
    res, err := s.PqC1Test.C1APIAppV1AppResourceTypeServiceGet(ctx, operations.C1APIAppV1AppResourceTypeServiceGetRequest{
        AppID: "excepturi",
        ID: "20592939-6fea-4759-aeb1-0faaa2352c59",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APIAppV1AppResourceTypeServiceGetResponse != nil {
        // handle response
    }
}
```

## C1APIAppV1AppsGet

Invokes the c1.api.app.v1.Apps.Get method.

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
    res, err := s.PqC1Test.C1APIAppV1AppsGet(ctx, operations.C1APIAppV1AppsGetRequest{
        ID: "55907aff-1a3a-42fa-9467-739251aa52c3",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APIAppV1GetAppResponse != nil {
        // handle response
    }
}
```

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
    res, err := s.PqC1Test.C1APIAuthV1AuthIntrospect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APIAuthV1IntrospectResponse != nil {
        // handle response
    }
}
```

## C1APIRequestcatalogV1RequestCatalogSearchServiceSearchEntitlements

Invokes the c1.api.requestcatalog.v1.RequestCatalogSearchService.SearchEntitlements method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"pq-c1-test"
	"pq-c1-test/pkg/models/shared"
)

func main() {
    s := pqc1test.New()

    ctx := context.Background()
    res, err := s.PqC1Test.C1APIRequestcatalogV1RequestCatalogSearchServiceSearchEntitlements(ctx, shared.C1APIRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest{
        EntitlementAlias: pqc1test.String("tenetur"),
        ExpandMask: &shared.C1APIAppV1AppEntitlementExpandMask{
            Paths: []string{
                "id",
                "possimus",
            },
        },
        PageSize: pqc1test.Float64(135.71),
        PageToken: pqc1test.String("quasi"),
        Query: pqc1test.String("error"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APIRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsResponse != nil {
        // handle response
    }
}
```

## C1APITaskV1TaskServiceCreateGrantTask

Invokes the c1.api.task.v1.TaskService.CreateGrantTask method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"pq-c1-test"
	"pq-c1-test/pkg/models/shared"
)

func main() {
    s := pqc1test.New()

    ctx := context.Background()
    res, err := s.PqC1Test.C1APITaskV1TaskServiceCreateGrantTask(ctx, shared.C1APITaskV1TaskServiceCreateGrantRequest{
        AppEntitlementID: pqc1test.String("temporibus"),
        AppID: pqc1test.String("laborum"),
        AppUserID: pqc1test.String("quasi"),
        Description: pqc1test.String("reiciendis"),
        ExpandMask: &shared.C1APITaskV1TaskExpandMask{
            Paths: []string{
                "vero",
                "nihil",
                "praesentium",
                "voluptatibus",
            },
        },
        GrantDuration: pqc1test.String("ipsa"),
        IdentityUserID: pqc1test.String("omnis"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APITaskV1TaskServiceCreateGrantResponse != nil {
        // handle response
    }
}
```

## C1APITaskV1TaskServiceCreateRevokeTask

Invokes the c1.api.task.v1.TaskService.CreateRevokeTask method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"pq-c1-test"
	"pq-c1-test/pkg/models/shared"
)

func main() {
    s := pqc1test.New()

    ctx := context.Background()
    res, err := s.PqC1Test.C1APITaskV1TaskServiceCreateRevokeTask(ctx, shared.C1APITaskV1TaskServiceCreateRevokeRequest{
        AppEntitlementID: pqc1test.String("voluptate"),
        AppID: pqc1test.String("cum"),
        AppUserID: pqc1test.String("perferendis"),
        Description: pqc1test.String("doloremque"),
        ExpandMask: &shared.C1APITaskV1TaskExpandMask{
            Paths: []string{
                "ut",
                "maiores",
            },
        },
        IdentityUserID: pqc1test.String("dicta"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APITaskV1TaskServiceCreateRevokeResponse != nil {
        // handle response
    }
}
```

## C1APITaskV1TaskServiceGet

Invokes the c1.api.task.v1.TaskService.Get method.

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
    res, err := s.PqC1Test.C1APITaskV1TaskServiceGet(ctx, operations.C1APITaskV1TaskServiceGetRequest{
        ID: "5471b5e6-e13b-499d-888e-1e91e450ad2a",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APITaskV1TaskServiceGetResponse != nil {
        // handle response
    }
}
```

## C1APIUserV1UserServiceGet

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
    res, err := s.PqC1Test.C1APIUserV1UserServiceGet(ctx, operations.C1APIUserV1UserServiceGetRequest{
        ID: "bd442698-02d5-402a-94bb-4f63c969e9a3",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APIUserV1UserServiceGetResponse != nil {
        // handle response
    }
}
```
