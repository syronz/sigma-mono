package table

import (
	"sigmamono/internal/core"
	"sigmamono/internal/types"
	"sigmamono/model"
	"sigmamono/repo"
	"sigmamono/service"
	"time"
)

// InsertCompanies for add required companies
func InsertCompanies(engine *core.Engine) {
	companyRepo := repo.ProvideCompanyRepo(engine)
	companyService := service.ProvideCompanyService(companyRepo)
	companies := []model.Company{
		{
			GormCol: types.GormCol{
				ID: 1001,
			},
			Name:       "Base",
			LegalName:  "Base",
			Key:        "08595aa978f4c6d40bb3f75950288117726588811287504fe9acc89ce752fe90290d2e2ec0c64efa4f8b9e94af7d574255bf591d8e5b02079c22385b2fcbfcf1ef7378bb2c3a1868a5e23a2ffdd6c0b3ccd4c959d7e63dbc18651521fa25823a54662a51c28b3df0ed431abaf72c27ac33ce338b9c3a3511d481e4048318b9c519ee36b5a40ba54f0b33f0f3b456a1efd774066eef92a3b6c6c043910888372a77c92174fb88cd8fa9ad81c6a3c2395e8c22238d4843cb8b9811d6fae7f0dc51ff6179181813cc1eb1682194bf3b85871458104c585c4276dc6fa4de14582f62d309f07d9e978cc8685b8423ce6719cc214a2ca19ac165f02c5c6a5cf0fe189e625144a39622fed128dd42da881d042041b330e8c4882145b9b676020217afc1c34a382dea459f21f81984fb93943ca976f01d678ff00899d5cb97240be8babc2549b8752a95a339e9b32b43caa1d74a16512b9bc986eea6daecdb4301c417fbd41866db813e54ae944f1ee699d89cf3bd723904c694b79c47da23ce5c16e36b9222b56d8f115f19d17f0808ca8c0329442b6776eb90701a9d16c19639e32a74f6bf39f309a220c4cb4fb28a306b280e9d5479e7841da228c4573c505c026a9a7994b6cf8cb8865f2ae7ee1de2468d8e8b024144e98f5a8da341dbc94c526334f9df16899320aefcc98e87c8d38c72b93bf50283bac6cb66ec99327fca03bc34196c53f91a7b85eb60386e4728462c593c0073d3215c8dc1a68ceba692294f9446179c4976279445b923ff8fe63931d6286e5efcf0b08513756773631261a96173f2eeda5470cf92c1443069d546b10deb6f139863d51af7404c72de8849c1831731cfe1e99374fc9e330ef7099a916451aeaa9e918e87a8f6d6182336c937833d978673252094c77749a221004c07abcf3b405db78dad97fc942dcf8ee2ac5970b95be37a944c6e6b13871462090864ea747bd96fee9a801369d089c9de490047bbf4c604ee9daa88525f949e8c83c3674dd4b306baffd9eaef7a3646e7ff8f40b6afec802d0e3da8b957c35d949ac8b510625aecf0734bfd5da0881a6c69eeffdc4e826fd287dafb20043896ff1edd395333fffb8776783bffde1ed5647973",
			Expiration: time.Now().AddDate(100, 0, 0),
			License:    "regular",
			Detail:     "",
			Phone:      "07505149171",
			Email:      "",
			Website:    "",
			Type:       "base",
			Code:       "",
		},
		{
			GormCol: types.GormCol{
				ID: 1002,
			},
			Name:       "for update 1",
			LegalName:  "for update 1",
			Key:        "08595aa978f4c",
			Expiration: time.Now().AddDate(100, 0, 0),
			License:    "for update 1",
			Detail:     "for update 1",
			Phone:      "for update 1",
			Email:      "for update 1",
			Website:    "for update 1",
			Type:       "base",
			Code:       "for update 1",
		},
		{
			GormCol: types.GormCol{
				ID: 1003,
			},
			Name:       "for update 2",
			LegalName:  "for update 2",
			Key:        "08595aa978f4c",
			Expiration: time.Now().AddDate(100, 0, 0),
			License:    "for update 2",
			Detail:     "for update 2",
			Phone:      "for update 2",
			Email:      "for update 2",
			Website:    "for update 2",
			Type:       "base",
			Code:       "for update 2",
		},
		{
			GormCol: types.GormCol{
				ID: 1004,
			},
			Name:       "for delete",
			LegalName:  "for delete",
			Key:        "for delete",
			Expiration: time.Now().AddDate(100, 0, 0),
			License:    "for delete",
			Detail:     "for delete",
			Phone:      "for delete",
			Email:      "for delete",
			Website:    "for delete",
			Type:       "base",
			Code:       "for delete",
		},
		{
			GormCol: types.GormCol{
				ID: 1005,
			},
			Name:       "for search 1",
			LegalName:  "for search 1",
			Key:        "for search 1",
			Expiration: time.Now().AddDate(100, 0, 0),
			License:    "for search 1",
			Detail:     "searchTerm1",
			Phone:      "for search 1",
			Email:      "for search 1",
			Website:    "for search 1",
			Type:       "base",
			Code:       "for search 1",
		},
		{
			GormCol: types.GormCol{
				ID: 1006,
			},
			Name:       "for search 2",
			LegalName:  "for search 2",
			Key:        "for search 2",
			Expiration: time.Now().AddDate(100, 0, 0),
			License:    "for search 2",
			Detail:     "searchTerm1",
			Phone:      "for search 2",
			Email:      "for search 2",
			Website:    "for search 2",
			Type:       "base",
			Code:       "for search 2",
		},
		{
			GormCol: types.GormCol{
				ID: 1007,
			},
			Name:       "for search 3",
			LegalName:  "for search 3",
			Key:        "for search 3",
			Expiration: time.Now().AddDate(100, 0, 0),
			License:    "for search 3",
			Detail:     "searchTerm1",
			Phone:      "for search 3",
			Email:      "for search 3",
			Website:    "for search 3",
			Type:       "base",
			Code:       "for search 3",
		},
		{
			GormCol: types.GormCol{
				ID: 1008,
			},
			Name:       "for search 4",
			LegalName:  "for search 4",
			Key:        "for search 4",
			Expiration: time.Now().AddDate(100, 0, 0),
			License:    "for search 4",
			Detail:     "searchTerm1",
			Phone:      "for search 4",
			Email:      "for search 4",
			Website:    "for search 4",
			Type:       "base",
			Code:       "for search 4",
		},
		{
			GormCol: types.GormCol{
				ID: 1009,
			},
			Name:       "for search 5",
			LegalName:  "for search 5",
			Key:        "for search 5",
			Expiration: time.Now().AddDate(100, 0, 0),
			License:    "for search 5",
			Detail:     "searchTerm1",
			Phone:      "for search 5",
			Email:      "for search 5",
			Website:    "for search 5",
			Type:       "base",
			Code:       "for search 5",
		},
		{
			GormCol: types.GormCol{
				ID: 1010,
			},
			Name:       "for search 6",
			LegalName:  "for search 6",
			Key:        "for search 6",
			Expiration: time.Now().AddDate(100, 0, 0),
			License:    "for search 6",
			Detail:     "searchTerm1",
			Phone:      "for search 6",
			Email:      "for search 6",
			Website:    "for search 6",
			Type:       "base",
			Code:       "for search 6",
		},
	}

	for _, v := range companies {
		if _, err := companyService.Save(v); err != nil {
			engine.ServerLog.Fatal(err)
		}
	}

}
