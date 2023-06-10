# C1APIAppV1AppEntitlement

The AppEntitlement message.

This message contains a oneof named max_grant_duration. Only a single field of the following list may be set at a time:
  - durationUnset
  - durationGrant



## Fields

| Field                                                                                                                                                                          | Type                                                                                                                                                                           | Required                                                                                                                                                                       | Description                                                                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `Alias`                                                                                                                                                                        | **string*                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                             | The alias field.                                                                                                                                                               |
| `AppID`                                                                                                                                                                        | **string*                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                             | The appId field.                                                                                                                                                               |
| `AppResourceID`                                                                                                                                                                | **string*                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                             | The appResourceId field.                                                                                                                                                       |
| `AppResourceTypeID`                                                                                                                                                            | **string*                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                             | The appResourceTypeId field.                                                                                                                                                   |
| `CertifyPolicyID`                                                                                                                                                              | **string*                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                             | The certifyPolicyId field.                                                                                                                                                     |
| `ComplianceFrameworkValueIds`                                                                                                                                                  | []*string*                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                             | The complianceFrameworkValueIds field.                                                                                                                                         |
| `CreatedAt`                                                                                                                                                                    | [*time.Time](https://pkg.go.dev/time#Time)                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                             | N/A                                                                                                                                                                            |
| `DeletedAt`                                                                                                                                                                    | [*time.Time](https://pkg.go.dev/time#Time)                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                             | N/A                                                                                                                                                                            |
| `Description`                                                                                                                                                                  | **string*                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                             | The description field.                                                                                                                                                         |
| `DisplayName`                                                                                                                                                                  | **string*                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                             | The displayName field.                                                                                                                                                         |
| `DurationGrant`                                                                                                                                                                | **string*                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                             | N/A                                                                                                                                                                            |
| `DurationUnset`                                                                                                                                                                | [*C1APIAppV1AppEntitlementDurationUnset](../../models/shared/c1apiappv1appentitlementdurationunset.md)                                                                         | :heavy_minus_sign:                                                                                                                                                             | N/A                                                                                                                                                                            |
| `GrantCount`                                                                                                                                                                   | **string*                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                             | The grantCount field.                                                                                                                                                          |
| `GrantPolicyID`                                                                                                                                                                | **string*                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                             | The grantPolicyId field.                                                                                                                                                       |
| `ID`                                                                                                                                                                           | **string*                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                             | The id field.                                                                                                                                                                  |
| `ProvisionerPolicy`                                                                                                                                                            | [*C1APIPolicyV1ProvisionPolicy](../../models/shared/c1apipolicyv1provisionpolicy.md)                                                                                           | :heavy_minus_sign:                                                                                                                                                             | The ProvisionPolicy message.<br/><br/>This message contains a oneof named typ. Only a single field of the following list may be set at a time:<br/>  - connector<br/>  - manual<br/>  - delegated<br/> |
| `RevokePolicyID`                                                                                                                                                               | **string*                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                             | The revokePolicyId field.                                                                                                                                                      |
| `RiskLevelValueID`                                                                                                                                                             | **string*                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                             | The riskLevelValueId field.                                                                                                                                                    |
| `Slug`                                                                                                                                                                         | **string*                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                             | The slug field.                                                                                                                                                                |
| `SystemBuiltin`                                                                                                                                                                | **bool*                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                             | The systemBuiltin field.                                                                                                                                                       |
| `UpdatedAt`                                                                                                                                                                    | [*time.Time](https://pkg.go.dev/time#Time)                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                             | N/A                                                                                                                                                                            |