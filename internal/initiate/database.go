package initiate

import (
	"log"
	"sigmamono/internal/core"

	"github.com/jinzhu/gorm"
)

// ConnectDB initiate the db connection by getting help from gorm
func ConnectDB(engine *core.Engine, printQueries bool) {
	var err error
	engine.DB, err = gorm.Open(engine.Env.Database.Data.Type, engine.Env.Database.Data.DSN)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// engine.DB.LogMode(false)

	// if gin.IsDebugging() {
	// 	engine.DB.LogMode(true)
	// }

	engine.DB.LogMode(engine.Env.Database.Data.Log)

	if printQueries {
		engine.DB.LogMode(true)
	}
}

// ConnectActivityDB initiate the db connection by getting help from gorm
func ConnectActivityDB(engine *core.Engine) {
	var err error
	engine.ActivityDB, err = gorm.Open(engine.Env.Database.Activity.Type,
		engine.Env.Database.Activity.DSN)
	if err != nil {
		log.Fatalln(err.Error())
	}

	engine.ActivityDB.LogMode(engine.Env.Database.Activity.Log)

	// engine.ActivityDB.LogMode(false)

	// if gin.IsDebugging() {
	// 	engine.ActivityDB.LogMode(true)
	// }
}
