package server

import (
	"fmt"
	"log"
	"net/http"
	"sigmamono/cmd/sync/determine"
	"sigmamono/internal/core"
	"sigmamono/internal/response"
	"sigmamono/internal/term"

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

	// No Route "Not Found"
	notFoundRoute(r, engine)

	rg := r.Group("/api/sync/v1")
	{
		determine.Route(*rg, engine)
	}

	if err := r.Run(fmt.Sprintf("%v:%v", engine.Env.Sync.ADDR, engine.Env.Sync.Port)); err != nil {
		log.Fatalln(err)
	}

	return r
}

func notFoundRoute(r *gin.Engine, engine *core.Engine) {
	r.NoRoute(func(c *gin.Context) {
		response.New(engine, c).Status(http.StatusNotFound).Error(term.Route_not_found).JSON()
	})
}
