package service

import (
	"errors"
	"sigmamono/internal/core"
	r "sigmamono/internal/core/access/resource"
	"sigmamono/internal/param"
	"sigmamono/internal/types"
	"sigmamono/model"
	"sigmamono/repo"
	"sigmamono/test/kernel"
	"testing"
)

func initRoleTest() (engine *core.Engine, roleServ RoleServ) {
	logQuery, debugLevel := initServiceTest()
	engine = kernel.StartMotor(logQuery, debugLevel)
	roleServ = ProvideRoleService(repo.ProvideRoleRepo(engine))

	return
}

func TestRoleCreate(t *testing.T) {
	_, roleServ := initRoleTest()
	regularParam := getRegularParam("roles.id asc")
	paramWithZeroCompanyID := regularParam
	paramWithZeroCompanyID.CompanyID = 0

	samples := []struct {
		in     model.Role
		params param.Param
		err    error
	}{
		{
			in: model.Role{
				Name:        "created 1",
				Resources:   r.SupperAccess,
				Description: "created 1",
			},
			params: regularParam,
			err:    nil,
		},
		{
			in: model.Role{
				Name:        "created 1",
				Resources:   r.SupperAccess,
				Description: "created 1",
			},
			params: regularParam,
			err:    errors.New("duplicate"),
		},
		{
			in: model.Role{
				Name:      "minimum fields",
				Resources: r.SupperAccess,
			},
			params: regularParam,
			err:    nil,
		},
		{
			in: model.Role{
				Name:        "long name: big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name big name",
				Resources:   r.SupperAccess,
				Description: "created 2",
			},
			params: regularParam,
			err:    errors.New("data too long for name"),
		},
		{
			in: model.Role{
				Resources:   r.SupperAccess,
				Description: "created 3",
			},
			params: regularParam,
			err:    errors.New("name is required"),
		},
		{
			in: model.Role{
				Name:      "created 4",
				Resources: r.SupperAccess,
			},
			params: paramWithZeroCompanyID,
			err:    errors.New("company_id is not exist"),
		},
	}

	for _, v := range samples {
		_, err := roleServ.Create(v.in, v.params)
		if (v.err == nil && err != nil) || (v.err != nil && err == nil) {
			t.Errorf("\nERROR FOR :::%+v::: \nRETURNS :::%+v:::, \nIT SHOULD BE :::%+v:::", v.in, err, v.err)
		}
	}

}

func TestRoleUpdate(t *testing.T) {
	_, roleServ := initRoleTest()

	samples := []struct {
		in  model.Role
		err error
	}{
		{
			in: model.Role{
				FixedCol: types.FixedCol{
					ID: 1001101000000005,
				},
				CompanyID:   1010,
				NodeCode:    101,
				Name:        "num 1 update",
				Resources:   r.SupperAccess,
				Description: "num 1 update",
			},
			err: nil,
		},
		{
			in: model.Role{
				FixedCol: types.FixedCol{
					ID: 1001101000000006,
				},
				CompanyID:   1010,
				NodeCode:    101,
				Name:        "num 2 update",
				Description: "num 2 update",
			},
			err: errors.New("resources are required"),
		},
	}

	for _, v := range samples {
		_, err := roleServ.Save(v.in)
		if (v.err == nil && err != nil) || (v.err != nil && err == nil) {
			t.Errorf("ERROR FOR ::::%+v::: \nRETURNS :::%+v:::, \nIT SHOULD BE :::%+v:::", v.in, err, v.err)
		}
	}

}

func TestRoleDelete(t *testing.T) {
	_, roleServ := initRoleTest()
	regularParam := getRegularParam("roles.id asc")

	samples := []struct {
		id  types.RowID
		err error
	}{
		{
			id:  1001101000000007,
			err: nil,
		},
		{
			id:  99999999,
			err: errors.New("record not found"),
		},
		{
			id:  1002101000000001,
			err: errors.New("don't have permission, out of scope"),
		},
	}

	for _, v := range samples {
		_, err := roleServ.Delete(v.id, regularParam)
		if (v.err == nil && err != nil) || (v.err != nil && err == nil) {
			t.Errorf("ERROR FOR ::::%+v::: \nRETURNS :::%+v:::, \nIT SHOULD BE :::%+v:::", v.id, err, v.err)
		}
	}
}

func TestRoleList(t *testing.T) {
	_, roleServ := initRoleTest()
	regularParam := getRegularParam("roles.id asc")
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
			count:  3,
		},
	}

	for _, v := range samples {
		data, err := roleServ.List(v.params)
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

func TestRoleExcel(t *testing.T) {
	_, roleServ := initRoleTest()
	regularParam := getRegularParam("roles.id asc")

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
		data, err := roleServ.Excel(v.params)
		if (v.err == nil && err != nil) || (v.err != nil && err == nil) || uint64(len(data)) < v.count {
			t.Errorf("FOR ::::%+v::: \nRETURNS :::%+v:::, \nIT SHOULD BE :::%+v::: \nErr :::%+v:::",
				v.params, uint64(len(data)), v.count, err)
		}
	}
}
