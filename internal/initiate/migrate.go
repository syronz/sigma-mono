package initiate

import (
	"sigmamono/internal/core"
	"sigmamono/model"
)

// Migrate the database for creating tables
func Migrate(engine *core.Engine) {
	engine.DB.AutoMigrate(&model.Role{})
	engine.DB.AutoMigrate(&model.Account{})
	engine.DB.AutoMigrate(&model.User{}).
		AddForeignKey("id", "accounts(id)", "CASCADE", "CASCADE").
		AddForeignKey("role_id", "roles(id)", "RESTRICT", "RESTRICT")

}
