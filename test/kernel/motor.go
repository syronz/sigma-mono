package kernel

import (
	"sigmamono/internal/core"
	"sigmamono/internal/initiate"
	"sigmamono/internal/logparam"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// StartMotor for generating engine special for TDD
func StartMotor(printQueries bool, debugLevel bool) *core.Engine {
	engine := LoadTestEnv()
	logparam.ServerLog(engine)
	initiate.LoadTerms(engine)
	if debugLevel {
		engine.Env.Log.ServerLog.Level = "trace"
	}
	logparam.ServerLog(engine)
	initiate.ConnectDB(engine, printQueries)
	initiate.ConnectActivityDB(engine)

	return engine
}
