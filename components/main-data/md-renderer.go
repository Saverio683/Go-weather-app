package maindata

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type MainDataRenderer struct {
	title, temp        *canvas.Text
	imgLink            string
	image              *canvas.Image
	wrapper, container *fyne.Container
}

func (renderer *MainDataRenderer) MinSize() fyne.Size {
	return renderer.container.Size()
}
func (renderer *MainDataRenderer) Layout(size fyne.Size) {
}
func (renderer *MainDataRenderer) Refresh() {

	renderer.title.TextSize = 18
	renderer.temp.TextSize = 50

	renderer.title.Move(fyne.NewPos(10, 20))

	renderer.image.FillMode = canvas.ImageFillContain
	renderer.image.Resize(fyne.NewSize(100, 100))
	renderer.image.Move(fyne.NewPos(310, 0))
	renderer.wrapper.Move(fyne.NewPos(10, 60))

	canvas.Refresh(renderer.container)
}
func (renderer *MainDataRenderer) Objects() []fyne.CanvasObject {
	return make([]fyne.CanvasObject, 0)
}
func (renderer *MainDataRenderer) Destroy() {}
