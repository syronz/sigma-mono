package helper

import (
	"sigmamono/internal/types"
	"testing"
)

func TestHeadID(t *testing.T) {

	samples := []struct {
		companyID types.RowID
		nodeCode  uint64
		result    types.RowID
	}{
		{1001, 101, 1001101},
		{1234, 234, 1234234},
		{1001, 0, 1001000},
		{0, 0, 0},
		{0, 1001, 1001},
	}

	for _, v := range samples {
		result := HeadID(v.companyID, v.nodeCode)
		if result != v.result {
			t.Errorf("for companyID: %v and nodeCode: %v, result should be %v, but it is %v",
				v.companyID, v.nodeCode, v.result, result)
		}
	}

}
