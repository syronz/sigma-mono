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
func StartMotor() *core.Engine {
	engine := LoadTestEnv()
	logparam.ServerLog(engine)
	initiate.LoadTerms(engine)
	initiate.ConnectDB(engine)
	initiate.ConnectActivityDB(engine)

	return engine
}
