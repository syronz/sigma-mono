package model

import (
	"radiusbilling/internal/core"
	"radiusbilling/internal/enum/action"
	"radiusbilling/internal/enum/lang"
	"radiusbilling/internal/term"
	"radiusbilling/internal/types"
	"radiusbilling/utils/helper"
	"regexp"
	"strings"
)

// User model
type User struct {
	ID       types.RowID `gorm:"not null;unique" json:"id"`
	RoleID   types.RowID `gorm:"index:role_id_idx" json:"role_id"`
	Username string      `gorm:"not null;unique" json:"username,omitempty"`
	Password string      `gorm:"not null" json:"password,omitempty"`
	Language string      `gorm:"type:varchar(2);default:'en'" json:"language,omitempty"`
	Email    string      `json:"email,omitempty"`
	Extra    interface{} `sql:"-" json:"user_extra,omitempty"`
	Error    error       `sql:"-" json:"user_error,omitempty"`
}

// Pattern returns the search pattern to be used inside the gorm's where
func (p User) Pattern() string {
	return `(
		users.username LIKE '%[1]v%%' OR
		users.id = '%[1]v' OR
		users.email LIKE '%[1]v%%' OR
		accounts.code LIKE '%[1]v' OR
		roles.name LIKE '%[1]v' OR
		accounts.status LIKE '%[1]v'
	)`
}

// Columns return list of total columns according to request, useful for inner joins
func (p User) Columns(variate string) (string, error) {
	full := []string{"users.id", "users.role_id", "users.username", "users.language",
		"users.email", "accounts.name", "accounts.status", "accounts.code", "accounts.type",
		"accounts.readonly", "accounts.score", "accounts.direction", "accounts.created_at",
		"accounts.updated_at", "accounts.deleted_at", "roles.name as role"}

	fieldError := core.NewFieldError(term.Error_in_nodes_url)

	if variate == "*" {
		return strings.Join(full, ","), nil
	}

	variates := strings.Split(variate, ",")
	for _, v := range variates {
		if ok, _ := helper.Includes(full, v); !ok {
			fieldError.Add(term.V_is_not_valid, v, strings.Join(full, ", "))
		}
	}
	if fieldError.HasError() {
		return "", fieldError
	}

	return variate, nil
}

// Validate check the type of
func (p *User) Validate(act action.Action) error {
	fieldError := core.NewFieldError(term.Error_in_users_form)

	switch act {
	case action.Create:

		if len(p.Password) < 8 {
			params := []interface{}{"password", 7}
			fieldError.Add(term.V_should_be_more_than_V_character, params, "password")
		}

		fallthrough

	case action.Update:

		if p.Username == "" {
			fieldError.Add(term.V_is_required, "Username", "username")
		}

		if p.RoleID == 0 {
			fieldError.Add(term.V_is_required, "Role", "role_id")
		}

		if ok, _ := helper.Includes(lang.Languages, p.Language); !ok {
			fieldError.Add(term.Accepted_values_are_v,
				strings.Join(lang.Languages, ", "), "language")
		}

		if p.Email != "" {
			re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
			if !re.MatchString(p.Email) {
				fieldError.Add(term.Email_is_not_valid, nil, "email")
			}
		}

	}

	if fieldError.HasError() {
		return fieldError
	}
	return nil

}
