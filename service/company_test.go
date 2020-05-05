package service

import (
	"flag"
	"sigmamono/internal/core"
	"sigmamono/model"
	"sigmamono/repo"
	"sigmamono/test/kernel"
	"testing"
	"time"
)

var log = flag.Bool("log", false, "Log the queries")

func initCompanyTest() (engine *core.Engine, companyServ CompanyServ) {
	engine = kernel.StartMotor()
	companyServ = ProvideCompanyService(repo.ProvideCompanyRepo(engine))
	return
}

func TestCompanyCreate(t *testing.T) {
	t.Log(*log)
	engine, companyServ := initCompanyTest()
	_ = engine
	t.Log("test run successfully")

	samples := []struct {
		in       model.Company
		hasError bool
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
			hasError: false,
		},
	}

	for _, v := range samples {
		_, err := companyServ.Save(v.in)
		t.Log(err)
	}

}
