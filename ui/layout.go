package ui

import (
	"weather-app/apptype"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func Setup(app *apptype.AppInit) {
	appContent := container.NewGridWithRows(
		3,
		app.State.Form,
		app.State.MainData,
		BuildWeatherDetails(app.State.Details),
	)

	app.State.Form.SetSubmitFunc(
		func(city, country string) {
			app.Window.SetContent(container.NewGridWithRows(
				3,
				app.State.Form,
				CreateSpinner(),
			))
			_, err := UpdateWindow(city, country, app)

			if err != nil {
				for i, detail := range app.State.Details {
					detail.SetFields("", "", i)
				}
				app.State.MainData.SetFields("Ops... something went wrong", "", "")
			}

			app.Window.SetContent(appContent)
		},
	)

	app.Window.Resize(fyne.NewSize(445, 400))
	app.Window.SetContent(appContent) //adding the ui elements to the window
}
