package response

import (
	"sigmamono/internal/enum/event"
	"sigmamono/repo"
	"sigmamono/service"
)

// Record is used for saving activity
func (r *Response) Record(ev event.Event, data ...interface{}) {
	activityServ := service.ProvideActivityService(repo.ProvideActivityRepo(r.Engine))
	activityServ.Record(r.Context, ev, data...)
}
