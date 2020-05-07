package model

import (
	"sigmamono/internal/consts"
	"sigmamono/internal/core"
	"sigmamono/internal/enum/action"
	"sigmamono/internal/enum/companytype"
	"sigmamono/internal/term"
	"sigmamono/internal/types"
	"sigmamono/utils/helper"
	"strings"
	"time"
)

// Company model
type Company struct {
	types.GormCol
	Name       string      `gorm:"not null" json:"name,omitempty"`
	LegalName  string      `gorm:"not null;unique" json:"legal_name,omitempty"`
	Key        string      `gorm:"type:text" json:"key,omitempty"`
	Expiration time.Time   `json:"expiration,omitempty"`
	License    string      `gorm:"unique" json:"license,omitempty"`
	Detail     string      `json:"detail,omitempty"`
	Phone      string      `gorm:"not null" json:"phone,omitempty"`
	Email      string      `gorm:"not null" json:"email,omitempty"`
	Website    string      `gorm:"not null" json:"website,omitempty"`
	Type       string      `gorm:"not null" json:"type,omitempty"`
	Code       string      `gorm:"not null" json:"code,omitempty"`
	Extra      interface{} `sql:"-" json:"extra_company,omitempty"`
	Error      error       `sql:"-" json:"error,omitempty"`
}

// Pattern returns the search pattern to be used inside the gorm's where
func (p Company) Pattern() string {
	return `(companies.name LIKE '%[1]v%%' OR
		companies.id = '%[1]v' OR
		companies.code LIKE '%[1]v' OR
		companies.legal_name LIKE '%[1]v' OR
		companies.detail LIKE '%%%[1]v%%' OR
		companies.website LIKE '%%%[1]v%%' OR
		companies.type LIKE '%%%[1]v%%' OR
		companies.phone LIKE '%[1]v%%')`
}

// Columns return list of total columns according to request, useful for inner joins
func (p Company) Columns(variate string) (string, error) {
	full := []string{"companies.id", "companies.name", "companies.legal_name",
		"companies.key", "companies.expiration", "companies.detail", "companies.phone",
		"companies.email", "companies.website", "companies.type", "companies.code"}

	return checkColumns(full, variate)
}

// Validate check the type of
func (p *Company) Validate(act action.Action) error {
	fieldError := core.NewFieldError(term.Error_in_companys_form)

	switch act {
	case action.Save:
		if p.Name == "" {
			fieldError.Add(term.V_is_required, "Name", "name")
		}

		if p.LegalName == "" {
			fieldError.Add(term.V_is_required, "Legal Name", "legal_name")
		}

		if ok, _ := helper.Includes(companytype.Types, p.Type); !ok {
			fieldError.Add(term.Accepted_values_for_type_are_v,
				strings.Join(companytype.Types, ", "), "type")
		}

		if p.ID < consts.MinCompanyID || p.ID > consts.MaxCompanyID {
			fieldError.Add(term.V_is_not_valid, "ID", "id")
		}

	}

	if fieldError.HasError() {
		return fieldError
	}
	return nil
}
