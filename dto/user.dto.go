package dto

import "radiusbilling/model"

// UserDto is used inside the list
type UserDto struct {
	model.User
	// model.Account
	// Role string `json:"role,omitempty"`
}
