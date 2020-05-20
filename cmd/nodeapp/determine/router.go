package determine

import (
	"sigmamono/internal/core"
	"sigmamono/internal/middleware"

	"github.com/gin-gonic/gin"
)

// Route trigger router and api methods
func Route(rg gin.RouterGroup, engine *core.Engine) {
	bondAPI := initBondAPI(engine)

	rg.POST("/activate/register/nodeapp", bondAPI.RegisterNode)

	rg.Use(middleware.AuthGuard(engine))

}
