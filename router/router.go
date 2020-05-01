package router

import (
	"radiusbilling/internal/core"
	// "radiusbilling/internal/middleware"

	"github.com/gin-gonic/gin"
)

// Route trigger router and api methods
func Route(rg gin.RouterGroup, engine *core.Engine) {
	userAPI := initUserAPI(engine)
	roleAPI := initRoleAPI(engine)

	// rg.Use(middleware.AuthGuard(engine))
	rg.GET("/username/:username", userAPI.FindByUsername)
	// rg.GET("/users", userAPI.List)
	// rg.GET("/users/:userID", userAPI.FindByID)
	rg.POST("/users", userAPI.Create)
	// rg.PUT("/users/:userID", userAPI.Update)
	// rg.DELETE("/users/:userID", userAPI.Delete)
	// rg.GET("excel/users", userAPI.Excel)

	rg.POST("/roles", roleAPI.Create)

	rg.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

}
