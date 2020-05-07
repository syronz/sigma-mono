package service

import (
	"encoding/json"
	"fmt"
	"sigmamono/internal/core"
	"sigmamono/internal/enum/event"
	"sigmamono/internal/types"
	"sigmamono/model"
	"sigmamono/repo"

	"github.com/gin-gonic/gin"
)

// RecordType is and int used as an enum
type RecordType int

const (
	read RecordType = iota
	writeBefore
	writeAfter
	writeBoth
)

// Activity for injecting auth repo
type Activity struct {
	Repo   repo.Activity
	Engine *core.Engine
}

// ProvideActivityService for activity is used in wire
func ProvideActivityService(p repo.Activity) Activity {
	return Activity{Repo: p, Engine: p.Engine}
}

// Save activity
func (p *Activity) Save(activity model.Activity) (createdActivity model.Activity, err error) {
	createdActivity, err = p.Repo.Save(activity)

	// p.Engine.CheckInfo(err, fmt.Sprintf("Failed in saving activity for %+v", activity))

	return
}

// Record will save the activity
func (p *Activity) Record(c *gin.Context, ev event.Event, data ...interface{}) {
	var userID types.RowID
	var username string

	recordType := p.findRecordType(data...)
	before, after := p.fillBeforeAfter(recordType, data...)

	if p.isRecordSetInEnvironment(recordType) {
		return
	}
	if userIDtmp, ok := c.Get("USER_ID"); ok {
		userID = userIDtmp.(types.RowID)
	}
	if usernameTmp, ok := c.Get("USERNAME"); ok {
		username = usernameTmp.(string)
	}

	activity := model.Activity{
		Event:    ev.String(),
		UserID:   userID,
		Username: username,
		IP:       c.ClientIP(),
		URI:      c.Request.RequestURI,
		Before:   string(before),
		After:    string(after),
	}

	_, err := p.Repo.Save(activity)
	p.Engine.CheckError(err, fmt.Sprintf("Failed in saving activity for %+v", activity))
}

func (p *Activity) fillBeforeAfter(recordType RecordType, data ...interface{}) (before, after []byte) {
	var err error
	if recordType == writeBefore || recordType == writeBoth {
		before, err = json.Marshal(data[0])
		p.Engine.CheckError(err, "error in encoding data to before-json")
	}
	if recordType == writeAfter || recordType == writeBoth {
		after, err = json.Marshal(data[1])
		p.Engine.CheckError(err, "error in encoding data to after-json")
	}

	return
}

func (p *Activity) findRecordType(data ...interface{}) RecordType {
	switch len(data) {
	case 0:
		return read
	case 2:
		return writeBoth
	default:
		if data[0] == nil {
			return writeAfter
		}
	}

	return writeBefore
}

func (p *Activity) isRecordSetInEnvironment(recordType RecordType) bool {
	switch recordType {
	case read:
		if !p.Engine.Env.Setting.RecordRead {
			return true
		}
	default:
		if !p.Engine.Env.Setting.RecordWrite {
			return true
		}
	}
	return false
}
