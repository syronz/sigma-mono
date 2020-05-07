package model

import (
	"sigmamono/internal/core"
	"sigmamono/internal/enum/action"
	"sigmamono/internal/term"
	"sigmamono/internal/types"
)

// Bond model
type Bond struct {
	types.GormCol
	CompanyID   types.RowID            `gorm:"not null;unique" json:"company_id,omitempty"`
	CompanyName string                 `gorm:"not null;unique" json:"company_name,omitempty"`
	NodeCode    uint64                 `json:"node_code,omitempty"`
	NodeName    string                 `json:"node_name,omitempty"`
	Key         string                 `gorm:"type:text" json:"key,omitempty"`
	MachineID   string                 `json:"machine_id,omitempty"`
	Detail      string                 `json:"detail,omitempty"`
	Error       error                  `sql:"-" json:"user_error,omitempty"`
	Extra       map[string]interface{} `sql:"-" json:"extra_bond,omitempty"`
}

// Pattern returns the search pattern to be used inside the gorm's where
func (p Bond) Pattern() string {
	return `(companies.name LIKE '%[1]v%%' OR
		companies.id = '%[1]v' OR
		companies.plan LIKE '%[1]v' OR
		companies.phone LIKE '%[1]v%%')`
}

// Columns return list of total columns according to request, useful for inner joins
func (p Bond) Columns(variate string) (string, error) {
	full := []string{"companies.id", "companies.name", "companies.legal_name",
		"companies.key", "companies.server_address", "companies.expiration", "companies.plan",
		"companies.detail", "companies.phone", "companies.email", "companies.website",
		"companies.type", "companies.code"}

	return checkColumns(full, variate)
}

// Validate check the type of
func (p *Bond) Validate(act action.Action) error {
	fieldError := core.NewFieldError(term.Error_in_companys_form)

	switch act {
	case action.Save:
		if p.CompanyName == "" {
			fieldError.Add(term.V_is_required, "Company Name", "company_name")
		}

		if p.CompanyID == 0 {
			fieldError.Add(term.V_is_required, "Company ID", "company_id")
		}

	}

	if fieldError.HasError() {
		return fieldError
	}
	return nil
}
