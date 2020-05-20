// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package determine

import (
	"sigmamono/api"
	"sigmamono/internal/core"
	"sigmamono/repo"
	"sigmamono/service"
)

// Injectors from wire.go:

func initBondAPI(e *core.Engine) api.BondAPI {
	bondRepo := repo.ProvideBondRepo(e)
	bondServ := service.ProvideBondService(bondRepo)
	bondAPI := api.ProvideBondAPI(bondServ)
	return bondAPI
}
