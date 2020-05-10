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
	params := param.Param{
		CompanyID: 1001,
	}

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
				Code:      102,
				Type:      nodetype.PC,
				Name:      "Simple PC",
				Status:    nodestatus.Active,
				MachineID: engine.Env.MachineID,
				Phone:     "07505149171",
			},
			params: param.Param{
				CompanyID: 1001,
			},
		},
		{
			node: model.Node{
				GormCol: types.GormCol{
					ID: 3,
				},
				Code:      101,
				Type:      nodetype.PC,
				Name:      "Simple PC",
				Status:    nodestatus.Active,
				MachineID: engine.Env.MachineID,
				Phone:     "07505149171",
			},
			params: param.Param{
				CompanyID: 1002,
			},
		},
		{
			node: model.Node{
				GormCol: types.GormCol{
					ID: 4,
				},
				Code:      102,
				Type:      nodetype.ServerPrivate,
				Name:      "minimum",
				Status:    nodestatus.Active,
				MachineID: "minimum machine ID",
				Phone:     "0750",
			},
			params: param.Param{
				CompanyID: 1002,
			},
		},
		{
			node: model.Node{
				GormCol: types.GormCol{
					ID: 5,
				},
				Code:      103,
				Type:      nodetype.PC,
				Name:      "node for update 1",
				Status:    nodestatus.Active,
				MachineID: "node for update 1",
				Phone:     "node for update 1",
			},
			params: param.Param{
				CompanyID: 1001,
			},
		},
		{
			node: model.Node{
				GormCol: types.GormCol{
					ID: 6,
				},
				Code:      104,
				Type:      nodetype.PC,
				Name:      "node for update 2",
				Status:    nodestatus.Active,
				MachineID: "node for update 2",
				Phone:     "node for update 2",
			},
			params: param.Param{
				CompanyID: 1001,
			},
		},
		{
			node: model.Node{
				GormCol: types.GormCol{
					ID: 7,
				},
				Code:      105,
				Type:      nodetype.PC,
				Name:      "node for delete 1",
				Status:    nodestatus.Active,
				MachineID: "node for delete 1",
				Phone:     "node for delete 1",
			},
			params: param.Param{
				CompanyID: 1001,
			},
		},
		{
			node: model.Node{
				GormCol: types.GormCol{
					ID: 8,
				},
				Code:      106,
				Type:      nodetype.PC,
				Name:      "for list 1",
				Status:    nodestatus.Active,
				MachineID: "for list 1",
				Phone:     "searchTerm1",
			},
			params: param.Param{
				CompanyID: 1001,
			},
		},
		{
			node: model.Node{
				GormCol: types.GormCol{
					ID: 9,
				},
				Code:      107,
				Type:      nodetype.PC,
				Name:      "for list 2",
				Status:    nodestatus.Active,
				MachineID: "for list 2",
				Phone:     "searchTerm1",
			},
			params: param.Param{
				CompanyID: 1001,
			},
		},
		{
			node: model.Node{
				GormCol: types.GormCol{
					ID: 10,
				},
				Code:      108,
				Type:      nodetype.PC,
				Name:      "for list 3",
				Status:    nodestatus.Active,
				MachineID: "for list 3",
				Phone:     "searchTerm1",
			},
			params: param.Param{
				CompanyID: 1001,
			},
		},
	}

	for _, v := range nodes {
		if _, err := nodeService.Save(v.node, v.params); err != nil {
			engine.ServerLog.Fatal(err)
		}
	}

}
