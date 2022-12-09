package form

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Form struct {
	widget.BaseWidget
	submitFunc func(city, country string)
}

func NewForm(submitFunc func(city, country string)) *Form {
	form := &Form{
		submitFunc: submitFunc,
	}

	form.ExtendBaseWidget(form)
	return form
}

func (f *Form) SetSubmitFunc(submitFunc func(city, country string)) {
	f.submitFunc = submitFunc
}

func (form *Form) CreateRenderer() fyne.WidgetRenderer {
	//creating widgets
	cityEntry := widget.NewEntry()
	countrEntry := widget.NewEntry()
	button := widget.NewButton("Search", func() { form.submitFunc(cityEntry.Text, countrEntry.Text) })

	//setting placeholder in text input
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

	container := container.New(
		layout.NewVBoxLayout(),
		container.NewWithoutLayout(
			cityEntry,
			countrEntry,
		),
		container.NewWithoutLayout(
			button,
		),
	)

	return &FormRenderer{
		cityEntry:   cityEntry,
		countrEntry: countrEntry,
		button:      button,
		container:   container,
	}
}
