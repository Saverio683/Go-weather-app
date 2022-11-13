package main

import (
	"weather-app/apptype"
	"weather-app/components/form"
	maindata "weather-app/components/main-data"
	weatherdetail "weather-app/components/weather-detail"
	"weather-app/ui"

	"fyne.io/fyne/v2/app"
)

func main() {
	weatherApp := app.New()
	window := weatherApp.NewWindow("Weather app")

	details := make([]*weatherdetail.WeatherDetail, 6)

	for i := 0; i < 6; i++ {
		details[i] = weatherdetail.NewWeatherDetail("", "", i)
	}

	state := apptype.State{
		Form:     form.NewForm(func(city, country string) {}),
		MainData: maindata.NewMainData("", "", ""),
		Details:  details,
		Loading:  false,
		Error:    false,
	}

	appInit := apptype.AppInit{
		Window: window,
		State:  &state,
	}

	ui.Setup(&appInit)

	appInit.Window.ShowAndRun()
}
