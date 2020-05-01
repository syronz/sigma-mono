package main

import (
	"radiusbilling/internal/initiate"
	"radiusbilling/internal/logparam"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {

	engine := initiate.LoadEnv()
	logparam.ServerLog(engine)
	logparam.APILog(engine)
	initiate.LoadTerms(engine)
	initiate.ConnectDB(engine)
	initiate.ConnectActivityDB(engine)
	initiate.Migrate(engine)

	engine.Debug("server started!")
}
