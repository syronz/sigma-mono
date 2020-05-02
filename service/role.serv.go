package service

import (
	"fmt"
	"net/http"
	"radiusbilling/internal/core"
	"radiusbilling/internal/enum/action"
	"radiusbilling/internal/param"
	"radiusbilling/internal/term"
	"radiusbilling/model"
	"radiusbilling/repo"
)

// RoleServ for injecting auth repo
type RoleServ struct {
	Repo   repo.RoleRepo
	Engine *core.Engine
}

// ProvideRoleService for role is used in wire
func ProvideRoleService(p repo.RoleRepo) RoleServ {
	return RoleServ{Repo: p, Engine: p.Engine}
}

// FindByID for getting role by it's id
func (p *RoleServ) FindByID(id uint64) (role model.Role, err error) {
	role, err = p.Repo.FindByID(id)
	p.Engine.CheckError(err, fmt.Sprintf("Role with id %v", id))

	return
}

// List of roles, it support pagination and search and return back count
func (p *RoleServ) List(params param.Param) (data map[string]interface{}, err error) {

	data = make(map[string]interface{})

	data["list"], err = p.Repo.List(params)
	p.Engine.CheckError(err, "roles list")
	if err != nil {
		return
	}

	data["count"], err = p.Repo.Count(params)
	p.Engine.CheckError(err, "roles count")

	return
}

// Create a role
func (p *RoleServ) Create(role model.Role, params param.Param) (createdRole model.Role, err error) {

	if err = role.Validate(action.Save); err != nil {
		p.Engine.CheckError(err, term.Validation_failed)
		return
	}

	createdRole, err = p.Repo.Create(role)

	p.Engine.CheckInfo(err, fmt.Sprintf("Failed in creating role for %+v", role))

	return
}

// Save a role, if it is exist update it, if not create it
func (p *RoleServ) Save(role model.Role) (savedRole model.Role, err error) {

	if err = role.Validate(action.Save); err != nil {
		p.Engine.CheckError(err, "validation failed")
		return
	}

	savedRole, err = p.Repo.Update(role)
	p.Engine.CheckInfo(err, fmt.Sprintf("Failed in updating role for %+v", role))

	return
}

// Delete role, it is soft delete
func (p *RoleServ) Delete(roleID uint64, params param.Param) (role model.Role, err error) {

	if role, err = p.FindByID(roleID); err != nil {
		return role, core.NewErrorWithStatus(err.Error(), http.StatusNotFound)
	}

	// rename unique key to prevent duplication
	role.Name = fmt.Sprintf("%v#%v", role.Name, role.ID)
	_, err = p.Save(role)
	p.Engine.CheckError(err, "rename role's name for prevent duplication")

	err = p.Repo.Delete(role)
	return
}

// HardDelete will delete the role permanently
func (p *RoleServ) HardDelete(roleID uint64) error {
	role, err := p.FindByID(roleID)
	if err != nil {
		return core.NewErrorWithStatus(err.Error(), http.StatusNotFound)
	}

	return p.Repo.HardDelete(role)
}

// Excel is used for export excel file
func (p *RoleServ) Excel(params param.Param) (roles []model.Role, err error) {
	params.Limit = p.Engine.Env.Setting.ExcelMaxRows
	params.Offset = 0
	params.Order = "roles.id ASC"

	roles, err = p.Repo.List(params)
	p.Engine.CheckError(err, "roles excel")

	return
}
