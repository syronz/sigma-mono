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

const thisCompany = "company"
const thisCompanies = "companies"

// CompanyAPI for injecting company service
type CompanyAPI struct {
	Service service.CompanyServ
	Engine  *core.Engine
}

// ProvideCompanyAPI for company is used in wire
func ProvideCompanyAPI(c service.CompanyServ) CompanyAPI {
	return CompanyAPI{Service: c, Engine: c.Engine}
}

// FindByID is used for fetch a company by it's id
func (p *CompanyAPI) FindByID(c *gin.Context) {
	resp := response.New(p.Engine, c)
	var err error
	var company model.Company

	if resp.CheckAccess(resource.CompanyRead) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	if company.ID, err = types.StrToRowID(c.Param("companyID")); err != nil {
		resp.Status(http.StatusNotAcceptable).Error(err).MessageT(term.Invalid_ID).JSON()
		// resp.Status(http.StatusNotAcceptable).Error(err).JSON()
		return
	}

	if company, err = p.Service.FindByID(company.ID); err != nil {
		resp.Error(err).JSON()
		return
	}

	p.Engine.Record(c, event.CompanyView)
	resp.Status(http.StatusOK).
		MessageT(term.V_info, thisCompany).
		JSON(company)
}

// List of companies
func (p *CompanyAPI) List(c *gin.Context) {
	resp := response.New(p.Engine, c)

	if resp.CheckAccess(resource.CompanyWrite) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	params := param.Get(c, p.Engine, thisCompanies)

	data, err := p.Service.List(params)
	if err != nil {
		resp.Status(http.StatusNotFound).Error(err).Message(term.Record_Not_Found).JSON()
		return
	}

	p.Engine.Record(c, event.CompanyList)
	resp.Status(http.StatusOK).
		MessageT(term.List_of_V, thisCompanies).
		JSON(data)
}

// Create company
func (p *CompanyAPI) Create(c *gin.Context) {
	var company model.Company
	resp := response.New(p.Engine, c)

	if resp.CheckAccess(resource.CompanyWrite) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	if err := c.ShouldBindJSON(&company); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, err)
		return
	}

	createdCompany, err := p.Service.Save(company)
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	p.Engine.Record(c, event.CompanyCreate, nil, company)

	resp.Status(http.StatusOK).
		MessageT(term.V_created_successfully, thisCompany).
		JSON(createdCompany)
}

// Update company
func (p *CompanyAPI) Update(c *gin.Context) {
	resp := response.New(p.Engine, c)
	var err error

	if resp.CheckAccess(resource.CompanyWrite) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	var company, companyBefore, companyUpdated model.Company

	company.ID, err = types.StrToRowID(c.Param("companyID"))

	if err != nil {
		resp.Error(term.Invalid_ID).JSON()
		return
	}

	if err = c.ShouldBindJSON(&company); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, err)
		return
	}

	if companyBefore, err = p.Service.FindByID(company.ID); err != nil {
		resp.Status(http.StatusNotFound).Error(term.Record_Not_Found).JSON()
		return
	}

	if companyUpdated, err = p.Service.Save(company); err != nil {
		resp.Error(err).JSON()
		return
	}

	p.Engine.Record(c, event.CompanyUpdate, companyBefore, company)

	resp.Status(http.StatusOK).
		MessageT(term.V_updated_successfully, thisCompany).
		JSON(companyUpdated)
}

// Delete company
func (p *CompanyAPI) Delete(c *gin.Context) {
	resp := response.New(p.Engine, c)
	var err error
	var company model.Company

	if resp.CheckAccess(resource.CompanyWrite) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	if company.ID, err = types.StrToRowID(c.Param("companyID")); err != nil {
		resp.Error(term.Invalid_ID).JSON()
		return
	}

	if company, err = p.Service.Delete(company.ID); err != nil {
		resp.Status(http.StatusInternalServerError).Error(err).JSON()
		return
	}

	p.Engine.Record(c, event.CompanyDelete, company)
	resp.Status(http.StatusOK).
		MessageT(term.V_deleted_successfully, thisCompany).
		JSON()
}

// Excel generate excel files based on search
func (p *CompanyAPI) Excel(c *gin.Context) {
	resp := response.New(p.Engine, c)

	if resp.CheckAccess(resource.CompanyExcel) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	params := param.Get(c, p.Engine, thisCompanies)
	companies, err := p.Service.Excel(params)
	if err != nil {
		resp.Status(http.StatusNotFound).Error(term.Record_Not_Found).JSON()
		return
	}

	p.Engine.Record(c, event.CompanyExcel)

	ex := excel.New("company")
	ex.AddSheet("Companies").
		AddSheet("Summary").
		Active("Companies").
		SetPageLayout("landscape", "A4").
		SetPageMargins(0.2).
		SetHeaderFooter().
		SetColWidth("B", "C", 15.3).
		SetColWidth("M", "M", 20).
		Active("Summary").
		SetColWidth("A", "D", 20).
		Active("Companies").
		WriteHeader("ID", "Name", "Legal Name", "Server Address", "Expiration", "Plan",
			"Detail", "Phone", "Email", "Website", "Type", "Code", "Updated At").
		SetSheetFields("ID", "Name", "LegalName", "ServerAddress", "Expiration", "Plan",
			"Detail", "Phone", "Email", "Website", "Type", "Code", "UpdatedAt").
		WriteData(companies).
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
