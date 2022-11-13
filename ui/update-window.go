package ui

import (
	"weather-app/apptype"
	weatherdetail "weather-app/components/weather-detail"
	"weather-app/utils"
)

func UpdateWindow(city, country string, app *apptype.AppInit) (bool, error) {
	res, err := utils.GetApiResults(city, country)

	if err != nil {
		return false, err
	} else {
		app.State.MainData.SetFields(res.MainData.Title, res.MainData.Temp, res.MainData.ImgLink)

		for i, detail := range app.State.Details {
			go func(i int, dt *weatherdetail.WeatherDetail) {
				key := res.WeatherDetails[i].Key
				value := res.WeatherDetails[i].Val
				dt.SetFields(key, value, i)
			}(i, detail)
		}

		return true, nil
	}
}
