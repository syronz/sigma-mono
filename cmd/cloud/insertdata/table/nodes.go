package table

import (
	"sigmamono/internal/core"
	"sigmamono/internal/enum/nodestatus"
	"sigmamono/internal/enum/nodetype"
	"sigmamono/internal/param"
	"sigmamono/internal/types"
	"sigmamono/model"
	"sigmamono/repo"
	"sigmamono/service"
)

// InsertNodes for add required nodes
func InsertNodes(engine *core.Engine) {
	nodeRepo := repo.ProvideNodeRepo(engine)
	nodeService := service.ProvideNodeService(nodeRepo)

	nodes := []struct {
		node   model.Node
		params param.Param
	}{
		{
			node: model.Node{
				GormCol: types.GormCol{
					ID: 1,
				},
				Code:      101,
				Type:      nodetype.Online,
				Name:      "Base",
				MachineID: engine.Env.MachineID,
				Status:    nodestatus.Active,
				Phone:     "07505149171",
			},
			params: param.Param{
				CompanyID: 1001,
			},
		},
		{
			node: model.Node{
				GormCol: types.GormCol{
					ID: 2,
				},
				Code:      101,
				Type:      nodetype.Online,
				Name:      "Base",
				MachineID: engine.Env.MachineID,
				Status:    nodestatus.Active,
				Phone:     "07505149171",
			},
			params: param.Param{
				CompanyID: 1002,
			},
		},
	}

	for _, v := range nodes {
		if _, err := nodeService.Save(v.node, v.params); err != nil {
			engine.ServerLog.Fatal(err)
		}

	}

}
