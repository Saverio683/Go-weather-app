package ui

func Setup(app *AppInit) {
	weatherDetailsContainer := BuildWeatherDetails(app.State)

	app.Window.SetContent(weatherDetailsContainer) //adding the ui elements to the window
}
