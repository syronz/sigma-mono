package repo

import (
	"sigmamono/internal/core"
	"sigmamono/internal/param"
	"sigmamono/internal/search"
	"sigmamono/internal/types"
	"sigmamono/model"
)

// CompanyRepo for injecting engine
type CompanyRepo struct {
	Engine *core.Engine
}

// ProvideCompanyRepo is used in wire
func ProvideCompanyRepo(engine *core.Engine) CompanyRepo {
	return CompanyRepo{Engine: engine}
}

// FindByID for company
func (p *CompanyRepo) FindByID(id types.RowID) (company model.Company, err error) {
	err = p.Engine.DB.First(&company, id.ToUint64()).Error
	return
}

// List of companies
func (p *CompanyRepo) List(params param.Param) (companies []model.Company, err error) {

	columns, err := model.Company{}.Columns(params.Select)
	if err != nil {
		return
	}

	err = p.Engine.DB.Select(columns).
		Where(search.Parse(params, model.Company{}.Pattern())).
		Order(params.Order).
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&companies).Error

	return
}

// Count of companies
func (p *CompanyRepo) Count(params param.Param) (count uint64, err error) {
	err = p.Engine.DB.Table("companies").
		Select(params.Select).
		Where("deleted_at is null").
		Where(search.Parse(params, model.Company{}.Pattern())).
		Count(&count).Error
	return
}

// Update CompanyRepo
func (p *CompanyRepo) Update(company model.Company) (u model.Company, err error) {
	err = p.Engine.DB.Save(&company).Error
	p.Engine.DB.Where("id = ?", company.ID).Find(&u)
	return
}

// Create CompanyRepo
func (p *CompanyRepo) Create(company model.Company) (u model.Company, err error) {
	err = p.Engine.DB.Create(&company).Scan(&u).Error
	return
}

// LastCompany of company table
func (p *CompanyRepo) LastCompany() (company model.Company, err error) {
	err = p.Engine.DB.Unscoped().Last(&company).Error
	return
}

// Delete company
func (p *CompanyRepo) Delete(company model.Company) (err error) {
	err = p.Engine.DB.Delete(&company).Error
	return
}

// HardDelete company
func (p *CompanyRepo) HardDelete(company model.Company) (err error) {
	err = p.Engine.DB.Unscoped().Delete(&company).Error
	return
}
