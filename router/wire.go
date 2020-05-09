// +build wireinject

package router

import (
	"sigmamono/api"
	"sigmamono/internal/core"
	"sigmamono/repo"
	"sigmamono/service"

	"github.com/google/wire"
)

func initSettingAPI(e *core.Engine) api.SettingAPI {
	wire.Build(repo.ProvideSettingRepo, service.ProvideSettingService, api.ProvideSettingAPI)
	return api.SettingAPI{}
}

func initUserAPI(engine *core.Engine) api.UserAPI {
	wire.Build(repo.ProvideUserRepo, service.ProvideUserService, api.ProvideUserAPI)
	return api.UserAPI{}
}

func initRoleAPI(e *core.Engine) api.RoleAPI {
	wire.Build(repo.ProvideRoleRepo, service.ProvideRoleService, api.ProvideRoleAPI)
	return api.RoleAPI{}
}

func initAccountAPI(e *core.Engine) api.AccountAPI {
	wire.Build(repo.ProvideAccountRepo, service.ProvideAccountService,
		api.ProvideAccountAPI)
	return api.AccountAPI{}
}
