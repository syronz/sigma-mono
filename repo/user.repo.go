package repo

import (
	"radiusbilling/dto"
	"radiusbilling/internal/core"
	"radiusbilling/internal/param"
	"radiusbilling/internal/search"
	"radiusbilling/internal/types"
	"radiusbilling/model"
)

// UserRepo for injecting engine
type UserRepo struct {
	Engine *core.Engine
}

// ProvideUserRepo is used in wire
func ProvideUserRepo(engine *core.Engine) UserRepo {
	return UserRepo{Engine: engine}
}

// FindByID for user
func (p *UserRepo) FindByID(id types.RowID) (user model.User, err error) {
	err = p.Engine.DB.First(&user, id.ToUint64()).Error
	return
}

// FindByUsername for user
func (p *UserRepo) FindByUsername(username string) (user model.User, err error) {
	err = p.Engine.DB.Where("username = ?", username).First(&user).Error
	return
}

// List of users
func (p *UserRepo) List(params param.Param) (users []model.User, err error) {
	columns, err := model.User{}.Columns(params.Select)
	if err != nil {
		return
	}

	var userDtos []dto.UserDto

	err = p.Engine.DB.Table("users").Select(columns).
		Joins("INNER JOIN accounts on accounts.id = users.id").
		Joins("INNER JOIN roles on roles.id = users.role_id").
		Where("accounts.deleted_at is null").
		Where(search.Parse(params, model.User{}.Pattern())).
		Order(params.Order).
		Limit(params.Limit).
		Offset(params.Offset).
		Scan(&userDtos).Error

	arr := make([]model.User, len(userDtos))
	for i, v := range userDtos {
		arr[i].ID = v.ID
		arr[i].Username = v.Username
		arr[i].RoleID = v.RoleID
		arr[i].Language = v.Language
		arr[i].Email = v.Email

		extra := make(map[string]interface{})
		// extra["role"] = v.Role
		arr[i].Extra = extra
	}

	users = arr

	p.Engine.Debug(userDtos)

	return
}

// Count of users
func (p *UserRepo) Count(params param.Param) (count uint64, err error) {
	err = p.Engine.DB.Table("users").
		Select(params.Select).
		Joins("INNER JOIN accounts on accounts.id = users.id").
		Joins("INNER JOIN roles on roles.id = users.role_id").
		Where("accounts.deleted_at is null").
		Where(search.Parse(params, model.User{}.Pattern())).
		Count(&count).Error
	return
}

// Update UserRepo
func (p *UserRepo) Update(user model.User) (u model.User, err error) {
	err = p.Engine.DB.Save(&user).Error
	p.Engine.DB.Where("id = ?", user.ID).Find(&u)
	return
}

// Create UserRepo
func (p *UserRepo) Create(user model.User) (u model.User, err error) {
	err = p.Engine.DB.Create(&user).Scan(&u).Error
	return
}

// Delete user
func (p *UserRepo) Delete(user model.User) (err error) {
	err = p.Engine.DB.Delete(&user).Error
	return
}
