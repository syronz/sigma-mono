package service

import (
	"fmt"
	"net/http"
	"sigmamono/internal/core"
	"sigmamono/internal/enum/action"
	"sigmamono/internal/param"
	"sigmamono/internal/types"
	"sigmamono/model"
	"sigmamono/repo"
)

// VersionServ for injecting auth repo
type VersionServ struct {
	Repo   repo.VersionRepo
	Engine *core.Engine
}

// ProvideVersionService for version is used in wire
func ProvideVersionService(p repo.VersionRepo) VersionServ {
	return VersionServ{Repo: p, Engine: p.Engine}
}

// FindByID for getting version by it's id
func (p *VersionServ) FindByID(id types.RowID) (version model.Version, err error) {
	version, err = p.Repo.FindByID(id)
	p.Engine.CheckInfo(err, fmt.Sprintf("Version with id %v", id))

	return
}

// List of versions, it support pagination and search and return back count
func (p *VersionServ) List(params param.Param) (data map[string]interface{}, err error) {

	data = make(map[string]interface{})

	if data["list"], err = p.Repo.List(params); err != nil {
		p.Engine.CheckError(err, "versions list")
		return
	}

	data["count"], err = p.Repo.Count(params)
	p.Engine.CheckError(err, "versions count")

	return
}

// Save version
func (p *VersionServ) Save(version model.Version) (createdVersion model.Version, err error) {

	if version.ID == 0 {
		version.ID, _ = p.LastID()
		version.ID++
	}

	if err = version.Validate(action.Save); err != nil {
		return
	}

	createdVersion, err = p.Repo.Save(version)
	p.Engine.CheckInfo(err, "version not saved", version)

	return
}

// LastID of versions table
func (p *VersionServ) LastID() (lastID types.RowID, err error) {
	version, err := p.Repo.LastVersion()
	lastID = version.ID
	return
}

// Delete version, it is soft delete
func (p *VersionServ) Delete(versionID types.RowID) (version model.Version, err error) {

	if version, err = p.FindByID(versionID); err != nil {
		return version, core.NewErrorWithStatus(err.Error(), http.StatusNotFound)
	}

	// rename unique key to prevent duplication
	version.Name = fmt.Sprintf("%v#%v", version.Name, version.ID)
	_, err = p.Save(version)
	p.Engine.CheckError(err, "rename version's name for prevent duplication")

	err = p.Repo.Delete(version)
	return
}

// Excel is used for export excel file
func (p *VersionServ) Excel(params param.Param) (versions []model.Version, err error) {
	params.Limit = p.Engine.Env.Setting.ExcelMaxRows
	params.Offset = 0
	params.Order = "versions.id ASC"

	versions, err = p.Repo.List(params)
	p.Engine.CheckError(err, "versions excel")

	return
}
