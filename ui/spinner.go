package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func CreateSpinner() *canvas.Image {
	r, _ := fyne.LoadResourceFromURLString("https://i.imgur.com/kEF8E2K.gif")
	spinner := canvas.NewImageFromResource(r)

	spinner.FillMode = canvas.ImageFillContain
	spinner.Resize(fyne.NewSize(100, 100))
	spinner.Move(fyne.NewPos(310, -5))

	return spinner
}
