package main

import (
	"weather-app/apptype"
	maindata "weather-app/components/main-data"
	weatherdetail "weather-app/components/weather-detail"
	"weather-app/ui"

	"fyne.io/fyne/v2/app"
)

/* type apiData struct {
	err     bool
	loading bool
	data    utils.RespBody
}

func handleSubmit(city, ct string) *apiData {
	resp, err := utils.MakeRequest(city, ct)

	if err != nil {
		return &apiData{err: true, loading: false, data: utils.RespBody{}}
	} else {
		return &apiData{err: false, loading: false, data: utils.RespBody{resp.Weather, resp.Main, resp.Sys}}
	}
}*/

func main() {
	weatherApp := app.New()
	window := weatherApp.NewWindow("Weather app")

	details := make([]*weatherdetail.WeatherDetail, 6)

	for i := 0; i < 6; i++ {
		details[i] = weatherdetail.NewWeatherDetail("", "", i)
	}

	state := apptype.State{
		MainData: maindata.NewMainData("Please search something", "", ""),
		Details:  details,
		Loading:  false,
		Error:    false,
	}

	appInit := ui.AppInit{
		Window: window,
		State:  &state,
	}

	ui.Setup(&appInit)

	appInit.Window.ShowAndRun()
	/*
		app := app.New()
		window := app.NewWindow("Current Weather")

		customForm := ui.CreateCustomForm(handleSubmit)
		apiResults := ui.CreateMainData("Cinisi", "20 °C")
		details := ui.CreateDetails(15, 18, 34, 1020, "6:12", "20:09")

		wrapper := container.NewGridWithRows(
			3,
			customForm,
			apiResults,
			details,
		)

		window.Resize(fyne.NewSize(445, 400))
		window.SetContent(wrapper)

		window.ShowAndRun()
	*/
}
