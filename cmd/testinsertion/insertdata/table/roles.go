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
		{
			FixedCol: types.FixedCol{
				ID: 1001101000000004,
			},
			CompanyID:   1001,
			NodeCode:    101,
			Name:        "for foreign 1",
			Resources:   r.SupperAccess,
			Description: "for foreign 1",
		},
		{
			FixedCol: types.FixedCol{
				ID: 1001101000000005,
			},
			CompanyID:   1001,
			NodeCode:    101,
			Name:        "for update 1",
			Resources:   r.SupperAccess,
			Description: "for update 1",
		},
		{
			FixedCol: types.FixedCol{
				ID: 1001101000000006,
			},
			CompanyID:   1001,
			NodeCode:    101,
			Name:        "for update 2",
			Resources:   r.SupperAccess,
			Description: "for update 2",
		},
		{
			FixedCol: types.FixedCol{
				ID: 1001101000000007,
			},
			CompanyID:   1001,
			NodeCode:    101,
			Name:        "for delete 1",
			Resources:   r.SupperAccess,
			Description: "for delete 1",
		},
		{
			FixedCol: types.FixedCol{
				ID: 1001101000000008,
			},
			CompanyID:   1001,
			NodeCode:    101,
			Name:        "for search 1",
			Resources:   r.SupperAccess,
			Description: "searchTerm1",
		},
		{
			FixedCol: types.FixedCol{
				ID: 1001101000000009,
			},
			CompanyID:   1001,
			NodeCode:    101,
			Name:        "for search 2",
			Resources:   r.SupperAccess,
			Description: "searchTerm1",
		},
		{
			FixedCol: types.FixedCol{
				ID: 1001101000000010,
			},
			CompanyID:   1001,
			NodeCode:    101,
			Name:        "for search 3",
			Resources:   r.SupperAccess,
			Description: "searchTerm1",
		},
		{
			FixedCol: types.FixedCol{
				ID: 1002101000000001,
			},
			CompanyID:   1002,
			NodeCode:    101,
			Name:        "for delete 2",
			Resources:   r.SupperAccess,
			Description: "for delete 2",
		},
	}

	for _, v := range roles {
		if _, err := roleService.Save(v); err != nil {
			engine.ServerLog.Fatal(err)
		}

	}

}
