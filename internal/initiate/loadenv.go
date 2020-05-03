package initiate

import (
	"log"
	"sigmamono/internal/core"

	envEngine "github.com/syronz/env/v6"
	"github.com/syronz/machineid"
)

// LoadEnv get variables from environment and put them inside engine.Env
func LoadEnv() *core.Engine {
	var engine core.Engine
	var err error
	if err = envEngine.Parse(&engine.Env); err != nil {
		log.Fatalln(err)
	}

	if engine.Env.MachineID, err = machineid.ProtectedID("SigmaMono"); err != nil {
		log.Fatal(err)
	}

	return &engine
}
