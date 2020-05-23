package service

import (
	"sigmamono/internal/core"
	"sigmamono/repo"
	"sigmamono/test/kernel"
	"testing"
	"time"
)

func initSyncSessionTest() (engine *core.Engine, syncSessionServ SyncSessionServ) {
	logQuery, debugLevel := initServiceTest()
	engine = kernel.StartMotor(logQuery, debugLevel)
	syncSessionServ = ProvideSyncSessionService(repo.ProvideSyncSessionRepo(engine))

	return
}

func TestSyncSessionGenerate(t *testing.T) {
	_, syncSessionServ := initSyncSessionTest()

	syncSessionServ.Generate(200, time.Now())

	t.Log("this is test")
}
