package weatherdetail

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type WeatherDetailRenderer struct {
	key, value    *canvas.Text
	container     *fyne.Container
	weatherDetail *WeatherDetail
}

func (renderer *WeatherDetailRenderer) MinSize() fyne.Size {
	return renderer.container.MinSize()
}
func (renderer *WeatherDetailRenderer) Layout(size fyne.Size) {
}
func (renderer *WeatherDetailRenderer) Refresh() {
	renderer.key.Text = renderer.weatherDetail.key
	renderer.value.Text = renderer.weatherDetail.value

	var (
		marginLeft float32 = 10
		y          float32
	)

	if renderer.weatherDetail.index < 3 {
		y = 35
	} else {
		y = 20
	}

	renderer.key.Move(fyne.NewPos(marginLeft, y))
	renderer.value.Move(fyne.NewPos(renderer.key.MinSize().Width+marginLeft, y))

	canvas.Refresh(renderer.value)
}
func (renderer *WeatherDetailRenderer) Objects() []fyne.CanvasObject {
	return renderer.container.Objects
}
func (renderer *WeatherDetailRenderer) Destroy() {}
