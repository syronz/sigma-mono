package determine

import (
	"sigmamono/internal/core"
	"sigmamono/internal/middleware"

	"github.com/gin-gonic/gin"
)

// Route trigger router and api methods
func Route(rg gin.RouterGroup, engine *core.Engine) {
	syncSessionAPI := initSyncSessionAPI(engine)

	rg.Use(middleware.AuthGuard(engine))
	rg.POST("initiate", syncSessionAPI.Initiate)

}
