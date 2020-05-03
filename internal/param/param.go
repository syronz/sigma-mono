package param

import (
	"errors"
	"sigmamono/internal/consts"
	"sigmamono/internal/term"
	"sigmamono/internal/types"
)

// Param for describing request's parameter
type Param struct {
	Pagination
	Search       string
	PreCondition string
	UserID       types.RowID
	CompanyID    types.RowID
	NodeCode     uint64
	Language     string
}

// Pagination is a struct, contains the fields which affected the front-end pagination
type Pagination struct {
	Select string
	Order  string
	Limit  uint64
	Offset uint64
}

// PrefixID returns the prefix for finding similar ides in the scope
func (p *Param) PrefixID() (rowID types.RowID, err error) {
	if p.CompanyID == 0 {
		err = errors.New(term.CompanyID_not_exist_in_context)
		return
	}

	if p.NodeCode == 0 {
		err = errors.New(term.NodeCode_not_exist_in_context)
		return
	}

	rowID = p.CompanyID*consts.CompanyRange + types.RowID(p.NodeCode)

	return
}
