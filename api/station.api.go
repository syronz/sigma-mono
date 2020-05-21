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

const thisStation = "station"
const thisStations = "stations"

// StationAPI for injecting station service
type StationAPI struct {
	Service service.StationServ
	Engine  *core.Engine
}

// ProvideStationAPI for station is used in wire
func ProvideStationAPI(c service.StationServ) StationAPI {
	return StationAPI{Service: c, Engine: c.Engine}
}

// FindByID is used for fetch a station by it's id
func (p *StationAPI) FindByID(c *gin.Context) {
	resp := response.New(p.Engine, c)
	var err error
	var station model.Station

	if resp.CheckAccess(resource.StationRead) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	if station.ID, err = types.StrToRowID(c.Param("stationID")); err != nil {
		resp.Status(http.StatusNotAcceptable).Error(err).MessageT(term.Invalid_ID).JSON()
		return
	}

	if station, err = p.Service.FindByID(station.ID); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(event.StationView)
	resp.Status(http.StatusOK).
		MessageT(term.V_info, thisStation).
		JSON(station)
}

// List of stations
func (p *StationAPI) List(c *gin.Context) {
	resp := response.New(p.Engine, c)

	if resp.CheckAccess(resource.StationWrite) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	params := param.Get(c, p.Engine, thisStations)

	data, err := p.Service.List(params)
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(event.StationList)
	resp.Status(http.StatusOK).
		MessageT(term.List_of_V, thisStations).
		JSON(data)
}

// Delete station
func (p *StationAPI) Delete(c *gin.Context) {
	resp := response.New(p.Engine, c)
	var err error
	var station model.Station

	if resp.CheckAccess(resource.StationWrite) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	if station.ID, err = types.StrToRowID(c.Param("stationID")); err != nil {
		resp.Error(term.Invalid_ID).JSON()
		return
	}

	if station, err = p.Service.Delete(station.ID); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(event.StationDelete, station)
	resp.Status(http.StatusOK).
		MessageT(term.V_deleted_successfully, thisStation).
		JSON()
}

// Excel generate excel files based on search
func (p *StationAPI) Excel(c *gin.Context) {
	resp := response.New(p.Engine, c)

	if resp.CheckAccess(resource.StationExcel) {
		resp.Status(http.StatusForbidden).Error(term.You_dont_have_permission).JSON()
		return
	}

	params := param.Get(c, p.Engine, thisStations)
	stations, err := p.Service.Excel(params)
	if err != nil {
		resp.Status(http.StatusNotFound).Error(term.Record_Not_Found).JSON()
		return
	}

	resp.Record(event.StationExcel)

	ex := excel.New("station")
	ex.AddSheet("Stations").
		AddSheet("Summary").
		Active("Stations").
		SetPageLayout("landscape", "A4").
		SetPageMargins(0.2).
		SetHeaderFooter().
		SetColWidth("B", "C", 15.3).
		SetColWidth("M", "M", 20).
		Active("Summary").
		SetColWidth("A", "D", 20).
		Active("Stations").
		WriteHeader("ID", "Name", "Legal Name", "Server Address", "Expiration", "Plan",
			"Detail", "Phone", "Email", "Website", "Type", "Code", "Updated At").
		SetSheetFields("ID", "Name", "LegalName", "ServerAddress", "Expiration", "Plan",
			"Detail", "Phone", "Email", "Website", "Type", "Code", "UpdatedAt").
		WriteData(stations).
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

// RegisterNode is used for communicate with cloud and if everyting fine create a station
// based on returned data
func (p *StationAPI) RegisterNode(c *gin.Context) {
	resp := response.New(p.Engine, c)
	var err error
	var node model.Node

	if err = c.ShouldBindJSON(&node); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, err)
		return
	}

	var station model.Station
	if station, err = p.Service.RegisterNode(node); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(event.NodeActivate, nil, station)

	resp.Status(http.StatusOK).
		MessageT(term.V_created_successfully, thisStation).
		JSON(station)

}
