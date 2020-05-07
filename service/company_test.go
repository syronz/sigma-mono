package service

import (
	"errors"
	"sigmamono/internal/core"
	"sigmamono/internal/param"
	"sigmamono/internal/types"
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
	_, companyServ := initCompanyTest()

	samples := []struct {
		in  model.Company
		err error
	}{
		{
			in: model.Company{
				Name:       "created 1",
				LegalName:  "created 1",
				Key:        "created 1",
				Expiration: time.Now().AddDate(1, 0, 0),
				License:    "created 1",
				Detail:     "",
				Phone:      "created 1",
				Email:      "created 1",
				Website:    "created 1",
				Type:       "multi branch with centeral finance",
				Code:       "created 1",
			},
			err: nil,
		},
		{
			in: model.Company{
				Name:       "created 1",
				LegalName:  "created 1",
				Key:        "created 1",
				Expiration: time.Now().AddDate(1, 0, 0),
				License:    "created 1",
				Detail:     "created 1",
				Phone:      "created 1",
				Email:      "created 1",
				Website:    "created 1",
				Type:       "multi branch with centeral finance",
				Code:       "created 1",
			},
			err: errors.New("duplicate"),
		},
		{
			in: model.Company{
				Name:       "minimum fields",
				LegalName:  "mini fields",
				Expiration: time.Now().AddDate(1, 0, 0),
				License:    "miniLicense",
				Type:       "base",
			},
			err: nil,
		},
		{
			in: model.Company{
				Name:       "long name: big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name",
				LegalName:  "long name",
				Expiration: time.Now().AddDate(1, 0, 0),
				License:    "long name",
				Type:       "base",
			},
			err: errors.New("data too long for name"),
		},
		{
			in: model.Company{
				Name:       "wrong type",
				LegalName:  "wrong type",
				Expiration: time.Now().AddDate(1, 0, 0),
				License:    "wrong type",
				Type:       "not existed type",
			},
			err: errors.New("data too long for name"),
		},
	}

	for _, v := range samples {
		_, err := companyServ.Save(v.in)
		if (v.err == nil && err != nil) || (v.err != nil && err == nil) {
			t.Errorf("\nERROR FOR :::%+v::: \nRETURNS :::%+v:::, \nIT SHOULD BE :::%+v:::", v.in, err, v.err)
		}
	}

}

func TestCompanyUpdate(t *testing.T) {
	_, companyServ := initCompanyTest()

	samples := []struct {
		in  model.Company
		err error
	}{
		{
			in: model.Company{
				GormCol: types.GormCol{
					ID: 1002,
				},
				Name:       "num 1 updated",
				LegalName:  "num 1 updated",
				Key:        "num 1 updated",
				Expiration: time.Now().AddDate(1, 0, 0),
				License:    "num 1 updated",
				Detail:     "num 1 updated",
				Phone:      "num 1 updated",
				Email:      "num 1 updated",
				Website:    "num 1 updated",
				Type:       "base",
				Code:       "num 1 updated",
			},
			err: nil,
		},
		{
			in: model.Company{
				GormCol: types.GormCol{
					ID: 1003,
				},
				Name:    "num 2 updated",
				Email:   "num 2 updated",
				Website: "num 2 updated",
				Type:    "base",
				Code:    "num 2 updated",
			},
			err: errors.New("legal_name is required"),
		},
	}

	for _, v := range samples {
		_, err := companyServ.Save(v.in)
		if (v.err == nil && err != nil) || (v.err != nil && err == nil) {
			t.Errorf("ERROR FOR ::::%+v::: \nRETURNS :::%+v:::, \nIT SHOULD BE :::%+v:::", v.in, err, v.err)
		}
	}

}

func TestCompanyDelete(t *testing.T) {
	_, companyServ := initCompanyTest()

	samples := []struct {
		id  types.RowID
		err error
	}{
		{
			id:  1004,
			err: nil,
		},
		{
			id:  99999999,
			err: errors.New("record not found"),
		},
	}

	for _, v := range samples {
		_, err := companyServ.Delete(v.id)
		if (v.err == nil && err != nil) || (v.err != nil && err == nil) {
			t.Errorf("ERROR FOR ::::%+v::: \nRETURNS :::%+v:::, \nIT SHOULD BE :::%+v:::", v.id, err, v.err)
		}
	}
}

func TestCompanyList(t *testing.T) {
	_, companyServ := initCompanyTest()
	regularParam := getRegularParam("companies.id asc")
	regularParam.Search = "searchTerm1"

	samples := []struct {
		params param.Param
		count  uint64
		err    error
	}{
		{
			params: param.Param{},
			err:    errors.New("error in url"),
			count:  0,
		},
		{
			params: regularParam,
			err:    nil,
			count:  6,
		},
	}

	for _, v := range samples {
		data, err := companyServ.List(v.params)
		var count uint64
		var ok bool
		if count, ok = data["count"].(uint64); !ok {
			count = 0
		}
		if (v.err == nil && err != nil) || (v.err != nil && err == nil) || count != v.count {
			t.Log(".........................", data)
			t.Errorf("FOR ::::%+v::: \nRETURNS :::%+v:::, \nIT SHOULD BE :::%+v:::", v.params, data["count"], v.count)
		}
	}
}

func TestCompanyExcel(t *testing.T) {
	_, companyServ := initCompanyTest()
	regularParam := getRegularParam("companies.id asc")

	samples := []struct {
		params param.Param
		count  uint64
		err    error
	}{
		{
			params: regularParam,
			err:    nil,
			count:  6,
		},
	}

	for _, v := range samples {
		data, err := companyServ.Excel(v.params)
		if (v.err == nil && err != nil) || (v.err != nil && err == nil) || uint64(len(data)) < v.count {
			t.Log(".........................", data)
			t.Errorf("FOR ::::%+v::: \nRETURNS :::%+v:::, \nIT SHOULD BE :::%+v::: \nErr :::%+v:::",
				v.params, uint64(len(data)), v.count, err)
		}
	}
}
