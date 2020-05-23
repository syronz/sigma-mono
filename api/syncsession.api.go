package api

import (
	"net/http"
	"sigmamono/internal/core"
	"sigmamono/internal/core/access/resource"
	"sigmamono/internal/param"
	"sigmamono/internal/response"
	"sigmamono/internal/term"
	"sigmamono/internal/types"
	"sigmamono/model"
	"sigmamono/service"
	"sigmamono/utils/excel"

	"github.com/gin-gonic/gin"
)

const thisSyncSession = "syncSession"
const thisSyncSessions = "syncSessions"

// SyncSessionAPI for injecting syncSession service
type SyncSessionAPI struct {
	Service service.SyncSessionServ
	Engine  *core.Engine
}

// ProvideSyncSessionAPI for syncSession is used in wire
func ProvideSyncSessionAPI(c service.SyncSessionServ) SyncSessionAPI {
	return SyncSessionAPI{Service: c, Engine: c.Engine}
}

// FindByID is used for fetch a syncSession by it's id
func (p *SyncSessionAPI) FindByID(c *gin.Context) {
	resp := response.New(p.Engine, c)
	var err error
	var syncSession model.SyncSession

	if resp.CheckAccess(resource.SyncSessionRead) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	if syncSession.ID, err = types.StrToRowID(c.Param("syncSessionID")); err != nil {
		resp.Error(term.Invalid_ID).JSON()
		return
	}

	if syncSession, err = p.Service.FindByID(syncSession.ID); err != nil {
		resp.Status(http.StatusNotFound).Error(err).MessageT(term.Record_Not_Found).JSON()
		return
	}

	resp.Status(http.StatusOK).
		MessageT(term.V_info, thisSyncSession).
		JSON(syncSession)
}

// DELETE BELOW
// List of syncSessions
func (p *SyncSessionAPI) List(c *gin.Context) {
	resp := response.New(p.Engine, c)

	if resp.CheckAccess(resource.SyncSessionWrite) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	params := param.Get(c, p.Engine, thisSyncSessions)

	data, err := p.Service.List(params)
	p.Engine.Debug(err)
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Status(http.StatusOK).
		MessageT(term.List_of_V, thisSyncSessions).
		JSON(data)
}

// Initiate syncSession
func (p *SyncSessionAPI) Initiate(c *gin.Context) {
	var syncSession model.SyncSession
	resp := response.New(p.Engine, c)

	if err := c.ShouldBindJSON(&syncSession); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, err)
		return
	}

	params := param.Get(c, p.Engine, thisSyncSessions)

	createdSyncSession, err := p.Service.Create(syncSession, params)
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	createdSyncSession.Delay = 1500

	resp.Status(http.StatusOK).
		MessageT(term.V_created_successfully, thisSyncSession).
		JSON(createdSyncSession)
}

// Delete syncSession
func (p *SyncSessionAPI) Delete(c *gin.Context) {
	resp := response.New(p.Engine, c)
	var err error
	var syncSession model.SyncSession

	if resp.CheckAccess(resource.SyncSessionWrite) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	if syncSession.ID, err = types.StrToRowID(c.Param("syncSessionID")); err != nil {
		resp.Error(term.Invalid_ID).JSON()
		return
	}

	params := param.Get(c, p.Engine, thisSyncSessions)

	if syncSession, err = p.Service.Delete(syncSession.ID, params); err != nil {
		resp.Status(http.StatusInternalServerError).Error(err).JSON()
		return
	}

	resp.Status(http.StatusOK).
		MessageT(term.V_deleted_successfully, thisSyncSession).
		JSON()
}

// Excel generate excel files based on search
func (p *SyncSessionAPI) Excel(c *gin.Context) {
	resp := response.New(p.Engine, c)

	if resp.CheckAccess(resource.SyncSessionExcel) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	params := param.Get(c, p.Engine, thisSyncSessions)
	syncSessions, err := p.Service.Excel(params)
	if err != nil {
		resp.Status(http.StatusNotFound).Error(term.Record_Not_Found).JSON()
		return
	}

	ex := excel.New("syncSession")
	ex.AddSheet("SyncSessions").
		AddSheet("Summary").
		Active("SyncSessions").
		SetPageLayout("landscape", "A4").
		SetPageMargins(0.2).
		SetHeaderFooter().
		SetColWidth("B", "C", 15.3).
		SetColWidth("M", "M", 20).
		Active("Summary").
		SetColWidth("A", "D", 20).
		Active("SyncSessions").
		WriteHeader("ID", "Name", "Legal Name", "Server Address", "Expiration", "Plan",
			"Detail", "Phone", "Email", "Website", "Type", "Code", "Updated At").
		SetSheetFields("ID", "Name", "LegalName", "ServerAddress", "Expiration", "Plan",
			"Detail", "Phone", "Email", "Website", "Type", "Code", "UpdatedAt").
		WriteData(syncSessions).
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
