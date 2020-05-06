package service

import (
	"errors"
	"sigmamono/internal/core"
	"sigmamono/model"
	"sigmamono/repo"
	"sigmamono/test/kernel"
	"testing"
	"time"
)

func initCompanyTest() (engine *core.Engine, companyServ CompanyServ) {
	logQuery, debugLevel := initServiceTest()
	engine = kernel.StartMotor(logQuery, debugLevel)
	companyServ = ProvideCompanyService(repo.ProvideCompanyRepo(engine))

	return
}

func TestCompanyCreate(t *testing.T) {
	engine, companyServ := initCompanyTest()
	_ = engine

	samples := []struct {
		in  model.Company
		err error
	}{
		{
			in: model.Company{
				Name:       "Sigma",
				LegalName:  "c1",
				Key:        "123456789",
				Expiration: time.Now().AddDate(1, 0, 0),
				License:    "regular",
				Detail:     "",
				Phone:      "07505149171",
				Email:      "info@erp14.com",
				Website:    "erp14.com",
				Type:       "multi branch with centeral finance",
				Code:       "9962",
			},
			err: nil,
		},
		{
			in: model.Company{
				Name:       "Sigma",
				LegalName:  "c1",
				Key:        "123456789",
				Expiration: time.Now().AddDate(1, 0, 0),
				License:    "regular",
				Detail:     "",
				Phone:      "07505149171",
				Email:      "info@erp14.com",
				Website:    "erp14.com",
				Type:       "multi branch with centeral finance",
				Code:       "9962",
			},
			err: errors.New("duplicate"),
		},
	}

	for _, v := range samples {
		_, err := companyServ.Save(v.in)
		if (v.err == nil && err != nil) || (v.err != nil && err == nil) {
			t.Errorf("ERROR FOR F::::%+v::: \nRETURNS :::%+v:::, \nIT SHOULD BE :::%+v:::", v.in, err, v.err)
		}
		// t.Log(err)
	}

}
