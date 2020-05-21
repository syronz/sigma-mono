package table

import (
	"sigmamono/internal/core"
	"sigmamono/internal/types"
	"sigmamono/model"
	"sigmamono/repo"
	"sigmamono/service"
)

// InsertStations for add required stations
func InsertStations(engine *core.Engine) {
	stationRepo := repo.ProvideStationRepo(engine)
	stationService := service.ProvideStationService(stationRepo)
	stations := []model.Station{
		{
			GormCol: types.GormCol{
				ID: 1,
			},
			CompanyID:   1001,
			CompanyName: "Base",
			NodeCode:    101,
			NodeName:    "Base-root",
			Key:         "306a3a8c78fec6845ab3a20c5272d7437d6589d6138e0648b7a1999ce155ac912a00282dc0941dfd1ad19394f22b591502be0e1cd95b00509c276a072e9da9a1bd767eed70684a3cfbea3d7cab8697e09a83cf5bd6eb68b94b64442fff708760573b2c0595dd66f0e71348bdf67175aa66c934db9b3c3747d886b459864bbf9445ec63e4a50cf2480e69f2f6b159f6e88026013de794f5b3c1c744c108d2672c77cc7778aada9bd9fdabd7c8ad9238018b732a881847c98b9544d1f8b2f2dc56a23528444b4a974ab13b2392eb31dfdc130d151f560d1576886aaa8947592f378850a27ecd97da996c59852e916b4cc9754f2ba7c8c167f1775b6258a2fd40ce350017ffc422fcd02ed7428b8c19042510e764e0c3882314b8b228545517afcf9712397eea47c829f41f8caf939b30ab26a24031dfa208cb85caca210eb2efbe234ee8702dc6fa3eeab32d18cda5821d4a557b9c9e8deea7d9bede1906c547fa871e67de816f02a1c2491bb2c9dc99fdbd78670795c3bc9741dd70ce0c16e235c02cb2648a135b1a862d530f9885082a432b3126ea93211fc81499976fb4292ef1ee3da50fff7193c11ce1db646e760c9c587fefd746a522c50737535f573fc07fc0bc9bd9e8845f2bb6bb1ab140d08c8c014113e9dc0cd5a5478ccc4d0d3938a58a11d9c723aea4c98ed59285de75ee35f651d2bbc09b31efc9602bcd5fe8361c6950f8187ad0ed303f681d24477e5030522482745cd5cba3d8bef5c7204e93421192492d739317bc70fa8fe13163d67b3f0faaf5e78c4b733477631261a73370a1ef805e2298c5c61762398712b65bb938139d668f49f0184b70858e4b93d3183498b6be957cf29b3654ab009c923402fbafc990dd82aff7d9482260c86f816a9c8124262797ca721bf421014b51fecd3a1c0eb0d7ac92f9922f998ab4fd0b76e80db079c21e6a6816824d665d593fba237c8e6feace844e3584d8ca8b4c004cb9f3cc53eb95ff8f0e5bc19ad48492641cdbb251b8f48db6ee7a314db6a08647e5fde9827d5c36abbe55cf599a9498b0116453eda5224afe0cabd84d3a35e3aadf4b826e97d697",
			MachineID:   "base",
			Detail:      "",
		},
	}

	for _, v := range stations {
		if _, err := stationService.Save(v); err != nil {
			engine.ServerLog.Fatal(err)
		}
	}

}
