package form

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Form struct {
	widget.BaseWidget
	CityPlaceHolder, CountryPlaceholder string
	submitFunc                          func(a, b string)
}

func NewForm(cityPlaceHolder, countryPlaceholder string, submitFunc func(a, b string)) *Form {
	form := &Form{
		CityPlaceHolder:    cityPlaceHolder,
		CountryPlaceholder: countryPlaceholder,
		submitFunc:         submitFunc,
	}
	form.ExtendBaseWidget(form)

	return form
}

func (form *Form) CreateRenderer() fyne.WidgetRenderer {
	//creating widgets
	cityEntry := widget.NewEntry()
	countrEntry := widget.NewEntry()
	button := widget.NewButton("Search", func() { form.submitFunc(cityEntry.Text, countrEntry.Text) })

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

/* func CreateCustomForm(onSubmitFunc func(a, b string)) *fyne.Container {
	//creating widgets
	cityEntry := widget.NewEntry()
	countrEntry := widget.NewEntry()
	button := widget.NewButton("Search", func() { onSubmitFunc(cityEntry.Text, countrEntry.Text) })

	//set placeholder in text input
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

	result := container.New(
		layout.NewVBoxLayout(),
		container.NewWithoutLayout(
			cityEntry,
			countrEntry,
		),
		container.NewWithoutLayout(
			button,
		),
	)

	return result
} */
