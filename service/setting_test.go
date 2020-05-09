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

func initSettingTest() (engine *core.Engine, settingServ SettingServ) {
	logQuery, debugLevel := initServiceTest()
	engine = kernel.StartMotor(logQuery, debugLevel)
	settingServ = ProvideSettingService(repo.ProvideSettingRepo(engine))

	return
}

func TestSettingUpdate(t *testing.T) {
	_, settingServ := initSettingTest()

	samples := []struct {
		in  model.Setting
		err error
	}{
		{
			in: model.Setting{
				FixedCol: types.FixedCol{
					ID: 1001101000000020,
				},
				CompanyID:   1001,
				Property:    "num 1 updated",
				Value:       "num 1 updated",
				Type:        "num 1 updated",
				Description: "num 1 updated",
			},
			err: nil,
		},
		{
			in: model.Setting{
				FixedCol: types.FixedCol{
					ID: 1001101000000021,
				},
				CompanyID:   1001,
				Value:       "num 2 updated",
				Type:        "num 2 updated",
				Description: "num 2 updated",
			},
			err: errors.New("property is required"),
		},
	}

	for _, v := range samples {
		_, err := settingServ.Save(v.in)
		if (v.err == nil && err != nil) || (v.err != nil && err == nil) {
			t.Errorf("ERROR FOR ::::%+v::: \nRETURNS :::%+v:::, \nIT SHOULD BE :::%+v:::", v.in, err, v.err)
		}
	}

}

func TestSettingList(t *testing.T) {
	_, settingServ := initSettingTest()
	regularParam := getRegularParam("settings.id asc")
	// regularParam.Search = "searchTerm1"
	regularParam.CompanyID = 1001

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
			count:  21,
		},
	}

	for _, v := range samples {
		data, err := settingServ.List(v.params)
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

func TestSettingExcel(t *testing.T) {
	_, settingServ := initSettingTest()
	regularParam := getRegularParam("settings.id asc")

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
		data, err := settingServ.Excel(v.params)
		if (v.err == nil && err != nil) || (v.err != nil && err == nil) || uint64(len(data)) < v.count {
			t.Errorf("FOR ::::%+v::: \nRETURNS :::%+v:::, \nIT SHOULD BE :::%+v::: \nErr :::%+v:::",
				v.params, uint64(len(data)), v.count, err)
		}
	}
}
