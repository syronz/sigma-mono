package main

import (
	"flag"
	"fmt"
	"sigmamono/cmd/testinsertion/determine"
	"sigmamono/cmd/testinsertion/insertdata"
	"sigmamono/internal/initiate"
	"sigmamono/internal/logparam"
	"sigmamono/test/kernel"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var noRest bool

func init() {
	flag.BoolVar(&noRest, "noReset", false, "by default it drop tables before migrate")
}

func main() {
	flag.Parse()
	fmt.Println(noRest)

	engine := kernel.LoadTestEnv()
	logparam.ServerLog(engine)
	logparam.APILog(engine)
	initiate.LoadTerms(engine)
	initiate.ConnectDB(engine)
	initiate.ConnectActivityDB(engine)
	determine.Migrate(engine, noRest)
	insertdata.Insert(engine)

	fmt.Println("Data has been reset successfully")

}
