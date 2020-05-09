package api

import (
	"net/http"
	"sigmamono/internal/core"
	"sigmamono/internal/core/access/resource"
	"sigmamono/internal/enum/event"
	"sigmamono/internal/param"
	"sigmamono/internal/response"
	"sigmamono/internal/term"
	"sigmamono/internal/types"
	"sigmamono/model"
	"sigmamono/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

const thisLicense = "license"
const thisLicenses = "licenses"

// LicenseAPI for injecting license service
type LicenseAPI struct {
	Service service.LicenseServ
	Engine  *core.Engine
}

// ProvideLicenseAPI for license is used in wire
func ProvideLicenseAPI(c service.LicenseServ) LicenseAPI {
	return LicenseAPI{Service: c, Engine: c.Engine}
}

// GeneratePublic create licenses according to the versions table, anyone can use it
func (p *LicenseAPI) GeneratePublic(c *gin.Context) {
	resp := response.New(p.Engine, c)

	if resp.CheckAccess(resource.LicenseWrite) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	var versionID types.RowID
	var countStr string
	var err error
	var count int
	var licenses []model.License

	if versionID, err = types.StrToRowID(c.Param("versionID")); err != nil {
		resp.Error(term.Invalid_ID).JSON()
		return
	}

	countStr = c.Param("count")
	if count, err = strconv.Atoi(countStr); err != nil {
		resp.Error(err).JSON()
	}

	if licenses, err = p.Service.GeneratePublic(versionID, count); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(event.LicenseCreate, nil, []interface{}{versionID, countStr})
	resp.Status(http.StatusOK).
		MessageT(term.V_created_successfully, thisLicense).
		JSON(licenses)
}

// GeneratePrivate create license just for one company and it is related to the
// company name, if they change the company's legal name the license not working
func (p *LicenseAPI) GeneratePrivate(c *gin.Context) {

}

// Update parse the license and if it is eligible update license colunn in company
func (p *LicenseAPI) Update(c *gin.Context) {
	resp := response.New(p.Engine, c)
	var err error
	var license model.License
	var companyKey string

	if err = c.ShouldBindJSON(&license); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, err)
		return
	}

	params := param.Get(c, p.Engine, thisLicense)

	if companyKey, err = p.Service.Update(license, params); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Status(http.StatusOK).
		MessageT(term.V_updated_successfully, thisLicense).
		JSON(companyKey)

}
