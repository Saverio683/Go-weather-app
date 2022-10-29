package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func CreateMainData(city, temp string) *fyne.Container {
	text := canvas.NewText("Current weather in "+city, color.White)
	text.TextSize = 18
	text.Move(fyne.NewPos(10, 20))

	temperature := canvas.NewText(temp, color.White)
	temperature.TextSize = 50

	//sample image
	r, _ := fyne.LoadResourceFromURLString("http://openweathermap.org/img/wn/09d@2x.png")
	image := canvas.NewImageFromResource(r)
	image.FillMode = canvas.ImageFillContain
	image.Resize(fyne.NewSize(100, 100))
	image.Move(fyne.NewPos(310, 0))

	//wrap the components to have them on the same line
	tempAndImg := container.NewWithoutLayout(
		temperature,
		image,
	)
	tempAndImg.Move(fyne.NewPos(10, 60))

	result := container.NewWithoutLayout(
		text,
		tempAndImg,
	)

	return result
}
