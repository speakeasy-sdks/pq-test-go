// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// C1APITaskV1TaskServiceCreateRevokeResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type C1APITaskV1TaskServiceCreateRevokeResponseExpanded struct {
	// The type of the serialized message.
	AtType *string `json:"@type,omitempty"`
}

// C1APITaskV1TaskServiceCreateRevokeResponse - The TaskServiceCreateRevokeResponse message.
type C1APITaskV1TaskServiceCreateRevokeResponse struct {
	// The expanded field.
	Expanded []C1APITaskV1TaskServiceCreateRevokeResponseExpanded `json:"expanded,omitempty"`
	// The TaskView message.
	TaskView *C1APITaskV1TaskView `json:"taskView,omitempty"`
}
