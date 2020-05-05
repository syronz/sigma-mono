package kernel

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sigmamono/internal/core"
)

// LoadTestEnv is used for testing environment
func LoadTestEnv() *core.Engine {
	engine := new(core.Engine)

	_, dir, _, _ := runtime.Caller(0)
	configJSON := filepath.Join(filepath.Dir(dir), "../..", "env", "tdd.json")

	jsonFile, err := os.Open(configJSON)

	if err != nil {
		log.Fatalln(err, "can't open the config file", dir+"/regularenvs.json")
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &engine.Env)
	if err != nil {
		log.Fatalln(err, "error in unmarshal JSON")
	}

	return engine
}
