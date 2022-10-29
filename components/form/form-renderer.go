package form

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type FormRenderer struct {
	cityEntry, countrEntry *widget.Entry
	button                 *widget.Button
	container              *fyne.Container
}

func (renderer *FormRenderer) MinSize() fyne.Size {
	return renderer.container.Size()
}
func (renderer *FormRenderer) Layout(size fyne.Size) {
}
func (renderer *FormRenderer) Refresh() {
	canvas.Refresh(renderer.button)
	canvas.Refresh(renderer.cityEntry)
	canvas.Refresh(renderer.countrEntry)
	canvas.Refresh(renderer.container)
}
func (renderer *FormRenderer) Objects() []fyne.CanvasObject {
	return renderer.container.Objects
}
func (renderer *FormRenderer) Destroy() {}
