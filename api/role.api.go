package api

import (
	"net/http"
	"radiusbilling/internal/core"
	"radiusbilling/internal/core/access/resource"
	"radiusbilling/internal/enum/event"
	"radiusbilling/internal/param"
	"radiusbilling/internal/response"
	"radiusbilling/internal/term"
	"radiusbilling/model"
	"radiusbilling/service"
	"radiusbilling/utils/excel"
	"strconv"

	"github.com/gin-gonic/gin"
)

const thisRole = "role"
const thisRoles = "roles"

// RoleAPI for injecting role service
type RoleAPI struct {
	Service service.RoleServ
	Engine  *core.Engine
}

// ProvideRoleAPI for role is used in wire
func ProvideRoleAPI(c service.RoleServ) RoleAPI {
	return RoleAPI{Service: c, Engine: c.Engine}
}

// FindByID is used for fetch a role by it's id
func (p *RoleAPI) FindByID(c *gin.Context) {
	resp := response.New(p.Engine, c)
	var err error
	var role model.Role

	if resp.CheckAccess(resource.RoleRead) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	if role.ID, err = strconv.ParseUint(c.Param("roleID"), 10, 64); err != nil {
		resp.Error(term.Invalid_ID).JSON()
		return
	}

	if role, err = p.Service.FindByID(role.ID); err != nil {
		resp.Status(http.StatusNotFound).Error(err).MessageT(term.Record_Not_Found).JSON()
		return
	}

	p.Engine.Record(c, event.RoleView)
	resp.Status(http.StatusOK).
		MessageT(term.V_info, thisRole).
		JSON(role)
}

// List of roles
func (p *RoleAPI) List(c *gin.Context) {
	resp := response.New(p.Engine, c)

	if resp.CheckAccess(resource.RoleWrite) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	params := param.Get(c, p.Engine, thisRoles)

	data, err := p.Service.List(params)
	p.Engine.Debug(err)
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	p.Engine.Record(c, event.RoleList)
	resp.Status(http.StatusOK).
		MessageT(term.List_of_V, thisRoles).
		JSON(data)
}

// Create role
func (p *RoleAPI) Create(c *gin.Context) {
	var role model.Role
	resp := response.New(p.Engine, c)

	if resp.CheckAccess(resource.RoleWrite) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	if err := c.ShouldBindJSON(&role); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, err)
		return
	}

	params := param.Get(c, p.Engine, thisRoles)

	createdRole, err := p.Service.Create(role, params)
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	p.Engine.Record(c, event.RoleCreate, nil, role)

	resp.Status(http.StatusOK).
		MessageT(term.V_created_successfully, thisRole).
		JSON(createdRole)
}

// Update role
func (p *RoleAPI) Update(c *gin.Context) {
	resp := response.New(p.Engine, c)
	var err error

	if resp.CheckAccess(resource.RoleWrite) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	var role, roleBefore, roleUpdated model.Role

	if role.ID, err = strconv.ParseUint(c.Param("roleID"), 10, 64); err != nil {
		resp.Error(term.Invalid_ID).JSON()
		return
	}

	if err = c.ShouldBindJSON(&role); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, err)
		return
	}

	if roleBefore, err = p.Service.FindByID(role.ID); err != nil {
		resp.Status(http.StatusNotFound).Error(term.Record_Not_Found).JSON()
		return
	}

	if roleUpdated, err = p.Service.Save(role); err != nil {
		resp.Error(err).JSON()
		return
	}

	p.Engine.Record(c, event.RoleUpdate, roleBefore, role)

	resp.Status(http.StatusOK).
		MessageT(term.V_updated_successfully, thisRole).
		JSON(roleUpdated)
}

// Delete role
func (p *RoleAPI) Delete(c *gin.Context) {
	resp := response.New(p.Engine, c)
	var err error
	var role model.Role

	if resp.CheckAccess(resource.RoleWrite) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	if role.ID, err = strconv.ParseUint(c.Param("roleID"), 10, 64); err != nil {
		resp.Error(term.Invalid_ID).JSON()
		return
	}

	params := param.Get(c, p.Engine, thisRoles)

	if role, err = p.Service.Delete(role.ID, params); err != nil {
		resp.Status(http.StatusInternalServerError).Error(err).JSON()
		return
	}

	p.Engine.Record(c, event.RoleDelete, role)
	resp.Status(http.StatusOK).
		MessageT(term.V_deleted_successfully, thisRole).
		JSON()
}

// Excel generate excel files based on search
func (p *RoleAPI) Excel(c *gin.Context) {
	resp := response.New(p.Engine, c)

	if resp.CheckAccess(resource.RoleExcel) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	params := param.Get(c, p.Engine, thisRoles)
	roles, err := p.Service.Excel(params)
	if err != nil {
		resp.Status(http.StatusNotFound).Error(term.Record_Not_Found).JSON()
		return
	}

	ex := excel.New("role")
	ex.AddSheet("Roles").
		AddSheet("Summary").
		Active("Roles").
		SetPageLayout("landscape", "A4").
		SetPageMargins(0.2).
		SetHeaderFooter().
		SetColWidth("B", "C", 15.3).
		SetColWidth("M", "M", 20).
		Active("Summary").
		SetColWidth("A", "D", 20).
		Active("Roles").
		WriteHeader("ID", "Name", "Legal Name", "Server Address", "Expiration", "Plan",
			"Detail", "Phone", "Email", "Website", "Type", "Code", "Updated At").
		SetSheetFields("ID", "Name", "LegalName", "ServerAddress", "Expiration", "Plan",
			"Detail", "Phone", "Email", "Website", "Type", "Code", "UpdatedAt").
		WriteData(roles).
		AddTable()

	buffer, downloadName, err := ex.Generate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &response.Result{
			Message: "Error in generating Excel file",
		})
		return
	}

	p.Engine.Record(c, event.RoleExcel)

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+downloadName)
	c.Data(http.StatusOK, "application/octet-stream", buffer.Bytes())

}
