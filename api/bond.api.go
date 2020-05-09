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

const thisBond = "bond"
const thisBonds = "bonds"

// BondAPI for injecting bond service
type BondAPI struct {
	Service service.BondServ
	Engine  *core.Engine
}

// ProvideBondAPI for bond is used in wire
func ProvideBondAPI(c service.BondServ) BondAPI {
	return BondAPI{Service: c, Engine: c.Engine}
}

// FindByID is used for fetch a bond by it's id
func (p *BondAPI) FindByID(c *gin.Context) {
	resp := response.New(p.Engine, c)
	var err error
	var bond model.Bond

	if resp.CheckAccess(resource.BondRead) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	if bond.ID, err = types.StrToRowID(c.Param("bondID")); err != nil {
		resp.Status(http.StatusNotAcceptable).Error(err).MessageT(term.Invalid_ID).JSON()
		return
	}

	if bond, err = p.Service.FindByID(bond.ID); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(event.BondView)
	resp.Status(http.StatusOK).
		MessageT(term.V_info, thisBond).
		JSON(bond)
}

// List of bonds
func (p *BondAPI) List(c *gin.Context) {
	resp := response.New(p.Engine, c)

	if resp.CheckAccess(resource.BondWrite) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	params := param.Get(c, p.Engine, thisBonds)

	data, err := p.Service.List(params)
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(event.BondList)
	resp.Status(http.StatusOK).
		MessageT(term.List_of_V, thisBonds).
		JSON(data)
}

// Delete bond
func (p *BondAPI) Delete(c *gin.Context) {
	resp := response.New(p.Engine, c)
	var err error
	var bond model.Bond

	if resp.CheckAccess(resource.BondWrite) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	if bond.ID, err = types.StrToRowID(c.Param("bondID")); err != nil {
		resp.Error(term.Invalid_ID).JSON()
		return
	}

	if bond, err = p.Service.Delete(bond.ID); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(event.BondDelete, bond)
	resp.Status(http.StatusOK).
		MessageT(term.V_deleted_successfully, thisBond).
		JSON()
}

// Excel generate excel files based on search
func (p *BondAPI) Excel(c *gin.Context) {
	resp := response.New(p.Engine, c)

	if resp.CheckAccess(resource.BondExcel) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	params := param.Get(c, p.Engine, thisBonds)
	bonds, err := p.Service.Excel(params)
	if err != nil {
		resp.Status(http.StatusNotFound).Error(term.Record_Not_Found).JSON()
		return
	}

	resp.Record(event.BondExcel)

	ex := excel.New("bond")
	ex.AddSheet("Bonds").
		AddSheet("Summary").
		Active("Bonds").
		SetPageLayout("landscape", "A4").
		SetPageMargins(0.2).
		SetHeaderFooter().
		SetColWidth("B", "C", 15.3).
		SetColWidth("M", "M", 20).
		Active("Summary").
		SetColWidth("A", "D", 20).
		Active("Bonds").
		WriteHeader("ID", "Name", "Legal Name", "Server Address", "Expiration", "Plan",
			"Detail", "Phone", "Email", "Website", "Type", "Code", "Updated At").
		SetSheetFields("ID", "Name", "LegalName", "ServerAddress", "Expiration", "Plan",
			"Detail", "Phone", "Email", "Website", "Type", "Code", "UpdatedAt").
		WriteData(bonds).
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

// RegisterNode is used for communicate with cloud and if everyting fine create a bond
// based on returned data
func (p *BondAPI) RegisterNode(c *gin.Context) {
	resp := response.New(p.Engine, c)
	var err error
	var node model.Node

	if err = c.ShouldBindJSON(&node); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, err)
		return
	}

	var bond model.Bond
	if bond, err = p.Service.RegisterNode(node); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(event.NodeActivate, nil, bond)

	resp.Status(http.StatusOK).
		MessageT(term.V_created_successfully, thisBond).
		JSON(bond)

}
