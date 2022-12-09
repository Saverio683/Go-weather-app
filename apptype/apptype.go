package apptype

import (
	"weather-app/components/form"
	maindata "weather-app/components/main-data"
	weatherdetail "weather-app/components/weather-detail"
)

type State struct {
	MainData *maindata.MainData
	Details  []*weatherdetail.WeatherDetail
	Form     *form.Form
}
