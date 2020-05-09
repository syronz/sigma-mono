package table

import (
	"fmt"
	"sigmamono/internal/consts"
	"sigmamono/internal/core"
	"sigmamono/internal/enum/feature"
	"sigmamono/internal/types"
	"sigmamono/model"
	"sigmamono/repo"
	"sigmamono/service"
	"strings"
)

// InsertVersions for add required versions
func InsertVersions(engine *core.Engine) {
	versionRepo := repo.ProvideVersionRepo(engine)
	versionService := service.ProvideVersionService(versionRepo)
	versions := []model.Version{
		{
			GormCol: types.GormCol{
				ID: consts.FreeVersionID,
			},
			Name: "Free",
			Features: fmt.Sprint(feature.Activity, ", ",
				feature.Autobackup, ", ",
				feature.BalanceSheet, ", ",
				feature.CashFlowStatement, ", ",
				feature.CurrencyNumber, ", ",
				feature.IncomeStatemnt, ", ",
				feature.Inventory),
			NodeCount:     5,
			LocationCount: 10,
			UserCount:     20,
			MonthExpire:   12,
			Description:   "Free version for 12 month",
		},
		{
			GormCol: types.GormCol{
				ID: 36,
			},
			Name:          "Enterprise",
			Features:      strings.Join(feature.Features, ", "),
			NodeCount:     consts.MaxNodeCode - consts.MinNodeCode,
			LocationCount: 999,
			UserCount:     9999,
			MonthExpire:   120,
			Description:   "Enterprise version for 120 month",
		},
		{
			GormCol: types.GormCol{
				ID: 55,
			},
			Name:          "for update 1",
			Features:      "for update 1",
			NodeCount:     8,
			LocationCount: 999,
			UserCount:     9999,
			MonthExpire:   120,
			Description:   "for update 1",
		},
		{
			GormCol: types.GormCol{
				ID: 66,
			},
			Name:          "for update 2",
			Features:      "for update 2",
			NodeCount:     8,
			LocationCount: 999,
			UserCount:     9999,
			MonthExpire:   120,
			Description:   "for update 2",
		},
		{
			GormCol: types.GormCol{
				ID: 77,
			},
			Name:          "for delete 1",
			Features:      "for delete 1",
			NodeCount:     8,
			LocationCount: 777,
			UserCount:     777,
			MonthExpire:   777,
			Description:   "for delete 1",
		},
		{
			GormCol: types.GormCol{
				ID: 91,
			},
			Name:          "for search 1",
			Features:      "for search 1",
			NodeCount:     10,
			LocationCount: 10,
			UserCount:     10,
			MonthExpire:   10,
			Description:   "searchTerm1",
		},
		{
			GormCol: types.GormCol{
				ID: 92,
			},
			Name:          "for search 2",
			Features:      "for search 2",
			NodeCount:     10,
			LocationCount: 10,
			UserCount:     10,
			MonthExpire:   10,
			Description:   "searchTerm1",
		},
		{
			GormCol: types.GormCol{
				ID: 93,
			},
			Name:          "for search 3",
			Features:      "for search 3",
			NodeCount:     10,
			LocationCount: 10,
			UserCount:     10,
			MonthExpire:   10,
			Description:   "searchTerm1",
		},
		{
			GormCol: types.GormCol{
				ID: 94,
			},
			Name:          "for search 4",
			Features:      "for search 4",
			NodeCount:     10,
			LocationCount: 10,
			UserCount:     10,
			MonthExpire:   10,
			Description:   "searchTerm1",
		},
		{
			GormCol: types.GormCol{
				ID: 95,
			},
			Name:          "for search 5",
			Features:      "for search 5",
			NodeCount:     10,
			LocationCount: 10,
			UserCount:     10,
			MonthExpire:   10,
			Description:   "searchTerm1",
		},
	}

	for _, v := range versions {
		if _, err := versionService.Save(v); err != nil {
			engine.ServerLog.Fatal(err)
		}

	}

}
