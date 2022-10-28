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
	//setting size
	renderer.cityEntry.Resize(fyne.NewSize(250, 40))
	renderer.countrEntry.Resize(fyne.NewSize(150, 40))
	renderer.button.Resize(fyne.NewSize(160, 40))

	//positioning
	renderer.cityEntry.Move(fyne.NewPos(10, 10))
	renderer.countrEntry.Move(fyne.NewPos(renderer.cityEntry.Size().Width+25, 10))
	renderer.button.Move(fyne.NewPos(10, 25))

	canvas.Refresh(renderer.container)
}
func (renderer *FormRenderer) Objects() []fyne.CanvasObject {
	return make([]fyne.CanvasObject, 0)
}
func (renderer *FormRenderer) Destroy() {}
