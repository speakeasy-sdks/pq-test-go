// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// C1APIRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type C1APIRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsResponseExpanded struct {
	// The type of the serialized message.
	AtType *string `json:"@type,omitempty"`
}

// C1APIRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsResponse - The RequestCatalogSearchServiceSearchEntitlementsResponse message.
type C1APIRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsResponse struct {
	// The expanded field.
	Expanded []C1APIRequestcatalogV1RequestCatalogSearchServiceSearchEntitlementsResponseExpanded `json:"expanded,omitempty"`
	// The list field.
	List []C1APIAppV1AppEntitlementWithUserBindings `json:"list,omitempty"`
	// The nextPageToken field.
	NextPageToken *string `json:"nextPageToken,omitempty"`
}
