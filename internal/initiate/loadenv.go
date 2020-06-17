package initiate

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
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

// LoadJSON is used for testing environment
func LoadJSON(jsonPath string) *core.Engine {
	engine := new(core.Engine)

	// _, dir, _, _ := runtime.Caller(0)
	// configJSON := filepath.Join(filepath.Dir(dir), "../..", "env", "tdd.json")

	jsonFile, err := os.Open(jsonPath)

	if err != nil {
		log.Fatalln(err, "can't open the config file", jsonPath)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &engine.Env)
	if err != nil {
		log.Fatalln(err, "error in unmarshal JSON")
	}

	if engine.Env.MachineID, err = machineid.ProtectedID("SigmaMono"); err != nil {
		log.Fatal(err)
	}

	return engine
}
