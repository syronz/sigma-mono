package service

import (
	"fmt"
	"net/http"
	"sigmamono/internal/core"
	"sigmamono/internal/enum/action"
	"sigmamono/internal/param"
	// "sigmamono/internal/term"
	"sigmamono/internal/types"
	"sigmamono/model"
	"sigmamono/repo"
	"sigmamono/utils/password"
)

// UserServ for injecting auth repo
type UserServ struct {
	Repo   repo.UserRepo
	Engine *core.Engine
}

// ProvideUserService for user is used in wire
func ProvideUserService(p repo.UserRepo) UserServ {
	return UserServ{Repo: p, Engine: p.Engine}
}

// FindByID for getting user by it's id
func (p *UserServ) FindByID(id types.RowID) (user model.User, err error) {
	user, err = p.Repo.FindByID(id)
	p.Engine.CheckError(err, fmt.Sprintf("User with id %v", id))

	// user.Account, err = p.getAccount(user.ID)

	return
}

// FindByUsername find user with username
func (p *UserServ) FindByUsername(username string) (user model.User, err error) {
	user, err = p.Repo.FindByUsername(username)
	p.Engine.CheckError(err, fmt.Sprintf("User with username %v", username))

	return
}

// List of users, it support pagination and search and return back count
func (p *UserServ) List(params param.Param) (data map[string]interface{}, err error) {

	data = make(map[string]interface{})

	data["list"], err = p.Repo.List(params)
	p.Engine.CheckError(err, "users list")
	if err != nil {
		return
	}

	data["count"], err = p.Repo.Count(params)
	p.Engine.CheckError(err, "users count")

	return
}

// Create a user
func (p *UserServ) Create(user model.User,
	params param.Param) (createdUser model.User, err error) {
	if err = user.Validate(action.Create); err != nil {
		p.Engine.Debug(err)
		return
	}

	user.Password, err = password.Hash(user.Password, p.Engine.Env.Setting.PasswordSalt)
	p.Engine.CheckError(err, fmt.Sprintf("Hashing password failed for %+v", user))

	createdUser, err = p.Repo.Create(user)
	p.Engine.CheckInfo(err, fmt.Sprintf("Failed in saving user for %+v", user))

	createdUser.Password = ""
	// createdUser.Account = account

	return
}

// Save user
func (p *UserServ) Save(user model.User) (createdUser model.User, err error) {

	if user.ID > 0 {
		// user.Account.ID = user.ID
		if err = user.Validate(action.Update); err != nil {
			p.Engine.Debug(err)
			return
		}
	} else {
		if err = user.Validate(action.Save); err != nil {
			p.Engine.Debug(err)
			return
		}
	}

	user.Password, err = password.Hash(user.Password, p.Engine.Env.Setting.PasswordSalt)
	p.Engine.CheckError(err, fmt.Sprintf("Hashing password failed for %+v", user))

	createdUser, err = p.Repo.Update(user)
	p.Engine.CheckInfo(err, fmt.Sprintf("Failed in saving user for %+v", user))

	createdUser.Password = ""
	// createdUser.Account = account

	return
}

// Delete user, it is soft delete
func (p *UserServ) Delete(userID types.RowID) (user model.User, err error) {
	if user, err = p.FindByID(userID); err != nil {
		return user, core.NewErrorWithStatus(err.Error(), http.StatusNotFound)
	}

	return
}

// Excel is used for export excel file
func (p *UserServ) Excel(params param.Param) (users []model.User, err error) {
	params.Limit = p.Engine.Env.Setting.ExcelMaxRows
	params.Offset = 0
	params.Order = "users.id ASC"

	users, err = p.Repo.List(params)
	p.Engine.CheckError(err, "users excel")

	return
}
