// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// C1APIPolicyV1ProvisionPolicy - The ProvisionPolicy message.
//
// This message contains a oneof named typ. Only a single field of the following list may be set at a time:
//   - connector
//   - manual
//   - delegated
type C1APIPolicyV1ProvisionPolicy struct {
	// The ConnectorProvision message.
	Connector map[string]interface{} `json:"connector,omitempty"`
	// The DelegatedProvision message.
	Delegated *C1APIPolicyV1DelegatedProvision `json:"delegated,omitempty"`
	// The ManualProvision message.
	Manual *C1APIPolicyV1ManualProvision `json:"manual,omitempty"`
}
