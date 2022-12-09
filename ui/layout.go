package ui

import (
	"weather-app/apptype"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func Setup(window fyne.Window, state *apptype.State) {
	appContent := container.NewGridWithRows( //the container of all components
		3,
		state.Form,
		state.MainData,
		BuildWeatherDetails(state.Details),
	)

	r, _ := fyne.LoadResourceFromURLString("https://i.imgur.com/sz4jAgM.png") //loading image
	image := canvas.NewImageFromResource(r)
	image.FillMode = canvas.ImageFillContain
	image.Resize(fyne.NewSize(130, 130))
	image.Move(fyne.NewPos(0, 50))

	state.Form.SetSubmitFunc( //setting the submit function
		func(city, country string) {
			window.SetContent(container.NewGridWithRows( //showing the error image
				3,
				state.Form,
				container.NewGridWithColumns( //loading
					3,
					layout.NewSpacer(),
					container.NewWithoutLayout(image),
				),
			))

			err := UpdateState(city, country, state) //updating state by making the api request
			if err != nil {                          //error handling, showing only the form and an error message
				state.MainData.SetFields(err.Error(), "", "")
				for _, detail := range state.Details {
					detail.SetFields("", "")
				}
			}

			window.SetContent(appContent) //updating window
		},
	)

	r, _ = fyne.LoadResourceFromURLString("https://i.imgur.com/Dzhl10j.png") //app icon

	window.Resize(fyne.NewSize(445, 400))
	window.SetIcon(r)
	window.SetFixedSize(true)
	window.SetContent(appContent) //adding the ui elements to the window
}
