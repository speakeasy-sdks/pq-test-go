// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// UserDirectoryStatus - The directoryStatus field.
type UserDirectoryStatus string

const (
	UserDirectoryStatusUnknown  UserDirectoryStatus = "UNKNOWN"
	UserDirectoryStatusEnabled  UserDirectoryStatus = "ENABLED"
	UserDirectoryStatusDisabled UserDirectoryStatus = "DISABLED"
	UserDirectoryStatusDeleted  UserDirectoryStatus = "DELETED"
)

func (e UserDirectoryStatus) ToPointer() *UserDirectoryStatus {
	return &e
}

func (e *UserDirectoryStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "UNKNOWN":
		fallthrough
	case "ENABLED":
		fallthrough
	case "DISABLED":
		fallthrough
	case "DELETED":
		*e = UserDirectoryStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UserDirectoryStatus: %v", v)
	}
}

// UserStatus - The status field.
type UserStatus string

const (
	UserStatusUnknown  UserStatus = "UNKNOWN"
	UserStatusEnabled  UserStatus = "ENABLED"
	UserStatusDisabled UserStatus = "DISABLED"
	UserStatusDeleted  UserStatus = "DELETED"
)

func (e UserStatus) ToPointer() *UserStatus {
	return &e
}

func (e *UserStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "UNKNOWN":
		fallthrough
	case "ENABLED":
		fallthrough
	case "DISABLED":
		fallthrough
	case "DELETED":
		*e = UserStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UserStatus: %v", v)
	}
}

// User - The User message.
type User struct {
	CreatedAt       *time.Time `json:"createdAt,omitempty"`
	DelegatedUserID *string    `json:"delegatedUserId,omitempty"`
	DeletedAt       *time.Time `json:"deletedAt,omitempty"`
	Department      *string    `json:"department,omitempty"`
	// The departmentSources field.
	DepartmentSources []UserAttributeMappingSource `json:"departmentSources,omitempty"`
	// The directoryIds field.
	DirectoryIds []string `json:"directoryIds,omitempty"`
	// The directoryStatus field.
	DirectoryStatus *UserDirectoryStatus `json:"directoryStatus,omitempty"`
	// The directoryStatusSources field.
	DirectoryStatusSources []UserAttributeMappingSource `json:"directoryStatusSources,omitempty"`
	DisplayName            *string                      `json:"displayName,omitempty"`
	Email                  *string                      `json:"email,omitempty"`
	EmploymentStatus       *string                      `json:"employmentStatus,omitempty"`
	// The employmentStatusSources field.
	EmploymentStatusSources []UserAttributeMappingSource `json:"employmentStatusSources,omitempty"`
	EmploymentType          *string                      `json:"employmentType,omitempty"`
	// The employmentTypeSources field.
	EmploymentTypeSources []UserAttributeMappingSource `json:"employmentTypeSources,omitempty"`
	ID                    *string                      `json:"id,omitempty"`
	JobTitle              *string                      `json:"jobTitle,omitempty"`
	// The jobTitleSources field.
	JobTitleSources []UserAttributeMappingSource `json:"jobTitleSources,omitempty"`
	// The managerIds field.
	ManagerIds []string `json:"managerIds,omitempty"`
	// The managerSources field.
	ManagerSources []UserAttributeMappingSource `json:"managerSources,omitempty"`
	// The roleIds field.
	RoleIds []string `json:"roleIds,omitempty"`
	// The status field.
	Status    *UserStatus `json:"status,omitempty"`
	UpdatedAt *time.Time  `json:"updatedAt,omitempty"`
}
