package main

import (
	"fmt"
	"weather-app/ui"
	"weather-app/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	prova, err := utils.MakeRequest("palermo", "it")

	if err != nil {
		panic(err)
	} else {
		fmt.Println(prova.Weather[0])
		fmt.Println(prova.Main)
		fmt.Println(prova.Sys)
		fmt.Println(prova.Dt)
	}

	app := app.New()
	window := app.NewWindow("Current Weather")

	customForm := ui.GenerateCustomForm()
	apiResults := ui.ShowApiResults("Cinisi", "20 °C")

	wrapper := container.NewGridWithRows(
		2,
		customForm,
		apiResults,
	)

	window.Resize(fyne.NewSize(445, 400))
	window.SetContent(wrapper)

	window.ShowAndRun()
}
