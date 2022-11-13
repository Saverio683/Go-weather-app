package ui

import (
	weatherdetail "weather-app/components/weather-detail"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

/* func milliToDate(m string, isMorning bool) string {
	var result string
	var hoursToAdd int

	if isMorning {
		hoursToAdd = 0
	} else {
		hoursToAdd = 12
	}

	t, err := strconv.ParseInt(m, 10, 64)
	if err != nil {
		result = "error"
	} else {
		hour := strconv.Itoa(time.UnixMilli(t).UTC().Hour() + hoursToAdd)
		minute := strconv.Itoa(time.UnixMilli(t).UTC().Minute())
		result = hour + ":" + minute
	}

	return result
} */

func BuildWeatherDetails(details []*weatherdetail.WeatherDetail) *fyne.Container {
	weatherDetails := make([]fyne.CanvasObject, len(details))

	for i, detail := range details {
		weatherDetails[i] = detail
	}

	result := container.New(
		layout.NewGridLayout(3),
		weatherDetails...,
	)

	return result
}
