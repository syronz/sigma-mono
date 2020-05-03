package server

import (
	"fmt"
	"log"
	"net/http"
	// "sigmamono/domain/accounting"
	// "sigmamono/domain/activation"
	// "sigmamono/domain/administration"
	// "sigmamono/domain/central"
	// "sigmamono/domain/core"
	// "sigmamono/domain/sync"
	"sigmamono/cmd/cloud/determine"
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

	// engine.Agg = make(chan types.Aggregate)

	// domains := types.Domains{
	// 	Accounting:     make(chan types.Aggregate),
	// 	Activation:     make(chan types.Aggregate),
	// 	Administration: make(chan types.Aggregate),
	// 	Central:        make(chan types.Aggregate),
	// 	Sync:           make(chan types.Aggregate),
	// }

	// go AggObserver(engine.Agg, domains)

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

	rg := r.Group("/api/cloud/v1")
	{
		router.Route(*rg, engine)
		determine.Route(*rg, engine)
		// accounting.Router(*rg, engine, domains.Accounting)
		// activation.Router(*rg, engine, domains.Activation)
		// administration.Router(*rg, engine, domains.Administration)
		// central.Router(*rg, engine, domains.Central)
		// sync.Router(*rg, engine, domains.Sync)
	}

	if err := r.Run(fmt.Sprintf("%v:%v", engine.Env.Cloud.ADDR, engine.Env.Cloud.Port)); err != nil {
		log.Fatalln(err)
	}

	return r
}

// AggObserver separate request among domains
// func AggObserver(aggChan chan types.Aggregate, domains types.Domains) {
// 	for agg := range aggChan {
// 		switch agg.Domain {
// 		case "accounting":
// 			domains.Accounting <- agg
// 		case "administration":
// 			domains.Administration <- agg
// 		case "central":
// 			domains.Central <- agg
// 		case "sync":
// 			domains.Sync <- agg
// 		}
// 	}
// }

func notFoundRoute(r *gin.Engine, engine *core.Engine) {
	r.NoRoute(func(c *gin.Context) {
		response.New(engine, c).Status(http.StatusNotFound).Error(term.Route_not_found).JSON()
	})
}
