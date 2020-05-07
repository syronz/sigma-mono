package api

import (
	"net/http"
	"sigmamono/internal/core"
	"sigmamono/internal/response"
	"sigmamono/model"
	"sigmamono/service"

	"github.com/gin-gonic/gin"
)

// Activity for injecting activity service
type Activity struct {
	Service service.Activity
	Engine  *core.Engine
}

// ProvideActivityAPI for activity is used in wire
func ProvideActivityAPI(c service.Activity) Activity {
	return Activity{Service: c, Engine: c.Engine}
}

// Create activity
func (p *Activity) Create(c *gin.Context) {
	var activity model.Activity
	resp := response.New(p.Engine, c)

	if err := c.ShouldBindJSON(&activity); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, err)
		return
	}

	// aggObj := types.Aggregate{Domain: "central"}
	// p.Engine.Agg <- aggObj

	createdActivity, err := p.Service.Save(activity)
	if err != nil {
		c.JSON(403, gin.H{"error": err.Error()})
		return
	}

	// _ = createdActivity
	resp.Status(203).
		Message("activity created successfully").
		JSON(createdActivity)
}
