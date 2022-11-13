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
	key, value string
	index      int
}

func NewWeatherDetail(k, v string, i int) *WeatherDetail {
	weatherDetail := &WeatherDetail{
		key:   k,
		value: v,
		index: i,
	}
	weatherDetail.ExtendBaseWidget(weatherDetail)

	return weatherDetail
}

func (w *WeatherDetail) SetFields(k, v string, i int) {
	w.key = k
	w.value = v
	w.index = i

	w.Refresh()
}

func (weatherDetail *WeatherDetail) CreateRenderer() fyne.WidgetRenderer {
	//creating widgets
	key := canvas.NewText(weatherDetail.key, color.White)
	value := canvas.NewText(weatherDetail.value, color.White)

	key.TextStyle.Bold = true

	container := container.NewWithoutLayout(
		key,
		value,
	)

	return &WeatherDetailRenderer{
		container:     container,
		key:           key,
		value:         value,
		weatherDetail: weatherDetail,
	}
}
