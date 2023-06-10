# User

### Available Operations

* [C1APIUserV1UserServiceGet](#c1apiuserv1userserviceget) - Invokes the c1.api.user.v1.UserService.Get method.

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
    res, err := s.User.C1APIUserV1UserServiceGet(ctx, "maiores")
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APIUserV1UserServiceGetResponse != nil {
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

**[*operations.C1APIUserV1UserServiceGetResponse](../../models/operations/c1apiuserv1userservicegetresponse.md), error**

