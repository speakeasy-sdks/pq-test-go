# Requestcatalog

### Available Operations

* [C1APIRequestcatalogV1RequestCatalogSearchServiceSearchEntitlements](#c1apirequestcatalogv1requestcatalogsearchservicesearchentitlements) - Invokes the c1.api.requestcatalog.v1.RequestCatalogSearchService.SearchEntitlements method.

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
    res, err := s.Requestcatalog.C1APIRequestcatalogV1RequestCatalogSearchServiceSearchEntitlements(ctx, shared.C1APIRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest{
        EntitlementAlias: pqc1test.String("debitis"),
        ExpandMask: &shared.C1APIAppV1AppEntitlementExpandMask{
            Paths: []string{
                "delectus",
            },
        },
        PageSize: pqc1test.Float64(2726.56),
        PageToken: pqc1test.String("suscipit"),
        Query: pqc1test.String("molestiae"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APIRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                            | Type                                                                                                                                                                                 | Required                                                                                                                                                                             | Description                                                                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                | :heavy_check_mark:                                                                                                                                                                   | The context to use for the request.                                                                                                                                                  |
| `request`                                                                                                                                                                            | [shared.C1APIRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsRequest](../../models/shared/c1apirequestcatalogv1requestcatalogsearchservicesearchentitlementsrequest.md) | :heavy_check_mark:                                                                                                                                                                   | The request object to use for the request.                                                                                                                                           |


### Response

**[*operations.C1APIRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsResponse](../../models/operations/c1apirequestcatalogv1requestcatalogsearchservicesearchentitlementsresponse.md), error**

