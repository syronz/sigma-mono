package main

import (
	"sigmamono/cmd/sync/determine"
	"sigmamono/cmd/sync/server"
	"sigmamono/internal/initiate"
	"sigmamono/internal/logparam"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	engine := initiate.LoadEnv()
	logparam.ServerLog(engine)
	initiate.LoadTerms(engine)
	initiate.ConnectDB(engine, false)
	defer engine.DB.Close()
	determine.Migrate(engine)
	initiate.Migrate(engine)
	// insertdata.Insert(engine)

	engine.Debug("sync server started!")
	server.Start(engine)
}
