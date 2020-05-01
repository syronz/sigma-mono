package logparam

import (
	"radiusbilling/internal/core"
)

// ServerLog connected to the engine
func ServerLog(engine *core.Engine) {
	// Server logs's params
	serverLogParam := LogParam{
		format:       engine.Env.Log.ServerLog.Format,
		output:       engine.Env.Log.ServerLog.Output,
		level:        engine.Env.Log.ServerLog.Level,
		JSONIndent:   engine.Env.Log.ServerLog.JSONIndent,
		showFileLine: true, // true means filename and line number should be printed
	}

	engine.ServerLog = initLog(serverLogParam)

}
