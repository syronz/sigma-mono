package service

import (
	"fmt"
	"net/http"
	"sigmamono/internal/core"
	"sigmamono/internal/enum/accountdirection"
	"sigmamono/internal/enum/action"
	"sigmamono/internal/param"
	"sigmamono/internal/term"
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
	if user, err = p.Repo.FindByID(id); err != nil {
		p.Engine.CheckError(err, fmt.Sprintf("User with id %v", id))
		return
	}

	user.Account, err = p.getAccount(id)

	return
}

func (p *UserServ) getAccount(id types.RowID) (account model.Account, err error) {
	accountServ := ProvideAccountService(repo.ProvideAccountRepo(p.Engine))

	if account, err = accountServ.FindByID(id); err != nil {
		err = core.NewErrorWithStatus(err.Error(), http.StatusInternalServerError).SetMsg(term.Error_in_getting_user)
		p.Engine.ServerLog.Error(err.Error())
		return
	}

	return
}

// FindByUsername find user with username
func (p *UserServ) FindByUsername(username string) (user model.User, err error) {
	user, err = p.Repo.FindByUsername(username)
	p.Engine.CheckError(err, fmt.Sprintf("User with username %v", username))

	user.Account, err = p.getAccount(user.ID)

	return
}

// List of users, it support pagination and search and return back count
func (p *UserServ) List(params param.Param) (data map[string]interface{}, err error) {

	data = make(map[string]interface{})

	p.Engine.Debug(params)

	data["list"], err = p.Repo.List(params)
	p.Engine.CheckError(err, "users list")
	if err != nil {
		return
	}

	data["count"], err = p.Repo.Count(params)
	p.Engine.CheckError(err, "users count")

	return
}

// Create user
func (p *UserServ) Create(user model.User,
	params param.Param) (createdUser model.User, err error) {

	if err = user.Validate(action.Create); err != nil {
		p.Engine.CheckError(err, "Failed in validation")
		return
	}

	original := p.Engine.DB
	tx := p.Engine.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	p.Engine.DB = tx

	user.Account.Direction = accountdirection.Direct

	accountServ := ProvideAccountService(repo.ProvideAccountRepo(p.Engine))
	if user.Account, err = accountServ.Create(user.Account, params); err != nil {
		p.Engine.DB = original
		tx.Rollback()
		return
	}
	user.ID = user.Account.ID

	user.Password, err = password.Hash(user.Password, p.Engine.Env.Setting.PasswordSalt)
	p.Engine.CheckError(err, fmt.Sprintf("Hashing password failed for %+v", user))

	if createdUser, err = p.Repo.Create(user); err != nil {
		tx.Rollback()
		p.Engine.CheckInfo(err, fmt.Sprintf("Failed in saving user for %+v", user))
	}
	tx.Commit()
	p.Engine.DB = original

	createdUser.Password = ""
	createdUser.Account = user.Account

	return
}

// Save user
func (p *UserServ) Save(user model.User) (createdUser model.User, err error) {

	if user.ID > 0 {
		user.Account.ID = user.ID
		if err = user.Validate(action.Update); err != nil {
			p.Engine.Debug(err)
			return
		}
	} else {
		if err = user.Validate(action.Create); err != nil {
			p.Engine.Debug(err)
			return
		}
	}

	original := p.Engine.DB
	tx := p.Engine.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	p.Engine.DB = tx

	accountServ := ProvideAccountService(repo.ProvideAccountRepo(p.Engine))
	if user.Account, err = accountServ.Save(user.Account); err != nil {
		p.Engine.DB = original
		tx.Rollback()
		return
	}

	user.ID = user.Account.ID

	user.Password, err = password.Hash(user.Password, p.Engine.Env.Setting.PasswordSalt)
	p.Engine.CheckError(err, fmt.Sprintf("Hashing password failed for %+v", user))

	if createdUser, err = p.Repo.Update(user); err != nil {
		tx.Rollback()
		p.Engine.CheckInfo(err, fmt.Sprintf("Failed in saving user for %+v", user))
	}
	tx.Commit()
	p.Engine.DB = original

	createdUser.Password = ""
	createdUser.Account = user.Account

	return
}

// Delete user, it is hard delete, by deleting account related to the user
func (p *UserServ) Delete(userID types.RowID, params param.Param) (user model.User, err error) {
	if user, err = p.FindByID(userID); err != nil {
		return user, core.NewErrorWithStatus(err.Error(), http.StatusNotFound)
	}

	accountServ := ProvideAccountService(repo.ProvideAccountRepo(p.Engine))
	if user.Account, err = accountServ.Delete(userID, params); err != nil {
		return
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
