package ui

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func onSubmitFunc(city, country string) {
	fmt.Println("submit")
}

func GenerateCustomForm() *fyne.Container {
	//creating widgets
	cityEntry := widget.NewEntry()
	countrEntry := widget.NewEntry()
	button := widget.NewButton("Search", func() { onSubmitFunc(cityEntry.Text, countrEntry.Text) })

	//set placeholder in text input
	cityEntry.SetPlaceHolder("City")
	countrEntry.SetPlaceHolder("Country (optional)")

	//setting size
	cityEntry.Resize(fyne.NewSize(250, 40))
	countrEntry.Resize(fyne.NewSize(150, 40))
	button.Resize(fyne.NewSize(160, 40))

	//positioning
	cityEntry.Move(fyne.NewPos(10, 10))
	countrEntry.Move(fyne.NewPos(cityEntry.Size().Width+25, 10))
	button.Move(fyne.NewPos(10, 25))

	customForm := container.New(
		layout.NewVBoxLayout(),
		container.NewWithoutLayout(
			cityEntry,
			countrEntry,
		),
		container.NewWithoutLayout(
			button,
		),
	)

	return customForm
}

func ShowApiResults(city, temp string) *fyne.Container {
	text := canvas.NewText("Current weather in "+city, color.White)
	text.TextSize = 18
	text.Move(fyne.NewPos(10, 0))

	temperature := canvas.NewText(temp, color.White)
	temperature.TextSize = 50

	//sample image
	r, _ := fyne.LoadResourceFromURLString("http://openweathermap.org/img/wn/09d@2x.png")
	image := canvas.NewImageFromResource(r)
	image.FillMode = canvas.ImageFillContain
	image.Resize(fyne.NewSize(100, 100))

	//wrap the components to have them on the same line
	tempAndImg := container.NewWithoutLayout(
		temperature,
		image,
	)

	tempAndImg.Move(fyne.NewPos(10, 40))
	image.Move(fyne.NewPos(300, 0))

	apiResults := container.NewWithoutLayout(
		text,
		tempAndImg,
	)

	return apiResults
}
