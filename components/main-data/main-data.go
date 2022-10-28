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
	Title, Temp *canvas.Text
	ImgLink     string
	Image       *canvas.Image
	Wrapper     *fyne.Container
	Container   *fyne.Container
}

func NewMainData(title, temp *canvas.Text, imgLink string, image *canvas.Image, wrapper, container *fyne.Container) *MainData {
	MainData := &MainData{
		Title:     title,
		Temp:      temp,
		ImgLink:   imgLink,
		Image:     image,
		Wrapper:   wrapper,
		Container: container,
	}

	MainData.ExtendBaseWidget(MainData)

	return MainData
}

func (MainData *MainData) CreateRenderer() fyne.WidgetRenderer {
	//creating widgets
	title := canvas.NewText(MainData.Title.Text, color.White)
	temp := canvas.NewText(MainData.Temp.Text, color.White)

	//"http://openweathermap.org/img/wn/09d@2x.png"
	r, _ := fyne.LoadResourceFromURLString(MainData.ImgLink)
	image := canvas.NewImageFromResource(r)

	wrapper := container.NewWithoutLayout(
		temp,
		image,
	)

	container := container.NewWithoutLayout(
		title,
		wrapper,
	)

	return &MainDataRenderer{
		title:     title,
		temp:      temp,
		imgLink:   MainData.ImgLink,
		image:     image,
		wrapper:   wrapper,
		container: container,
	}
}
