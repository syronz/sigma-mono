package repo

import (
	"fmt"
	"sigmamono/internal/core"
	"sigmamono/internal/param"
	"sigmamono/internal/search"
	"sigmamono/internal/types"
	"sigmamono/model"

	"github.com/jinzhu/gorm"
)

// SyncSessionRepo for injecting engine
type SyncSessionRepo struct {
	Engine *core.Engine
}

// ProvideSyncSessionRepo is used in wire
func ProvideSyncSessionRepo(engine *core.Engine) SyncSessionRepo {
	return SyncSessionRepo{Engine: engine}
}

// FindByID for syncSession
func (p *SyncSessionRepo) FindByID(id types.RowID) (syncSession model.SyncSession, err error) {
	err = p.Engine.DB.First(&syncSession, id.ToUint64()).Error
	return
}

// DELETE BELOW
// List of syncSessions
func (p *SyncSessionRepo) List(params param.Param) (syncSessions []model.SyncSession, err error) {
	columns, err := model.SyncSession{}.Columns(params.Select)
	if err != nil {
		return
	}

	err = p.Engine.DB.Select(columns).
		Joins("INNER JOIN companies on companies.id = syncSessions.company_id").
		Where("syncSessions.company_id = ?", params.CompanyID).
		Where(search.Parse(params, model.SyncSession{}.Pattern())).
		Order(params.Order).
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&syncSessions).Error

	return
}

// Count of syncSessions
func (p *SyncSessionRepo) Count(params param.Param) (count uint64, err error) {
	err = p.Engine.DB.Table("syncSessions").
		Select(params.Select).
		Where("deleted_at is null").
		Where("syncSessions.company_id = ?", params.CompanyID).
		Where(search.Parse(params, model.SyncSession{}.Pattern())).
		Count(&count).Error
	return
}

// Update SyncSessionRepo
func (p *SyncSessionRepo) Update(syncSession model.SyncSession) (u model.SyncSession, err error) {
	err = p.Engine.DB.Save(&syncSession).Error
	p.Engine.DB.Where("id = ?", syncSession.ID).Find(&u)
	return
}

// Create SyncSessionRepo
func (p *SyncSessionRepo) Create(syncSession model.SyncSession) (u model.SyncSession, err error) {
	err = p.Engine.DB.Create(&syncSession).Scan(&u).Error
	return
}

// LastSyncSession of syncSession table
func (p *SyncSessionRepo) LastSyncSession(prefix types.RowID) (syncSession model.SyncSession, err error) {
	err = p.Engine.DB.Unscoped().Where("id LIKE ?", fmt.Sprintf("%v%%", prefix)).
		Last(&syncSession).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}
	return
}

// Delete syncSession
func (p *SyncSessionRepo) Delete(syncSession model.SyncSession) (err error) {
	err = p.Engine.DB.Delete(&syncSession).Error
	return
}

// HardDelete syncSession
func (p *SyncSessionRepo) HardDelete(syncSession model.SyncSession) (err error) {
	err = p.Engine.DB.Unscoped().Delete(&syncSession).Error
	return
}
