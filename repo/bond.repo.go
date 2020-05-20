package repo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"sigmamono/internal/core"
	"sigmamono/internal/param"
	"sigmamono/internal/search"
	"sigmamono/internal/types"
	"sigmamono/model"
)

// BondRepo for injecting engine
type BondRepo struct {
	Engine *core.Engine
}

// ProvideBondRepo is used in wire
func ProvideBondRepo(engine *core.Engine) BondRepo {
	return BondRepo{Engine: engine}
}

// FindByID for bond
func (p *BondRepo) FindByID(id types.RowID) (bond model.Bond, err error) {
	err = p.Engine.DB.First(&bond, id.ToUint64()).Error
	return
}

// FindByCompanyID for bond
func (p *BondRepo) FindByCompanyID(companyID types.RowID) (bond model.Bond, err error) {
	err = p.Engine.DB.Where("company_id = ?", companyID.ToUint64()).First(&bond).Error
	return
}

// List of bonds
func (p *BondRepo) List(params param.Param) (bonds []model.Bond, err error) {

	columns, err := model.Bond{}.Columns(params.Select)
	if err != nil {
		return
	}

	err = p.Engine.DB.Select(columns).
		Where(search.Parse(params, model.Bond{}.Pattern())).
		Order(params.Order).
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&bonds).Error

	return
}

// Count of bonds
func (p *BondRepo) Count(params param.Param) (count uint64, err error) {
	err = p.Engine.DB.Table("bonds").
		Select(params.Select).
		Where("deleted_at is null").
		Where(search.Parse(params, model.Bond{}.Pattern())).
		Count(&count).Error
	return
}

// Update BondRepo
func (p *BondRepo) Update(bond model.Bond) (u model.Bond, err error) {
	err = p.Engine.DB.Save(&bond).Error
	p.Engine.DB.Where("id = ?", bond.ID).Find(&u)
	return
}

// Create BondRepo
func (p *BondRepo) Create(bond model.Bond) (u model.Bond, err error) {
	err = p.Engine.DB.Create(&bond).Scan(&u).Error
	return
}

// LastBond of bond table
func (p *BondRepo) LastBond() (bond model.Bond, err error) {
	err = p.Engine.DB.Unscoped().Last(&bond).Error
	return
}

// Delete bond
func (p *BondRepo) Delete(bond model.Bond) (err error) {
	err = p.Engine.DB.Delete(&bond).Error
	return
}

// HardDelete bond
func (p *BondRepo) HardDelete(bond model.Bond) (err error) {
	err = p.Engine.DB.Unscoped().Delete(&bond).Error
	return
}

// RegisterNode communicate with cloud server for register and activate the node
func (p *BondRepo) RegisterNode(node model.Node,
	serverAddress string) (bond model.Bond, err error) {

	var payloadBytes []byte
	if payloadBytes, err = json.Marshal(node); err != nil {
		p.Engine.CheckError(err, "error in marshal json")
		return
	}

	body := bytes.NewReader(payloadBytes)

	var req *http.Request
	if req, err = http.NewRequest("POST", serverAddress, body); err != nil {
		p.Engine.CheckError(err, "error in creating request")
		return
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	p.Engine.Debug(serverAddress)

	var resp *http.Response
	if resp, err = client.Do(req); err != nil {
		p.Engine.CheckError(err, "error in sending request")
		return
	}
	defer resp.Body.Close()

	var respBody []uint8
	if respBody, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}

	errorStruct := struct {
		Message string `json:"message"`
		Error   string `json:"error"`
	}{}

	if resp.StatusCode != 200 {
		if err = json.Unmarshal([]byte(respBody), &errorStruct); err != nil {
			return
		}

		err = errors.New(errorStruct.Message)
		return
	}

	respMap := struct {
		Data    model.Bond `json:"data"`
		Message string     `json:"message"`
	}{}
	if err = json.Unmarshal([]byte(respBody), &respMap); err != nil {
		return
	}

	bond = respMap.Data

	p.Engine.Debug("^^^^^^^^^^^^^^^^^:", bond)

	return
}
