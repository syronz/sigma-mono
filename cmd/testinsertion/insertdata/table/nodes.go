package table

import (
	"sigmamono/internal/core"
	"sigmamono/internal/enum/nodestatus"
	"sigmamono/internal/enum/nodetype"
	"sigmamono/internal/types"
	"sigmamono/model"
	"sigmamono/repo"
	"sigmamono/service"
)

// InsertNodes for add required nodes
func InsertNodes(engine *core.Engine) {
	nodeRepo := repo.ProvideNodeRepo(engine)
	nodeService := service.ProvideNodeService(nodeRepo)
	nodes := []model.Node{
		{
			GormCol: types.GormCol{
				ID: 1,
			},
			Code:      101,
			Type:      nodetype.Online,
			CompanyID: 1001,
			Name:      "Base",
			Status:    nodestatus.Active,
			MachineID: engine.Env.MachineID,
			Phone:     "07505149171",
		},
		{
			GormCol: types.GormCol{
				ID: 2,
			},
			Code:      102,
			Type:      nodetype.PC,
			CompanyID: 1001,
			Name:      "Simple PC",
			Status:    nodestatus.Active,
			MachineID: "local machine ID",
			Phone:     "07505149171",
		},
		{
			GormCol: types.GormCol{
				ID: 3,
			},
			Code:      101,
			Type:      nodetype.PC,
			CompanyID: 1002,
			Name:      "Simple PC",
			Status:    nodestatus.Active,
			MachineID: engine.Env.MachineID,
			Phone:     "07505149171",
		},
		{
			GormCol: types.GormCol{
				ID: 4,
			},
			Code:      102,
			Type:      nodetype.ServerPrivate,
			CompanyID: 1002,
			Name:      "minimum",
			Status:    nodestatus.Active,
			MachineID: "miniumu machine ID",
			Phone:     "075",
		},
		{
			GormCol: types.GormCol{
				ID: 5,
			},
			Code:      103,
			Type:      nodetype.PC,
			CompanyID: 1001,
			Name:      "node for update 1",
			Status:    nodestatus.Active,
			MachineID: "node for update 1",
			Phone:     "node for update 1",
		},
		{
			GormCol: types.GormCol{
				ID: 6,
			},
			Code:      104,
			Type:      nodetype.PC,
			CompanyID: 1001,
			Name:      "node for update 2",
			Status:    nodestatus.Active,
			MachineID: "node for update 2",
			Phone:     "node for update 2",
		},
		{
			GormCol: types.GormCol{
				ID: 7,
			},
			Code:      105,
			Type:      nodetype.PC,
			CompanyID: 1001,
			Name:      "node for delete 1",
			Status:    nodestatus.Active,
			MachineID: "node for delete 1",
			Phone:     "node for delete 1",
		},
		{
			GormCol: types.GormCol{
				ID: 8,
			},
			Code:      106,
			Type:      nodetype.PC,
			CompanyID: 1001,
			Name:      "for list 1",
			Status:    nodestatus.Active,
			MachineID: "for list 1",
			Phone:     "searchTerm1",
		},
		{
			GormCol: types.GormCol{
				ID: 9,
			},
			Code:      107,
			Type:      nodetype.PC,
			CompanyID: 1001,
			Name:      "for list 2",
			Status:    nodestatus.Active,
			MachineID: "for list 2",
			Phone:     "searchTerm1",
		},
		{
			GormCol: types.GormCol{
				ID: 10,
			},
			Code:      108,
			Type:      nodetype.PC,
			CompanyID: 1001,
			Name:      "for list 3",
			Status:    nodestatus.Active,
			MachineID: "for list 3",
			Phone:     "searchTerm1",
		},
	}

	for _, v := range nodes {
		if _, err := nodeService.Save(v); err != nil {
			engine.ServerLog.Fatal(err)
		}

	}

}
