package main

import (
	"flag"
	"sigmamono/cmd/cloud/determine"
	"sigmamono/cmd/cloud/insertdata"
	"sigmamono/cmd/cloud/server"
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
	defer engine.DB.Close()
	initiate.ConnectActivityDB(engine)
	determine.Migrate(engine)
	initiate.Migrate(engine)
	insertdata.Insert(engine)

	engine.Debug("cloud server started!")
	server.Start(engine)

}
