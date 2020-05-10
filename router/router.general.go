package router

import (
	"sigmamono/internal/core"
	"sigmamono/internal/middleware"

	"github.com/gin-gonic/gin"
)

// Route trigger router and api methods
func Route(rg gin.RouterGroup, engine *core.Engine) {
	userAPI := initUserAPI(engine)
	roleAPI := initRoleAPI(engine)
	accountAPI := initAccountAPI(engine)
	settingAPI := initSettingAPI(engine)
	authAPI := initAuthAPI(engine)

	rg.POST("/login", authAPI.Login)

	rg.Use(middleware.AuthGuard(engine))

	rg.GET("/settings", settingAPI.List)
	rg.GET("/settings/:settingID", settingAPI.FindByID)
	rg.PUT("/settings/:settingID", settingAPI.Update)
	rg.GET("excel/settings", settingAPI.Excel)

	rg.GET("/username/:username", userAPI.FindByUsername)
	rg.GET("/users", userAPI.List)
	rg.GET("/users/:userID", userAPI.FindByID)
	rg.POST("/users", userAPI.Create)
	rg.PUT("/users/:userID", userAPI.Update)
	rg.DELETE("/users/:userID", userAPI.Delete)
	rg.GET("excel/users", userAPI.Excel)

	rg.GET("/roles", roleAPI.List)
	rg.GET("/roles/:roleID", roleAPI.FindByID)
	rg.POST("/roles", roleAPI.Create)
	rg.PUT("/roles/:roleID", roleAPI.Update)
	rg.DELETE("/roles/:roleID", roleAPI.Delete)
	rg.GET("excel/roles", roleAPI.Excel)

	rg.GET("/accounts", accountAPI.List)
	rg.POST("/accounts", accountAPI.Create)
	rg.PUT("/accounts/:accountID", accountAPI.Update)
	rg.DELETE("/accounts/:accountID", accountAPI.Delete)
	rg.GET("/accounts/:accountID", accountAPI.FindByID)
	rg.GET("/excel/accounts", accountAPI.Excel)

}
