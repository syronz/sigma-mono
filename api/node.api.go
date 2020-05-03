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

const thisNode = "node"
const thisNodes = "nodes"

// NodeAPI for injecting node service
type NodeAPI struct {
	Service service.NodeServ
	Engine  *core.Engine
}

// ProvideNodeAPI for node is used in wire
func ProvideNodeAPI(c service.NodeServ) NodeAPI {
	return NodeAPI{Service: c, Engine: c.Engine}
}

// FindByID is used for fetch a node by it's id
func (p *NodeAPI) FindByID(c *gin.Context) {
	resp := response.New(p.Engine, c)
	var err error
	var node model.Node

	if resp.CheckAccess(resource.NodeRead) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	if node.ID, err = types.StrToRowID(c.Param("nodeID")); err != nil {
		resp.Error(term.Invalid_ID).JSON()
		return
	}

	if node, err = p.Service.FindByID(node.ID); err != nil {
		resp.Status(http.StatusNotFound).Error(err).MessageT(term.Record_Not_Found).JSON()
		return
	}

	p.Engine.Record(c, event.NodeView)
	resp.Status(http.StatusOK).
		MessageT(term.V_info, thisNode).
		JSON(node)
}

// List of nodes
func (p *NodeAPI) List(c *gin.Context) {
	resp := response.New(p.Engine, c)

	if resp.CheckAccess(resource.NodeWrite) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	params := param.Get(c, p.Engine, thisNodes)

	data, err := p.Service.List(params)
	p.Engine.Debug(err)
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	p.Engine.Record(c, event.NodeList)
	resp.Status(http.StatusOK).
		MessageT(term.List_of_V, thisNodes).
		JSON(data)
}

// Create node
func (p *NodeAPI) Create(c *gin.Context) {
	var node model.Node
	resp := response.New(p.Engine, c)

	if resp.CheckAccess(resource.NodeWrite) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	if err := c.ShouldBindJSON(&node); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, err)
		return
	}

	createdNode, err := p.Service.Save(node)
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	p.Engine.Record(c, event.NodeCreate, nil, node)

	resp.Status(http.StatusOK).
		MessageT(term.V_created_successfully, thisNode).
		JSON(createdNode)
}

// Update node
func (p *NodeAPI) Update(c *gin.Context) {
	resp := response.New(p.Engine, c)
	var err error

	if resp.CheckAccess(resource.NodeWrite) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	var node, nodeBefore, nodeUpdated model.Node

	node.ID, err = types.StrToRowID(c.Param("nodeID"))
	if err != nil {
		resp.Error(term.Invalid_ID).JSON()
		return
	}

	if err = c.ShouldBindJSON(&node); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, err)
		return
	}

	if nodeBefore, err = p.Service.FindByID(node.ID); err != nil {
		resp.Status(http.StatusNotFound).Error(term.Record_Not_Found).JSON()
		return
	}

	if nodeUpdated, err = p.Service.Save(node); err != nil {
		resp.Error(err).JSON()
		return
	}

	p.Engine.Record(c, event.NodeUpdate, nodeBefore, node)

	resp.Status(http.StatusOK).
		MessageT(term.V_updated_successfully, thisNode).
		JSON(nodeUpdated)
}

// Delete node
func (p *NodeAPI) Delete(c *gin.Context) {
	resp := response.New(p.Engine, c)
	var err error
	var node model.Node

	if resp.CheckAccess(resource.NodeWrite) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	if node.ID, err = types.StrToRowID(c.Param("nodeID")); err != nil {
		resp.Error(term.Invalid_ID).JSON()
		return
	}

	if node, err = p.Service.Delete(node.ID); err != nil {
		resp.Status(http.StatusInternalServerError).Error(err).JSON()
		return
	}

	p.Engine.Record(c, event.NodeDelete, node)
	resp.Status(http.StatusOK).
		MessageT(term.V_deleted_successfully, thisNode).
		JSON()
}

// Excel generate excel files based on search
func (p *NodeAPI) Excel(c *gin.Context) {
	resp := response.New(p.Engine, c)

	if resp.CheckAccess(resource.NodeExcel) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	params := param.Get(c, p.Engine, thisNodes)
	nodes, err := p.Service.Excel(params)
	if err != nil {
		resp.Status(http.StatusNotFound).Error(term.Record_Not_Found).JSON()
		return
	}

	ex := excel.New("node")
	ex.AddSheet("Nodes").
		AddSheet("Summary").
		Active("Nodes").
		SetPageLayout("landscape", "A4").
		SetPageMargins(0.2).
		SetHeaderFooter().
		SetColWidth("B", "C", 15.3).
		SetColWidth("M", "M", 20).
		Active("Summary").
		SetColWidth("A", "D", 20).
		Active("Nodes").
		WriteHeader("ID", "Name", "Legal Name", "Server Address", "Expiration", "Plan",
			"Detail", "Phone", "Email", "Website", "Type", "Code", "Updated At").
		SetSheetFields("ID", "Name", "LegalName", "ServerAddress", "Expiration", "Plan",
			"Detail", "Phone", "Email", "Website", "Type", "Code", "UpdatedAt").
		WriteData(nodes).
		AddTable()

	buffer, downloadName, err := ex.Generate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &response.Result{
			Message: "Error in generating Excel file",
		})
		return
	}

	p.Engine.Record(c, event.NodeExcel)

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+downloadName)
	c.Data(http.StatusOK, "application/octet-stream", buffer.Bytes())

}

// Activate nodes
func (p *NodeAPI) Activate(c *gin.Context) {
	resp := response.New(p.Engine, c)
	var err error
	var node model.Node

	if err = c.ShouldBindJSON(&node); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, err)
		return
	}

	var bond model.Bond
	if bond, err = p.Service.Activate(node); err != nil {
		resp.Error(err).JSON()
		return
	}

	p.Engine.Record(c, event.NodeActivate, nil, node)

	resp.Status(http.StatusOK).
		MessageT(term.V_created_successfully, thisNode).
		JSON(bond)
}
