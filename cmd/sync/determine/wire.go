// +build wireinject

package determine

import (
	"sigmamono/api"
	"sigmamono/internal/core"
	"sigmamono/repo"
	"sigmamono/service"

	"github.com/google/wire"
)

func initSyncSessionAPI(e *core.Engine) api.SyncSessionAPI {
	wire.Build(repo.ProvideSyncSessionRepo, service.ProvideSyncSessionService,
		api.ProvideSyncSessionAPI)
	return api.SyncSessionAPI{}
}
