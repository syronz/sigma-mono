package determine

import (
	"sigmamono/internal/core"
	"sigmamono/model"
)

// Migrate the database for creating tables
func Migrate(engine *core.Engine) {
	engine.DB.AutoMigrate(&model.Version{})
	engine.DB.AutoMigrate(&model.Company{})
	engine.DB.AutoMigrate(&model.Node{}).
		AddForeignKey("company_id", "companies(id)", "RESTRICT", "RESTRICT")
	// engine.DB.Set("gorm:table_options", "ENGINE=Archive").AutoMigrate(&model.LicenseHistory{})
	engine.DB.AutoMigrate(&model.Activation{})
}
