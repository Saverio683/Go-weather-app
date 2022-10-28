package ui

import (
	"weather-app/apptype"
	weatherdetail "weather-app/components/weather-detail"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func BuildWeatherDetails(state *apptype.State) *fyne.Container {
	weatherDetails := make([]fyne.CanvasObject, len(state.Details))

	for i := 0; i < len(state.Details); i++ {
		key := state.Details[i].Key
		val := state.Details[i].Value

		switch key {
		case "temp min":
		case "temp max":
			val += "°C"
		case "humidity":
			val += "%"
		case "pressure":
			val += "mb"
		case "sunrise":
		case "sunset":
		}

		w := weatherdetail.NewWeatherDetail(state.Details[i].Key, val, i)
		weatherDetails[i] = w
	}

	result := container.New(
		layout.NewGridLayout(3),
		weatherDetails...,
	)

	return result
}
