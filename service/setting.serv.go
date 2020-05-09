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

// SettingServ for injecting auth repo
type SettingServ struct {
	Repo   repo.SettingRepo
	Engine *core.Engine
}

// ProvideSettingService for setting is used in wire
func ProvideSettingService(p repo.SettingRepo) SettingServ {
	return SettingServ{Repo: p, Engine: p.Engine}
}

// FindByID for getting setting by it's id
func (p *SettingServ) FindByID(id types.RowID) (setting model.Setting, err error) {
	setting, err = p.Repo.FindByID(id)
	p.Engine.CheckInfo(err, fmt.Sprintf("Setting with id %v", id))

	return
}

// FindByProperty find setting with property
func (p *SettingServ) FindByProperty(property string) (setting model.Setting, err error) {
	setting, err = p.Repo.FindByProperty(property)
	p.Engine.CheckError(err, fmt.Sprintf("Setting with property %v", property))

	return
}

// List returns setting's property, it support pagination and search and return back count
func (p *SettingServ) List(params param.Param) (data map[string]interface{}, err error) {

	data = make(map[string]interface{})

	data["list"], err = p.Repo.List(params)
	p.Engine.CheckError(err, "settings list")
	if err != nil {
		return
	}

	data["count"], err = p.Repo.Count(params)
	p.Engine.CheckError(err, "settings count")

	return
}

// Save setting
func (p *SettingServ) Save(setting model.Setting) (savedSetting model.Setting, err error) {
	if err = setting.Validate(action.Save); err != nil {
		p.Engine.Debug(err)
		return
	}

	savedSetting, err = p.Repo.Save(setting)

	return
}

// Update setting
func (p *SettingServ) Update(setting model.Setting) (savedSetting model.Setting, err error) {
	savedSetting, err = p.Repo.Update(setting)

	return
}

// Delete setting, it is soft delete
func (p *SettingServ) Delete(settingID types.RowID) (setting model.Setting, err error) {
	if setting, err = p.FindByID(settingID); err != nil {
		return setting, core.NewErrorWithStatus(err.Error(), http.StatusNotFound)
	}

	return
}

// Excel is used for export excel file
func (p *SettingServ) Excel(params param.Param) (settings []model.Setting, err error) {
	params.Limit = p.Engine.Env.Setting.ExcelMaxRows
	params.Offset = 0
	params.Order = "settings.id ASC"

	settings, err = p.Repo.List(params)
	p.Engine.CheckError(err, "settings excel")

	return
}
