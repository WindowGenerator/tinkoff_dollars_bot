package commands

import (
	"testing"

	"github.com/WidowGenerator/tinkoff_dollars_bot/internal/db/enums"
	interfaces "github.com/WidowGenerator/tinkoff_dollars_bot/internal/interfaces"
	assert "github.com/stretchr/testify/assert"
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

func TestSimpleParseArguments(t *testing.T) {
	fs, _ := getAtmParser()

	err := fs.Parse([]string{})

	if err != nil {
		t.Errorf("Failed parsing nullable args: %v", err)
	}
}

func TestParseCityArgument(t *testing.T) {
	fs, inputAtmInfo := getAtmParser()

	err := fs.Parse([]string{"--city", "Yekaterinburg"})

	if err != nil {
		t.Errorf("Failed parsing city: %v", err)
	}

	assert.Equal(t, *inputAtmInfo.City, enums.CityYekaterinburg)
}

func TestParseEmptyCityArgument(t *testing.T) {
	fs, inputAtmInfo := getAtmParser()

	err := fs.Parse([]string{})

	if err != nil {
		t.Errorf("Failed parsing city: %v", err)
	}

	assert.Equal(t, *inputAtmInfo.City, enums.CityMoscow)
}

func TestParseCurrencyArgument(t *testing.T) {
	fs, inputAtmInfo := getAtmParser()

	err := fs.Parse([]string{"--currency", "USD"})

	if err != nil {
		t.Errorf("Failed parsing nullable args: %v", err)
	}

	assert.Equal(t, *inputAtmInfo.Currency, enums.CurrencyUSD)
}

func TestParseEmptyCurrencyArgument(t *testing.T) {
	fs, inputAtmInfo := getAtmParser()

	err := fs.Parse([]string{})

	if err != nil {
		t.Errorf("Failed parsing nullable args: %v", err)
	}

	assert.Equal(t, *inputAtmInfo.Currency, enums.CurrencyAny)
}
