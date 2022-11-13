package maindata

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type MainDataRenderer struct {
	mainData    *MainData
	title, temp *canvas.Text
	image       *canvas.Image
	container   *fyne.Container
}

func (renderer *MainDataRenderer) MinSize() fyne.Size {
	return renderer.container.MinSize()
}
func (renderer *MainDataRenderer) Layout(size fyne.Size) {
}
func (renderer *MainDataRenderer) Refresh() {
	renderer.title.Text = renderer.mainData.Title
	renderer.temp.Text = renderer.mainData.Temp

	r, _ := fyne.LoadResourceFromURLString(renderer.mainData.ImgLink)
	renderer.image.Resource = r

	go canvas.Refresh(renderer.image)
	go canvas.Refresh(renderer.temp)
	go canvas.Refresh(renderer.title)
}
func (renderer *MainDataRenderer) Objects() []fyne.CanvasObject {
	return renderer.container.Objects
}
func (renderer *MainDataRenderer) Destroy() {}
