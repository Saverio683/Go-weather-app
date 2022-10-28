package apptype

import weatherdetail "weather-app/components/weather-detail"

type FormData struct {
	CityField, CountryField string
}

type MainData struct {
	City, Temp, ImgLink string
}

type State struct {
	FormData
	MainData
	Details []*weatherdetail.WeatherDetail
	Loading bool
	Error   bool
}
