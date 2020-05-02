package initiate

import (
	"radiusbilling/internal/core"
	"radiusbilling/model"
)

// Migrate the database for creating tables
func Migrate(engine *core.Engine) {
	engine.DB.AutoMigrate(&model.User{})
	engine.DB.AutoMigrate(&model.Role{})

}
