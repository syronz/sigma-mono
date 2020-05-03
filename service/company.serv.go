package service

import (
	"fmt"
	"net/http"
	"sigmamono/internal/consts"
	"sigmamono/internal/core"
	"sigmamono/internal/enum/action"
	"sigmamono/internal/param"
	"sigmamono/internal/types"
	"sigmamono/model"
	"sigmamono/repo"
)

// CompanyServ for injecting auth repo
type CompanyServ struct {
	Repo   repo.CompanyRepo
	Engine *core.Engine
}

// ProvideCompanyService for company is used in wire
func ProvideCompanyService(p repo.CompanyRepo) CompanyServ {
	return CompanyServ{Repo: p, Engine: p.Engine}
}

// FindByID for getting company by it's id
func (p *CompanyServ) FindByID(id types.RowID) (company model.Company, err error) {
	company, err = p.Repo.FindByID(id)
	p.Engine.CheckError(err, fmt.Sprintf("Company with id %v", id))

	return
}

// List of companies, it support pagination and search and return back count
func (p *CompanyServ) List(params param.Param) (data map[string]interface{}, err error) {

	data = make(map[string]interface{})

	data["list"], err = p.Repo.List(params)
	if err != nil {
		return
	}
	p.Engine.CheckError(err, "companies list")

	data["count"], err = p.Repo.Count(params)
	p.Engine.CheckError(err, "companies count")

	return
}

// Save company
func (p *CompanyServ) Save(company model.Company) (savedCompany model.Company, err error) {

	if company.ID == 0 {
		company.ID, _ = p.LastID()
		company.ID++
		if company.ID < consts.MinCompanyID {
			company.ID = consts.MinCompanyID
		}

		savedCompany, err = p.create(company)
	} else {
		savedCompany, err = p.update(company)
	}

	return
}

func (p *CompanyServ) create(company model.Company) (result model.Company, err error) {

	if err = company.Validate(action.Save); err != nil {
		p.Engine.CheckError(err, "validation failed")
		return
	}

	result, err = p.Repo.Create(company)
	p.Engine.CheckInfo(err, fmt.Sprintf("Failed in creating company for %+v", company))

	return
}

func (p *CompanyServ) update(company model.Company) (result model.Company, err error) {

	if err = company.Validate(action.Save); err != nil {
		p.Engine.CheckError(err, "validation failed")
		return
	}

	result, err = p.Repo.Update(company)
	p.Engine.CheckInfo(err, fmt.Sprintf("Failed in updating company for %+v", company))

	return
}

// LastID of companies table
func (p *CompanyServ) LastID() (lastID types.RowID, err error) {
	company, err := p.Repo.LastCompany()
	lastID = company.ID
	return
}

// Delete company, it is soft delete
func (p *CompanyServ) Delete(companyID types.RowID) (company model.Company, err error) {
	if company, err = p.FindByID(companyID); err != nil {
		return company, core.NewErrorWithStatus(err.Error(), http.StatusNotFound)
	}

	// rename unique key to prevent duplication
	company.LegalName = fmt.Sprintf("%v#%v", company.LegalName, company.ID)
	_, err = p.Save(company)
	p.Engine.CheckError(err, "rename company's legal name for prevent duplication")

	err = p.Repo.Delete(company)
	return
}

// HardDelete will delete the company permanently
func (p *CompanyServ) HardDelete(companyID types.RowID) error {
	company, err := p.FindByID(companyID)
	if err != nil {
		return core.NewErrorWithStatus(err.Error(), http.StatusNotFound)
	}

	return p.Repo.HardDelete(company)
}

// Excel is used for export excel file
func (p *CompanyServ) Excel(params param.Param) (companies []model.Company, err error) {
	params.Limit = p.Engine.Env.Setting.ExcelMaxRows
	params.Offset = 0
	params.Order = "id ASC"

	companies, err = p.Repo.List(params)
	p.Engine.CheckError(err, "companies excel")

	return
}
