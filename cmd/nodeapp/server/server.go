package server

import (
	"fmt"
	"log"
	"net/http"
	"sigmamono/cmd/nodeapp/determine"
	"sigmamono/internal/core"
	"sigmamono/internal/middleware"
	"sigmamono/internal/response"
	"sigmamono/internal/term"
	"sigmamono/router"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

// Start initiate the server
func Start(engine *core.Engine) *gin.Engine {

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "127.0.0.1"
		},
		//MaxAge: 12 * time.Hour,
	}))
	r.Use(middleware.APILogger(engine))

	// No Route "Not Found"
	notFoundRoute(r, engine)

	rg := r.Group("/api/nodeapp/v1")
	{
		router.Route(*rg, engine)
		determine.Route(*rg, engine)
	}

	if err := r.Run(fmt.Sprintf("%v:%v", engine.Env.NodeApp.ADDR, engine.Env.NodeApp.Port)); err != nil {
		log.Fatalln(err)
	}

	return r
}

func notFoundRoute(r *gin.Engine, engine *core.Engine) {
	r.NoRoute(func(c *gin.Context) {
		response.New(engine, c).Status(http.StatusNotFound).Error(term.Route_not_found).JSON()
	})
}
