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
	"sigmamono/utils/excel"

	"github.com/gin-gonic/gin"
)

const thisVersion = "version"
const thisVersions = "versions"

// VersionAPI for injecting version service
type VersionAPI struct {
	Service service.VersionServ
	Engine  *core.Engine
}

// ProvideVersionAPI for version is used in wire
func ProvideVersionAPI(c service.VersionServ) VersionAPI {
	return VersionAPI{Service: c, Engine: c.Engine}
}

// FindByID is used for fetch a version by it's id
func (p *VersionAPI) FindByID(c *gin.Context) {
	resp := response.New(p.Engine, c)
	var err error
	var version model.Version

	if resp.CheckAccess(resource.VersionRead) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	if version.ID, err = types.StrToRowID(c.Param("versionID")); err != nil {
		resp.Status(http.StatusNotAcceptable).Error(err).MessageT(term.Invalid_ID).JSON()
		return
	}

	if version, err = p.Service.FindByID(version.ID); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(event.VersionView)
	resp.Status(http.StatusOK).
		MessageT(term.V_info, thisVersion).
		JSON(version)
}

// List of versions
func (p *VersionAPI) List(c *gin.Context) {
	resp := response.New(p.Engine, c)

	if resp.CheckAccess(resource.VersionWrite) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	params := param.Get(c, p.Engine, thisVersions)

	data, err := p.Service.List(params)
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(event.VersionList)
	resp.Status(http.StatusOK).
		MessageT(term.List_of_V, thisVersions).
		JSON(data)
}

// Create version
func (p *VersionAPI) Create(c *gin.Context) {
	var version model.Version
	resp := response.New(p.Engine, c)

	if resp.CheckAccess(resource.VersionWrite) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	if err := c.ShouldBindJSON(&version); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, err)
		return
	}

	createdVersion, err := p.Service.Save(version)
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(event.VersionCreate, nil, version)
	resp.Status(http.StatusOK).
		MessageT(term.V_created_successfully, thisVersion).
		JSON(createdVersion)
}

// Update version
func (p *VersionAPI) Update(c *gin.Context) {
	resp := response.New(p.Engine, c)
	var err error

	if resp.CheckAccess(resource.VersionWrite) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	var version, versionBefore, versionUpdated model.Version

	version.ID, err = types.StrToRowID(c.Param("versionID"))
	if err != nil {
		resp.Error(term.Invalid_ID).JSON()
		return
	}

	if err = c.ShouldBindJSON(&version); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, err)
		return
	}

	if versionBefore, err = p.Service.FindByID(version.ID); err != nil {
		resp.Status(http.StatusNotFound).Error(term.Record_Not_Found).JSON()
		return
	}

	if versionUpdated, err = p.Service.Save(version); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(event.VersionUpdate, versionBefore, version)
	resp.Status(http.StatusOK).
		MessageT(term.V_updated_successfully, thisVersion).
		JSON(versionUpdated)
}

// Delete version
func (p *VersionAPI) Delete(c *gin.Context) {
	resp := response.New(p.Engine, c)
	var err error
	var version model.Version

	if resp.CheckAccess(resource.VersionWrite) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	if version.ID, err = types.StrToRowID(c.Param("versionID")); err != nil {
		resp.Error(term.Invalid_ID).JSON()
		return
	}

	if version, err = p.Service.Delete(version.ID); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(event.VersionDelete, version)
	resp.Status(http.StatusOK).
		MessageT(term.V_deleted_successfully, thisVersion).
		JSON()
}

// Excel generate excel files based on search
func (p *VersionAPI) Excel(c *gin.Context) {
	resp := response.New(p.Engine, c)

	if resp.CheckAccess(resource.VersionExcel) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	params := param.Get(c, p.Engine, thisVersions)

	versions, err := p.Service.Excel(params)
	if err != nil {
		resp.Status(http.StatusNotFound).Error(term.Record_Not_Found).JSON()
		return
	}

	resp.Record(event.VersionExcel)

	ex := excel.New("version")
	ex.AddSheet("Versions").
		AddSheet("Summary").
		Active("Versions").
		SetPageLayout("landscape", "A4").
		SetPageMargins(0.2).
		SetHeaderFooter().
		SetColWidth("B", "C", 15.3).
		SetColWidth("M", "M", 20).
		Active("Summary").
		SetColWidth("A", "D", 20).
		Active("Versions").
		WriteHeader("ID", "Name", "Features", "Node Count", "Location Count", "User Count",
			"Month Expire", "Description").
		SetSheetFields("ID", "Name", "Features", "NodeCount", "LocationCount", "UserCount",
			"MonthExpire", "Description").
		WriteData(versions).
		AddTable()

	buffer, downloadName, err := ex.Generate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &response.Result{
			Message: "Error in generating Excel file",
		})
		return
	}

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+downloadName)
	c.Data(http.StatusOK, "application/octet-stream", buffer.Bytes())

}
