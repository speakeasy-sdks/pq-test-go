# App

### Available Operations

* [C1APIAppV1AppEntitlementUserBindingServiceListAppUsersForIdentityWithGrant](#c1apiappv1appentitlementuserbindingservicelistappusersforidentitywithgrant) - Invokes the c1.api.app.v1.AppEntitlementUserBindingService.ListAppUsersForIdentityWithGrant method.
* [C1APIAppV1AppEntitlementsGet](#c1apiappv1appentitlementsget) - Invokes the c1.api.app.v1.AppEntitlements.Get method.
* [C1APIAppV1AppResourceServiceGet](#c1apiappv1appresourceserviceget) - Invokes the c1.api.app.v1.AppResourceService.Get method.
* [C1APIAppV1AppResourceTypeServiceGet](#c1apiappv1appresourcetypeserviceget) - Invokes the c1.api.app.v1.AppResourceTypeService.Get method.
* [C1APIAppV1AppsGet](#c1apiappv1appsget) - Invokes the c1.api.app.v1.Apps.Get method.

## C1APIAppV1AppEntitlementUserBindingServiceListAppUsersForIdentityWithGrant

Invokes the c1.api.app.v1.AppEntitlementUserBindingService.ListAppUsersForIdentityWithGrant method.

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
    res, err := s.App.C1APIAppV1AppEntitlementUserBindingServiceListAppUsersForIdentityWithGrant(ctx, "quibusdam", "unde", "nulla")
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APIAppV1ListAppUsersForIdentityWithGrantResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `appEntitlementID`                                    | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |
| `appID`                                               | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |
| `identityUserID`                                      | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.C1APIAppV1AppEntitlementUserBindingServiceListAppUsersForIdentityWithGrantResponse](../../models/operations/c1apiappv1appentitlementuserbindingservicelistappusersforidentitywithgrantresponse.md), error**


## C1APIAppV1AppEntitlementsGet

Invokes the c1.api.app.v1.AppEntitlements.Get method.

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
    res, err := s.App.C1APIAppV1AppEntitlementsGet(ctx, "corrupti", "illum")
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APIAppV1GetAppEntitlementResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `appID`                                               | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.C1APIAppV1AppEntitlementsGetResponse](../../models/operations/c1apiappv1appentitlementsgetresponse.md), error**


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
    res, err := s.App.C1APIAppV1AppResourceServiceGet(ctx, "vel", "error", "deserunt")
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APIAppV1AppResourceServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `appID`                                               | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |
| `appResourceTypeID`                                   | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.C1APIAppV1AppResourceServiceGetResponse](../../models/operations/c1apiappv1appresourceservicegetresponse.md), error**


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
    res, err := s.App.C1APIAppV1AppResourceTypeServiceGet(ctx, "suscipit", "iure")
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APIAppV1AppResourceTypeServiceGetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `appID`                                               | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.C1APIAppV1AppResourceTypeServiceGetResponse](../../models/operations/c1apiappv1appresourcetypeservicegetresponse.md), error**


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
    res, err := s.App.C1APIAppV1AppsGet(ctx, "magnam")
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APIAppV1GetAppResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `id`                                                  | *string*                                              | :heavy_check_mark:                                    | N/A                                                   |


### Response

**[*operations.C1APIAppV1AppsGetResponse](../../models/operations/c1apiappv1appsgetresponse.md), error**

