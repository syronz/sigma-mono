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
			CompanyID: 1001,
			Code:      101,
			Type:      nodetype.Online,
			Name:      "Base",
			MachineID: "111111",
			Status:    nodestatus.Active,
			Phone:     "07505149171",
		},
		{
			GormCol: types.GormCol{
				ID: 2,
			},
			CompanyID: 1002,
			Code:      101,
			Type:      nodetype.Online,
			Name:      "Base",
			MachineID: "111111",
			Status:    nodestatus.Active,
			Phone:     "07505149171",
		},
	}

	for _, v := range nodes {
		if _, err := nodeService.Save(v); err != nil {
			engine.ServerLog.Fatal(err)
		}

	}

}
