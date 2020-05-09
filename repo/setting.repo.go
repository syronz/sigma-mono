package repo

import (
	"sigmamono/internal/core"
	"sigmamono/internal/param"
	"sigmamono/internal/search"
	"sigmamono/internal/types"
	"sigmamono/model"
)

// SettingRepo for injecting engine
type SettingRepo struct {
	Engine *core.Engine
}

// ProvideSettingRepo is used in wire
func ProvideSettingRepo(engine *core.Engine) SettingRepo {
	return SettingRepo{Engine: engine}
}

// FindByID for setting
func (p *SettingRepo) FindByID(id types.RowID) (setting model.Setting, err error) {
	err = p.Engine.DB.First(&setting, id.ToUint64()).Error
	return
}

// FindByProperty for setting
func (p *SettingRepo) FindByProperty(property string) (setting model.Setting, err error) {
	err = p.Engine.DB.Where("property = ?", property).First(&setting).Error
	return
}

// List of settings
func (p *SettingRepo) List(params param.Param) (settings []model.Setting, err error) {
	columns, err := model.Setting{}.Columns(params.Select)
	if err != nil {
		return
	}

	err = p.Engine.DB.Table("settings").Select(columns).
		Where("settings.company_id = ?", params.CompanyID).
		Where(search.Parse(params, model.Setting{}.Pattern())).
		Order(params.Order).
		Limit(params.Limit).
		Offset(params.Offset).
		Scan(&settings).Error

	return
}

// Count of settings
func (p *SettingRepo) Count(params param.Param) (count uint64, err error) {
	err = p.Engine.DB.Table("settings").
		Select(params.Select).
		Where("settings.company_id = ?", params.CompanyID).
		Where(search.Parse(params, model.Setting{}.Pattern())).
		Count(&count).Error
	return
}

// Save SettingRepo
func (p *SettingRepo) Save(setting model.Setting) (u model.Setting, err error) {
	err = p.Engine.DB.Save(&setting).Error
	p.Engine.DB.Where("id = ?", setting.ID).Find(&u)
	return
}

// Update SettingRepo
func (p *SettingRepo) Update(setting model.Setting) (u model.Setting, err error) {
	setting.Property = ""
	setting.Type = ""
	setting.Description = ""
	err = p.Engine.DB.Model(&setting).Updates(&setting).Error
	p.Engine.DB.Where("id = ?", setting.ID).Find(&u)
	return
}

// Delete setting
func (p *SettingRepo) Delete(setting model.Setting) (err error) {
	err = p.Engine.DB.Delete(&setting).Error
	return
}
