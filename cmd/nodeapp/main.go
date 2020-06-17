package main

import (
	"flag"
	"sigmamono/cmd/nodeapp/determine"
	"sigmamono/cmd/nodeapp/insertdata"
	"sigmamono/cmd/nodeapp/server"
	"sigmamono/internal/core"
	"sigmamono/internal/initiate"
	"sigmamono/internal/logparam"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var pathJSON string

func init() {
	flag.StringVar(&pathJSON, "json", "", "load environment from a JSON file")
}

func main() {
	flag.Parse()

	var engine *core.Engine
	if pathJSON == "" {
		engine = initiate.LoadEnv()
	} else {
		engine = initiate.LoadJSON(pathJSON)
	}
	logparam.ServerLog(engine)
	logparam.APILog(engine)
	initiate.LoadTerms(engine)
	initiate.ConnectDB(engine, false)
	initiate.ConnectActivityDB(engine)
	determine.Migrate(engine)
	initiate.Migrate(engine)
	insertdata.Insert(engine)

	engine.Debug("nodeapp server started!")
	server.Start(engine)

}
