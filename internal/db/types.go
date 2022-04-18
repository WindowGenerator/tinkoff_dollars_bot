package db

import (
	enums "github.com/WidowGenerator/tinkoff_dollars_bot/internal/db/enums"
	interfaces "github.com/WidowGenerator/tinkoff_dollars_bot/internal/interfaces"
)

type BoundsWithZoom struct {
	Bounds *interfaces.Bounds
	Zoom   uint32
}

var Bank2BankId = map[enums.Bank]string{
	enums.BankTinkoff:    "tcs",
	enums.BankSber:       "11242",
	enums.BankVTB:        "11249",
	enums.BankAlpha:      "11250",
	enums.BankRaiffaizen: "11241",
	enums.BankGasProm:    "11371",
}

var ALL_BANKS = getAllBanks()

func getAllBanks() []string {
	v := make([]string, 0, len(Bank2BankId))

	for _, value := range Bank2BankId {
		v = append(v, value)
	}
	return v
}

var Cities2GS = map[enums.City]BoundsWithZoom{
	enums.CityYekaterinburg: {
		Bounds: &interfaces.Bounds{
			BottomLeft: interfaces.GeographicCoodrinates{
				Lat: 56.63827841259033,
				Lng: 60.33014140617972,
			},
			TopRight: interfaces.GeographicCoodrinates{
				Lat: 57.005757763347255,
				Lng: 60.9948142577422,
			},
		},
		Zoom: 11,
	},
	enums.CityMoscow: {
		Bounds: &interfaces.Bounds{
			BottomLeft: interfaces.GeographicCoodrinates{
				Lat: 56.63827841259033,
				Lng: 60.33014140617972,
			},
			TopRight: interfaces.GeographicCoodrinates{
				Lat: 57.005757763347255,
				Lng: 60.9948142577422,
			},
		},
		Zoom: 9,
	},
}
