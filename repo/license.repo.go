package repo

import (
	"sigmamono/internal/core"
	"sigmamono/internal/types"
	"sigmamono/model"
)

// LicenseRepo for injecting engine
type LicenseRepo struct {
	Engine *core.Engine
}

// ProvideLicenseRepo is used in wire
func ProvideLicenseRepo(engine *core.Engine) LicenseRepo {
	return LicenseRepo{Engine: engine}
}

// FindByID for license
func (p *LicenseRepo) FindByID(id types.RowID) (license model.License, err error) {
	err = p.Engine.DB.First(&license, id.ToUint64()).Error
	return
}

// Create activation row for record usage of license
func (p *LicenseRepo) Create(activation model.Activation) (a model.Activation, err error) {
	err = p.Engine.DB.Create(&activation).Scan(&a).Error
	return
}
