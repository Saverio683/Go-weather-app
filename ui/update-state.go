package ui

import (
	"weather-app/apptype"
	weatherdetail "weather-app/components/weather-detail"
	"weather-app/utils"
)

func UpdateState(city, country string, state *apptype.State) error { //the function state asynchronously, returns only the eventual error
	res, err := utils.GetApiResults(city, country) //making the api request

	if err != nil { //error handling
		return err
	} else {
		//updating main datas
		state.MainData.SetFields(res.MainData.Title, res.MainData.Temp, res.MainData.ImgLink)

		//updating details
		for i, detail := range state.Details {
			go func(i int, dt *weatherdetail.WeatherDetail) {
				key := res.WeatherDetails[i].Key
				value := res.WeatherDetails[i].Val
				dt.SetFields(key, value)
			}(i, detail)
		}

		return nil //no errors
	}
}
