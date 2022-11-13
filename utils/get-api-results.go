package utils

import (
	"log"
	"net/http"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Detail struct {
	Key, Val string
	Index    int
}
type MainData struct {
	Title, Temp, ImgLink string
}

type ApiResult struct {
	WeatherDetails []Detail
	MainData       MainData
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

	md := MainData{
		Title:   "Current weather in " + cases.Title(language.English, cases.Compact).String(city),
		ImgLink: "http://openweathermap.org/img/wn/" + GetStringFromAny(parsedResp.Weather[0]["icon"]) + "@2x.png",
		Temp:    GetStringFromAny(parsedResp.Main["temp"]) + "°C",
	}

	wd := append(make([]Detail, 0),
		Detail{Key: "temp min: ", Val: GetStringFromAny(parsedResp.Main["temp_min"]) + "°C"},
		Detail{Key: "temp max: ", Val: GetStringFromAny(parsedResp.Main["temp_max"]) + "°C"},
		Detail{Key: "feels like: ", Val: GetStringFromAny(parsedResp.Main["feels_like"]) + "°C"},
		Detail{Key: "humidity: ", Val: GetStringFromAny(parsedResp.Main["humidity"]) + "%"},
		Detail{Key: "sunrise: ", Val: /* GetStringFromAny(parsedResp.Sys["sunrise"]) */ "6:00"},
		Detail{Key: "sunset: ", Val: /* GetStringFromAny(parsedResp.Sys["sunset"]) */ "19:45"},
	)

	return ApiResult{
		MainData:       md,
		WeatherDetails: wd,
	}, nil
}
