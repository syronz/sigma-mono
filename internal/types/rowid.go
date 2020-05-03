package types

import (
	"database/sql/driver"
	"fmt"
	"sigmamono/internal/consts"
	"strconv"
)

// RowID is used for defining the type of id
type RowID uint64

// ToUint64 convert the value to uint, useful for gorm
func (r *RowID) ToUint64() uint64 {
	return uint64(*r)
}

// ToString convert the value to string
func (r *RowID) ToString() string {
	return fmt.Sprint(*r)
}

// RowIDPointer return a pointer to the RowID
func RowIDPointer(num uint64) *RowID {
	rowID := RowID(num)
	return &rowID
}

// StrToRowID convert string number to RowID
func StrToRowID(strNum string) (rowID RowID, err error) {
	tmpID, err := strconv.ParseUint(strNum, 10, 64)
	rowID = RowID(tmpID)
	return
}

// Split break the id to three part, companyID, nodeCode and tailID
func (r RowID) Split() (companyID RowID, nodeCode uint64, tailID RowID) {
	tailID = r % consts.IDRange
	r = r / consts.IDRange

	nodeCode = r.ToUint64() % consts.NodeRange
	r = r / consts.NodeRange

	companyID = r

	return
}

// Value is used to save to the database
func (r RowID) Value() (driver.Value, error) {
	// TODO: instead of sprintf use number conversion
	return fmt.Sprint(r), nil
}

/*
// TODO: I don't know what is the usage of this
// https://github.com/jinzhu/gorm/issues/47

func (r *RowID) Scan(value interface{}) error {
	value, _ = value.(uint64)
	*r = RowID(value)
	return nil
}
*/

// func (r *RowID) Scan(value interface{}) error {
// 	fmt.Println("**************************", value)
// 	result, _ := value.(RowID)
// 	*r = RowID(result)
// 	return nil
// }
