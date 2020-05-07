package core

import (
	"encoding/json"
	"fmt"
	"sigmamono/env"
	"sigmamono/internal/enum/lang"
	"sigmamono/internal/term"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	goaes "github.com/syronz/goAES"
)

// Engine to keep all database connections and
// logs configuration and environments and etc
type Engine struct {
	DB         *gorm.DB
	ActivityDB *gorm.DB
	ServerLog  *logrus.Logger
	APILog     *logrus.Logger
	Env        env.Environment
	Dict       term.Dict
	AES        goaes.BuildModel
}

// Debug print struct with details with logrus ability
func (e *Engine) Debug(objs ...interface{}) {
	for _, v := range objs {
		parts := make(map[string]interface{}, 2)
		parts["type"] = fmt.Sprintf("%T", v)
		parts["value"] = v
		dataInJSON, _ := json.Marshal(parts)

		e.ServerLog.Debug(string(dataInJSON))
	}
}

// CheckError print all errors which happened inside the services, mainly they just have
// an error and a message
func (e *Engine) CheckError(err error, message string, data ...interface{}) {
	if err != nil {
		e.ServerLog.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Error(message)
		if data != nil {
			e.Debug(data...)
		}
	}
}

// CheckInfo print all errors which happened inside the services, mainly they just have
// an error and a message
func (e *Engine) CheckInfo(err error, message string, data ...interface{}) {
	if err != nil {
		e.ServerLog.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Info(message)
		if data != nil {
			e.Debug(data...)
		}
	}
}

// T Translating the term
func (e *Engine) T(str string, language lang.Language, params ...interface{}) string {
	return e.Dict.Translate(str, language, params...)
}

// SafeT Translating the term and if the word won't exist return false
func (e *Engine) SafeT(str string, language lang.Language, params ...interface{}) (string, bool) {
	return e.Dict.SafeTranslate(str, language, params...)
}
