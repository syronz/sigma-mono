package determine

import (
	"sigmamono/internal/core"
	"sigmamono/internal/middleware"

	"github.com/gin-gonic/gin"
)

// Route trigger router and api methods
func Route(rg gin.RouterGroup, engine *core.Engine) {
	versionAPI := initVersionAPI(engine)
	companyAPI := initCompanyAPI(engine)
	nodeAPI := initNodeAPI(engine)
	licenseAPI := initLicenseAPI(engine)
	registerAPI := initRegisterAPI(engine)

	// route no need to token
	rg.POST("/register", registerAPI.Register)
	rg.POST("/activate/nodeapp", nodeAPI.Activate)

	rg.Use(middleware.AuthGuard(engine))

	// route need to token
	rg.GET("/versions", versionAPI.List)
	rg.POST("/versions", versionAPI.Create)
	rg.PUT("/versions/:versionID", versionAPI.Update)
	rg.DELETE("/versions/:versionID", versionAPI.Delete)
	rg.GET("/versions/:versionID", versionAPI.FindByID)
	rg.GET("/excel/versions", versionAPI.Excel)

	rg.POST("/licenses/public/:versionID/:count", licenseAPI.GeneratePublic)
	// rg.POST("/licenses/private", licenseAPI.GeneratePrivate)
	rg.PUT("/licenses", licenseAPI.Update)

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
