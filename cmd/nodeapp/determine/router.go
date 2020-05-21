package determine

import (
	"sigmamono/internal/core"
	"sigmamono/internal/middleware"

	"github.com/gin-gonic/gin"
)

// Route trigger router and api methods
func Route(rg gin.RouterGroup, engine *core.Engine) {
	stationAPI := initStationAPI(engine)

	rg.POST("/activate/register/nodeapp", stationAPI.RegisterNode)

	rg.Use(middleware.AuthGuard(engine))

}
