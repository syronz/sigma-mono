package helper

import (
	"errors"
	"sigmamono/internal/consts"
	"sigmamono/internal/term"
	"sigmamono/internal/types"

	"github.com/gin-gonic/gin"
)

// HeadID combine companyID and nodeID for finding the last id
func HeadID(companyID types.RowID, nodeCode uint64) types.RowID {
	return companyID*consts.CompanyRange + types.RowID(nodeCode)
}

// MinID return minimum possible id for specific companyID and nodeID
func MinID(companyID types.RowID, nodeCode uint64) types.RowID {
	return (companyID*consts.CompanyRange + types.RowID(nodeCode)) * consts.IDRange
}

// GetHeadID extract companyID and nodeCode from gin's context
func GetHeadID(c *gin.Context) (rowID types.RowID, err error) {
	companyIDTmp, ok := c.Get("COMPANY_ID")
	if !ok {
		err = errors.New(term.CompanyID_not_exist_in_context)
		return
	}

	nodeCodeTmp, ok := c.Get("NODE_CODE")
	if !ok {
		err = errors.New(term.NodeCode_not_exist_in_context)
		return
	}

	companyID := companyIDTmp.(types.RowID)
	nodeCode := nodeCodeTmp.(types.RowID)

	rowID = companyID*consts.CompanyRange + nodeCode

	return
}

// PrefixMinID return minimum possible id for specific companyID and nodeID
func PrefixMinID(prefix types.RowID) types.RowID {
	return (prefix) * consts.IDRange
}
