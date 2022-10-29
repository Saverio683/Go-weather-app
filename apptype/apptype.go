package apptype

import (
	weatherdetail "weather-app/components/weather-detail"
)

type MainData struct {
	City, Temp, ImgLink string
}

type State struct {
	City, Country string
	Details       []*weatherdetail.WeatherDetail
	Loading       bool
	Error         bool
}
