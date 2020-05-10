package model

import (
	"sigmamono/internal/core"
	"sigmamono/internal/enum/action"
	"sigmamono/internal/term"
)

// Auth model
type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Validate check the type of fields for auth
func (p *Auth) Validate(act action.Action) error {
	fieldError := core.NewFieldError(term.Error_in_login_form)

	switch act {
	case action.Login:
		if p.Username == "" {
			fieldError.Add(term.V_is_required, "Username", "username")
		}

		if p.Password == "" {
			fieldError.Add(term.V_is_required, "Password", "password")
		}
	}

	if fieldError.HasError() {
		return fieldError
	}
	return nil
}
