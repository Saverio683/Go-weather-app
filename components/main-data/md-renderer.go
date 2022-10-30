package maindata

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type MainDataRenderer struct {
	title, temp        *canvas.Text
	image              *canvas.Image
	wrapper, container *fyne.Container
}

func (renderer *MainDataRenderer) MinSize() fyne.Size {
	return renderer.container.Size()
}
func (renderer *MainDataRenderer) Layout(size fyne.Size) {
}
func (renderer *MainDataRenderer) Refresh() {
	canvas.Refresh(renderer.image)
	canvas.Refresh(renderer.temp)
	canvas.Refresh(renderer.title)
	canvas.Refresh(renderer.wrapper)
	canvas.Refresh(renderer.container)
}
func (renderer *MainDataRenderer) Objects() []fyne.CanvasObject {
	return renderer.container.Objects
}
func (renderer *MainDataRenderer) Destroy() {}
