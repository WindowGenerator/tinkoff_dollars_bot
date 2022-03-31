package commands

import (
	"fmt"
	"testing"

	interfaces "github.com/WidowGenerator/tinkoff_dollars_bot/internal/interfaces"
)

var testValues = interfaces.TinkoffATMBody{
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
	Filters: &interfaces.Filters{
		ShowUnavailable: true,
		Currencies:      []string{"USD"},
	},
	Zoom: 11,
}

func TestGetAllAtms(t *testing.T) {
}

func TestSimpleParseArguments(t *testing.T) {
	tinkoffATMBody, err := parseGetAllAtmsArgs("")

	fs.Parse([]string{})
	fmt.Printf("{city: %v}\n\n", *inputAtmInfo.City)

	fs.Parse([]string{"--city", "Yerevan"})
	fmt.Printf("{city: %v}\n\n", *inputAtmInfo.City)

	fs.Parse([]string{"--currencies", "USD"})
	fmt.Println(*inputAtmInfo.Currencies)
	t.Fail()
}
