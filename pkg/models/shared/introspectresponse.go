// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// IntrospectResponse - The IntrospectResponse message.
type IntrospectResponse struct {
	AccessTokenID *string `json:"accessTokenId,omitempty"`
	// The features field.
	Features []string `json:"features,omitempty"`
	// The permissions field.
	Permissions []string `json:"permissions,omitempty"`
	PrincipleID *string  `json:"principleId,omitempty"`
	// The roles field.
	Roles    []string `json:"roles,omitempty"`
	TenantID *string  `json:"tenantId,omitempty"`
	UserID   *string  `json:"userId,omitempty"`
}
