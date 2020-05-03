package main

import (
	"sigmamono/cmd/cloud/determine"
	"sigmamono/cmd/cloud/insertdata"
	"sigmamono/cmd/cloud/server"
	"sigmamono/internal/initiate"
	"sigmamono/internal/logparam"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {

	// engine := setup.LoadEnv()
	// logparam.ServerLog(engine)
	// logparam.APILog(engine)
	// setup.LoadTerms(engine)
	// setup.AES(engine)
	// setup.ConnectDB(engine)
	// setup.ConnectActivityDB(engine)
	// migrate.Migrate(engine)

	// go insertdata.Insert(engine)

	// server.Initialize(engine)

	engine := initiate.LoadEnv()
	logparam.ServerLog(engine)
	logparam.APILog(engine)
	initiate.LoadTerms(engine)
	initiate.ConnectDB(engine)
	initiate.ConnectActivityDB(engine)
	determine.Migrate(engine)
	initiate.Migrate(engine)
	insertdata.Insert(engine)

	engine.Debug("server started!")
	server.Start(engine)

}
