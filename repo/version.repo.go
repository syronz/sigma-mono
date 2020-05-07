package repo

import (
	"sigmamono/internal/core"
	"sigmamono/internal/param"
	"sigmamono/internal/search"
	"sigmamono/internal/types"
	"sigmamono/model"
)

// VersionRepo for injecting engine
type VersionRepo struct {
	Engine *core.Engine
}

// ProvideVersionRepo is used in wire
func ProvideVersionRepo(engine *core.Engine) VersionRepo {
	return VersionRepo{Engine: engine}
}

// FindByID for version
func (p *VersionRepo) FindByID(id types.RowID) (version model.Version, err error) {
	err = p.Engine.DB.First(&version, id.ToUint64()).Error
	return
}

// List of versions
func (p *VersionRepo) List(params param.Param) (versions []model.Version, err error) {
	columns, err := model.Version{}.Columns(params.Select)
	if err != nil {
		return
	}

	err = p.Engine.DB.Select(columns).
		Where(search.Parse(params, model.Version{}.Pattern())).
		Order(params.Order).
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&versions).Error

	return
}

// Count of versions
func (p *VersionRepo) Count(params param.Param) (count uint64, err error) {
	err = p.Engine.DB.Table("versions").
		Select(params.Select).
		Where("deleted_at is null").
		Where(search.Parse(params, model.Version{}.Pattern())).
		Count(&count).Error
	return
}

// Save VersionRepo
func (p *VersionRepo) Save(version model.Version) (u model.Version, err error) {
	// TODO, check if it works fine or not
	err = p.Engine.DB.Save(&version).Scan(&u).Error
	return
}

// LastVersion of version table
func (p *VersionRepo) LastVersion() (version model.Version, err error) {
	err = p.Engine.DB.Unscoped().Last(&version).Error
	return
}

// Delete version
func (p *VersionRepo) Delete(version model.Version) (err error) {
	err = p.Engine.DB.Delete(&version).Error
	return
}
