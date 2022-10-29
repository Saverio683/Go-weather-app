package weatherdetail

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type WeatherDetailRenderer struct {
	key, val *canvas.Text
	index    int
	wrapper  *fyne.Container
}

func (renderer *WeatherDetailRenderer) MinSize() fyne.Size {
	return renderer.wrapper.Size()
}
func (renderer *WeatherDetailRenderer) Layout(size fyne.Size) {
}
func (renderer *WeatherDetailRenderer) Refresh() {
	canvas.Refresh(renderer.key)
	canvas.Refresh(renderer.val)
	canvas.Refresh(renderer.wrapper)
}
func (renderer *WeatherDetailRenderer) Objects() []fyne.CanvasObject {
	return renderer.wrapper.Objects
}
func (renderer *WeatherDetailRenderer) Destroy() {}
