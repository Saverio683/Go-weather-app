package weatherdetail

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type WeatherDetail struct {
	widget.BaseWidget
	Key, Value string
	Index      int
}

func NewWeatherDetail(key, value string, index int) *WeatherDetail {
	weatherDetail := &WeatherDetail{
		Key:   key,
		Value: value,
		Index: index,
	}
	weatherDetail.ExtendBaseWidget(weatherDetail)

	return weatherDetail
}

func (weatherDetail *WeatherDetail) CreateRenderer() fyne.WidgetRenderer {
	//creating widgets
	key := canvas.NewText(weatherDetail.Key+": ", color.White)
	value := canvas.NewText(weatherDetail.Value, color.White)

	key.TextStyle.Bold = true

	var marginLeft float32 = 10
	var height float32

	if weatherDetail.Index < 3 {
		height = 20
	} else {
		height = 40
	}

	key.Move(fyne.NewPos(marginLeft, height))
	value.Move(fyne.NewPos(key.MinSize().Width+marginLeft, height))

	wrapper := container.NewWithoutLayout(
		key,
		value,
	)

	return &WeatherDetailRenderer{
		key:     key,
		val:     value,
		index:   weatherDetail.Index,
		wrapper: wrapper,
	}
}
