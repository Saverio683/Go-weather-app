package ui

import (
	"strconv"
	"time"
	"weather-app/apptype"
	weatherdetail "weather-app/components/weather-detail"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func milliToDate(m string, isMorning bool) string {
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
		/* 		hours := int((t / (1000 * 3600)) % 24)
		   		minutes := int((t / (1000 * 60)) % 60) */
		hour := strconv.Itoa(time.UnixMilli(t).UTC().Hour() + hoursToAdd)
		minute := strconv.Itoa(time.UnixMilli(t).UTC().Minute())
		result = hour + ":" + minute
	}

	return result
}

func BuildWeatherDetails(state *apptype.State) *fyne.Container {
	weatherDetails := make([]fyne.CanvasObject, len(state.Details))

	for i := 0; i < len(state.Details); i++ {
		key := state.Details[i].Key
		val := state.Details[i].Value

		switch key {
		case "humidity":
			val += "%"
		case "sunrise":
			val = milliToDate(val, true)
		case "sunset":
			val = milliToDate(val, false)
		default:
			val += "°C"
		}

		w := weatherdetail.NewWeatherDetail(state.Details[i].Key, val, i)
		weatherDetails[i] = w
	}

	result := container.New(
		layout.NewGridLayout(3),
		weatherDetails...,
	)

	return result
}
