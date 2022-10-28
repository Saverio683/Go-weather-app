package ui

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func CreateMainData(city, temp string) *fyne.Container {
	text := canvas.NewText("Current weather in "+city, color.White)
	text.TextSize = 18
	text.Move(fyne.NewPos(10, 20))

	temperature := canvas.NewText(temp, color.White)
	temperature.TextSize = 50

	//sample image
	r, _ := fyne.LoadResourceFromURLString("http://openweathermap.org/img/wn/09d@2x.png")
	image := canvas.NewImageFromResource(r)
	image.FillMode = canvas.ImageFillContain
	image.Resize(fyne.NewSize(100, 100))
	image.Move(fyne.NewPos(310, 0))

	//wrap the components to have them on the same line
	tempAndImg := container.NewWithoutLayout(
		temperature,
		image,
	)
	tempAndImg.Move(fyne.NewPos(10, 60))

	result := container.NewWithoutLayout(
		text,
		tempAndImg,
	)

	return result
}

func createDetail(k, v string, isSecondRow bool) *fyne.Container {
	key := canvas.NewText(k+": ", color.White)
	val := canvas.NewText(v, color.White)

	key.TextStyle.Bold = true

	var marginLeft float32 = 10
	var height float32

	if isSecondRow {
		height = 20
	} else {
		height = 40
	}

	key.Move(fyne.NewPos(marginLeft, height))
	val.Move(fyne.NewPos(key.MinSize().Width+marginLeft, height))

	result := container.NewWithoutLayout(
		key,
		val,
	)

	return result
}

func CreateDetails(t_min, t_max, h, p int, sr, st string) *fyne.Container {
	temp_min := createDetail("temp min", strconv.Itoa(t_min)+"°C", false)
	temp_max := createDetail("temp max", strconv.Itoa(t_max)+"°C", false)
	humidity := createDetail("humidity", strconv.Itoa(h)+"%", false)
	pressure := createDetail("pressure", strconv.Itoa(p)+"mb", true)
	sunrise := createDetail("sunrise", sr, true)
	sunset := createDetail("sunset", st, true)

	result := container.New(
		layout.NewGridLayout(3),
		temp_min,
		temp_max,
		humidity,
		pressure,
		sunrise,
		sunset,
	)

	return result
}
