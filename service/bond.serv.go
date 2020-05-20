package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sigmamono/internal/aes"
	"sigmamono/internal/core"
	"sigmamono/internal/enum/action"
	"sigmamono/internal/param"
	"sigmamono/internal/term"
	"sigmamono/internal/types"
	"sigmamono/model"
	"sigmamono/repo"
)

// BondServ for injecting auth repo
type BondServ struct {
	Repo   repo.BondRepo
	Engine *core.Engine
}

// ProvideBondService for bond is used in wire
func ProvideBondService(p repo.BondRepo) BondServ {
	return BondServ{Repo: p, Engine: p.Engine}
}

func (p *BondServ) setExtra(bond *model.Bond) (err error) {
	bond.Extra = make(map[string]interface{})
	var companyKey model.CompanyKey
	if companyKey, err = p.ParseCompanyKey(bond.Key); err != nil {
		p.Engine.CheckError(err, "error in setExtra when try to call ParseCompanyKey")
		return
	}
	bond.Extra["company_key"] = companyKey

	return
}

// FindByID for getting bond by it's id
func (p *BondServ) FindByID(id types.RowID) (bond model.Bond, err error) {
	bond, err = p.Repo.FindByID(id)
	p.Engine.CheckError(err, fmt.Sprintf("Bond with id %v", id))
	p.setExtra(&bond)

	return
}

// FindByCompanyID for getting bond by it's id
func (p *BondServ) FindByCompanyID(companyID types.RowID) (bond model.Bond, err error) {
	bond, err = p.Repo.FindByCompanyID(companyID)
	p.Engine.CheckError(err, fmt.Sprintf("Bond with companyID %v", companyID))
	p.setExtra(&bond)

	return
}

// List of bonds, it support pagination and search and return back count
func (p *BondServ) List(params param.Param) (data map[string]interface{}, err error) {

	data = make(map[string]interface{})

	data["bonds"], err = p.Repo.List(params)
	if err != nil {
		return
	}
	p.Engine.CheckError(err, "bonds list")

	data["count"], err = p.Repo.Count(params)
	p.Engine.CheckError(err, "bonds count")

	return
}

// Save bond
func (p *BondServ) Save(bond model.Bond) (savedBond model.Bond, err error) {

	if bond.ID == 0 {
		savedBond, err = p.create(bond)
	} else {
		savedBond, err = p.update(bond)
	}

	return
}

func (p *BondServ) create(bond model.Bond) (result model.Bond, err error) {

	if err = bond.Validate(action.Save); err != nil {
		p.Engine.CheckError(err, "validation failed")
		return
	}

	result, err = p.Repo.Create(bond)
	p.Engine.CheckInfo(err, fmt.Sprintf("Failed in creating bond for %+v", bond))

	return
}

func (p *BondServ) update(bond model.Bond) (result model.Bond, err error) {

	if err = bond.Validate(action.Save); err != nil {
		p.Engine.CheckError(err, "validation failed")
		return
	}

	result, err = p.Repo.Update(bond)
	p.Engine.CheckInfo(err, fmt.Sprintf("Failed in updating bond for %+v", bond))

	return
}

// LastID of bonds table
func (p *BondServ) LastID() (lastID types.RowID, err error) {
	bond, err := p.Repo.LastBond()
	lastID = bond.ID
	return
}

// Delete bond, it is soft delete
func (p *BondServ) Delete(bondID types.RowID) (bond model.Bond, err error) {
	if bond, err = p.FindByID(bondID); err != nil {
		return bond, core.NewErrorWithStatus(err.Error(), http.StatusNotFound)
	}

	// rename unique key to prevent duplication
	_, err = p.Save(bond)
	p.Engine.CheckError(err, "rename bond's legal name for prevent duplication")

	err = p.Repo.Delete(bond)
	return
}

// HardDelete will delete the bond permanently
func (p *BondServ) HardDelete(bondID types.RowID) error {
	bond, err := p.FindByID(bondID)
	if err != nil {
		return core.NewErrorWithStatus(err.Error(), http.StatusNotFound)
	}

	return p.Repo.HardDelete(bond)
}

// Excel is used for export excel file
func (p *BondServ) Excel(params param.Param) (bonds []model.Bond, err error) {
	params.Limit = p.Engine.Env.Setting.ExcelMaxRows
	params.Offset = 0
	params.Order = "id ASC"

	bonds, err = p.Repo.List(params)
	p.Engine.CheckError(err, "bonds excel")

	return
}

// ParseCompanyKey transfer an encrypted text to CompanyKey struct
func (p *BondServ) ParseCompanyKey(key string) (companyKey model.CompanyKey, err error) {

	var companyKeyDeprycated string
	if companyKeyDeprycated, err = aes.DecryptTwice(key); err != nil {
		return
	}

	if err = json.Unmarshal([]byte(companyKeyDeprycated), &companyKey); err != nil {
		return
	}

	if companyKey.CompanyID == 0 {
		err = fmt.Errorf(term.Company_key_is_not_valid)
	}

	return
}

// GetCompanyKey decrypt company-key and return back info
func (p *BondServ) GetCompanyKey(companyID types.RowID) (key model.CompanyKey, err error) {
	// var bond model.Bond
	// if bond, err = p.FindByCompanyID(companyID); err != nil {
	// 	return
	// }

	// var companyKeyJSON []byte
	// companyKeyJSON , err

	return
}

// RegisterNode try to parse compnay_code and find server address, if everyting fine save
// the bond
func (p *BondServ) RegisterNode(node model.Node) (bond model.Bond, err error) {
	var companyKey model.CompanyKey
	if companyKey, err = p.ParseCompanyKey(node.Extra["company_key"].(string)); err != nil {
		p.Engine.CheckError(err, "Error in parsing company_key in RegisterNode")
		return
	}

	node.Extra["machine_id"] = p.Engine.Env.MachineID
	node.CompanyID = companyKey.CompanyID

	p.Engine.Debug(node)

	if bond, err = p.Repo.RegisterNode(node,
		companyKey.ServerAddress+"/activate/nodeapp"); err != nil {
		p.Engine.CheckError(err, "Error in sending acitvation request for node")
		return
	}

	if bond, err = p.Save(bond); err != nil {
		return
	}

	p.Engine.Debug(bond, err)

	return
}
