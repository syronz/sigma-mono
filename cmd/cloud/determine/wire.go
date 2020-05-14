// +build wireinject

package determine

import (
	"sigmamono/api"
	"sigmamono/internal/core"
	"sigmamono/repo"
	"sigmamono/service"

	"github.com/google/wire"
)

func initVersionAPI(e *core.Engine) api.VersionAPI {
	wire.Build(repo.ProvideVersionRepo, service.ProvideVersionService,
		api.ProvideVersionAPI)
	return api.VersionAPI{}
}

func initLicenseAPI(e *core.Engine) api.LicenseAPI {
	wire.Build(repo.ProvideLicenseRepo, service.ProvideLicenseService,
		api.ProvideLicenseAPI)
	return api.LicenseAPI{}
}

func initRegisterAPI(e *core.Engine) api.RegisterAPI {
	wire.Build(service.ProvideRegisterService, api.ProvideRegisterAPI)
	return api.RegisterAPI{}
}

func initCompanyAPI(e *core.Engine) api.CompanyAPI {
	wire.Build(repo.ProvideCompanyRepo, service.ProvideCompanyService,
		api.ProvideCompanyAPI)
	return api.CompanyAPI{}
}

func initNodeAPI(e *core.Engine) api.NodeAPI {
	wire.Build(repo.ProvideNodeRepo, service.ProvideNodeService,
		api.ProvideNodeAPI)
	return api.NodeAPI{}
}
