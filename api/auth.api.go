package api

import (
	"net/http"
	"sigmamono/internal/core"
	"sigmamono/internal/enum/event"
	"sigmamono/internal/response"
	"sigmamono/internal/term"
	"sigmamono/model"
	"sigmamono/service"

	"github.com/gin-gonic/gin"
)

// AuthAPI for injecting auth service
type AuthAPI struct {
	Service service.AuthServ
	Engine  *core.Engine
}

// ProvideAuthAPI for auth used in wire
func ProvideAuthAPI(p service.AuthServ) AuthAPI {
	return AuthAPI{Service: p, Engine: p.Engine}
}

// Login auth
func (p *AuthAPI) Login(c *gin.Context) {
	var auth model.Auth
	resp := response.New(p.Engine, c)

	if err := c.BindJSON(&auth); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, err)
		return
	}

	user, err := p.Service.Login(auth)
	if err != nil {
		resp.Error(err).JSON()
		// p.Engine.Record(c, term.Auth_login_failed, auth.Username, len(auth.Password))
		resp.Record(term.Auth_login_failed, auth.Username, len(auth.Password))
		return
	}

	// p.Engine.Record(c, event.Login, nil, user)

	resp.Record(event.Login, nil, user)
	resp.Status(http.StatusOK).
		Message(term.User_loged_in_successfully).
		JSON(user)
}
