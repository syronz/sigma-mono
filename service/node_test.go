package service

import (
	"sigmamono/internal/core"
	"sigmamono/model"
	"sigmamono/repo"
	"sigmamono/test/kernel"
	"testing"
)

func initNodeTest() (engine *core.Engine, nodeServ NodeServ) {
	logQuery, debugLevel := initServiceTest()
	engine = kernel.StartMotor(logQuery, debugLevel)
	nodeServ = ProvideNodeService(repo.ProvideNodeRepo(engine))

	return
}

func TestNodeCreate(t *testing.T) {
	engine, nodeServ := initNodeTest()
	_ = engine

	samples := []struct {
		in       model.Node
		hasError bool
	}{}

	for _, v := range samples {
		_, err := nodeServ.Save(v.in)
		t.Log(err)
	}

}
