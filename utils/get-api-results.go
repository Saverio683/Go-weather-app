package utils

import (
	"log"
	"net/http"
	"weather-app/apptype"
)

type ApiResult struct {
	WeatherDetails []apptype.Detail
	MainData       apptype.MainData
}

const id = "17e57c3ab6fdc68fb851ae80c2f9c4b6"

func GetApiResults(city, country string) (ApiResult, error) { //makes the request to the OpenWeatherMap API
	url := "https://api.openweathermap.org/data/2.5/weather?q=" + city + "," + country + "&units=metric&appid=" + id
	resp, err := http.Get(url)

	if err != nil {
		log.Println("error retrieving weather api response", err)
		return ApiResult{}, err
	}
	defer resp.Body.Close()

	parsedResp, err := ParseJson(resp.Body)
	if err != nil {
		log.Println("error decoding response body")
		return ApiResult{}, err
	}

	md := apptype.MainData{
		Title:   "Current weather in " + city,
		ImgLink: "http://openweathermap.org/img/wn/" + GetStringFromAny(parsedResp.Weather[0]["icon"]) + "@2x.png",
		Temp:    GetStringFromAny(parsedResp.Main["temp"]),
	}

	wd := append(make([]apptype.Detail, 6),
		apptype.Detail{Key: "temp min:", Val: GetStringFromAny(parsedResp.Main["temp_min"]) + "°C"},
		apptype.Detail{Key: "temp max:", Val: GetStringFromAny(parsedResp.Main["temp_max"]) + "°C"},
		apptype.Detail{Key: "feels like:", Val: GetStringFromAny(parsedResp.Main["feels_like"]) + "°C"},
		apptype.Detail{Key: "humidity:", Val: GetStringFromAny(parsedResp.Main["humidity"]) + "°%"},
		apptype.Detail{Key: "sunrise:", Val: GetStringFromAny(parsedResp.Main["sunrise"])},
		apptype.Detail{Key: "sunset:", Val: GetStringFromAny(parsedResp.Main["sunset"])},
	)

	/* md := apptype.MainData{



	} */
	/* wd = append(wd,
		weatherdetail.NewWeatherDetail("temp min", GetStringFromAny(parsedResp.Main["temp_min"])+"°C", 0),
		weatherdetail.NewWeatherDetail("temp max", GetStringFromAny(parsedResp.Main["temp_max"])+"°C", 1),
		weatherdetail.NewWeatherDetail("feels like", GetStringFromAny(parsedResp.Main["feels_like"])+"°C", 2),
		weatherdetail.NewWeatherDetail("humidity", GetStringFromAny(parsedResp.Main["humidity"])+"%", 3),
		weatherdetail.NewWeatherDetail("sunrise", GetStringFromAny(parsedResp.Sys["sunrise"]), 4),
		weatherdetail.NewWeatherDetail("sunset", GetStringFromAny(parsedResp.Sys["sunset"]), 5),
	) */

	return ApiResult{
		WeatherDetails: wd,
		MainData:       md,
	}, nil
}
