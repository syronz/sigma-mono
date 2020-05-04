package table

import (
	"sigmamono/internal/core"
	r "sigmamono/internal/core/access/resource"
	"sigmamono/internal/types"
	"sigmamono/model"
	"sigmamono/repo"
	"sigmamono/service"
	"strings"
)

// InsertRoles for add required roles
func InsertRoles(engine *core.Engine) {
	roleRepo := repo.ProvideRoleRepo(engine)
	roleService := service.ProvideRoleService(roleRepo)
	roles := []model.Role{
		{
			FixedCol: types.FixedCol{
				ID: 1001101000000001,
			},
			CompanyID: 1001,
			NodeCode:  101,
			Name:      "Super-Admin",
			Resources: strings.Join([]string{
				r.SupperAccess,
				r.CompanyRead, r.CompanyWrite, r.CompanyExcel,
				r.NodeRead, r.NodeWrite, r.NodeExcel,
				r.VersionRead, r.VersionWrite, r.VersionExcel,
				r.LicenseWrite,
				r.SettingRead, r.SettingWrite, r.SettingExcel,
				r.UserNames, r.UserWrite, r.UserRead, r.UserReport,
				r.ActivitySelf, r.ActivityAll,
				r.AccountNames, r.AccountRead, r.AccountWrite, r.AccountExcel,
				r.RoleRead, r.RoleWrite, r.RoleExcel,
			}, ", "),
			Description: "super-admin has all privileges - do not edit",
		},
		{
			FixedCol: types.FixedCol{
				ID: 1001101000000002,
			},
			CompanyID: 1001,
			NodeCode:  101,
			Name:      "Admin",
			Resources: strings.Join([]string{
				r.CompanyRead, r.CompanyWrite, r.CompanyExcel,
				r.UserNames, r.UserWrite, r.UserRead, r.UserReport,
				r.NodeRead, r.NodeWrite, r.NodeExcel,
				r.ActivitySelf, r.ActivityAll,
				r.AccountNames, r.AccountRead, r.AccountWrite, r.AccountExcel,
				r.RoleRead, r.RoleWrite,
			}, ", "),
			Description: "admin has all privileges - do not edit",
		},
		{
			FixedCol: types.FixedCol{
				ID: 1001101000000003,
			},
			CompanyID:   1001,
			NodeCode:    101,
			Name:        "Cashier",
			Resources:   strings.Join([]string{r.ActivitySelf}, ", "),
			Description: "cashier has all privileges - after migration reset",
		},
	}
	// roles[0].ID = types.RowID(1)
	// roles[1].ID = types.RowID(2)
	// roles[2].ID = types.RowID(3)

	for _, v := range roles {
		if _, err := roleService.Save(v); err != nil {
			engine.ServerLog.Fatal(err)
		}

	}

}
