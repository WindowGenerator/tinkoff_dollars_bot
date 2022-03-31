package tasks

import (
	interfaces "github.com/WidowGenerator/tinkoff_dollars_bot/internal/interfaces"
)

var city2Cache = map[string]interfaces.TinkoffATM{}

func notify() {

}

func updateATMDiff(newAtmData *interfaces.TinkoffATM, oldAtmData **interfaces.TinkoffATM) {

}

// func pipeline(city string) {
// 	tinkoffATM := interfaces.NewTinkkoffAtmAPI(testValues.Bounds, testValues.Filters, testValues.Zoom)

// 	err := tinkoffATM.GetATMs()

// 	if _, ok := city2Cache[city]; !ok {
// 		city2Cache[city] =
// 	}
// }
