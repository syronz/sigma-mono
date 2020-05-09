package table

import (
	"sigmamono/internal/core"
	"sigmamono/internal/types"
	"sigmamono/model"
	"sigmamono/repo"
	"sigmamono/service"
)

// InsertSettings for add required settings
func InsertSettings(engine *core.Engine) {
	settingRepo := repo.ProvideSettingRepo(engine)
	settingService := service.ProvideSettingService(settingRepo)
	settings := []model.Setting{
		{
			FixedCol: types.FixedCol{
				ID: 1001101000000001,
			},
			CompanyID:   1001,
			Property:    "buy_invoice_location_selection_level",
			Value:       "invoice",
			Type:        "string",
			Description: "select destination and source in level of item or invoice",
		},
		{
			FixedCol: types.FixedCol{
				ID: 1001101000000002,
			},
			CompanyID:   1001,
			Property:    "sell_invoice_location_selection_level",
			Value:       "item",
			Type:        "string",
			Description: "select destination and source in level of item or invoice",
		},
		{
			FixedCol: types.FixedCol{
				ID: 1001101000000003,
			},
			CompanyID:   1001,
			Property:    "transfer_invoice_location_selection_level",
			Value:       "invoice",
			Type:        "string",
			Description: "select destination and source in level of item or invoice",
		},
		{
			FixedCol: types.FixedCol{
				ID: 1001101000000004,
			},
			CompanyID:   1001,
			Property:    "default_language",
			Value:       "ku",
			Type:        "string",
			Description: "in case of user JWT not specified this value has been used",
		},
		{
			FixedCol: types.FixedCol{
				ID: 1001101000000005,
			},
			CompanyID:   1001,
			Property:    "company_logo",
			Value:       "public/logo.png",
			Type:        "string",
			Description: "path of logo, if branch logo won’t defined use this logo for invoices",
		},
		{
			FixedCol: types.FixedCol{
				ID: 1001101000000006,
			},
			CompanyID:   1001,
			Property:    "inventory_method",
			Value:       "instant",
			Type:        "string",
			Description: "shipping/instant, there is two way for inventory, the first one is locking system, which is useful if we lock the items then transfer them as out, the other one as soon as inventory applied the inventory will saved",
		},
		{
			FixedCol: types.FixedCol{
				ID: 1001101000000007,
			},
			CompanyID:   1001,
			Property:    "shipping_level",
			Value:       "invoice",
			Type:        "string",
			Description: "invoice/item, it is used for affect the inventory",
		},
		{
			FixedCol: types.FixedCol{
				ID: 1001101000000008,
			},
			CompanyID:   1001,
			Property:    "invoice_number_pattern",
			Value:       "location_year_series",
			Type:        "string",
			Description: "location_year_series, location_series, series, year_series, fullyear_series, location_fullyear_series",
		},
		{
			FixedCol: types.FixedCol{
				ID: 1001101000000009,
			},
			CompanyID:   1001,
			Property:    "shared_warehouse",
			Value:       "no",
			Type:        "bool",
			Description: "shared warehouse mean that a location can has a access to other location’s inventory. In case we choose true, for each branch we should define location_priority. In case of false each branch just has access to it’s inventory",
		},
		{
			FixedCol: types.FixedCol{
				ID: 1001101000000010,
			},
			CompanyID:   1001,
			Property:    "default_discount_account",
			Value:       "100110100000000032",
			Type:        "rowid",
			Description: "use that account for collect all discounts",
		},
		{
			FixedCol: types.FixedCol{
				ID: 1001101000000011,
			},
			CompanyID:   1001,
			Property:    "default_income_account",
			Value:       "100110100000000033",
			Type:        "rowid",
			Description: "",
		},
		{
			FixedCol: types.FixedCol{
				ID: 1001101000000012,
			},
			CompanyID:   1001,
			Property:    "default_cost_of_goods_sold",
			Value:       "100110100000000034",
			Type:        "string",
			Description: "",
		},
		{
			FixedCol: types.FixedCol{
				ID: 1001101000000013,
			},
			CompanyID:   1001,
			Property:    "balance_sheet_asset_depth",
			Value:       "3",
			Type:        "int",
			Description: "it is used for the total calculation of the balance sheet",
		},
		{
			FixedCol: types.FixedCol{
				ID: 1001101000000014,
			},
			CompanyID:   1001,
			Property:    "balance_sheet_equity_depth",
			Value:       "3",
			Type:        "int",
			Description: "",
		},
		{
			FixedCol: types.FixedCol{
				ID: 1001101000000015,
			},
			CompanyID:   1001,
			Property:    "balance_sheet_liability_depth",
			Value:       "3",
			Type:        "int",
			Description: "",
		},
		{
			FixedCol: types.FixedCol{
				ID: 1001101000000016,
			},
			CompanyID:   1001,
			Property:    "variation_consume",
			Value:       "fifo",
			Type:        "string",
			Description: "fifo/lifo, it is used for type of consuming from variation_qty",
		},
		{
			FixedCol: types.FixedCol{
				ID: 1001101000000017,
			},
			CompanyID:   1001,
			Property:    "accept_negative_qty_for_locations",
			Value:       "no",
			Type:        "bool",
			Description: "yes/no, it is for describing negative items for branch, useful for cases ordering from third parties happened",
		},
		{
			FixedCol: types.FixedCol{
				ID: 1001101000000018,
			},
			CompanyID:   1001,
			Property:    "default_currency",
			Value:       "$",
			Type:        "string",
			Description: "usd/iqd and all different currencies, it is string and should be exist in the currencies code",
		},
		{
			FixedCol: types.FixedCol{
				ID: 1001101000000019,
			},
			CompanyID:   1001,
			Property:    "max_different_sync_time_for_edit",
			Value:       "24",
			Type:        "int",
			Description: "difference time for letting user edit element in hour",
		},
	}

	for _, v := range settings {
		if _, err := settingService.Save(v); err != nil {
			engine.ServerLog.Fatal(err)
		}

	}

}
