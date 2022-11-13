package apptype

import (
	"weather-app/components/form"
	maindata "weather-app/components/main-data"
	weatherdetail "weather-app/components/weather-detail"

	"fyne.io/fyne/v2"
)

type State struct {
	MainData *maindata.MainData
	Details  []*weatherdetail.WeatherDetail
	Form     *form.Form
	Loading  bool
	Error    bool
}

type AppInit struct {
	Window fyne.Window
	State  *State
}
