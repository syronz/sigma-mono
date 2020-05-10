package main

import (
	"sigmamono/cmd/nodeapp/determine"
	"sigmamono/cmd/nodeapp/insertdata"
	"sigmamono/cmd/nodeapp/server"
	"sigmamono/internal/initiate"
	"sigmamono/internal/logparam"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {

	engine := initiate.LoadEnv()
	logparam.ServerLog(engine)
	logparam.APILog(engine)
	initiate.LoadTerms(engine)
	initiate.ConnectDB(engine, false)
	initiate.ConnectActivityDB(engine)
	determine.Migrate(engine)
	initiate.Migrate(engine)
	insertdata.Insert(engine)

	engine.Debug("server started!")
	server.Start(engine)

}
