package utils

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Detail struct {
	Key, Val string
}
type MainData struct {
	Title, Temp, ImgLink string
}

type ApiResult struct {
	WeatherDetails []Detail
	MainData       MainData
}

func milliToDate(m string) string { //to convert the sunrise and sunset fields, given by the api in milliseconds
	t, _ := strconv.ParseInt(m, 10, 64)

	hour := strconv.Itoa(time.Unix(t, 0).Hour())
	minute := strconv.Itoa(time.Unix(t, 0).Minute())

	return hour + ":" + minute
}

const id = "17e57c3ab6fdc68fb851ae80c2f9c4b6"

func GetApiResults(city, country string) (ApiResult, error) { //makes the request to the OpenWeatherMap API
	url := "https://api.openweathermap.org/data/2.5/weather?q=" + city + "," + country + "&units=metric&appid=" + id
	resp, err := http.Get(url)

	if err != nil { //error retrieving weather api response
		log.Print(err)
		return ApiResult{}, errors.New("Ops... something went wrong")
	}
	defer resp.Body.Close()

	parsedResp, err := ParseJson(resp.Body)
	if err != nil { //error decoding response body
		log.Print(err)
		return ApiResult{}, errors.New("Ops... something went wrong")
	}

	icon := GetStringFromAny(parsedResp.Weather[0]["icon"])
	if icon == "" { //error when no result found
		return ApiResult{}, errors.New("No result found")
	}

	md := MainData{
		Title:   "Current weather in " + cases.Title(language.English, cases.Compact).String(city),
		ImgLink: "http://openweathermap.org/img/wn/" + icon + "@2x.png",
		Temp:    GetStringFromAny(parsedResp.Main["temp"]) + "°C",
	}

	wd := append(make([]Detail, 0),
		Detail{Key: "temp min: ", Val: GetStringFromAny(parsedResp.Main["temp_min"]) + "°C"},
		Detail{Key: "temp max: ", Val: GetStringFromAny(parsedResp.Main["temp_max"]) + "°C"},
		Detail{Key: "feels like: ", Val: GetStringFromAny(parsedResp.Main["feels_like"]) + "°C"},
		Detail{Key: "humidity: ", Val: GetStringFromAny(parsedResp.Main["humidity"]) + "%"},
		Detail{Key: "sunrise: ", Val: milliToDate(GetStringFromAny(parsedResp.Sys["sunrise"]))},
		Detail{Key: "sunset: ", Val: milliToDate(GetStringFromAny(parsedResp.Sys["sunset"]))},
	)

	return ApiResult{
		MainData:       md,
		WeatherDetails: wd,
	}, nil
}
