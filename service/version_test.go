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
)

func initVersionTest() (engine *core.Engine, versionServ VersionServ) {
	logQuery, debugLevel := initServiceTest()
	engine = kernel.StartMotor(logQuery, debugLevel)
	versionServ = ProvideVersionService(repo.ProvideVersionRepo(engine))

	return
}

func TestVersionCreate(t *testing.T) {
	_, versionServ := initVersionTest()

	samples := []struct {
		in  model.Version
		err error
	}{
		{
			in: model.Version{
				Name:          "created 1",
				Features:      "created 1",
				NodeCount:     10,
				LocationCount: 10,
				UserCount:     10,
				MonthExpire:   10,
				Description:   "created 1",
			},
			err: nil,
		},
		{
			in: model.Version{
				Name:          "created 1",
				Features:      "created 1",
				NodeCount:     10,
				LocationCount: 10,
				UserCount:     10,
				MonthExpire:   10,
				Description:   "created 1",
			},
			err: errors.New("duplicate"),
		},
		{
			in: model.Version{
				Name:          "minimum fields",
				Features:      "mini fields",
				LocationCount: 10,
				UserCount:     10,
			},
			err: nil,
		},
		{
			in: model.Version{
				Name:          "long name: big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name",
				Features:      "long name",
				LocationCount: 10,
				UserCount:     10,
			},
			err: errors.New("data too long for name"),
		},
		{
			in: model.Version{
				Name:          "wrong type",
				Features:      "wrong type",
				LocationCount: 13,
				UserCount:     13,
			},
			err: nil,
		},
		{
			in: model.Version{
				Features:      "wrong type",
				LocationCount: 13,
				UserCount:     13,
			},
			err: errors.New("name is required"),
		},
	}

	for _, v := range samples {
		_, err := versionServ.Save(v.in)
		if (v.err == nil && err != nil) || (v.err != nil && err == nil) {
			t.Errorf("\nERROR FOR :::%+v::: \nRETURNS :::%+v:::, \nIT SHOULD BE :::%+v:::", v.in, err, v.err)
		}
	}

}

func TestVersionUpdate(t *testing.T) {
	_, versionServ := initVersionTest()

	samples := []struct {
		in   model.Version
		err  error
		name string
	}{
		{
			in: model.Version{
				GormCol: types.GormCol{
					ID: 55,
				},
				Name:          "num 1 updated",
				Features:      "num 1 updated",
				NodeCount:     10,
				LocationCount: 10,
				UserCount:     10,
				MonthExpire:   10,
				Description:   "num 1 updated: description",
			},
			name: "num 1 updated",
			err:  nil,
		},
		// {
		// 	in: model.Version{
		// 		GormCol: types.GormCol{
		// 			ID: 1003,
		// 		},
		// 		Name: "num 2 updated",
		// 	},
		// 	err: errors.New("legal_name is required"),
		// },
	}

	for _, v := range samples {
		result, err := versionServ.Save(v.in)
		if (v.err == nil && err != nil) || (v.err != nil && err == nil) {
			t.Errorf("ERROR FOR ::::%+v::: \nRETURNS :::%+v:::, \nIT SHOULD BE :::%+v:::", v.in, err, v.err)
		}

		if v.name != "" && v.name != result.Name {
			t.Errorf("Name are not same, it supposed to be %q, but it is %q", v.name, result.Name)
		}

	}

}

func TestVersionDelete(t *testing.T) {
	_, versionServ := initVersionTest()

	samples := []struct {
		id  types.RowID
		err error
	}{
		{
			id:  77,
			err: nil,
		},
		{
			id:  99999999,
			err: errors.New("record not found"),
		},
	}

	for _, v := range samples {
		_, err := versionServ.Delete(v.id)
		if (v.err == nil && err != nil) || (v.err != nil && err == nil) {
			t.Errorf("ERROR FOR ::::%+v::: \nRETURNS :::%+v:::, \nIT SHOULD BE :::%+v:::", v.id, err, v.err)
		}
	}
}

func TestVersionList(t *testing.T) {
	_, versionServ := initVersionTest()
	regularParam := getRegularParam("versions.id asc")
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
			count:  5,
		},
	}

	for _, v := range samples {
		data, err := versionServ.List(v.params)
		var count uint64
		var ok bool
		if count, ok = data["count"].(uint64); !ok {
			count = 0
		}
		if (v.err == nil && err != nil) || (v.err != nil && err == nil) || count != v.count {
			t.Errorf("FOR ::::%+v::: \nRETURNS :::%+v:::, \nIT SHOULD BE :::%+v:::", v.params, data["count"], v.count)
		}
	}
}

func TestVersionExcel(t *testing.T) {
	_, versionServ := initVersionTest()
	regularParam := getRegularParam("versions.id asc")

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
		data, err := versionServ.Excel(v.params)
		if (v.err == nil && err != nil) || (v.err != nil && err == nil) || uint64(len(data)) < v.count {
			t.Errorf("FOR ::::%+v::: \nRETURNS :::%+v:::, \nIT SHOULD BE :::%+v::: \nErr :::%+v:::",
				v.params, uint64(len(data)), v.count, err)
		}
	}
}
