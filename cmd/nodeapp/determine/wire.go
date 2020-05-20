// +build wireinject

package determine

import (
	"sigmamono/api"
	"sigmamono/internal/core"
	"sigmamono/repo"
	"sigmamono/service"

	"github.com/google/wire"
)

func initBondAPI(e *core.Engine) api.BondAPI {
	wire.Build(repo.ProvideBondRepo, service.ProvideBondService, api.ProvideBondAPI)
	return api.BondAPI{}
}
