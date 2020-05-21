package insertdata

import (
	"sigmamono/internal/core"
)

// Insert is used for add static rows to database
func Insert(engine *core.Engine) {

	if engine.Env.Setting.AutoMigrate {
		// table.InsertVersions(engine)
		// table.InsertCompanies(engine)
		// table.InsertNodes(engine)
		// table.InsertStations(engine)
		// table.InsertSettings(engine)
		// table.InsertRoles(engine)
		// table.InsertAccounts(engine)
		// table.InsertUsers(engine)
	}

}
