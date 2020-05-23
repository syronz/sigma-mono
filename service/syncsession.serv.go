package service

import (
	"fmt"
	"net/http"
	"sigmamono/internal/core"
	"sigmamono/internal/enum/action"
	"sigmamono/internal/param"
	"sigmamono/internal/term"
	"sigmamono/internal/types"
	"sigmamono/model"
	"sigmamono/repo"
	"sigmamono/utils/helper"
	"time"
)

// SyncSessionServ for injecting auth repo
type SyncSessionServ struct {
	Repo   repo.SyncSessionRepo
	Engine *core.Engine
}

// ProvideSyncSessionService for syncSession is used in wire
func ProvideSyncSessionService(p repo.SyncSessionRepo) SyncSessionServ {
	return SyncSessionServ{Repo: p, Engine: p.Engine}
}

// FindByID for getting syncSession by it's id
func (p *SyncSessionServ) FindByID(id types.RowID) (syncSession model.SyncSession, err error) {
	syncSession, err = p.Repo.FindByID(id)
	p.Engine.CheckError(err, fmt.Sprintf("SyncSession with id %v", id))

	return
}

// DELETE BELOW
// List of syncSessions, it support pagination and search and return back count
func (p *SyncSessionServ) List(params param.Param) (data map[string]interface{}, err error) {

	data = make(map[string]interface{})

	data["list"], err = p.Repo.List(params)
	p.Engine.CheckError(err, "syncSessions list")
	if err != nil {
		return
	}

	data["count"], err = p.Repo.Count(params)
	p.Engine.CheckError(err, "syncSessions count")

	return
}

// Create a syncSession
func (p *SyncSessionServ) Create(syncSession model.SyncSession, params param.Param) (createdSyncSession model.SyncSession, err error) {

	var prefix, lastID types.RowID
	if prefix, err = params.PrefixID(); err != nil {
		p.Engine.CheckError(err, "company_id is not exist in JWT's token", params)
		return
	}

	if lastID, err = p.LastID(prefix); err != nil {
		p.Engine.CheckError(err, "error in getting the lastID")
		return
	}

	syncSession.ID = lastID + 1

	syncSession.CompanyID = params.CompanyID
	syncSession.NodeCode = params.NodeCode

	if err = syncSession.Validate(action.Save); err != nil {
		p.Engine.CheckError(err, term.Validation_failed)
		return
	}

	createdSyncSession, err = p.Repo.Create(syncSession)

	p.Engine.CheckInfo(err, fmt.Sprintf("Failed in creating syncSession for %+v", syncSession))

	return
}

// Save a syncSession, if it is exist update it, if not create it
func (p *SyncSessionServ) Save(syncSession model.SyncSession) (savedSyncSession model.SyncSession, err error) {

	syncSession.CompanyID, syncSession.NodeCode, _ = syncSession.ID.Split()

	if err = syncSession.Validate(action.Save); err != nil {
		p.Engine.CheckError(err, "validation failed")
		return
	}

	savedSyncSession, err = p.Repo.Update(syncSession)
	p.Engine.CheckInfo(err, fmt.Sprintf("Failed in updating syncSession for %+v", syncSession))

	return
}

// LastID of syncSessions table
func (p *SyncSessionServ) LastID(prefix types.RowID) (lastID types.RowID, err error) {
	syncSession, err := p.Repo.LastSyncSession(prefix)
	lastID = syncSession.ID
	if lastID < helper.PrefixMinID(prefix) {
		lastID = helper.PrefixMinID(prefix)
	}
	return
}

// Delete syncSession, it is soft delete
func (p *SyncSessionServ) Delete(syncSessionID types.RowID, params param.Param) (syncSession model.SyncSession, err error) {

	if syncSession, err = p.FindByID(syncSessionID); err != nil {
		return
	}

	if syncSession.CompanyID != params.CompanyID {
		return syncSession, core.NewErrorWithStatus(term.You_dont_have_permission_out_of_scope, http.StatusForbidden)
	}

	err = p.Repo.Delete(syncSession)
	return
}

// HardDelete will delete the syncSession permanently
func (p *SyncSessionServ) HardDelete(syncSessionID types.RowID) error {
	syncSession, err := p.FindByID(syncSessionID)
	if err != nil {
		return core.NewErrorWithStatus(err.Error(), http.StatusNotFound)
	}

	return p.Repo.HardDelete(syncSession)
}

// Excel is used for export excel file
func (p *SyncSessionServ) Excel(params param.Param) (syncSessions []model.SyncSession, err error) {
	params.Limit = p.Engine.Env.Setting.ExcelMaxRows
	params.Offset = 0
	params.Order = "syncSessions.id ASC"

	syncSessions, err = p.Repo.List(params)
	p.Engine.CheckError(err, "syncSessions excel")

	return
}

// Generate gob file for sync data
func (p *SyncSessionServ) Generate(companyID types.RowID, lastSyncDate time.Time) (err error) {
	fmt.Println("it is working fine")

	return
}
