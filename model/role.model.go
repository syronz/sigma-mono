package model

import (
	"sigmamono/internal/core"
	"sigmamono/internal/enum/action"
	"sigmamono/internal/term"
	"sigmamono/internal/types"
	"sigmamono/utils/helper"
	"strings"
)

// Role model
type Role struct {
	types.FixedCol
	Name        string `gorm:"not null;unique_index:idx_companyID_name" json:"name,omitempty"`
	Resources   string `gorm:"type:text" json:"resources,omitempty"`
	Description string `json:"description,omitempty"`
	Error       error  `sql:"-" json:"user_error,omitempty"`
}

// Pattern returns the search pattern to be used inside the gorm's where
func (p Role) Pattern() string {
	return `(roles.name LIKE '%[1]v%%' OR
		roles.id = '%[1]v' OR
		roles.resources LIKE '%%%[1]v%%')`
}

// Columns return list of total columns according to request, useful for inner joins
func (p Role) Columns(variate string) (string, error) {
	full := []string{"roles.id", "roles.company_id", "roles.name", "roles.description", "roles.resources",
		"roles.created_at", "roles.updated_at"}

	fieldError := core.NewFieldError(term.Error_in_url)

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

// Validate check the type of fields
func (p *Role) Validate(act action.Action) error {
	fieldError := core.NewFieldError(term.Error_in_role_form)

	switch act {
	case action.Create:
		if p.Name == "" {
			fieldError.Add(term.V_is_required, "Name", "name")
		}

		if len(p.Name) < 5 {
			fieldError.Add(term.Name_at_least_be_5_character, nil, "name")
		}

		if p.Resources == "" {
			fieldError.Add(term.V_is_required, "Resources", "resources")
		}

	}

	if fieldError.HasError() {
		return fieldError
	}
	return nil

}
