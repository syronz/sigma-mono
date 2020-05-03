// +build wireinject

package router

import (
	"sigmamono/api"
	"sigmamono/internal/core"
	"sigmamono/repo"
	"sigmamono/service"

	"github.com/google/wire"
)

func initRoleAPI(e *core.Engine) api.RoleAPI {
	wire.Build(repo.ProvideRoleRepo, service.ProvideRoleService, api.ProvideRoleAPI)
	return api.RoleAPI{}
}

func initUserAPI(engine *core.Engine) api.UserAPI {
	wire.Build(repo.ProvideUserRepo, service.ProvideUserService, api.ProvideUserAPI)
	return api.UserAPI{}
}

func initAccountAPI(e *core.Engine) api.AccountAPI {
	wire.Build(repo.ProvideAccountRepo, service.ProvideAccountService,
		api.ProvideAccountAPI)
	return api.AccountAPI{}
}
