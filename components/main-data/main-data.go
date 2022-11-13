package maindata

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type MainData struct {
	widget.BaseWidget
	Title, Temp, ImgLink string
}

func NewMainData(title, temp, imgLink string) *MainData {
	MainData := &MainData{
		Title:   title,
		Temp:    temp,
		ImgLink: imgLink,
	}

	MainData.ExtendBaseWidget(MainData)

	return MainData
}

func (m *MainData) SetFields(title, temp, imgLink string) {
	m.Title = title
	m.Temp = temp
	m.ImgLink = imgLink

	m.Refresh()
}

func (MainData *MainData) CreateRenderer() fyne.WidgetRenderer {
	//creating widgets
	title := canvas.NewText(MainData.Title, color.White)
	title.TextSize = 18
	title.Move(fyne.NewPos(10, 20))

	temp := canvas.NewText(MainData.Temp, color.White)
	temp.TextSize = 50

	r, _ := fyne.LoadResourceFromURLString(MainData.ImgLink)
	image := canvas.NewImageFromResource(r)

	image.FillMode = canvas.ImageFillContain
	image.Resize(fyne.NewSize(100, 100))
	image.Move(fyne.NewPos(310, -5))

	wrapper := container.NewWithoutLayout(
		temp,
		image,
	)
	wrapper.Move(fyne.NewPos(10, 60))

	container := container.NewWithoutLayout(
		title,
		wrapper,
	)

	return &MainDataRenderer{
		title:     title,
		temp:      temp,
		image:     image,
		container: container,
		mainData:  MainData,
	}
}
