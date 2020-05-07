package determine

import (
	"sigmamono/internal/core"
	"sigmamono/model"
)

// Migrate the database for creating tables
func Migrate(engine *core.Engine, noReset bool) {

	if !noReset {
		dropTable(engine)
	}

	engine.DB.AutoMigrate(&model.Company{})
	engine.DB.AutoMigrate(&model.Node{}).
		AddForeignKey("company_id", "companies(id)", "RESTRICT", "RESTRICT")

	engine.DB.AutoMigrate(&model.Role{})
	engine.DB.AutoMigrate(&model.Account{})
	engine.DB.AutoMigrate(&model.User{}).
		AddForeignKey("id", "accounts(id)", "CASCADE", "CASCADE").
		AddForeignKey("role_id", "roles(id)", "RESTRICT", "RESTRICT")
}

func dropTable(engine *core.Engine) {
	var err error
	if err = engine.DB.DropTable(&model.User{}).Error; err != nil {
		engine.ServerLog.Error(err)
	}
	if err = engine.DB.DropTable(&model.Account{}).Error; err != nil {
		engine.ServerLog.Error(err)
	}
	if err = engine.DB.DropTable(&model.Role{}).Error; err != nil {
		engine.ServerLog.Error(err)
	}
	if err = engine.DB.DropTable(&model.Node{}).Error; err != nil {
		engine.ServerLog.Error(err)
	}
	if err = engine.DB.DropTable(&model.Company{}).Error; err != nil {
		engine.ServerLog.Error(err)
	}
}
