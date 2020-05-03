package repo

import (
	"fmt"
	"sigmamono/internal/core"
	"sigmamono/internal/param"
	"sigmamono/internal/search"
	"sigmamono/internal/types"
	"sigmamono/model"
)

// AccountRepo for injecting engine
type AccountRepo struct {
	Engine *core.Engine
}

// ProvideAccountRepo is used in wire
func ProvideAccountRepo(engine *core.Engine) AccountRepo {
	return AccountRepo{Engine: engine}
}

// FindByID for account
func (p *AccountRepo) FindByID(id types.RowID) (account model.Account, err error) {
	err = p.Engine.DB.First(&account, id.ToUint64()).Error
	return
}

// List of accounts
func (p *AccountRepo) List(params param.Param) (accounts []model.Account, err error) {
	columns, err := model.Account{}.Columns(params.Select)
	if err != nil {
		return
	}

	err = p.Engine.DB.Select(columns).
		Joins("INNER JOIN companies on companies.id = accounts.company_id").
		Where(search.Parse(params, model.Account{}.Pattern())).
		Order(params.Order).
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&accounts).Error

	return
}

// Count of accounts
func (p *AccountRepo) Count(params param.Param) (count uint64, err error) {
	err = p.Engine.DB.Table("accounts").
		Select(params.Select).
		Where("deleted_at is null").
		Where(search.Parse(params, model.Account{}.Pattern())).
		Count(&count).Error
	return
}

// Save AccountRepo
func (p *AccountRepo) Save(account model.Account) (u model.Account, err error) {
	err = p.Engine.DB.Save(&account).Error
	p.Engine.DB.Where("id = ?", account.ID).Find(&u)
	return
}

// Create AccountRepo
func (p *AccountRepo) Create(account model.Account) (u model.Account, err error) {
	err = p.Engine.DB.Create(&account).Scan(&u).Error
	return
}

// LastAccount of account table
func (p *AccountRepo) LastAccount(prefix types.RowID) (account model.Account, err error) {
	err = p.Engine.DB.Unscoped().Where("id LIKE ?", fmt.Sprintf("%v%%", prefix)).Last(&account).Error
	return
}

// Delete account
func (p *AccountRepo) Delete(account model.Account) (err error) {
	err = p.Engine.DB.Delete(&account).Error
	return
}

// HardDelete account
func (p *AccountRepo) HardDelete(account model.Account) (err error) {
	err = p.Engine.DB.Unscoped().Delete(&account).Error
	return
}
