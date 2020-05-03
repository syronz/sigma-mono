// +build wireinject

package router

import (
	"sigmamono/api"
	"sigmamono/internal/core"
	"sigmamono/repo"
	"sigmamono/service"

	"github.com/google/wire"
)

func initRoleAPI(engine *core.Engine) api.RoleAPI {
	wire.Build(repo.ProvideRoleRepo, service.ProvideRoleService, api.ProvideRoleAPI)
	return api.RoleAPI{}
}

func initUserAPI(engine *core.Engine) api.UserAPI {
	wire.Build(repo.ProvideUserRepo, service.ProvideUserService, api.ProvideUserAPI)
	return api.UserAPI{}
}

// func initSettingAPI(engine *core.Engine) api.SettingAPI {
// 	wire.Build(repo.ProvideSettingRepo, service.ProvideSettingService, api.ProvideSettingAPI)
// 	return api.SettingAPI{}
// }

// func initCompanyAPI(e *core.Engine) api.CompanyAPI {
// 	wire.Build(repo.ProvideCompanyRepo, service.ProvideCompanyService,
// 		api.ProvideCompanyAPI)
// 	return api.CompanyAPI{}
// }
