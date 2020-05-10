package table

import (
	"sigmamono/internal/consts"
	"sigmamono/internal/core"
	"sigmamono/internal/enum/accountdirection"
	"sigmamono/internal/enum/accountstatus"
	"sigmamono/internal/enum/accounttype"
	"sigmamono/internal/enum/lang"
	"sigmamono/internal/types"
	"sigmamono/model"
	"sigmamono/repo"
	"sigmamono/service"
)

// InsertUsers for add required users
func InsertUsers(engine *core.Engine) {
	userRepo := repo.ProvideUserRepo(engine)
	userService := service.ProvideUserService(userRepo)
	users := []model.User{
		{
			ID:       1001101000000002,
			RoleID:   1001101000000001,
			Username: consts.NodeAppUsername,
			Password: consts.NodeAppPassword,
			Language: string(lang.Ku),
			Account: model.Account{
				FixedCol: types.FixedCol{
					ID: 1001101000000002,
				},
				ParentID:  types.RowIDPointer(1001101000000001),
				CompanyID: 1001,
				NodeCode:  101,
				Direction: accountdirection.Direct,
				Type:      accounttype.Asset,
				Name:      engine.Env.Cloud.SuperAdminUsername,
				Code:      110001,
				Status:    accountstatus.Active,
			},
		},
		/*
			{
				FixedCol: types.FixedCol{
					ID: 2,
				},
				RoleID:   2,
				Name:     "Admin",
				Username: "admin",
				Password: "omega",
			},
		*/
	}

	for _, v := range users {
		if _, err := userService.Save(v); err != nil {
			engine.ServerLog.Fatal(err)
		}

	}

}
