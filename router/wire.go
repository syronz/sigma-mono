// +build wireinject

package router

import (
	"radiusbilling/api"
	"radiusbilling/internal/core"
	"radiusbilling/repo"
	"radiusbilling/service"

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
