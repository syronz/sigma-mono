package repo

import (
	"fmt"
	"radiusbilling/internal/core"
	"radiusbilling/internal/param"
	"radiusbilling/internal/search"
	"radiusbilling/internal/types"
	"radiusbilling/model"

	"github.com/jinzhu/gorm"
)

// RoleRepo for injecting engine
type RoleRepo struct {
	Engine *core.Engine
}

// ProvideRoleRepo is used in wire
func ProvideRoleRepo(engine *core.Engine) RoleRepo {
	return RoleRepo{Engine: engine}
}

// FindByID for role
func (p *RoleRepo) FindByID(id uint64) (role model.Role, err error) {
	err = p.Engine.DB.First(&role, id).Error
	return
}

// List of roles
func (p *RoleRepo) List(params param.Param) (roles []model.Role, err error) {
	columns, err := model.Role{}.Columns(params.Select)
	if err != nil {
		return
	}

	err = p.Engine.DB.Select(columns).
		Joins("INNER JOIN companies on companies.id = roles.company_id").
		Where(search.Parse(params, model.Role{}.Pattern())).
		Order(params.Order).
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&roles).Error

	return
}

// Count of roles
func (p *RoleRepo) Count(params param.Param) (count uint64, err error) {
	err = p.Engine.DB.Table("roles").
		Select(params.Select).
		Where("deleted_at is null").
		Where(search.Parse(params, model.Role{}.Pattern())).
		Count(&count).Error
	return
}

// Update RoleRepo
func (p *RoleRepo) Update(role model.Role) (u model.Role, err error) {
	err = p.Engine.DB.Save(&role).Error
	p.Engine.DB.Where("id = ?", role.ID).Find(&u)
	return
}

// Create RoleRepo
func (p *RoleRepo) Create(role model.Role) (u model.Role, err error) {
	err = p.Engine.DB.Create(&role).Scan(&u).Error
	return
}

// LastRole of role table
func (p *RoleRepo) LastRole(prefix types.RowID) (role model.Role, err error) {
	err = p.Engine.DB.Unscoped().Where("id LIKE ?", fmt.Sprintf("%v%%", prefix)).
		Last(&role).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}
	return
}

// Delete role
func (p *RoleRepo) Delete(role model.Role) (err error) {
	err = p.Engine.DB.Delete(&role).Error
	return
}

// HardDelete role
func (p *RoleRepo) HardDelete(role model.Role) (err error) {
	err = p.Engine.DB.Unscoped().Delete(&role).Error
	return
}
