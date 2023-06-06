// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UserServiceGetResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type UserServiceGetResponseExpanded struct {
	// The type of the serialized message.
	AtType *string `json:"@type,omitempty"`
}

// UserServiceGetResponse - The UserServiceGetResponse message.
type UserServiceGetResponse struct {
	// The expanded field.
	Expanded []UserServiceGetResponseExpanded `json:"expanded,omitempty"`
	// The UserView message.
	UserView *UserView `json:"userView,omitempty"`
}
