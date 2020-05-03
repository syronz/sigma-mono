package model

import (
	"fmt"
	"sigmamono/internal/consts"
	"sigmamono/internal/core"
	"sigmamono/internal/enum/accountdirection"
	"sigmamono/internal/enum/accountstatus"
	"sigmamono/internal/enum/accounttype"
	"sigmamono/internal/enum/action"
	"sigmamono/internal/term"
	"sigmamono/internal/types"
	"sigmamono/utils/helper"
	"strings"
)

// Account model
type Account struct {
	types.FixedCol
	CompanyID types.RowID  `gorm:"unique_index:idx_companyID_name,idx_companyID_code" json:"company_id,omitempty"`
	NodeCode  uint64       `json:"node_code,omitempty"`
	ParentID  *types.RowID `gorm:"index:account_account_idx" json:"parent_id,omitempty"`
	Name      string       `gorm:"not null;unique_index:idx_companyID_name" json:"name,omitempty"`
	Status    string       `gorm:"not null;default:'active'" json:"status,omitempty"`
	Code      uint         `gorm:"unique_index:idx_companyID_code" json:"code,omitempty"`
	Type      string       `json:"type,omitempty"`
	Readonly  bool         `json:"readonly,omitempty"`
	Score     int          `json:"score,omitempty"`
	Direction string       `json:"direction,omitempty"`
	Extra     interface{}  `sql:"-" json:"account_extra,omitempty"`
	Error     error        `sql:"-" json:"account_error,omitempty"`
}

// Pattern returns the search pattern to be used inside the gorm's where
func (p Account) Pattern() string {
	return `(accounts.name LIKE '%[1]v%%' OR
		accounts.status LIKE '%[1]v' OR
		accounts.company_id LIKE '%[1]v' OR
		accounts.node_code LIKE '%[1]v' OR
		accounts.type LIKE '%[1]v' OR
		accounts.direction LIKE '%%%[1]v%%' OR
		accounts.code LIKE '%[1]v%%')`
}

// Columns return list of total columns according to request, useful for inner joins
func (p Account) Columns(variate string) (string, error) {
	full := []string{"accounts.id", "accounts.company_id", "accounts.name",
		"accounts.code", "accounts.type", "accounts.created_at", "accounts.updated_at",
	}

	fieldError := core.NewFieldError(term.Error_in_accounts_url)

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
func (p *Account) Validate(act action.Action) error {
	fieldError := core.NewFieldError(term.Error_in_accounts_form)

	switch act {
	case action.Save:
		if p.Name == "" {
			fieldError.Add(term.V_is_required, term.Name, "name")
		}

		if ok, _ := helper.Includes(accounttype.Types, p.Type); !ok {
			fieldError.Add(term.Accepted_values_for_type_are_v,
				strings.Join(accounttype.Types, ", "), "type")
		}

		if ok, _ := helper.Includes(accountstatus.Statuses, p.Status); !ok {
			fieldError.Add(term.Accepted_values_for_status_are_v,
				strings.Join(accountstatus.Statuses, ", "), "status")
		}

		if ok, _ := helper.Includes(accountdirection.Directions, p.Direction); !ok {
			fieldError.Add(term.Accepted_values_are_v,
				strings.Join(accountdirection.Directions, ", "), "direction")
		}

		if p.NodeCode < consts.MinNodeCode || p.NodeCode > consts.MaxNodeCode {
			fieldError.Add(term.Accepted_values_are_v,
				fmt.Sprintf("%v to %v", consts.MinNodeCode, consts.MaxNodeCode), "node_code")
		}

		if p.Code < 1 {
			fieldError.Add(term.V_is_required, term.Code, term.Code)
		}

	}

	if fieldError.HasError() {
		return fieldError
	}
	return nil
}
