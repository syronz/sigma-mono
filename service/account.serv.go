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

	"github.com/jinzhu/gorm"
)

// AccountServ for injecting auth repo
type AccountServ struct {
	Repo   repo.AccountRepo
	Engine *core.Engine
}

// ProvideAccountService for account is used in wire
func ProvideAccountService(p repo.AccountRepo) AccountServ {
	return AccountServ{Repo: p, Engine: p.Engine}
}

// FindByID for getting account by it's id
func (p *AccountServ) FindByID(id types.RowID) (account model.Account, err error) {
	account, err = p.Repo.FindByID(id)
	p.Engine.CheckError(err, fmt.Sprintf("Account with id %v", id))

	return
}

// List of accounts, it support pagination and search and return back count
func (p *AccountServ) List(params param.Param) (data map[string]interface{}, err error) {

	data = make(map[string]interface{})

	if data["list"], err = p.Repo.List(params); err != nil {
		p.Engine.CheckError(err, "accounts list")
		return
	}

	data["count"], err = p.Repo.Count(params)
	p.Engine.CheckError(err, "accounts count")

	return
}

// Create an account
func (p *AccountServ) Create(account model.Account,
	params param.Param) (createdAccount model.Account, err error) {

	var prefix, lastID types.RowID

	if prefix, err = params.PrefixID(); err != nil {
		return
	}

	if lastID, err = p.LastID(prefix); err != nil {
		p.Engine.Debug(lastID, err.Error())
		return
	}

	account.ID = lastID + 1

	account.CompanyID = params.CompanyID
	account.NodeCode = params.NodeCode

	if err = account.Validate(action.Save); err != nil {
		p.Engine.CheckError(err, term.Validation_failed)
		return
	}

	createdAccount, err = p.Repo.Create(account)

	p.Engine.CheckInfo(err, fmt.Sprintf("Failed in creating account for %+v", account))

	return

}

// Save account, if it is exist update it otherwise create it
func (p *AccountServ) Save(account model.Account) (savedAccount model.Account, err error) {

	account.CompanyID, account.NodeCode, _ = account.ID.Split()

	if err = account.Validate(action.Save); err != nil {
		p.Engine.CheckError(err, "validation failed", account)
		return
	}

	savedAccount, err = p.Repo.Save(account)
	p.Engine.CheckInfo(err, fmt.Sprintf("Failed in saving account for %+v", account))

	return
}

// LastID of accounts table according to companyID and nodeID
func (p *AccountServ) LastID(prefix types.RowID) (lastID types.RowID, err error) {
	account, err := p.Repo.LastAccount(prefix)

	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}

	lastID = account.ID
	if lastID < helper.PrefixMinID(prefix) {
		lastID = helper.PrefixMinID(prefix)
	}
	return
}

// Delete account, it is soft delete
func (p *AccountServ) Delete(accountID types.RowID, params param.Param) (account model.Account, err error) {

	if account, err = p.FindByID(accountID); err != nil {
		p.Engine.ServerLog.Error(err.Error())
		return
	}

	if account.CompanyID != params.CompanyID {
		err = core.NewErrorWithStatus(term.You_dont_have_permission_out_of_scope, http.StatusForbidden)
		p.Engine.ServerLog.Error(err.Error())
		return
	}

	if err = p.Repo.Delete(account); err != nil {
		p.Engine.ServerLog.Error(err.Error())
		return
		// return account, core.NewErrorWithStatus(err.Error(), http.StatusNotFound).
		// 	SetMsg(term.Error_in_deleting_account)
	}

	return
}

// HardDelete will delete the account permanently
func (p *AccountServ) HardDelete(accountID types.RowID) error {
	account, err := p.FindByID(accountID)
	if err != nil {
		return core.NewErrorWithStatus(err.Error(), http.StatusNotFound)
	}

	return p.Repo.HardDelete(account)
}

// Excel is used for export excel file
func (p *AccountServ) Excel(params param.Param) (accounts []model.Account, err error) {
	params.Limit = p.Engine.Env.Setting.ExcelMaxRows
	params.Offset = 0
	params.Order = "accounts.id ASC"

	accounts, err = p.Repo.List(params)
	p.Engine.CheckError(err, "accounts excel")

	return
}
