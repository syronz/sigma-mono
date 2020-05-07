package service

import (
	"errors"
	"sigmamono/internal/core"
	"sigmamono/internal/param"
	"sigmamono/internal/types"
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
		in  model.Node
		err error
	}{
		{
			in: model.Node{
				Type:      "PC",
				CompanyID: 1001,
				Name:      "created node 1",
				Status:    "active",
				MachineID: "machine id for created 1",
				Phone:     "07505149171",
			},
			err: nil,
		},
		{
			in: model.Node{
				Type:      "PC",
				CompanyID: 1001,
				Name:      "created node 1",
				Status:    "active",
				MachineID: "machine id for created 1",
				Phone:     "07505149171",
			},
			err: errors.New("duplicate"),
		},
		{
			in: model.Node{
				Type:      "check validation",
				CompanyID: 1001,
				Name:      "created node 3",
				Status:    "active",
				MachineID: "machine id for created 3",
				Phone:     "07505149171",
			},
			err: errors.New("validation"),
		},
	}

	for _, v := range samples {
		_, err := nodeServ.Save(v.in)
		if (v.err == nil && err != nil) || (v.err != nil && err == nil) {
			t.Errorf("\nERROR FOR :::%+v::: \nRETURNS :::%+v:::, \nIT SHOULD BE :::%+v:::", v.in, err, v.err)
		}
	}
}

func TestNodeUpdate(t *testing.T) {
	_, nodeServ := initNodeTest()

	samples := []struct {
		in  model.Node
		err error
	}{
		{
			in: model.Node{
				GormCol: types.GormCol{
					ID: 5,
				},
				Code:      166, // code can be changed
				Type:      "server-private",
				Status:    "active",
				CompanyID: 1001,
				Name:      "node 1 updated",
				MachineID: "machine id 1 updated",
				Phone:     "node 1 updated",
			},
			err: nil,
		},
		{
			in: model.Node{
				GormCol: types.GormCol{
					ID: 6,
				},
				Code:      177, // code can be changed
				CompanyID: 1001,
				Name:      "node 2 not updated",
				MachineID: "machine id not 2 updated",
				Phone:     "node 2 not updated",
			},
			err: errors.New("validation"),
		},
		{
			in: model.Node{
				GormCol: types.GormCol{
					ID: 6,
				},
				Code:      9999999, // wrong code
				Type:      "online",
				Status:    "active",
				CompanyID: 1001,
				Name:      "node 2 not updated",
				MachineID: "machine id not 2 updated",
				Phone:     "node 2 not updated",
			},
			err: errors.New("validation"),
		},
	}

	for _, v := range samples {
		_, err := nodeServ.Save(v.in)
		if (v.err == nil && err != nil) || (v.err != nil && err == nil) {
			t.Errorf("ERROR FOR ::::%+v::: \nRETURNS :::%+v:::, \nIT SHOULD BE :::%+v:::", v.in, err, v.err)
		}
	}
}

func TestNodeDelete(t *testing.T) {
	_, nodeServ := initNodeTest()

	samples := []struct {
		id  types.RowID
		err error
	}{
		{
			id:  7,
			err: nil,
		},
		{
			id:  99999999,
			err: errors.New("record not found"),
		},
	}

	for _, v := range samples {
		_, err := nodeServ.Delete(v.id)
		if (v.err == nil && err != nil) || (v.err != nil && err == nil) {
			t.Errorf("ERROR FOR ::::%+v::: \nRETURNS :::%+v:::, \nIT SHOULD BE :::%+v:::", v.id, err, v.err)
		}
	}
}

func TestNodeList(t *testing.T) {
	_, nodeServ := initNodeTest()
	regularParam := getRegularParam("nodes.id asc")
	regularParam.Search = "searchTerm1"

	samples := []struct {
		params param.Param
		count  uint64
		err    error
	}{
		{
			params: param.Param{},
			err:    errors.New("error in url"),
			count:  0,
		},
		{
			params: regularParam,
			err:    nil,
			count:  3,
		},
	}

	for _, v := range samples {
		data, err := nodeServ.List(v.params)
		var count uint64
		var ok bool
		if count, ok = data["count"].(uint64); !ok {
			count = 0
		}
		if (v.err == nil && err != nil) || (v.err != nil && err == nil) || count != v.count {
			t.Log(".........................", data)
			t.Errorf("FOR ::::%+v::: \nRETURNS :::%+v:::, \nIT SHOULD BE :::%+v:::", v.params, data["count"], v.count)
		}
	}
}

func TestNodeExcel(t *testing.T) {
	_, nodeServ := initNodeTest()
	regularParam := getRegularParam("nodes.id asc")

	samples := []struct {
		params param.Param
		count  uint64
		err    error
	}{
		{
			params: regularParam,
			err:    nil,
			count:  6,
		},
	}

	for _, v := range samples {
		data, err := nodeServ.Excel(v.params)
		if (v.err == nil && err != nil) || (v.err != nil && err == nil) || uint64(len(data)) < v.count {
			t.Log(".........................", data)
			t.Errorf("FOR ::::%+v::: \nRETURNS :::%+v:::, \nIT SHOULD BE :::%+v::: \nErr :::%+v:::",
				v.params, uint64(len(data)), v.count, err)
		}
	}
}
