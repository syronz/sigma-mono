package repo

import (
	"sigmamono/internal/core"
	"sigmamono/model"
)

// Activity for injecting engine
type Activity struct {
	Engine *core.Engine
}

// ProvideActivityRepo is used in wire
func ProvideActivityRepo(engine *core.Engine) Activity {
	return Activity{Engine: engine}
}

// Save Activity
func (p *Activity) Save(activity model.Activity) (u model.Activity, err error) {
	err = p.Engine.ActivityDB.Save(&activity).Error
	return
}
