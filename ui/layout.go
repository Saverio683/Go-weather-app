package ui

import (
	"weather-app/components/form"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func Setup(app *AppInit) {
	form := form.NewForm(
		func(city, country string) {
			UpdateWindow(city, country, app)
		},
	)

	appContainer := container.NewGridWithRows(
		3,
		form,
		app.State.MainData,
		BuildWeatherDetails(app.State.Details),
	)

	/* if app.State.City != "" {
		appContainer = container.NewGridWithRows(
			3,
			form,
			layout.NewSpacer(),
			layout.NewSpacer(),
		)
	} else {
		message := canvas.NewText("Please search something", color.White)
		message.Move(fyne.NewPos(125, 200))
		message.TextStyle.Bold = true

		if app.State.Error {
			message.Text = "Ops... something went wrong"
		}

		appContainer = container.NewWithoutLayout(
			form,
			message,
			layout.NewSpacer(),
		)
	} */

	app.Window.Resize(fyne.NewSize(445, 400))
	app.Window.SetContent(appContainer) //adding the ui elements to the window
}
