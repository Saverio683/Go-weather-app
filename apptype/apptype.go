package apptype

import (
	maindata "weather-app/components/main-data"
	weatherdetail "weather-app/components/weather-detail"
)

type State struct {
	MainData *maindata.MainData
	Details  []*weatherdetail.WeatherDetail
	Loading  bool
	Error    bool
}

type Detail struct {
	Key, Val string
	Index    int
}

type MainData struct {
	Title, Temp, ImgLink string
}

func SetDetails(details []Detail, state *State) {
	for i, detail := range state.Details {
		detail.Key = details[i].Key
		detail.Value = details[i].Val
		detail.Index = details[i].Index

		detail.Refresh()
	}
}

func SetMainData(data MainData, state *State) {
	state.MainData.Title = data.Title
	state.MainData.Temp = data.Temp
	state.MainData.ImgLink = data.ImgLink

	state.MainData.Refresh()
}
