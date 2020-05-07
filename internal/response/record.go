package response

import (
	"fmt"
	"sigmamono/internal/enum/event"
	"sigmamono/repo"
	"sigmamono/service"
)

func (r *Response) Record(ev event.Event, data ...interface{}) {
	activityServ := service.ProvideActivityService(repo.ProvideActivityRepo(r.Engine))
	activityServ.Record(r.Context, ev, data...)
	fmt.Println("this is simple record........................", activityServ)
}
