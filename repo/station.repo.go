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

// StationRepo for injecting engine
type StationRepo struct {
	Engine *core.Engine
}

// ProvideStationRepo is used in wire
func ProvideStationRepo(engine *core.Engine) StationRepo {
	return StationRepo{Engine: engine}
}

// FindByID for station
func (p *StationRepo) FindByID(id types.RowID) (station model.Station, err error) {
	err = p.Engine.DB.First(&station, id.ToUint64()).Error
	return
}

// FindByCompanyID for station
func (p *StationRepo) FindByCompanyID(companyID types.RowID) (station model.Station, err error) {
	err = p.Engine.DB.Where("company_id = ?", companyID.ToUint64()).First(&station).Error
	return
}

// List of stations
func (p *StationRepo) List(params param.Param) (stations []model.Station, err error) {

	columns, err := model.Station{}.Columns(params.Select)
	if err != nil {
		return
	}

	err = p.Engine.DB.Select(columns).
		Where(search.Parse(params, model.Station{}.Pattern())).
		Order(params.Order).
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&stations).Error

	return
}

// Count of stations
func (p *StationRepo) Count(params param.Param) (count uint64, err error) {
	err = p.Engine.DB.Table("stations").
		Select(params.Select).
		Where("deleted_at is null").
		Where(search.Parse(params, model.Station{}.Pattern())).
		Count(&count).Error
	return
}

// Update StationRepo
func (p *StationRepo) Update(station model.Station) (u model.Station, err error) {
	err = p.Engine.DB.Save(&station).Error
	p.Engine.DB.Where("id = ?", station.ID).Find(&u)
	return
}

// Create StationRepo
func (p *StationRepo) Create(station model.Station) (u model.Station, err error) {
	err = p.Engine.DB.Create(&station).Scan(&u).Error
	return
}

// LastStation of station table
func (p *StationRepo) LastStation() (station model.Station, err error) {
	err = p.Engine.DB.Unscoped().Last(&station).Error
	return
}

// Delete station
func (p *StationRepo) Delete(station model.Station) (err error) {
	err = p.Engine.DB.Delete(&station).Error
	return
}

// HardDelete station
func (p *StationRepo) HardDelete(station model.Station) (err error) {
	err = p.Engine.DB.Unscoped().Delete(&station).Error
	return
}

// RegisterNode communicate with cloud server for register and activate the node
func (p *StationRepo) RegisterNode(node model.Node,
	serverAddress string) (station model.Station, err error) {

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
		Data    model.Station `json:"data"`
		Message string        `json:"message"`
	}{}
	if err = json.Unmarshal([]byte(respBody), &respMap); err != nil {
		return
	}

	station = respMap.Data

	p.Engine.Debug("^^^^^^^^^^^^^^^^^:", station)

	return
}
