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

var noReset bool
var logQuery bool

func init() {
	flag.BoolVar(&noReset, "noReset", false, "by default it drop tables before migrate")
	flag.BoolVar(&logQuery, "logQuery", false, "print queries in gorm")
}

func main() {
	flag.Parse()

	engine := kernel.LoadTestEnv()
	logparam.ServerLog(engine)
	logparam.APILog(engine)
	initiate.LoadTerms(engine)
	initiate.ConnectDB(engine, logQuery)
	initiate.ConnectActivityDB(engine)
	determine.Migrate(engine, noReset)
	insertdata.Insert(engine)

	if noReset {
		fmt.Println("Data has been migrated successfully (no reset)")
	} else {
		fmt.Println("Data has been reset successfully")
	}

}
