package service

import (
	"fmt"
	"net/http"
	"sigmamono/internal/consts"
	"sigmamono/internal/core"
	"sigmamono/internal/enum/action"
	"sigmamono/internal/enum/nodestatus"
	"sigmamono/internal/param"
	"sigmamono/internal/term"
	"sigmamono/internal/types"
	"sigmamono/model"
	"sigmamono/repo"
)

// NodeServ for injecting auth repo
type NodeServ struct {
	Repo   repo.NodeRepo
	Engine *core.Engine
}

// ProvideNodeService for node is used in wire
func ProvideNodeService(p repo.NodeRepo) NodeServ {
	return NodeServ{Repo: p, Engine: p.Engine}
}

// FindByID for getting node by it's id
func (p *NodeServ) FindByID(id types.RowID) (node model.Node, err error) {
	node, err = p.Repo.FindByID(id)
	p.Engine.CheckInfo(err, fmt.Sprintf("Node with id %v", id))

	return
}

// List of nodes, it support pagination and search and return back count
func (p *NodeServ) List(params param.Param) (data map[string]interface{}, err error) {

	data = make(map[string]interface{})

	if data["list"], err = p.Repo.List(params); err != nil {
		p.Engine.CheckError(err, "nodes list", params)
		return
	}

	data["count"], err = p.Repo.Count(params)
	p.Engine.CheckError(err, "nodes count", params)

	return
}

// Save node
func (p *NodeServ) Save(node model.Node, params param.Param) (savedNode model.Node, err error) {
	node.CompanyID = params.CompanyID

	if node.ID > 0 {
		savedNode, err = p.update(node)
	} else {
		savedNode, err = p.create(node)
	}

	return
}

func (p *NodeServ) create(node model.Node) (result model.Node, err error) {

	lastCode, err := p.LastCode(node.CompanyID)
	if err != nil {
		p.Engine.CheckError(err, "error in returning lastCode for node")
		return
	}

	node.Code = lastCode + 1

	if node.Code < consts.MinNodeCode {
		node.Code = consts.MinNodeCode
	}

	if err = node.Validate(action.Save); err != nil {
		p.Engine.CheckInfo(err, "validation failed for saving node", node)
		return
	}

	result, err = p.Repo.Create(node)
	p.Engine.CheckError(err, "node not created", node)

	return
}

func (p *NodeServ) update(node model.Node) (result model.Node, err error) {

	if err = node.Validate(action.Save); err != nil {
		p.Engine.CheckInfo(err, "validation failed for saving node", node)
		return
	}

	result, err = p.Repo.Update(node)
	p.Engine.CheckError(err, "node not updated", node)

	return
}

// LastCode of nodes table
func (p *NodeServ) LastCode(companyID types.RowID) (lastCode uint64, err error) {
	node, err := p.Repo.LastNode(companyID)
	if node.Code < consts.MinNodeCode {
		node.Code = consts.MinNodeCode
	}
	lastCode = node.Code
	return
}

// Delete node, it is soft delete
func (p *NodeServ) Delete(nodeID types.RowID, params param.Param) (node model.Node, err error) {

	if node, err = p.FindByID(nodeID); err != nil {
		return node, core.NewErrorWithStatus(err.Error(), http.StatusNotFound)
	}

	// rename unique key to prevent duplication
	node.Name = fmt.Sprintf("%v#%v", node.Name, node.ID)
	_, err = p.Save(node, params)
	p.Engine.CheckError(err, "rename node's name for prevent duplication")

	err = p.Repo.Delete(node)
	return
}

// HardDelete will delete the node permanently
func (p *NodeServ) HardDelete(nodeID types.RowID) error {
	node, err := p.FindByID(nodeID)
	if err != nil {
		return core.NewErrorWithStatus(err.Error(), http.StatusNotFound)
	}

	return p.Repo.HardDelete(node)
}

// Excel is used for export excel file
func (p *NodeServ) Excel(params param.Param) (nodes []model.Node, err error) {
	params.Limit = p.Engine.Env.Setting.ExcelMaxRows
	params.Offset = 0
	params.Order = "nodes.id ASC"

	nodes, err = p.Repo.List(params)
	p.Engine.CheckError(err, "nodes excel")

	return
}

// Activate is used when a nodeapp send activation request to the cloud
func (p *NodeServ) Activate(node model.Node) (bond model.Bond, err error) {

	// companyKeyTmp := connector.New().
	// 	Domain(domains.Central).
	// 	Entity("Bond").
	// 	Method("ParseCompanyKey").
	// 	Args(node.Extra["company_key"]).
	// 	SendReceive(p.Engine)

	var companyKey model.CompanyKey
	// var ok bool
	// if companyKey, ok = companyKeyTmp.(model.CompanyKey); !ok {
	// 	return
	// }

	if companyKey.Error != nil {
		p.Engine.CheckError(companyKey.Error, "error in parsing company-key")
		err = fmt.Errorf(term.Company_key_is_not_valid)
		return
	}

	nodeToSave := model.Node{
		CompanyID: companyKey.CompanyID,
		Type:      node.Type,
		Name:      node.Name,
		MachineID: node.Extra["machine_id"].(string),
		Status:    nodestatus.Inactive,
		Phone:     node.Phone,
	}

	var createdNode model.Node
	if createdNode, err = p.create(nodeToSave); err != nil {
		return
	}

	bond.CompanyID = companyKey.CompanyID
	bond.CompanyName = companyKey.CompanyName
	bond.NodeCode = createdNode.Code
	bond.NodeName = createdNode.Name
	bond.Key = node.Extra["company_key"].(string)
	bond.MachineID = node.MachineID

	return
}
