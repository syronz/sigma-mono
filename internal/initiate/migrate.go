package initiate

import (
	"sigmamono/internal/core"
	"sigmamono/model"
)

// Migrate the database for creating tables
func Migrate(engine *core.Engine) {
	engine.DB.AutoMigrate(&model.User{})
	engine.DB.AutoMigrate(&model.Role{})

}
