package ui

import (
	weatherdetail "weather-app/components/weather-detail"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

// the function creates the container that wraps the weather details
func BuildWeatherDetails(details []*weatherdetail.WeatherDetail) *fyne.Container {
	weatherDetails := make([]fyne.CanvasObject, len(details))

	for i, detail := range details {
		weatherDetails[i] = detail
	}

	result := container.New(
		layout.NewGridLayout(3),
		weatherDetails...,
	)

	return result
}
