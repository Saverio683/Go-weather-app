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

func (MainData *MainData) CreateRenderer() fyne.WidgetRenderer {
	//creating widgets
	title := canvas.NewText(MainData.Title, color.White)
	temp := canvas.NewText(MainData.Temp, color.White)

	title.TextSize = 18
	temp.TextSize = 50

	title.Move(fyne.NewPos(10, 20))

	//"http://openweathermap.org/img/wn/09d@2x.png"
	r, _ := fyne.LoadResourceFromURLString(MainData.ImgLink)
	image := canvas.NewImageFromResource(r)

	image.FillMode = canvas.ImageFillContain
	image.Resize(fyne.NewSize(100, 100))
	image.Move(fyne.NewPos(310, 0))

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
		wrapper:   wrapper,
		container: container,
	}
}
