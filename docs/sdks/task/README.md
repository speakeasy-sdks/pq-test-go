# Task

### Available Operations

* [C1APITaskV1TaskActionsServiceApprove](#c1apitaskv1taskactionsserviceapprove) - Invokes the c1.api.task.v1.TaskActionsService.Approve method.
* [C1APITaskV1TaskActionsServiceComment](#c1apitaskv1taskactionsservicecomment) - Invokes the c1.api.task.v1.TaskActionsService.Comment method.
* [C1APITaskV1TaskActionsServiceDeny](#c1apitaskv1taskactionsservicedeny) - Invokes the c1.api.task.v1.TaskActionsService.Deny method.
* [C1APITaskV1TaskSearchServiceSearch](#c1apitaskv1tasksearchservicesearch) - Invokes the c1.api.task.v1.TaskSearchService.Search method.
* [C1APITaskV1TaskServiceCreateGrantTask](#c1apitaskv1taskservicecreategranttask) - Invokes the c1.api.task.v1.TaskService.CreateGrantTask method.
* [C1APITaskV1TaskServiceCreateRevokeTask](#c1apitaskv1taskservicecreaterevoketask) - Invokes the c1.api.task.v1.TaskService.CreateRevokeTask method.
* [C1APITaskV1TaskServiceGet](#c1apitaskv1taskserviceget) - Invokes the c1.api.task.v1.TaskService.Get method.

## C1APITaskV1TaskActionsServiceApprove

Invokes the c1.api.task.v1.TaskActionsService.Approve method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"pq-c1-test"
	"pq-c1-test/pkg/models/operations"
	"pq-c1-test/pkg/models/shared"
)

func main() {
    s := pqc1test.New()

    ctx := context.Background()
    res, err := s.Task.C1APITaskV1TaskActionsServiceApprove(ctx, "minus", &shared.C1APITaskV1TaskActionsServiceApproveRequestInput{
        Comment: pqc1test.String("placeat"),
        ExpandMask: &shared.C1APITaskV1TaskExpandMask{
            Paths: []string{
                "iusto",
                "excepturi",
                "nisi",
            },
        },
        PolicyStepID: pqc1test.String("recusandae"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APITaskV1TaskActionsServiceApproveResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                           | Type                                                                                                                                | Required                                                                                                                            | Description                                                                                                                         |
| ----------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                                               | :heavy_check_mark:                                                                                                                  | The context to use for the request.                                                                                                 |
| `taskID`                                                                                                                            | *string*                                                                                                                            | :heavy_check_mark:                                                                                                                  | N/A                                                                                                                                 |
| `c1APITaskV1TaskActionsServiceApproveRequestInput`                                                                                  | [*shared.C1APITaskV1TaskActionsServiceApproveRequestInput](../../models/shared/c1apitaskv1taskactionsserviceapproverequestinput.md) | :heavy_minus_sign:                                                                                                                  | N/A                                                                                                                                 |


### Response

**[*operations.C1APITaskV1TaskActionsServiceApproveResponse](../../models/operations/c1apitaskv1taskactionsserviceapproveresponse.md), error**


## C1APITaskV1TaskActionsServiceComment

Invokes the c1.api.task.v1.TaskActionsService.Comment method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"pq-c1-test"
	"pq-c1-test/pkg/models/operations"
	"pq-c1-test/pkg/models/shared"
)

func main() {
    s := pqc1test.New()

    ctx := context.Background()
    res, err := s.Task.C1APITaskV1TaskActionsServiceComment(ctx, "temporibus", &shared.C1APITaskV1TaskActionsServiceCommentRequestInput{
        Comment: pqc1test.String("ab"),
        ExpandMask: &shared.C1APITaskV1TaskExpandMask{
            Paths: []string{
                "veritatis",
                "deserunt",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APITaskV1TaskActionsServiceCommentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                           | Type                                                                                                                                | Required                                                                                                                            | Description                                                                                                                         |
| ----------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                                               | :heavy_check_mark:                                                                                                                  | The context to use for the request.                                                                                                 |
| `taskID`                                                                                                                            | *string*                                                                                                                            | :heavy_check_mark:                                                                                                                  | N/A                                                                                                                                 |
| `c1APITaskV1TaskActionsServiceCommentRequestInput`                                                                                  | [*shared.C1APITaskV1TaskActionsServiceCommentRequestInput](../../models/shared/c1apitaskv1taskactionsservicecommentrequestinput.md) | :heavy_minus_sign:                                                                                                                  | N/A                                                                                                                                 |


### Response

**[*operations.C1APITaskV1TaskActionsServiceCommentResponse](../../models/operations/c1apitaskv1taskactionsservicecommentresponse.md), error**


## C1APITaskV1TaskActionsServiceDeny

Invokes the c1.api.task.v1.TaskActionsService.Deny method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"pq-c1-test"
	"pq-c1-test/pkg/models/operations"
	"pq-c1-test/pkg/models/shared"
)

func main() {
    s := pqc1test.New()

    ctx := context.Background()
    res, err := s.Task.C1APITaskV1TaskActionsServiceDeny(ctx, "perferendis", &shared.C1APITaskV1TaskActionsServiceDenyRequestInput{
        Comment: pqc1test.String("ipsam"),
        ExpandMask: &shared.C1APITaskV1TaskExpandMask{
            Paths: []string{
                "sapiente",
                "quo",
                "odit",
                "at",
            },
        },
        PolicyStepID: pqc1test.String("at"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APITaskV1TaskActionsServiceDenyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                     | Type                                                                                                                          | Required                                                                                                                      | Description                                                                                                                   |
| ----------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                                                         | :heavy_check_mark:                                                                                                            | The context to use for the request.                                                                                           |
| `taskID`                                                                                                                      | *string*                                                                                                                      | :heavy_check_mark:                                                                                                            | N/A                                                                                                                           |
| `c1APITaskV1TaskActionsServiceDenyRequestInput`                                                                               | [*shared.C1APITaskV1TaskActionsServiceDenyRequestInput](../../models/shared/c1apitaskv1taskactionsservicedenyrequestinput.md) | :heavy_minus_sign:                                                                                                            | N/A                                                                                                                           |


### Response

**[*operations.C1APITaskV1TaskActionsServiceDenyResponse](../../models/operations/c1apitaskv1taskactionsservicedenyresponse.md), error**


## C1APITaskV1TaskSearchServiceSearch

Invokes the c1.api.task.v1.TaskSearchService.Search method.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"pq-c1-test"
	"pq-c1-test/pkg/models/shared"
	"pq-c1-test/pkg/types"
)

func main() {
    s := pqc1test.New()

    ctx := context.Background()
    res, err := s.Task.C1APITaskV1TaskSearchServiceSearch(ctx, shared.C1APITaskV1TaskSearchRequest{
        AccessReviewIds: []string{
            "molestiae",
            "quod",
            "quod",
            "esse",
        },
        AccountOwnerIds: []string{
            "porro",
            "dolorum",
            "dicta",
        },
        ActorID: pqc1test.String("nam"),
        AppEntitlementIds: []string{
            "occaecati",
            "fugit",
            "deleniti",
        },
        AppResourceIds: []string{
            "optio",
            "totam",
            "beatae",
            "commodi",
        },
        AppResourceTypeIds: []string{
            "modi",
            "qui",
        },
        AppUserSubjectIds: []string{
            "cum",
            "esse",
            "ipsum",
            "excepturi",
        },
        ApplicationIds: []string{
            "perferendis",
        },
        AssigneesInIds: []string{
            "natus",
            "sed",
        },
        CreatedAfter: types.MustTimeFromString("2022-07-22T16:55:44.795Z"),
        CreatedBefore: types.MustTimeFromString("2022-03-24T20:42:46.563Z"),
        CurrentStep: shared.C1APITaskV1TaskSearchRequestCurrentStepTaskSearchCurrentStepProvision.ToPointer(),
        ExcludeAppEntitlementIds: []string{
            "fuga",
            "in",
            "corporis",
            "iste",
        },
        ExcludeIds: []string{
            "saepe",
            "quidem",
        },
        ExpandMask: &shared.C1APITaskV1TaskExpandMask{
            Paths: []string{
                "ipsa",
            },
        },
        IncludeDeleted: pqc1test.Bool(false),
        MyWorkUserIds: []string{
            "est",
            "mollitia",
            "laborum",
            "dolores",
        },
        OpenerIds: []string{
            "corporis",
        },
        PageSize: pqc1test.Float64(1289.26),
        PageToken: pqc1test.String("nobis"),
        PreviouslyActedOnIds: []string{
            "omnis",
            "nemo",
        },
        Query: pqc1test.String("minima"),
        Refs: []shared.C1APITaskV1TaskRef{
            shared.C1APITaskV1TaskRef{
                ID: pqc1test.String("07aff1a3-a2fa-4946-b739-251aa52c3f5a"),
            },
            shared.C1APITaskV1TaskRef{
                ID: pqc1test.String("d019da1f-fe78-4f09-bb00-74f15471b5e6"),
            },
            shared.C1APITaskV1TaskRef{
                ID: pqc1test.String("e13b99d4-88e1-4e91-a450-ad2abd442698"),
            },
        },
        SortBy: shared.C1APITaskV1TaskSearchRequestSortByTaskSearchSortByUnspecified.ToPointer(),
        SubjectIds: []string{
            "assumenda",
        },
        TaskStates: []shared.C1APITaskV1TaskSearchRequestTaskStates{
            shared.C1APITaskV1TaskSearchRequestTaskStatesTaskStateUnspecified,
            shared.C1APITaskV1TaskSearchRequestTaskStatesTaskStateUnspecified,
        },
        TaskTypes: []shared.C1APITaskV1TaskType{
            shared.C1APITaskV1TaskType{
                Certify: &shared.C1APITaskV1TaskTypeCertify{
                    AccessReviewID: pqc1test.String("excepturi"),
                    AccessReviewSelection: pqc1test.String("tempora"),
                    AppEntitlementID: pqc1test.String("facilis"),
                    AppID: pqc1test.String("tempore"),
                    AppUserID: pqc1test.String("labore"),
                    IdentityUserID: pqc1test.String("delectus"),
                    Outcome: shared.C1APITaskV1TaskTypeCertifyOutcomeCertifyOutcomeDecertified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-03-31T00:30:19.135Z"),
                },
                Grant: &shared.C1APITaskV1TaskTypeGrant{
                    AppEntitlementID: pqc1test.String("sint"),
                    AppID: pqc1test.String("aliquid"),
                    AppUserID: pqc1test.String("provident"),
                    GrantDuration: pqc1test.String("necessitatibus"),
                    IdentityUserID: pqc1test.String("sint"),
                    Outcome: shared.C1APITaskV1TaskTypeGrantOutcomeGrantOutcomeError.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-02-09T13:58:59.361Z"),
                },
                Revoke: &shared.C1APITaskV1TaskTypeRevoke{
                    AppEntitlementID: pqc1test.String("a"),
                    AppID: pqc1test.String("dolorum"),
                    AppUserID: pqc1test.String("in"),
                    IdentityUserID: pqc1test.String("in"),
                    Outcome: shared.C1APITaskV1TaskTypeRevokeOutcomeRevokeOutcomeCancelled.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2020-11-26T01:41:04.216Z"),
                    Source: &shared.C1APITaskV1TaskRevokeSource{
                        Expired: &shared.C1APITaskV1TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-09-14T10:27:07.590Z"),
                        },
                        NonUsage: &shared.C1APITaskV1TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2020-07-23T21:23:35.691Z"),
                            LastLogin: types.MustTimeFromString("2022-08-09T06:36:34.417Z"),
                        },
                        Request: &shared.C1APITaskV1TaskRevokeSourceRequest{
                            RequestUserID: pqc1test.String("laborum"),
                        },
                        Review: &shared.C1APITaskV1TaskRevokeSourceReview{
                            AccessReviewID: pqc1test.String("accusamus"),
                            CertTicketID: pqc1test.String("non"),
                        },
                    },
                },
            },
            shared.C1APITaskV1TaskType{
                Certify: &shared.C1APITaskV1TaskTypeCertify{
                    AccessReviewID: pqc1test.String("occaecati"),
                    AccessReviewSelection: pqc1test.String("enim"),
                    AppEntitlementID: pqc1test.String("accusamus"),
                    AppID: pqc1test.String("delectus"),
                    AppUserID: pqc1test.String("quidem"),
                    IdentityUserID: pqc1test.String("provident"),
                    Outcome: shared.C1APITaskV1TaskTypeCertifyOutcomeCertifyOutcomeError.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2021-12-31T00:47:48.012Z"),
                },
                Grant: &shared.C1APITaskV1TaskTypeGrant{
                    AppEntitlementID: pqc1test.String("deleniti"),
                    AppID: pqc1test.String("sapiente"),
                    AppUserID: pqc1test.String("amet"),
                    GrantDuration: pqc1test.String("deserunt"),
                    IdentityUserID: pqc1test.String("nisi"),
                    Outcome: shared.C1APITaskV1TaskTypeGrantOutcomeGrantOutcomeDenied.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2021-10-15T07:59:26.631Z"),
                },
                Revoke: &shared.C1APITaskV1TaskTypeRevoke{
                    AppEntitlementID: pqc1test.String("molestiae"),
                    AppID: pqc1test.String("perferendis"),
                    AppUserID: pqc1test.String("nihil"),
                    IdentityUserID: pqc1test.String("magnam"),
                    Outcome: shared.C1APITaskV1TaskTypeRevokeOutcomeRevokeOutcomeError.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-06-04T18:23:50.695Z"),
                    Source: &shared.C1APITaskV1TaskRevokeSource{
                        Expired: &shared.C1APITaskV1TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-08-14T00:52:14.624Z"),
                        },
                        NonUsage: &shared.C1APITaskV1TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2021-07-03T02:32:39.849Z"),
                            LastLogin: types.MustTimeFromString("2022-02-14T08:45:31.579Z"),
                        },
                        Request: &shared.C1APITaskV1TaskRevokeSourceRequest{
                            RequestUserID: pqc1test.String("aspernatur"),
                        },
                        Review: &shared.C1APITaskV1TaskRevokeSourceReview{
                            AccessReviewID: pqc1test.String("architecto"),
                            CertTicketID: pqc1test.String("magnam"),
                        },
                    },
                },
            },
            shared.C1APITaskV1TaskType{
                Certify: &shared.C1APITaskV1TaskTypeCertify{
                    AccessReviewID: pqc1test.String("et"),
                    AccessReviewSelection: pqc1test.String("excepturi"),
                    AppEntitlementID: pqc1test.String("ullam"),
                    AppID: pqc1test.String("provident"),
                    AppUserID: pqc1test.String("quos"),
                    IdentityUserID: pqc1test.String("sint"),
                    Outcome: shared.C1APITaskV1TaskTypeCertifyOutcomeCertifyOutcomeUnspecified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2021-01-23T15:47:23.464Z"),
                },
                Grant: &shared.C1APITaskV1TaskTypeGrant{
                    AppEntitlementID: pqc1test.String("mollitia"),
                    AppID: pqc1test.String("ad"),
                    AppUserID: pqc1test.String("eum"),
                    GrantDuration: pqc1test.String("dolor"),
                    IdentityUserID: pqc1test.String("necessitatibus"),
                    Outcome: shared.C1APITaskV1TaskTypeGrantOutcomeGrantOutcomeUnspecified.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2022-11-26T12:00:10.052Z"),
                },
                Revoke: &shared.C1APITaskV1TaskTypeRevoke{
                    AppEntitlementID: pqc1test.String("iure"),
                    AppID: pqc1test.String("doloribus"),
                    AppUserID: pqc1test.String("debitis"),
                    IdentityUserID: pqc1test.String("eius"),
                    Outcome: shared.C1APITaskV1TaskTypeRevokeOutcomeRevokeOutcomeCancelled.ToPointer(),
                    OutcomeTime: types.MustTimeFromString("2021-08-05T03:52:18.835Z"),
                    Source: &shared.C1APITaskV1TaskRevokeSource{
                        Expired: &shared.C1APITaskV1TaskRevokeSourceExpired{
                            ExpiredAt: types.MustTimeFromString("2022-11-25T10:00:44.006Z"),
                        },
                        NonUsage: &shared.C1APITaskV1TaskRevokeSourceNonUsage{
                            ExpiresAt: types.MustTimeFromString("2022-01-30T09:19:56.236Z"),
                            LastLogin: types.MustTimeFromString("2022-04-15T07:14:46.128Z"),
                        },
                        Request: &shared.C1APITaskV1TaskRevokeSourceRequest{
                            RequestUserID: pqc1test.String("nihil"),
                        },
                        Review: &shared.C1APITaskV1TaskRevokeSourceReview{
                            AccessReviewID: pqc1test.String("repellat"),
                            CertTicketID: pqc1test.String("quibusdam"),
                        },
                    },
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APITaskV1TaskSearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [shared.C1APITaskV1TaskSearchRequest](../../models/shared/c1apitaskv1tasksearchrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.C1APITaskV1TaskSearchServiceSearchResponse](../../models/operations/c1apitaskv1tasksearchservicesearchresponse.md), error**


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
    res, err := s.Task.C1APITaskV1TaskServiceCreateGrantTask(ctx, shared.C1APITaskV1TaskServiceCreateGrantRequest{
        AppEntitlementID: pqc1test.String("sed"),
        AppID: pqc1test.String("saepe"),
        AppUserID: pqc1test.String("pariatur"),
        Description: pqc1test.String("accusantium"),
        ExpandMask: &shared.C1APITaskV1TaskExpandMask{
            Paths: []string{
                "praesentium",
            },
        },
        GrantDuration: pqc1test.String("natus"),
        IdentityUserID: pqc1test.String("magni"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APITaskV1TaskServiceCreateGrantResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [shared.C1APITaskV1TaskServiceCreateGrantRequest](../../models/shared/c1apitaskv1taskservicecreategrantrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.C1APITaskV1TaskServiceCreateGrantTaskResponse](../../models/operations/c1apitaskv1taskservicecreategranttaskresponse.md), error**


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
    res, err := s.Task.C1APITaskV1TaskServiceCreateRevokeTask(ctx, shared.C1APITaskV1TaskServiceCreateRevokeRequest{
        AppEntitlementID: pqc1test.String("sunt"),
        AppID: pqc1test.String("quo"),
        AppUserID: pqc1test.String("illum"),
        Description: pqc1test.String("pariatur"),
        ExpandMask: &shared.C1APITaskV1TaskExpandMask{
            Paths: []string{
                "ea",
                "excepturi",
                "odit",
                "ea",
            },
        },
        IdentityUserID: pqc1test.String("accusantium"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APITaskV1TaskServiceCreateRevokeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [shared.C1APITaskV1TaskServiceCreateRevokeRequest](../../models/shared/c1apitaskv1taskservicecreaterevokerequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.C1APITaskV1TaskServiceCreateRevokeTaskResponse](../../models/operations/c1apitaskv1taskservicecreaterevoketaskresponse.md), error**


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
    res, err := s.Task.C1APITaskV1TaskServiceGet(ctx, "ab")
    if err != nil {
        log.Fatal(err)
    }

    if res.C1APITaskV1TaskServiceGetResponse != nil {
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

**[*operations.C1APITaskV1TaskServiceGetResponse](../../models/operations/c1apitaskv1taskservicegetresponse.md), error**

