package model

import (
	"sigmamono/internal/core"
	"sigmamono/internal/enum/action"
	"sigmamono/internal/term"
	"sigmamono/internal/types"
)

// Setting model
type Setting struct {
	types.FixedCol
	CompanyID   types.RowID `gorm:"unique_index:idx_companyID_property" json:"company_id,omitempty"`
	Property    string      `gorm:"not null;unique_index:idx_companyID_property" json:"property,omitempty"`
	Value       string      `gorm:"type:text" json:"value,omitempty"`
	Type        string      `json:"type,omitempty"`
	Description string      `json:"description,omitempty"`
	Error       error       `sql:"-" json:"user_error,omitempty"`
}

// Pattern returns the search pattern to be used inside the gorm's where
func (p Setting) Pattern() string {
	return `(
		settings.property LIKE '%[1]v%%' OR
		settings.ID = '%[1]v' OR
		settings.company_id LIKE '%[1]v' OR
		settings.value LIKE '%[1]v' OR
		settings.type LIKE '%[1]v' OR
		settings.description LIKE '%[1]v'
	)`
}

// Columns return list of total columns according to request, useful for inner joins
func (p Setting) Columns(variate string) (string, error) {
	full := []string{"settings.id", "settings.company_id", "settings.property", "settings.value",
		"settings.type", "settings.description"}

	return checkColumns(full, variate)
}

// Validate check the type of fields
func (p *Setting) Validate(act action.Action) error {
	fieldError := core.NewFieldError(term.Error_in_role_form)

	switch act {
	case action.Save:
		if p.Property == "" {
			fieldError.Add(term.V_is_required, "Property", "property")
		}
	}

	if fieldError.HasError() {
		return fieldError
	}
	return nil

}
