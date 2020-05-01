package logparam

import (
	"radiusbilling/internal/core"
)

// APILog is used inside internal.core.StartEngine and test.core.StartEngine
func APILog(engine *core.Engine) {
	apiLogParams := LogParam{
		format:       engine.Env.Log.APILog.Format,
		output:       engine.Env.Log.APILog.Output,
		level:        engine.Env.Log.APILog.Level,
		JSONIndent:   engine.Env.Log.APILog.JSONIndent,
		showFileLine: false,
	}

	engine.APILog = initLog(apiLogParams)
}
