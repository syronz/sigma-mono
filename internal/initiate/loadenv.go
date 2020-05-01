package initiate

import (
	"log"
	"radiusbilling/internal/core"

	envEngine "github.com/syronz/env/v6"
)

// LoadEnv get variables from environment and put them inside engine.Env
func LoadEnv() *core.Engine {
	var engine core.Engine
	var err error
	if err = envEngine.Parse(&engine.Env); err != nil {
		log.Fatalln(err)
	}

	return &engine
}
