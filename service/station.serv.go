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

// StationServ for injecting auth repo
type StationServ struct {
	Repo   repo.StationRepo
	Engine *core.Engine
}

// ProvideStationService for station is used in wire
func ProvideStationService(p repo.StationRepo) StationServ {
	return StationServ{Repo: p, Engine: p.Engine}
}

func (p *StationServ) setExtra(station *model.Station) (err error) {
	station.Extra = make(map[string]interface{})
	var companyKey model.CompanyKey
	if companyKey, err = p.ParseCompanyKey(station.Key); err != nil {
		p.Engine.CheckError(err, "error in setExtra when try to call ParseCompanyKey")
		return
	}
	station.Extra["company_key"] = companyKey

	return
}

// FindByID for getting station by it's id
func (p *StationServ) FindByID(id types.RowID) (station model.Station, err error) {
	station, err = p.Repo.FindByID(id)
	p.Engine.CheckError(err, fmt.Sprintf("Station with id %v", id))
	p.setExtra(&station)

	return
}

// FindByCompanyID for getting station by it's id
func (p *StationServ) FindByCompanyID(companyID types.RowID) (station model.Station, err error) {
	station, err = p.Repo.FindByCompanyID(companyID)
	p.Engine.CheckError(err, fmt.Sprintf("Station with companyID %v", companyID))
	p.setExtra(&station)

	return
}

// List of stations, it support pagination and search and return back count
func (p *StationServ) List(params param.Param) (data map[string]interface{}, err error) {

	data = make(map[string]interface{})

	data["stations"], err = p.Repo.List(params)
	if err != nil {
		return
	}
	p.Engine.CheckError(err, "stations list")

	data["count"], err = p.Repo.Count(params)
	p.Engine.CheckError(err, "stations count")

	return
}

// Save station
func (p *StationServ) Save(station model.Station) (savedStation model.Station, err error) {

	if station.ID == 0 {
		savedStation, err = p.create(station)
	} else {
		savedStation, err = p.update(station)
	}

	return
}

func (p *StationServ) create(station model.Station) (result model.Station, err error) {

	if err = station.Validate(action.Save); err != nil {
		p.Engine.CheckError(err, "validation failed")
		return
	}

	result, err = p.Repo.Create(station)
	p.Engine.CheckInfo(err, fmt.Sprintf("Failed in creating station for %+v", station))

	return
}

func (p *StationServ) update(station model.Station) (result model.Station, err error) {

	if err = station.Validate(action.Save); err != nil {
		p.Engine.CheckError(err, "validation failed")
		return
	}

	result, err = p.Repo.Update(station)
	p.Engine.CheckInfo(err, fmt.Sprintf("Failed in updating station for %+v", station))

	return
}

// LastID of stations table
func (p *StationServ) LastID() (lastID types.RowID, err error) {
	station, err := p.Repo.LastStation()
	lastID = station.ID
	return
}

// Delete station, it is soft delete
func (p *StationServ) Delete(stationID types.RowID) (station model.Station, err error) {
	if station, err = p.FindByID(stationID); err != nil {
		return station, core.NewErrorWithStatus(err.Error(), http.StatusNotFound)
	}

	// rename unique key to prevent duplication
	_, err = p.Save(station)
	p.Engine.CheckError(err, "rename station's legal name for prevent duplication")

	err = p.Repo.Delete(station)
	return
}

// HardDelete will delete the station permanently
func (p *StationServ) HardDelete(stationID types.RowID) error {
	station, err := p.FindByID(stationID)
	if err != nil {
		return core.NewErrorWithStatus(err.Error(), http.StatusNotFound)
	}

	return p.Repo.HardDelete(station)
}

// Excel is used for export excel file
func (p *StationServ) Excel(params param.Param) (stations []model.Station, err error) {
	params.Limit = p.Engine.Env.Setting.ExcelMaxRows
	params.Offset = 0
	params.Order = "id ASC"

	stations, err = p.Repo.List(params)
	p.Engine.CheckError(err, "stations excel")

	return
}

// ParseCompanyKey transfer an encrypted text to CompanyKey struct
func (p *StationServ) ParseCompanyKey(key string) (companyKey model.CompanyKey, err error) {

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
func (p *StationServ) GetCompanyKey(companyID types.RowID) (key model.CompanyKey, err error) {
	// var station model.Station
	// if station, err = p.FindByCompanyID(companyID); err != nil {
	// 	return
	// }

	// var companyKeyJSON []byte
	// companyKeyJSON , err

	return
}

// RegisterNode try to parse compnay_code and find server address, if everyting fine save
// the station
func (p *StationServ) RegisterNode(node model.Node) (station model.Station, err error) {
	var companyKey model.CompanyKey
	if companyKey, err = p.ParseCompanyKey(node.Extra["company_key"].(string)); err != nil {
		p.Engine.CheckError(err, "Error in parsing company_key in RegisterNode")
		return
	}

	node.Extra["machine_id"] = p.Engine.Env.MachineID
	node.CompanyID = companyKey.CompanyID

	p.Engine.Debug(node)

	if station, err = p.Repo.RegisterNode(node,
		companyKey.ServerAddress+"/activate/nodeapp"); err != nil {
		p.Engine.CheckError(err, "Error in sending acitvation request for node")
		return
	}

	if station, err = p.Save(station); err != nil {
		return
	}

	p.Engine.Debug(station, err)

	return
}
