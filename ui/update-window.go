package ui

import (
	"fmt"
	"weather-app/apptype"
	"weather-app/utils"
)

func UpdateWindow(city, country string, app *AppInit) {
	res, err := utils.GetApiResults(city, country)

	if err != nil {
	} else {
		apptype.SetDetails(res.WeatherDetails, app.State)
		apptype.SetMainData(res.MainData, app.State)

		fmt.Println(000)
	}

	/* var appContainer *fyne.Container

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
	) */

	//app.Window.SetContent(appContainer)
}
