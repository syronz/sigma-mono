// +build wireinject

package determine

import (
	"sigmamono/api"
	"sigmamono/internal/core"
	"sigmamono/repo"
	"sigmamono/service"

	"github.com/google/wire"
)

func initStationAPI(e *core.Engine) api.StationAPI {
	wire.Build(repo.ProvideStationRepo, service.ProvideStationService, api.ProvideStationAPI)
	return api.StationAPI{}
}
