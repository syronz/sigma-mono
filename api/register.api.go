package api

import (
	"net/http"
	"sigmamono/dto"
	"sigmamono/internal/core"
	"sigmamono/internal/enum/event"
	"sigmamono/internal/response"
	"sigmamono/internal/term"
	"sigmamono/service"

	"github.com/gin-gonic/gin"
)

// RegisterAPI for injecting level 2 of domain
type RegisterAPI struct {
	Service service.RegisterServ
	Engine  *core.Engine
}

// ProvideRegisterAPI returns level 2 of domain inside the wire
func ProvideRegisterAPI(c service.RegisterServ) RegisterAPI {
	return RegisterAPI{Service: c, Engine: c.Engine}
}

// Register is used for defining company, node and user for the first time
func (p *RegisterAPI) Register(c *gin.Context) {
	var register dto.Register
	var err error
	resp := response.New(p.Engine, c)

	if err := c.ShouldBindJSON(&register); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, err)
		p.Engine.CheckError(err, "error in binding to register")
		return
	}

	register.Company.LegalName = register.Company.Name
	register.Node.Phone = register.Company.Phone

	if register, err = p.Service.Register(register); err != nil {
		resp.Error(err).JSON()
		return
	}

	// _ = resp

	// c.JSON(200, gin.H{"e": err, "register": register, "company": company})

	resp.Record(event.CompanyRegister, nil, register)
	resp.Status(http.StatusOK).
		MessageT(term.V_created_successfully, "company").
		JSON(register)

}
