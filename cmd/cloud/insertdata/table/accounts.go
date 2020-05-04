package table

import (
	"sigmamono/internal/core"
	"sigmamono/internal/enum/accounttype"
	"sigmamono/internal/types"
	"sigmamono/model"
	"sigmamono/repo"
	"sigmamono/service"
)

// InsertAccounts for add required accounts
func InsertAccounts(engine *core.Engine) {
	accountRepo := repo.ProvideAccountRepo(engine)
	accountService := service.ProvideAccountService(accountRepo)
	accounts := []model.Account{
		{
			FixedCol: types.FixedCol{
				ID: 1001101000000001,
			},
			CompanyID: 1001,
			NodeCode:  101,
			Name:      accounttype.Asset,
			Status:    "active",
			Code:      122,
			Type:      "asset",
			Readonly:  true,
			Score:     1,
			Direction: "direct",
		},
	}

	for _, v := range accounts {
		if _, err := accountService.Save(v); err != nil {
			engine.ServerLog.Fatal(err)
		}

	}

}
