package model

import (
	"sigmamono/internal/consts"
	"sigmamono/internal/core"
	"sigmamono/internal/enum/action"
	"sigmamono/internal/enum/nodestatus"
	"sigmamono/internal/enum/nodetype"
	"sigmamono/internal/term"
	"sigmamono/internal/types"
	"sigmamono/utils/helper"
	"strings"
)

// Node model
type Node struct {
	types.GormCol
	CompanyID   types.RowID            `gorm:"unique_index:idx_companyID_nodeCode,idx_companyID_machineID,idx_companyID_nodeName" json:"company_id,omitempty"`
	Code        uint64                 `gorm:"not null;unique_index:idx_companyID_nodeCode" json:"code,omitempty"`
	Type        string                 `gorm:"not null" json:"type,omitempty"`
	Name        string                 `gorm:"not null;unique_index:idx_companyID_nodeName" json:"name,omitempty"`
	MachineID   string                 `gorm:"not null;unique_index:idx_companyID_machineID" json:"machine_id,omitempty"`
	Status      string                 `gorm:"default:'inactive'" json:"status,omitempty"`
	Phone       string                 `gorm:"not null" json:"phone,omitempty"`
	Extra       map[string]interface{} `sql:"-" json:"extra_node,omitempty"`
	Error       error                  `sql:"-" json:"error,omitempty"`
	CompanyName string                 `sql:"-" json:"company_name,omitempty"`
}

// Pattern returns the search pattern to be used inside the gorm's where
func (p Node) Pattern() string {
	return `(nodes.name LIKE '%[1]v%%' OR
		nodes.phone LIKE '%[1]v%%')`
}

// Columns return list of total columns according to request, useful for inner joins
func (p Node) Columns(variate string) (string, error) {
	full := []string{"nodes.id", "nodes.company_id", "nodes.name", "nodes.phone",
		"nodes.created_at", "nodes.updated_at", "companies.name as company_name"}

	return checkColumns(full, variate)
}

// Validate check the type of
func (p *Node) Validate(act action.Action) error {
	fieldError := core.NewFieldError(term.Error_in_nodes_form)

	switch act {
	case action.Save:
		if p.Name == "" {
			fieldError.Add(term.V_is_required, term.Name, term.Name)
		}

		if p.Phone == "" {
			fieldError.Add(term.V_is_required, term.Phone, term.Phone)
		}

		if p.CompanyID == 0 {
			fieldError.Add(term.V_is_required, term.CompanyID, term.CompanyID)
		}

		if p.MachineID == "" {
			fieldError.Add(term.V_is_required, term.MachineID, term.MachineID)
		}

		if p.Code < consts.MinNodeCode || p.Code > consts.MaxNodeCode {
			fieldError.Add(term.V_is_not_valid, term.Code, term.Code)
		}

		if ok, _ := helper.Includes(nodetype.Types, p.Type); !ok {
			fieldError.Add(term.Accepted_values_are_v,
				strings.Join(nodetype.Types, ", "), term.Type)
		}

		// if p.Status != "" {
		if ok, _ := helper.Includes(nodestatus.Statuses, p.Status); !ok {
			fieldError.Add(term.Accepted_values_for_status_are_v,
				strings.Join(nodestatus.Statuses, ", "), "status")
		}
		// }

	}

	if fieldError.HasError() {
		return fieldError
	}
	return nil
}
