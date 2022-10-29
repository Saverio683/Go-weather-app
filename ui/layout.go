package ui

import (
	"weather-app/components/form"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func Setup(app *AppInit) {
	weatherDetails := BuildWeatherDetails(app.State)
	form := form.NewForm(app.State.City, app.State.Country, func(a, b string) {})

	container := container.NewGridWithRows(
		3,
		form,
		layout.NewSpacer(),
		weatherDetails,
	)

	app.Window.Resize(fyne.NewSize(445, 400))
	app.Window.SetContent(container) //adding the ui elements to the window
}
