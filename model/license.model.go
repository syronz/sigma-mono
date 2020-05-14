package model

import (
	"sigmamono/internal/core"
	"sigmamono/internal/enum/action"
	"sigmamono/internal/term"
	"sigmamono/internal/types"
	"time"
)

// License model
type License struct {
	Name   string `json:"name,omitempty"`
	Key    string `json:"key,omitempty"`
	Serial string `json:"serial,omitempty"`
	Count  int    `json:"count,omitempty"`
}

// Activation keeps usage of licenses to prevent duplication
type Activation struct {
	License   string      `gorm:"primary_key" json:"license"`
	UsedAt    time.Time   `json:"used_at"`
	CompanyID types.RowID `json:"company_id"`
}

// Validate check the type of
func (p *License) Validate(act action.Action) error {
	fieldError := core.NewFieldError(term.Error_in_versions_form)

	switch act {
	case action.Create:
		if p.Count > 9999 {
			params := []interface{}{1, 9999}
			fieldError.Add(term.Range_is_v_to_v, params, "count")
		}
	}

	if fieldError.HasError() {
		return fieldError
	}
	return nil
}
