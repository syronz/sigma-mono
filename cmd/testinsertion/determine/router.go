package determine

import (
	"sigmamono/internal/core"
	"sigmamono/internal/middleware"

	"github.com/gin-gonic/gin"
)

// Route trigger router and api methods
func Route(rg gin.RouterGroup, engine *core.Engine) {
	companyAPI := initCompanyAPI(engine)
	nodeAPI := initNodeAPI(engine)

	rg.Use(middleware.AuthGuard(engine))
	rg.GET("/companies", companyAPI.List)
	rg.POST("/companies", companyAPI.Create)
	rg.PUT("/companies/:companyID", companyAPI.Update)
	rg.DELETE("/companies/:companyID", companyAPI.Delete)
	rg.GET("/companies/:companyID", companyAPI.FindByID)
	rg.GET("excel/companies", companyAPI.Excel)

	rg.GET("/nodes", nodeAPI.List)
	rg.POST("/nodes", nodeAPI.Create)
	rg.PUT("/nodes/:nodeID", nodeAPI.Update)
	rg.DELETE("/nodes/:nodeID", nodeAPI.Delete)
	rg.GET("/nodes/:nodeID", nodeAPI.FindByID)
	rg.GET("/excel/nodes", nodeAPI.Excel)
}
