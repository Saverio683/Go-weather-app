package ui

import (
	"weather-app/apptype"
	weatherdetail "weather-app/components/weather-detail"

	"fyne.io/fyne/v2"
)

type AppInit struct {
	Window  fyne.Window
	State   *apptype.State
	Details []*weatherdetail.WeatherDetail
}
